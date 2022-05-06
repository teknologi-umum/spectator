using Minio;
using InfluxDB.Client;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Options;
using Spectator.Repositories;
using Spectator.RepositoryDALs.Internals;
using Spectator.RepositoryDALs.Mapper;

namespace Spectator.RepositoryDALs {
	public static class ServiceCollectionExtensions {
		public static IServiceCollection AddRepositoryDALs(this IServiceCollection services) {
			services.AddSingleton(serviceProvider => {
				var options = serviceProvider.GetRequiredService<IOptions<MinioOptions>>().Value;
				return new MinioClient()
					.WithEndpoint(options.Url ?? throw new InvalidOperationException("MinioOptions:Url is required"))
					.WithCredentials(
						accessKey: options.AccessKey ?? throw new InvalidOperationException("MinioOptions:AccessKey is required"),
						secretKey: options.SecretKey ?? throw new InvalidOperationException("MinioOptions:SecretKey is required")
					);
			});
			services.AddSingleton(serviceProvider => {
				var options = serviceProvider.GetRequiredService<IOptions<InfluxDbOptions>>().Value;
				return InfluxDBClientFactory.Create(
					url: options.Url ?? throw new InvalidOperationException("InfluxDbOptions:Url is required"),
					token: options.Token ?? throw new InvalidOperationException("InfluxDbOptions:Token is required")
				);
			});
			services.AddSingleton<IDomainObjectMapper, DomainObjectMapper>();
			services.AddTransient<IQuestionRepository, QuestionRepositoryDAL>();
			services.AddTransient<ISessionEventRepository, SessionEventRepositoryDAL>();
			services.AddTransient<IInputEventRepository, InputEventRepositoryDAL>();
			return services;
		}
	}
}
