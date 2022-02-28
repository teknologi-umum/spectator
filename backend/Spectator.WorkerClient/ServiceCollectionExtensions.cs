using Microsoft.Extensions.DependencyInjection;

namespace Spectator.WorkerClient {
	public static class ServiceCollectionExtensions {
		public static IServiceCollection AddWorkerClient(this IServiceCollection services) {
			services.AddSingleton<WorkerClient>();
			return services;
		}
	}
}
