using System;
using FluentAssertions;
using InfluxDB.Client.Api.Domain;
using Spectator.DomainEvents.InputDomain;
using Spectator.RepositoryDALs.Mapper;
using Xunit;

namespace Spectator.RepositoryDALs.Tests {
	public class InputEventObjectMapperTests {
		[Fact]
		public void CanSerializeKeystrokeEvent() {
			var sessionId = Guid.NewGuid();
			var timestamp = DateTimeOffset.UtcNow;
			var @event = new KeystrokeEvent(
				SessionId: sessionId,
				Timestamp: timestamp,
				KeyChar: "V",
				Shift: true,
				Alt: false,
				Control: false,
				Meta: false,
				UnrelatedKey: false
			);
			var mapper = new DomainObjectMapper();
			var pointData = mapper.ConvertToPointData(@event, WritePrecision.Ns);
			var lineProtocol = pointData.ToLineProtocol();
			lineProtocol.Should().StartWith("keystroke,");
			lineProtocol.Should().Contain($"session_id={sessionId}");
			lineProtocol.Should().Contain("key_char=\"V\"");
			lineProtocol.Should().Contain("shift=true");
			lineProtocol.Should().Contain("alt=false");
			lineProtocol.Should().Contain("control=false");
			lineProtocol.Should().Contain("meta=false");
			lineProtocol.Should().Contain("unrelated_key=false");
		}
	}
}
