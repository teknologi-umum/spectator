using System;
using System.Threading.Tasks;
using FluentAssertions;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
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

				bool isCelsius(std::string unit) { return unit.compare(CELCIUS) == 0; }
				bool isFahrenheit(std::string unit) { return unit.compare(FAHRENHEIT) == 0; }
				bool isKelvin(std::string unit) { return unit.compare(KELVIN) == 0; }
			
				int calculateTemperature(int temp, std::string from, std::string to) {
					if (isCelsius(a) && isFahrenheit(b)) return (n * 9 / 5) + 32;
					if (isCelsius(a) && isKelvin(b)) return n + 273.15;
					if (isFahrenheit(a) && isCelsius(b)) return (n - 32) * 5 / 9;
					if (isFahrenheit(a) && isKelvin(b)) return (n - 32) * 5 / 9 + 273.15;
					if (isKelvin(a) && isCelsius(b)) return n - 273.15;
					if (isKelvin(a) && isFahrenheit(b)) return (n - 273.15) * 9 / 5 + 32;
					return n;
				}
			";

			var submissionServices = ServiceProvider.GetRequiredService<SubmissionServices>();

			var submission = await submissionServices.EvaluateSubmissionAsync(
				questionNumber: 2,
				locale: Locale.EN,
				language: Language.CPP,
				solution: code,
				scratchPad: "wkwkwkwk"
			);

			submission.Accepted.Should().BeTrue();
		}
	}
}
