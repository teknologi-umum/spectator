using Microsoft.Extensions.DependencyInjection;

namespace Spectator.VideoClient {
	public static class ServiceCollectionExtensions {
		public static IServiceCollection AddVideoClient(this IServiceCollection services) {
			services.AddSingleton<VideoServices>();
			return services;
		}
	}
}
