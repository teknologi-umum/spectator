using Microsoft.Extensions.DependencyInjection;

namespace Spectator.PoormansAuth {
	public static class ServiceCollectionExtensions {
		public static IServiceCollection AddPoormansAuth(this IServiceCollection services) {
			services.AddTransient<PoormansAuthentication>();
			return services;
		}
	}
}
