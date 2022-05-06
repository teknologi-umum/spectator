using InfluxDB.Client;
using InfluxDB.Client.Core.Exceptions;
namespace Spectator.RepositoryDALs;

public static class InfluxDbInitialization {
	public static async Task InitializeAsync(InfluxDBClient client, InfluxDbOptions options) {
		var organizations = await client.GetOrganizationsApi()
			.FindOrganizationsAsync(limit: 1, org: options.Org);

		if (organizations.Count == 0) {
			throw new Exception("Organization was not found");
		}

		try {
			var bucket = await client.GetBucketsApi().FindBucketByNameAsync(options.InputEventsBucket);
			if (bucket == null) {
				throw new NotFoundException("Bucket was not found");
			}
		} catch (NotFoundException) {
			await client.GetBucketsApi()
				.CreateBucketAsync(options.InputEventsBucket, organizations[0]);
		}

		try {
			var bucket = await client.GetBucketsApi().FindBucketByNameAsync(options.SessionEventsBucket);
			if (bucket == null) {
				throw new NotFoundException("Bucket was not found");
			}
		} catch (NotFoundException) {
			await client.GetBucketsApi()
				.CreateBucketAsync(options.SessionEventsBucket, organizations[0]);
		}
	}
}
