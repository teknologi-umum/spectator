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
				language: "c",
				code: @"
					#include <stdio.h>

					int main() {
						return 0;
					}
				",
				cancellationToken: timeoutSource.Token
			);

			executeResult.Run.Code.Should().Be(0);
			executeResult.Run.Stdout.Should().BeEmpty();
			executeResult.Run.Stderr.Should().BeEmpty();
			executeResult.Run.Output.Should().BeEmpty();

			// Wait 500ms to avoid HTTP 429
			await Task.Delay(TimeSpan.FromMilliseconds(500));

			executeResult = await pistonClient.ExecuteAsync(
				language: "c",
				code: @"
					#include <stdio.h>

					int main() {
						printf(""Hello world"");
						return 0;
					}
				",
				cancellationToken: timeoutSource.Token
			);

			executeResult.Run.Code.Should().Be(0);
			executeResult.Run.Stdout.Should().Be("Hello world");
			executeResult.Run.Stderr.Should().BeEmpty();
			executeResult.Run.Output.Should().Be("Hello world");

			// Wait 500ms to avoid HTTP 429
			await Task.Delay(TimeSpan.FromMilliseconds(500));

			executeResult = await pistonClient.ExecuteAsync(
				language: "c",
				code: @"
					#include <stdio.h>

					int main() {
						return 1;
					}
				",
				cancellationToken: timeoutSource.Token
			);

			executeResult.Run.Code.Should().Be(1);
			executeResult.Run.Stdout.Should().BeEmpty();
			executeResult.Run.Stderr.Should().BeEmpty();
			executeResult.Run.Output.Should().BeEmpty();
		}

		[Fact]
		public async Task CanReturnSyntaxErrorAsync() {
			var pistonClient = ServiceProvider.GetRequiredService<PistonClient>();

			// Only wait piston API for 5 seconds to save github CI quota
			using var timeoutSource = new CancellationTokenSource(TimeSpan.FromSeconds(5));

			// Wait 500ms to avoid HTTP 429
			await Task.Delay(TimeSpan.FromMilliseconds(500));

			var executeResult = await pistonClient.ExecuteAsync(
				language: "c",
				code: @"
					#include <stdio.h>

					int main() {
						return 0
					}
				",
				cancellationToken: timeoutSource.Token
			);

			executeResult.Run.Code.Should().Be(127);
			executeResult.Run.Stdout.Should().BeEmpty();
			executeResult.Run.Stderr.Should().Be($"/piston/packages/gcc/{executeResult.Version}/run: line 6: ./a.out: No such file or directory\n");
			executeResult.Run.Output.Should().Be($"/piston/packages/gcc/{executeResult.Version}/run: line 6: ./a.out: No such file or directory\n");
		}
	}
}
