using System;
using FluentAssertions;
using InfluxDB.Client.Api.Domain;
using Spectator.DomainEvents.SessionDomain;
using Spectator.RepositoryDALs.Mapper;
using Xunit;

namespace Spectator.RepositoryDALs.Tests {
	public class ObjectMapperTests {
		[Fact]
		public void CanSerializeSessionStartedEvent() {
			var sessionId = Guid.NewGuid();
			var timestamp = DateTimeOffset.UtcNow;
			var @event = new SessionStartedEvent(
				SessionId: sessionId,
				Timestamp: timestamp
			);
			var mapper = new DomainObjectMapper();
			var pointData = mapper.ConvertToPointData(@event, WritePrecision.Ns);
			var lineProtocol = pointData.ToLineProtocol();
			lineProtocol.Should().BeEmpty();
		}
	}
}
