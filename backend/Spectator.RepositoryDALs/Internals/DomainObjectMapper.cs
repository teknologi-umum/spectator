using InfluxDB.Client;
using InfluxDB.Client.Api.Domain;
using InfluxDB.Client.Core.Flux.Domain;
using InfluxDB.Client.Writes;
using Spectator.DomainModels.EventDomain;

namespace Spectator.RepositoryDALs.Internals {
	internal class DomainObjectMapper : IDomainObjectMapper {
		private static readonly Dictionary<Type, EventMapper> EVENT_MAPPER_BY_TYPE = new();
		private static readonly object EVENT_MAPPERS_GATE = new();

		public T ConvertToEntity<T>(FluxRecord fluxRecord) {
			if (typeof(T) == typeof(EventBase)) {
				return GetEventMapper<T>().ConvertToEntity<T>(fluxRecord);
			} else {
				throw new InvalidOperationException($"DomainObjectMapper doesn't support {typeof(T)}");
			}
		}

		public object ConvertToEntity(FluxRecord fluxRecord, Type type) => throw new NotImplementedException();

		public PointData ConvertToPointData<T>(T entity, WritePrecision precision) {
			if (entity is EventBase @event) {
				return GetEventMapper<T>().ConvertToPointData(@event);
			} else {
				throw new InvalidOperationException($"DomainObjectMapper doesn't support {typeof(T)}");
			}
		}

		private static EventMapper GetEventMapper<T>() {
			lock (EVENT_MAPPERS_GATE) {
				if (EVENT_MAPPER_BY_TYPE.TryGetValue(typeof(T), out var eventMapper)) {
					return eventMapper;
				}
				eventMapper = EventMapper.For<T>();
				EVENT_MAPPER_BY_TYPE.Add(typeof(T), eventMapper);
				return eventMapper;
			}
		}
	}
}
