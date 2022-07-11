using InfluxDB.Client;
using InfluxDB.Client.Core.Exceptions;
using Microsoft.Extensions.Options;

namespace Spectator.RepositoryDALs;

public class InfluxDbInitializer {
	private readonly InfluxDBClient _influxDBClient;
	private readonly InfluxDbOptions _influxDbOptions;

	public InfluxDbInitializer(
		InfluxDBClient influxDBClient,
		IOptions<InfluxDbOptions> optionsAccessor
	) {
		_influxDBClient = influxDBClient;
		_influxDbOptions = optionsAccessor.Value;
	}

	public async Task InitializeAsync() {
		var organizations = await _influxDBClient.GetOrganizationsApi()
			.FindOrganizationsAsync(limit: 1, org: _influxDbOptions.Org);

		if (organizations.Count == 0) {
			throw new NotFoundException("Organization was not found");
		}

		try {
			var bucket = await _influxDBClient.GetBucketsApi().FindBucketByNameAsync(_influxDbOptions.InputEventsBucket);
			if (bucket == null) {
				throw new NotFoundException("Bucket was not found");
			}
		} catch (NotFoundException) {
			await _influxDBClient.GetBucketsApi()
				.CreateBucketAsync(_influxDbOptions.InputEventsBucket, organizations[0]);
		}

		try {
			var bucket = await _influxDBClient.GetBucketsApi().FindBucketByNameAsync(_influxDbOptions.SessionEventsBucket);
			if (bucket == null) {
				throw new NotFoundException("Bucket was not found");
			}
		} catch (NotFoundException) {
			await _influxDBClient.GetBucketsApi()
				.CreateBucketAsync(_influxDbOptions.SessionEventsBucket, organizations[0]);
		}
	}
}
