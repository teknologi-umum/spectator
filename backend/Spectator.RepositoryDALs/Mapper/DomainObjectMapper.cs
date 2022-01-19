using InfluxDB.Client;
using InfluxDB.Client.Api.Domain;
using InfluxDB.Client.Core.Flux.Domain;
using InfluxDB.Client.Writes;
using Spectator.DomainEvents.SessionDomain;

namespace Spectator.RepositoryDALs.Mapper {
	public class DomainObjectMapper : IDomainObjectMapper {
		private static readonly Dictionary<Type, SessionEventMapper> SESSION_EVENT_MAPPER_BY_TYPE = new();
		private static readonly object SESSION_EVENT_MAPPERS_GATE = new();

		public T ConvertToEntity<T>(FluxRecord fluxRecord) {
			if (typeof(SessionEventBase).IsAssignableFrom(typeof(T))) {
				var eventType = SessionEventMapper.EventTypeFor(fluxRecord);
				return (T)(object)GetSessionEventMapper(eventType).ConvertToEntity(fluxRecord);
			} else {
				throw new InvalidOperationException($"DomainObjectMapper doesn't support {typeof(T)}");
			}
		}

		public object ConvertToEntity(FluxRecord fluxRecord, Type type) {
			if (typeof(SessionEventBase).IsAssignableFrom(type)) {
				var eventType = SessionEventMapper.EventTypeFor(fluxRecord);
				return GetSessionEventMapper(eventType).ConvertToEntity(fluxRecord);
			} else {
				throw new InvalidOperationException($"DomainObjectMapper doesn't support {type}");
			}
		}

		public PointData ConvertToPointData<T>(T entity, WritePrecision precision) {
			if (entity is SessionEventBase @event) {
				return GetSessionEventMapper(entity.GetType()).ConvertToPointData(@event);
			} else {
				throw new InvalidOperationException($"DomainObjectMapper doesn't support {typeof(T)}");
			}
		}

		private static SessionEventMapper GetSessionEventMapper(Type eventType) {
			lock (SESSION_EVENT_MAPPERS_GATE) {
				if (SESSION_EVENT_MAPPER_BY_TYPE.TryGetValue(eventType, out var eventMapper)) {
					return eventMapper;
				}
				eventMapper = SessionEventMapper.For(eventType);
				SESSION_EVENT_MAPPER_BY_TYPE.Add(eventType, eventMapper);
				return eventMapper;
			}
		}
	}
}
