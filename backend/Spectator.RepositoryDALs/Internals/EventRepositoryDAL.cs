using InfluxDB.Client;
using InfluxDB.Client.Api.Domain;
using Spectator.DomainModels.EventDomain;
using Spectator.Repositories;

namespace Spectator.RepositoryDALs.Internals {
	internal class EventRepositoryDAL : IEventRepository {
		private readonly InfluxDBClient _db;
		private readonly IDomainObjectMapper _mapper;

		public EventRepositoryDAL(
			InfluxDBClient db,
			IDomainObjectMapper mapper
		) {
			_db = db;
			_mapper = mapper;
		}

		public IAsyncEnumerable<EventBase> GetAllEventsAsync(Guid sessionId, CancellationToken cancellationToken) {
			return _db.GetQueryApi(_mapper).QueryAsyncEnumerable<EventBase>($"select * from events where sessionId = '{sessionId:N}'", cancellationToken);
		}

		public Task AddEventAsync(EventBase @event) {
			return _db.GetWriteApiAsync(_mapper).WriteMeasurementAsync(WritePrecision.Ms, @event);
		}

		public Task AddEventsAsync(EventBase[] events) {
			return _db.GetWriteApiAsync(_mapper).WriteMeasurementsAsync(WritePrecision.Ms, events);
		}
	}
}
