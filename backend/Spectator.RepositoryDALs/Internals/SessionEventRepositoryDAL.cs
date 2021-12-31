using InfluxDB.Client;
using InfluxDB.Client.Api.Domain;
using Spectator.DomainEvents.SessionDomain;
using Spectator.Repositories;

namespace Spectator.RepositoryDALs.Internals {
	internal class SessionEventRepositoryDAL : ISessionEventRepository {
		private readonly InfluxDBClient _db;
		private readonly IDomainObjectMapper _mapper;

		public SessionEventRepositoryDAL(
			InfluxDBClient db,
			IDomainObjectMapper mapper
		) {
			_db = db;
			_mapper = mapper;
		}

		public IAsyncEnumerable<SessionEventBase> GetAllEventsAsync(Guid sessionId, CancellationToken cancellationToken) {
			return _db.GetQueryApi(_mapper).QueryAsyncEnumerable<SessionEventBase>($"select * from events where sessionId = '{sessionId:N}'", cancellationToken);
		}

		public Task AddEventAsync(SessionEventBase @event) {
			return _db.GetWriteApiAsync(_mapper).WriteMeasurementAsync(WritePrecision.Ns, @event);
		}

		public Task AddEventsAsync(SessionEventBase[] events) {
			return _db.GetWriteApiAsync(_mapper).WriteMeasurementsAsync(WritePrecision.Ns, events);
		}
	}
}
