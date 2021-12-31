using InfluxDB.Client;
using InfluxDB.Client.Api.Domain;
using InfluxDB.Client.Core.Flux.Domain;
using InfluxDB.Client.Writes;
using Spectator.DomainEvents.SessionDomain;

namespace Spectator.RepositoryDALs.Internals {
	internal class DomainObjectMapper : IDomainObjectMapper {
		private static readonly Dictionary<Type, SessionEventMapper> SESSION_EVENT_MAPPER_BY_TYPE = new();
		private static readonly object SESSION_EVENT_MAPPERS_GATE = new();

		public T ConvertToEntity<T>(FluxRecord fluxRecord) {
			if (typeof(SessionEventBase).IsAssignableFrom(typeof(T))) {
				return GetSessionEventMapper<T>().ConvertToEntity<T>(fluxRecord);
			} else {
				throw new InvalidOperationException($"DomainObjectMapper doesn't support {typeof(T)}");
			}
		}

		public object ConvertToEntity(FluxRecord fluxRecord, Type type) => throw new NotImplementedException();

		public PointData ConvertToPointData<T>(T entity, WritePrecision precision) {
			if (entity is SessionEventBase @event) {
				return GetSessionEventMapper<T>().ConvertToPointData(@event);
			} else {
				throw new InvalidOperationException($"DomainObjectMapper doesn't support {typeof(T)}");
			}
		}

		private static SessionEventMapper GetSessionEventMapper<T>() {
			lock (SESSION_EVENT_MAPPERS_GATE) {
				if (SESSION_EVENT_MAPPER_BY_TYPE.TryGetValue(typeof(T), out var eventMapper)) {
					return eventMapper;
				}
				eventMapper = SessionEventMapper.For<T>();
				SESSION_EVENT_MAPPER_BY_TYPE.Add(typeof(T), eventMapper);
				return eventMapper;
			}
		}
	}
}
