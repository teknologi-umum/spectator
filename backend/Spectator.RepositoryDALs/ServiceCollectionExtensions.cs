using Minio;
using InfluxDB.Client;
using InfluxDB.Client.Core.Exceptions;
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
			services.AddSingleton(async serviceProvider => {
				var options = serviceProvider.GetRequiredService<IOptions<InfluxDbOptions>>().Value;
				var client = InfluxDBClientFactory.Create(
					url: options.Url ?? throw new InvalidOperationException("InfluxDbOptions:Url is required"),
					token: options.Token ?? throw new InvalidOperationException("InfluxDbOptions:Token is required")
				);
				await InfluxDbInitialization.InitializeAsync(client, options);
				return client;
			});
			services.AddSingleton<IDomainObjectMapper, DomainObjectMapper>();
			services.AddTransient<IQuestionRepository, QuestionRepositoryDAL>();
			services.AddTransient<ISessionEventRepository, SessionEventRepositoryDAL>();
			services.AddTransient<IInputEventRepository, InputEventRepositoryDAL>();
			return services;
		}
	}
}
