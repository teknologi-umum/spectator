using System.Runtime.CompilerServices;
using InfluxDB.Client;
using InfluxDB.Client.Api.Domain;
using Microsoft.Extensions.Options;
using Spectator.DomainEvents.SessionDomain;
using Spectator.DomainModels.ExamReportDomain;
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

		public async IAsyncEnumerable<Guid> GetAllSessionIdsAsync(AdministratorSession administratorSession, [EnumeratorCancellation] CancellationToken cancellationToken) {
			// Authorize administrator session
			if (administratorSession == null) throw new UnauthorizedAccessException();

			// Get all session_started events
			var sessionEventsAsync = _db.GetQueryApi(_mapper).QueryAsyncEnumerable<SessionEventBase>($@"
				from(bucket: ""{_bucket}"")
				  |> range(start: 0)
				  |> pivot(rowKey:[""_time""], columnKey: [""_field""], valueColumn: ""_value"")
				  |> filter(fn: (r) => r[""_measurement""] == ""session_started"")
			", _org, cancellationToken);

			// For each session_started events, return only its SessionId
			await foreach (var sessionEvent in sessionEventsAsync) {
				yield return sessionEvent.SessionId;
			}
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
			return _db.GetWriteApiAsync(_mapper).WriteMeasurementAsync(@event, bucket: _bucket, org: _org, precision: WritePrecision.Ns);
		}

		public Task AddEventsAsync(SessionEventBase[] events) {
			return _db.GetWriteApiAsync(_mapper).WriteMeasurementsAsync(events, bucket: _bucket, org: _org, precision: WritePrecision.Ns);
		}
	}
}
