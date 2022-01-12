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

namespace Spectator.DomainServices.Tests {
	public class SubmissionServicesTests {
		private IServiceProvider ServiceProvider { get; }

		public SubmissionServicesTests() {
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
				const std::string CELCIUS = ""Celcius"";
				const std::string FAHRENHEIT = ""Fahrenheit"";
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
				solution: code,
				scratchPad: "wkwkwkwk",
				cancellationToken: timeoutSource.Token
			);

			submission.Accepted.Should().BeTrue();
			submission.Language.Should().Be(Language.CPP);
			submission.Solution.Should().Be(code);
			submission.ScratchPad.Should().Be("wkwkwkwk");
			submission.TestResults.Length.Should().Be(10);
			submission.TestResults.Should().AllBeOfType<PassingTestResult>();
		}

		[Fact]
		public async Task CanRejectIncorrectSolutionAsync() {
			const string code = @"
				const std::string CELCIUS = ""Celcius"";
				const std::string FAHRENHEIT = ""Fahrenheit"";
				const std::string KELVIN = ""Kelvin"";

				bool isCelcius(std::string unit) { return unit.compare(CELCIUS) == 0; }
				bool isFahrenheit(std::string unit) { return unit.compare(FAHRENHEIT) == 0; }
				bool isKelvin(std::string unit) { return unit.compare(KELVIN) == 0; }

				int calculateTemperature(int n, std::string from, std::string to) {
					if (isCelcius(from) && isFahrenheit(to)) return (n * 9 / 5) + 32;
					if (isCelcius(from) && isKelvin(to)) return n + 272.15; //                   <-- wrong formula
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
				solution: code,
				scratchPad: "wlwlwlwl",
				cancellationToken: timeoutSource.Token
			);

			submission.Accepted.Should().BeFalse();
			submission.Language.Should().Be(Language.CPP);
			submission.Solution.Should().Be(code);
			submission.ScratchPad.Should().Be("wlwlwlwl");
			submission.TestResults.Length.Should().Be(10);
			submission.TestResults[0].Should().BeOfType<PassingTestResult>();
			submission.TestResults[1].Should().BeOfType<PassingTestResult>();
			submission.TestResults[2].Should().BeOfType<FailingTestResult>().Which.ExpectedStdout.Should().Be("273");
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
				const std::string CELCIUS = ""Celcius"";
				const std::string FAHRENHEIT = ""Fahrenheit"";
				const std::string KELVIN = ""Kelvin"";

				bool isCelcius(std::string unit) { return unit.compare(CELCIUS) == 0; }
				bool isFahrenheit(std::string unit) { return unit.compare(FAHRENHEIT) == 0; }
				bool isKelvin(std::string unit) { return unit.compare(KELVIN) == 0; }

				int calculateTemperature(int n, std::string from, std::string to) {
					if (isCelcius(from) && isFahrenheit(to)) return (n * 9 / 5) + 32;
					if (isCelcius(NULL) && isKelvin(to)) return n + 273.15; //                   <-- pass NULL pointer
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
				solution: code,
				scratchPad: "wkwkwkk",
				cancellationToken: timeoutSource.Token
			);

			submission.Accepted.Should().BeFalse();
			submission.Language.Should().Be(Language.CPP);
			submission.Solution.Should().Be(code);
			submission.ScratchPad.Should().Be("wkwkwkk");
			submission.TestResults.Length.Should().Be(1);
			submission.TestResults[0].Should().BeOfType<RuntimeErrorResult>().Which.Stderr.Should().Contain("basic_string::_M_construct null not valid");
		}
	}
}
