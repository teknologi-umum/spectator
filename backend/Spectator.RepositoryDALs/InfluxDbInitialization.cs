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
			 await client.GetBucketsApi().FindBucketByNameAsync(options.InputEventsBucket);
		 } catch (NotFoundException) {
			 await client.GetBucketsApi()
				 .CreateBucketAsync(options.InputEventsBucket, organizations[0]);
		 }

		 try {
			 await client.GetBucketsApi().FindBucketByNameAsync(options.SessionEventsBucket);
		 } catch (NotFoundException) {
			 await client.GetBucketsApi()
				 .CreateBucketAsync(options.SessionEventsBucket, organizations[0]);
		 }
	 }
}
