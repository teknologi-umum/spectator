using InfluxDB.Client;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Options;
using Spectator.Repositories;
using Spectator.RepositoryDALs.Internals;

namespace Spectator.RepositoryDALs {
	public static class ServiceCollectionExtensions {
		public static IServiceCollection AddRepositoryDALs(this IServiceCollection services) {
			services.AddSingleton(serviceProvider => {
				var options = serviceProvider.GetRequiredService<IOptions<InfluxDbOptions>>().Value;
				return InfluxDBClientFactory.Create(
					url: options.Url ?? throw new InvalidOperationException("InfluxDbOptions:Url is required"),
					token: options.Token ?? throw new InvalidOperationException("InfluxDbOptions:Token is required")
				);
			});
			services.AddSingleton<IDomainObjectMapper, DomainObjectMapper>();
			services.AddTransient<ISessionEventRepository, SessionEventRepositoryDAL>();
			return services;
		}
	}
}
