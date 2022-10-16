#if DEBUG
using System;
using System.Threading;
using System.Threading.Tasks;
using FluentAssertions;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
using Spectator.DomainModels.SubmissionDomain;
using Spectator.DomainServices.PistonDomain;
using Spectator.Piston;
using Spectator.Primitives;
using Spectator.RepositoryDALs;
using Xunit;
using Xunit.Abstractions;

namespace Spectator.DomainServices.Tests {
	[Collection("PistonConsumer")]
	public class SubmissionServicesTests {
		private readonly ITestOutputHelper _testOutputHelper;
		private IServiceProvider ServiceProvider { get; }

		public SubmissionServicesTests(ITestOutputHelper testOutputHelper) {
			_testOutputHelper = testOutputHelper;
			var configuration = new ConfigurationBuilder()
				.AddKeyPerFile("/run/secrets", optional: true)
				.AddEnvironmentVariables("ASPNETCORE_")
				.AddUserSecrets<SubmissionServicesTests>(optional: true)
				.Build();
			var services = new ServiceCollection();
			services.AddSingleton<IConfiguration>(configuration);
			services.Configure<PistonOptions>(configuration.GetSection("PistonOptions"));
			services.AddHttpClient();
			services.AddMemoryCache();
			services.AddPistonClient();
			services.AddRepositoryDALs();
			services.AddDomainServices();
			ServiceProvider = services.BuildServiceProvider();
		}

		[Fact]
		public async Task CanAcceptCorrectSolutionAsync() {
			const string code = @"
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

	return n";

			var submissionServices = ServiceProvider.GetRequiredService<SubmissionServices>();

			// Only wait piston API for 5 seconds to save github CI quota
			using var timeoutSource = new CancellationTokenSource(TimeSpan.FromSeconds(5));

			var submission = await submissionServices.EvaluateSubmissionAsync(
				questionNumber: 2,
				locale: Locale.EN,
				language: Language.Python,
				directives: "",
				solution: code,
				scratchPad: "wkwkwkwk",
				cancellationToken: timeoutSource.Token
			);

			submission.Accepted.Should().BeTrue();
			submission.Language.Should().Be(Language.Python);
			submission.Solution.Should().Be(code);
			submission.ScratchPad.Should().Be("wkwkwkwk");
			submission.TestResults.Length.Should().Be(10);
			submission.TestResults.Should().AllBeOfType<PassingTestResult>();
		}

		[Fact]
		public async Task CanRejectIncorrectSolutionAsync() {
			const string code = @"
def calculateTemperature(n, a, b):
	if a == ""Celcius"" and b == ""Fahrenheit"":
		return (n * 9 / 5) + 32
	elif a == ""Celcius"" and b == ""Kelvin"":
		return n + 200 #                       <-- wrong formula
	elif a == ""Fahrenheit"" and b == ""Celcius"":
		return (n - 32) * 5 / 9
	elif a == ""Fahrenheit"" and b == ""Kelvin"":
		return (n - 32) * 5 / 9 + 273.15
	elif a == ""Kelvin"" and b == ""Celcius"":
		return n - 273.15
	elif a == ""Kelvin"" and b == ""Fahrenheit"":
		return (n - 273.15) * 9 / 5 + 32

	return n
			";

			var submissionServices = ServiceProvider.GetRequiredService<SubmissionServices>();

			// Only wait piston API for 5 seconds to save github CI quota
			using var timeoutSource = new CancellationTokenSource(TimeSpan.FromSeconds(5));

			var submission = await submissionServices.EvaluateSubmissionAsync(
				questionNumber: 2,
				locale: Locale.EN,
				language: Language.Python,
				directives: "",
				solution: code,
				scratchPad: "wlwlwlwl",
				cancellationToken: timeoutSource.Token
			);

			submission.Accepted.Should().BeFalse();
			submission.Language.Should().Be(Language.Python);
			submission.Solution.Should().Be(code);
			submission.ScratchPad.Should().Be("wlwlwlwl");
			submission.TestResults.Length.Should().Be(10);
			submission.TestResults[0].Should().BeOfType<PassingTestResult>();
			submission.TestResults[1].Should().BeOfType<PassingTestResult>();
			submission.TestResults[2].Should().BeOfType<FailingTestResult>().Which.ExpectedStdout.Should().Be("273.15");
			submission.TestResults[3].Should().BeOfType<PassingTestResult>();
			submission.TestResults[4].Should().BeOfType<PassingTestResult>();
		}

