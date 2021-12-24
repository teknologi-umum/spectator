using InfluxDB.Client;
using InfluxDB.Client.Api.Domain;
using InfluxDB.Client.Core.Flux.Domain;
using InfluxDB.Client.Writes;

namespace Spectator.RepositoryDALs.Internals {
	internal class DomainObjectMapper : IDomainObjectMapper {
		public T ConvertToEntity<T>(FluxRecord fluxRecord) => throw new NotImplementedException();
		public object ConvertToEntity(FluxRecord fluxRecord, Type type) => throw new NotImplementedException();
		public PointData ConvertToPointData<T>(T entity, WritePrecision precision) => throw new NotImplementedException();
	}
}
