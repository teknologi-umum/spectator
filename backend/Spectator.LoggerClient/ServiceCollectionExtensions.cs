using Microsoft.Extensions.DependencyInjection;

namespace Spectator.LoggerClient {
	public static class ServiceCollectionExtensions {
		public static IServiceCollection AddLoggerClient(this IServiceCollection services) {
			services.AddSingleton<LoggerServices>();
			return services;
		}
	}
}
