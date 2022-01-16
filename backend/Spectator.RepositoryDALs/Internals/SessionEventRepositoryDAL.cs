using InfluxDB.Client;
using InfluxDB.Client.Api.Domain;
using InfluxDB.Client.Linq;
using Microsoft.Extensions.Options;
using Spectator.DomainEvents.SessionDomain;
using Spectator.Repositories;

namespace Spectator.RepositoryDALs.Internals {
	internal class SessionEventRepositoryDAL : ISessionEventRepository {
		private readonly InfluxDBClient _db;
		private readonly string _bucket;
		private readonly string _org;
		private readonly IDomainObjectMapper _mapper;

		public SessionEventRepositoryDAL(
			InfluxDBClient db,
			IOptions<InfluxDbOptions> influxDbOptionsAccessor,
			IDomainObjectMapper mapper
		) {
			_db = db;
			var influxDbOptions = influxDbOptionsAccessor.Value;
			_bucket = influxDbOptions.SessionEventsBucket ?? throw new InvalidOperationException("InfluxDbOptions:SessionEventsBucket is required");
			_org = influxDbOptions.Org ?? throw new InvalidOperationException("InfluxDbOptions:Org is required");
			_mapper = mapper;
		}

		public IAsyncEnumerable<SessionEventBase> GetAllEventsAsync(Guid sessionId, CancellationToken cancellationToken) {
			return _db.GetQueryApi(_mapper).QueryAsyncEnumerable<SessionEventBase>($@"
				from(bucket: ""{_bucket}"")
				  |> range(start: 0)
				  |> pivot(rowKey:[""_time""], columnKey: [""_field""], valueColumn: ""_value"")
				  |> filter(fn: (r) => r[""session_id""] == ""{sessionId}"")
			", _org, cancellationToken);
		}

		public Task AddEventAsync(SessionEventBase @event) {
			return _db.GetWriteApiAsync(_mapper).WriteMeasurementAsync(_bucket, _org, WritePrecision.Ns, @event);
		}

		public Task AddEventsAsync(SessionEventBase[] events) {
			return _db.GetWriteApiAsync(_mapper).WriteMeasurementsAsync(_bucket, _org, WritePrecision.Ns, events);
		}
	}
}
