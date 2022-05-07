#if DEBUG
using System;
using System.Threading;
using System.Threading.Tasks;
using FluentAssertions;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
using Xunit;

namespace Spectator.Piston.Tests {
	[Collection("PistonConsumer")]
	public class PistonClientTests {
		private IServiceProvider ServiceProvider { get; }

		public PistonClientTests() {
			var configuration = new ConfigurationBuilder()
				.AddKeyPerFile("/run/secrets", optional: true)
				.AddEnvironmentVariables("ASPNETCORE_")
				.AddUserSecrets<PistonClientTests>(optional: true)
				.Build();
			var services = new ServiceCollection();
			services.AddSingleton<IConfiguration>(configuration);
			services.Configure<PistonOptions>(configuration.GetSection("PistonOptions"));
			services.AddHttpClient();
			services.AddPistonClient();
			ServiceProvider = services.BuildServiceProvider();
		}

		[Fact]
		public async Task CanExecuteCCodeAsync() {
			var pistonClient = ServiceProvider.GetRequiredService<PistonClient>();

			// Only wait piston API for 10 seconds to save github CI quota
			using var timeoutSource = new CancellationTokenSource(TimeSpan.FromSeconds(10));

			// Wait 500ms to avoid HTTP 429
			await Task.Delay(TimeSpan.FromMilliseconds(500));

			var executeResult = await pistonClient.ExecuteAsync(
				language: "C",
				version: "9.3.0",
				code: @"#include <stdio.h>

						int main() {
							return 0;
						}",
				cancellationToken: timeoutSource.Token
			);

			executeResult.Compile.ExitCode.Should().Be(0);
			executeResult.Compile.Stdout.Should().BeEmpty();
			executeResult.Compile.Stderr.Should().BeEmpty();
			executeResult.Runtime.ExitCode.Should().Be(0);
			executeResult.Runtime.Stderr.Should().BeEmpty();
			executeResult.Runtime.Stdout.Should().BeEmpty();

			// Wait 500ms to avoid HTTP 429
			await Task.Delay(TimeSpan.FromMilliseconds(500));

			executeResult = await pistonClient.ExecuteAsync(
				language: "C",
				version: "9.3.0",
				code: @"#include <stdio.h>

						int main() {
							printf(""Hello world"");
							return 0;
						}",
				cancellationToken: timeoutSource.Token
			);

			executeResult.Compile.ExitCode.Should().Be(0);
			executeResult.Compile.Stdout.Should().BeEmpty();
			executeResult.Compile.Stderr.Should().BeEmpty();
			executeResult.Runtime.ExitCode.Should().Be(0);
			executeResult.Runtime.Stderr.Should().BeEmpty();
			executeResult.Runtime.Stdout.Should().Be("Hello world");

			// Wait 500ms to avoid HTTP 429
			await Task.Delay(TimeSpan.FromMilliseconds(500));

			executeResult = await pistonClient.ExecuteAsync(
				language: "C",
				version: "9.3.0",
				code: @"#include <stdio.h>

						int main() {
							return 1;
						}",
				cancellationToken: timeoutSource.Token
			);

			executeResult.Compile.ExitCode.Should().Be(0);
			executeResult.Compile.Stdout.Should().BeEmpty();
			executeResult.Compile.Stderr.Should().BeEmpty();
			executeResult.Runtime.ExitCode.Should().Be(1);
			executeResult.Runtime.Stderr.Should().BeEmpty();
			executeResult.Runtime.Stdout.Should().BeEmpty();
		}

		[Fact]
		public async Task CanReturnSyntaxErrorAsync() {
			var pistonClient = ServiceProvider.GetRequiredService<PistonClient>();

			// Only wait piston API for 5 seconds to save github CI quota
			using var timeoutSource = new CancellationTokenSource(TimeSpan.FromSeconds(5));

			// Wait 500ms to avoid HTTP 429
			await Task.Delay(TimeSpan.FromMilliseconds(500));

			var executeResult = await pistonClient.ExecuteAsync(
				language: "C",
				version: "9.3.0",
				code: @"#include <stdio.h>

						int main() {
							return 0
						}",
				cancellationToken: timeoutSource.Token
			);

			executeResult.Compile.ExitCode.Should().Be(1);
			executeResult.Compile.Stdout.Should().BeEmpty();
			executeResult.Compile.Stderr.Should().Be("code.c: In function 'main':\ncode.c:4:16: error: expected ';' before '}' token\n    4 |        return 0\n      |                ^\n      |                ;\n    5 |       }\n      |       ~         \n");
		}