		[Fact]
		public async Task CanRejectSolutionWithCompileErrorAsync() {
			const string code = @"
				const std::string CELCIUS = ""Celcius"";
				const std::string FAHRENHEIT = ""Fahrenheit"" //                                 <-- missing semicolon
				const std::string KELVIN = ""Kelvin"";

				bool isCelcius(std::string unit) { return unit.compare(CELCIUS) == 0; }
				bool isFahrenheit(std::string unit) { return unit.compare(FAHRENHEIT) == 0; }
				bool isKelvin(std::string unit) { return unit.compare(KELVIN) == 0; }

				int calculateTemperature(int n, std::string from, std::string to) {
					if (isCelcius(from) && isFahrenheit(to)) return (n * 9 / 5) + 32;
					if (isCelcius(from) && isKelvin(to)) return n + 273.15;
					if (isFahrenheit(from) && isCelcius(to)) return (n - 32) * 5 / 9;
					if (isFahrenheit(from) && isKelvin(to)) return (n - 32) * 5 / 9 + 273.15;
					if (isKelvin(from) && isCelcius(to)) return n - 273.15;
					if (isKelvin(from) && isFahrenheit(to)) return (n - 273.15) * 9 / 5 + 32;
					return n;
				}
			";

			var submissionServices = ServiceProvider.GetRequiredService<SubmissionServices>();

			// Only wait piston API for 5 seconds to save github CI quota
			using var timeoutSource = new CancellationTokenSource(TimeSpan.FromSeconds(5));

			var submission = await submissionServices.EvaluateSubmissionAsync(
				questionNumber: 2,
				locale: Locale.EN,
				language: Language.CPP,
				directives: "",
				solution: code,
				scratchPad: "wkwkwk",
				cancellationToken: timeoutSource.Token
			);

			submission.Accepted.Should().BeFalse();
			submission.Language.Should().Be(Language.CPP);
			submission.Solution.Should().Be(code);
			submission.ScratchPad.Should().Be("wkwkwk");
			submission.TestResults.Length.Should().Be(1);
			submission.TestResults[0].Should().BeOfType<CompileErrorResult>().Which.Stderr.Should().Contain("error: expected ',' or ';'");
		}

		[Fact]
		public async Task CanRejectSolutionWithRuntimeErrorAsync() {
			const string code = @"
function calculateTemperature(n, from, to) {
	return a / 0;
}
			";

			var submissionServices = ServiceProvider.GetRequiredService<SubmissionServices>();

			// Only wait piston API for 5 seconds to save github CI quota
			using var timeoutSource = new CancellationTokenSource(TimeSpan.FromSeconds(5));

			var submission = await submissionServices.EvaluateSubmissionAsync(
				questionNumber: 2,
				locale: Locale.EN,
				language: Language.Javascript,
				directives: "",
				solution: code,
				scratchPad: "wkwkwkk",
				cancellationToken: timeoutSource.Token
			);

			submission.Accepted.Should().BeFalse();
			submission.Language.Should().Be(Language.Javascript);
			submission.Solution.Should().Be(code);
			submission.ScratchPad.Should().Be("wkwkwkk");
			submission.TestResults.Length.Should().Be(1);
			submission.TestResults[0].Should().BeOfType<RuntimeErrorResult>().Which.Stderr.Should().Contain("ReferenceError: a is not defined");
		}

		[Fact]
		public async Task CanCheckFirstQuestionUsingHardcodedCheckAsync() {
			const string correctCode = @"
def printLyrics():
    print(""Twinkle twinkle little star\nHow I wonder what you are\nUp above the world so high\nLike a diamond in the sky\nTwinkle twinkle little star\nHow I wonder what you are"")
";

			const string incorrectCode = @"
def printLyrics():
    print(""Twinkle twinkle little star\nHow I wonder what you are\nUp  above the world so high\nLike a diamond in the sky\nTwinkle twinkle little star\nHow I wonder what you are"")
";

			var submissionServices = ServiceProvider.GetRequiredService<SubmissionServices>();

			// Only wait piston API for 5 seconds to save github CI quota
			using var timeoutSource = new CancellationTokenSource(TimeSpan.FromSeconds(5));

			var submission = await submissionServices.EvaluateSubmissionAsync(
				questionNumber: 1,
				locale: Locale.EN,
				language: Language.Python,
				directives: "",
				solution: correctCode,
				scratchPad: "wkwkwkk",
				cancellationToken: timeoutSource.Token
			);

			submission.Accepted.Should().BeTrue();
			submission.Language.Should().Be(Language.Python);
			submission.Solution.Should().Be(correctCode);
			submission.ScratchPad.Should().Be("wkwkwkk");
			submission.TestResults.Length.Should().Be(1);
			submission.TestResults[0].Should().BeOfType<PassingTestResult>();

			submission = await submissionServices.EvaluateSubmissionAsync(
				questionNumber: 1,
				locale: Locale.EN,
				language: Language.Python,
				directives: "",
				solution: incorrectCode,
				scratchPad: "hahaha",
				cancellationToken: timeoutSource.Token
			);

			submission.Accepted.Should().BeFalse();
			submission.Language.Should().Be(Language.Python);
			submission.Solution.Should().Be(incorrectCode);
			submission.ScratchPad.Should().Be("hahaha");
			submission.TestResults.Length.Should().Be(1);
			submission.TestResults[0].Should().BeOfType<FailingTestResult>();
		}

