using InfluxDB.Client;
using InfluxDB.Client.Api.Domain;
using InfluxDB.Client.Core.Flux.Domain;
using InfluxDB.Client.Writes;
using Spectator.DomainEvents.InputDomain;
using Spectator.DomainEvents.SessionDomain;

namespace Spectator.RepositoryDALs.Mapper {
	public class DomainObjectMapper : IDomainObjectMapper {
		private static readonly Dictionary<Type, SessionEventMapper> SESSION_EVENT_MAPPER_BY_TYPE = new();
		private static readonly Dictionary<Type, InputEventMapper> INPUT_EVENT_MAPPER_BY_TYPE = new();
		private static readonly object SESSION_EVENT_MAPPERS_GATE = new();
		private static readonly object INPUT_EVENT_MAPPERS_GATE = new();

		public T ConvertToEntity<T>(FluxRecord fluxRecord) {
			if (typeof(SessionEventBase).IsAssignableFrom(typeof(T))) {
				var eventType = SessionEventMapper.EventTypeFor(fluxRecord);
				return (T)(object)GetSessionEventMapper(eventType).ConvertToEntity(fluxRecord);
			} else if (typeof(InputEventBase).IsAssignableFrom(typeof(T))) {
				var eventType = InputEventMapper.EventTypeFor(fluxRecord);
				return (T)(object)GetInputEventMapper(eventType).ConvertToEntity(fluxRecord);
			} else {
				throw new InvalidOperationException($"DomainObjectMapper doesn't support {typeof(T)}");
			}
		}

		public object ConvertToEntity(FluxRecord fluxRecord, Type type) {
			if (typeof(SessionEventBase).IsAssignableFrom(type)) {
				var eventType = SessionEventMapper.EventTypeFor(fluxRecord);
				return GetSessionEventMapper(eventType).ConvertToEntity(fluxRecord);
			} else if (typeof(InputEventBase).IsAssignableFrom(type)) {
				var eventType = InputEventMapper.EventTypeFor(fluxRecord);
				return GetInputEventMapper(eventType).ConvertToEntity(fluxRecord);
			} else {
				throw new InvalidOperationException($"DomainObjectMapper doesn't support {type}");
			}
		}

		public PointData ConvertToPointData<T>(T entity, WritePrecision precision) {
			if (entity is SessionEventBase sessionEvent) {
				return GetSessionEventMapper(entity.GetType()).ConvertToPointData(sessionEvent);
			} else if (entity is InputEventBase inputEvent) {
				return GetInputEventMapper(entity.GetType()).ConvertToPointData(inputEvent);
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

		private static InputEventMapper GetInputEventMapper(Type eventType) {
			lock (INPUT_EVENT_MAPPERS_GATE) {
				if (INPUT_EVENT_MAPPER_BY_TYPE.TryGetValue(eventType, out var eventMapper)) {
					return eventMapper;
				}
				eventMapper = InputEventMapper.For(eventType);
				INPUT_EVENT_MAPPER_BY_TYPE.Add(eventType, eventMapper);
				return eventMapper;
			}
		}
	}
}
