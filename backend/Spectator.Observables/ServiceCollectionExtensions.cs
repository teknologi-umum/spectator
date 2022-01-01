using Microsoft.Extensions.DependencyInjection;
using Spectator.Observables.SessionDomain;

namespace Spectator.Observables {
	public static class ServiceCollectionExtensions {
		public static IServiceCollection AddObservables(this IServiceCollection services) {
			services.AddSingleton<SessionSilo>();
			return services;
		}
	}
}