		[Fact]
		public async Task CanCheckHelloWorldQuestionUsingHardcodedCheckAsync() {
			const string correctCode = @"
def helloWorld():
    print(""Hello world"")
";

			const string incorrectCode = @"
def helloWorld():
    print(""Haha"")
";

			var submissionServices = ServiceProvider.GetRequiredService<SubmissionServices>();

			// Only wait piston API for 5 seconds to save github CI quota
			using var timeoutSource = new CancellationTokenSource(TimeSpan.FromSeconds(5));

			var submission = await submissionServices.EvaluateSubmissionAsync(
				questionNumber: 0,
				locale: Locale.EN,
				language: Language.Python,
				directives: "",
				solution: correctCode,
				scratchPad: "Lorem ipsum",
				cancellationToken: timeoutSource.Token
			);

			submission.Accepted.Should().BeTrue();
			submission.Language.Should().Be(Language.Python);
			submission.Solution.Should().Be(correctCode);
			submission.ScratchPad.Should().Be("Lorem ipsum");
			submission.TestResults.Length.Should().Be(1);
			submission.TestResults[0].Should().BeOfType<PassingTestResult>();

			submission = await submissionServices.EvaluateSubmissionAsync(
				questionNumber: 0,
				locale: Locale.EN,
				language: Language.Python,
				directives: "",
				solution: incorrectCode,
				scratchPad: "Lorem ipsum",
				cancellationToken: timeoutSource.Token
			);

			submission.Accepted.Should().BeFalse();
			submission.Language.Should().Be(Language.Python);
			submission.Solution.Should().Be(incorrectCode);
			submission.ScratchPad.Should().Be("Lorem ipsum");
			submission.TestResults.Length.Should().Be(1);
			submission.TestResults[0].Should().BeOfType<FailingTestResult>();
		}

		[Fact]
		public async Task CanReturnFailingTestOnInvalidInput() {
			const string code = @"
# `calculateTemperature` is a function that accepts 3 arguments as its input:
# `temp` as integer, `from` as string, `to` as string. It returns a float as
# its output (it does not accept manual user inputs and print output).
# If there are any errors during running the test, please recheck your code and read the instructions carefully.
temp = int(input(""Temp : ""))
def calculateTemperature(temp, From, To):
	if  From == ""Celcius"" and To == ""Fahrenheit"":
		temp = (temp*9/5) + 32
		print(temp)
	elif From == ""Fahrenheit"" and To == ""Celcius"":
		temp = (temp - 32) * 5/9
		print(temp)
		# write your code here
";

			var submissionServices = ServiceProvider.GetRequiredService<SubmissionServices>();

			// Only wait piston API for 30 seconds to save github CI quota
			using var timeoutSource = new CancellationTokenSource(TimeSpan.FromSeconds(30));

			var submission = await submissionServices.EvaluateSubmissionAsync(
				questionNumber: 2,
				locale: Locale.EN,
				language: Language.Python,
				directives: "",
				solution: code,
				scratchPad: "Lorem ipsum",
				cancellationToken: timeoutSource.Token
			);

			submission.Accepted.Should().BeFalse();
			submission.Language.Should().Be(Language.Python);
			submission.Solution.Should().Be(code);
			submission.ScratchPad.Should().Be("Lorem ipsum");
			submission.TestResults.Length.Should().Be(1);
			submission.TestResults[0].Should().BeOfType<InvalidInputResult>();
		}
	}
}
#endif
