using System;
using FluentAssertions;
using InfluxDB.Client.Api.Domain;
using InfluxDB.Client.Core.Flux.Domain;
using Spectator.DomainEvents.InputDomain;
using Spectator.Primitives;
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

		[Fact]
		public void CanSerializeMouseMovedEvent() {
			var sessionId = Guid.NewGuid();
			var timestamp = DateTimeOffset.UtcNow;
			var @event = new MouseMovedEvent(
				SessionId: sessionId,
				Timestamp: timestamp,
				X: 100,
				Y: 200,
				Direction: MouseDirection.Right
			);
			var mapper = new DomainObjectMapper();
			var pointData = mapper.ConvertToPointData(@event, WritePrecision.Ns);
			var lineProtocol = pointData.ToLineProtocol();
			lineProtocol.Should().StartWith("mouse_moved,");
			lineProtocol.Should().Contain($"session_id={sessionId}");
			lineProtocol.Should().Contain("x=100");
			lineProtocol.Should().Contain("y=200");
			lineProtocol.Should().Contain("direction=\"Right\"");
		}

		[Fact]
		public void CanDeserializeMouseMovedEvent() {
			var sessionId = Guid.NewGuid();
			var fluxRecord = new FluxRecord(1) {
				Values = {
					{ "_measurement", "mouse_moved" },
					{ "session_id", sessionId.ToString() },
					{ "x", 100 },
					{ "y", 200 },
					{ "direction", "Right" }
				}
			};
			var mapper = new DomainObjectMapper();
			var @event = mapper.ConvertToEntity<MouseMovedEvent>(fluxRecord);
			@event.SessionId.Should().Be(sessionId);
			@event.X.Should().Be(100);
			@event.Y.Should().Be(200);
			@event.Direction.Should().Be(MouseDirection.Right);
		}
	}
}
