using InfluxDB.Client;
using Microsoft.Extensions.DependencyInjection;
using Spectator.Repositories;
using Spectator.RepositoryDALs.Internals;

namespace Spectator.RepositoryDALs {
	public static class ServiceCollectionExtensions {
		public static IServiceCollection AddRepositoryDALs(this IServiceCollection services) {
			services.AddSingleton<IDomainObjectMapper, DomainObjectMapper>();
			services.AddTransient<IEventRepository, EventRepositoryDAL>();
			return services;
		}
	}
}
