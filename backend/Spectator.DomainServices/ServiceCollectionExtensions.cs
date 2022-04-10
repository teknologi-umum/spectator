using Microsoft.Extensions.DependencyInjection;
using Spectator.DomainServices.InputDomain;
using Spectator.DomainServices.PistonDomain;
using Spectator.DomainServices.QuestionDomain;
using Spectator.DomainServices.SessionDomain;

namespace Spectator.DomainServices {
	public static class ServiceCollectionExtensions {
		public static IServiceCollection AddDomainServices(this IServiceCollection services) {
			services.AddTransient<QuestionServices>();
			services.AddTransient<SessionServices>();
			services.AddTransient<InputServices>();
			services.AddTransient<SubmissionServices>();
			return services;
		}
	}
}