		[Fact]
		public async Task CanExecuteTestsAsync() {
			var pistonClient = ServiceProvider.GetRequiredService<PistonClient>();

			using var timeoutSource = new CancellationTokenSource(TimeSpan.FromSeconds(5));

			var testResult = await pistonClient.ExecuteTestsAsync(language: Primitives.Language.Python,
				testCode: @"import random as __random
from decimal import Decimal

 def calculateTemperature(n, a, b):
        if a == ""Celcius"" and b == ""Fahrenheit"":
            return (n * 9 / 5) + 32

		elif a == ""Celcius"" and b == ""Kelvin"":
            return n + 273.15

		elif a == ""Fahrenheit"" and b == ""Celcius"":
            return (n - 32) * 5 / 9

		elif a == ""Fahrenheit"" and b == ""Kelvin"":
            return (n - 32) * 5 / 9 + 273.15

		elif a == ""Kelvin"" and b == ""Celcius"":
            return n - 273.15

		elif a == ""Kelvin"" and b == ""Fahrenheit"":
            return (n - 273.15) * 9 / 5 + 32


		return n

def main():
    testCases = [
        {
            ""got"": calculateTemperature(100, ""Celcius"", ""Fahrenheit""),

			""expected"": 212

		},
        {
            ""got"": calculateTemperature(212, ""Fahrenheit"", ""Kelvin""),
			""expected"": 373

		},
        {
            ""got"": calculateTemperature(0, ""Celcius"", ""Kelvin""),
            ""expected"": 273.15
        },
        	""got"": calculateTemperature(0, ""Celcius"", ""Fahrenheit""),
            ""expected"": 32
		},
        {
			""got"": calculateTemperature(0, ""Kelvin"", ""Fahrenheit""),
            ""expected"": -459.67

		}
    ]

    def workingAnswer(n, a, b):
        if a == ""Celcius"" and b == ""Fahrenheit"":
            return (n * 9 / 5) + 32

		elif a == ""Celcius"" and b == ""Kelvin"":
            return n + 273.15

		elif a == ""Fahrenheit"" and b == ""Celcius"":
            return (n - 32) * 5 / 9

		elif a == ""Fahrenheit"" and b == ""Kelvin"":
            return (n - 32) * 5 / 9 + 273.15

		elif a == ""Kelvin"" and b == ""Celcius"":
            return n - 273.15

		elif a == ""Kelvin"" and b == ""Fahrenheit"":
            return (n - 273.15) * 9 / 5 + 32


		return n


	temperatures = [""Celcius"", ""Fahrenheit"", ""Kelvin""]

	for _ in range(5):

		fromTemperature = __random.choice(temperatures)

		toTemperature = __random.choice(temperatures)

		n = __random.randint(-500, 500)

		expected = workingAnswer(n, fromTemperature, toTemperature)

		got = calculateTemperature(n, fromTemperature, toTemperature)

		testCases.append({ ""expected"": expected, ""got"": got })

    for i, test in enumerate(testCases):

		if round(float(test[""got""]), 2) == round(float(test[""expected""]), 2):
            print(f'# {i+1} PASSING')

		else:
            print(f""# {i+1} FAILED"")

			print(f""> EXPECTED { round(float(test['expected']), 2) }"")

			print(f""> GOT { round(float(test['got']), 2) }"")

if __name__ == ""__main__"":
    main()
",
			cancellationToken: timeoutSource.Token);
			testResult.Length.Equals(10);
			Console.WriteLine(testResult);
		}
	}
}
#endif
