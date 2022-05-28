using Microsoft.Extensions.DependencyInjection;

namespace Spectator.Piston {
	public static class ServiceCollectionExtensions {
		public static IServiceCollection AddPistonClient(this IServiceCollection services) {
			services.AddSingleton<PistonClient>();
			return services;
		}
	}
}
