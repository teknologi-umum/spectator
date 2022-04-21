using InfluxDB.Client;
using InfluxDB.Client.Api.Domain;
using Microsoft.Extensions.Options;
using Spectator.DomainEvents.InputDomain;
using Spectator.Repositories;

namespace Spectator.RepositoryDALs.Internals {
	internal class InputEventRepositoryDAL : IInputEventRepository {
		private readonly InfluxDBClient _db;
		private readonly string _bucket;
		private readonly string _org;
		private readonly IDomainObjectMapper _mapper;

		public InputEventRepositoryDAL(
			InfluxDBClient db,
			IOptions<InfluxDbOptions> influxDbOptionsAccessor,
			IDomainObjectMapper mapper
		) {
			_db = db;
			var influxDbOptions = influxDbOptionsAccessor.Value;
			_bucket = influxDbOptions.InputEventsBucket ?? throw new InvalidOperationException("InfluxDbOptions:InputEventsBucket is required");
			_org = influxDbOptions.Org ?? throw new InvalidOperationException("InfluxDbOptions:Org is required");
			_mapper = mapper;
		}

		public IAsyncEnumerable<InputEventBase> GetAllEventsAsync(Guid sessionId, CancellationToken cancellationToken) {
			return _db.GetQueryApi(_mapper).QueryAsyncEnumerable<InputEventBase>($@"
				from(bucket: ""{_bucket}"")
				  |> range(start: 0)
				  |> pivot(rowKey:[""_time""], columnKey: [""_field""], valueColumn: ""_value"")
				  |> filter(fn: (r) => r[""session_id""] == ""{sessionId}"")
			", _org, cancellationToken);
		}

		public Task AddEventAsync(InputEventBase @event) {
			return _db.GetWriteApiAsync(_mapper).WriteMeasurementAsync(_bucket, _org, WritePrecision.Ns, @event);
		}

		public Task AddEventsAsync(InputEventBase[] events) {
			return _db.GetWriteApiAsync(_mapper).WriteMeasurementsAsync(_bucket, _org, WritePrecision.Ns, events);
		}
	}
}
