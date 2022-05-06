using System;
using FluentAssertions;
using InfluxDB.Client.Api.Domain;
using InfluxDB.Client.Core.Flux.Domain;
using Spectator.DomainEvents.SessionDomain;
using Spectator.Primitives;
using Spectator.RepositoryDALs.Mapper;
using Xunit;

namespace Spectator.RepositoryDALs.Tests {
	public class SessionEventObjectMapperTests {
		[Fact]
		public void CanSerializeSessionStartedEvent() {
			var sessionId = Guid.NewGuid();
			var timestamp = DateTimeOffset.UtcNow;
			var @event = new SessionStartedEvent(
				SessionId: sessionId,
				Timestamp: timestamp,
				Locale: Locale.ID
			);
			var mapper = new DomainObjectMapper();
			var pointData = mapper.ConvertToPointData(@event, WritePrecision.Ns);
			var lineProtocol = pointData.ToLineProtocol();
			lineProtocol.Should().StartWith("session_started,");
			lineProtocol.Should().Contain($"session_id={sessionId}");
			lineProtocol.Should().Contain("locale=\"ID\"");
		}

		[Fact]
		public void CanDeserializeSessionStartedEvent() {
			var sessionId = Guid.NewGuid();
			var fluxRecord = new FluxRecord(1) {
				Values = {
					{ "_measurement", "session_started" },
					{ "session_id", sessionId.ToString() },
					{ "locale", "ID" }
				}
			};
			var mapper = new DomainObjectMapper();
			var @event = mapper.ConvertToEntity<SessionStartedEvent>(fluxRecord);
			@event.SessionId.Should().Be(sessionId);
			@event.Locale.Should().Be(Locale.ID);
		}

		[Fact]
		public void CannotSerializeNonEventClasses() {
			var exc = new Exception();
			var mapper = new DomainObjectMapper();
			new Action(() => mapper.ConvertToPointData(exc, WritePrecision.Ns)).Should().Throw<InvalidOperationException>()
				.And.Message.Should().Be("DomainObjectMapper doesn't support System.Exception");
		}

		[Fact]
		public void CannotDeserializeToNonEventClasses() {
			var fluxRecord = new FluxRecord(1) {
				Values = {
					{ "_measurement", "exception" },
					{ "message", "Your app fucked up" }
				}
			};
			var mapper = new DomainObjectMapper();
			new Action(() => mapper.ConvertToEntity<Exception>(fluxRecord)).Should().Throw<InvalidOperationException>()
				.And.Message.Should().Be("DomainObjectMapper doesn't support System.Exception");
		}

		[Fact]
		public void CannotDeserializeToEntityOfWrongType() {
			var sessionId = Guid.NewGuid();
			var fluxRecord = new FluxRecord(1) {
				Values = {
					{ "_measurement", "session_started" },
					{ "session_id", sessionId.ToString() },
					{ "locale", "ID" }
				}
			};
			var mapper = new DomainObjectMapper();
			new Action(() => mapper.ConvertToEntity<ExamStartedEvent>(fluxRecord)).Should().Throw<InvalidCastException>();
		}

		[Fact]
		public void CanSerializeBeforeExamSAMSubmittedEvent() {
			var sessionId = Guid.NewGuid();
			var timestamp = DateTimeOffset.UtcNow;
			var @event = new BeforeExamSAMSubmittedEvent(
				SessionId: sessionId,
				Timestamp: timestamp,
				SelfAssessmentManikin: new SelfAssessmentManikin(2, 3)
			);
			var mapper = new DomainObjectMapper();
			var pointData = mapper.ConvertToPointData(@event, WritePrecision.Ns);
			var lineProtocol = pointData.ToLineProtocol();
			lineProtocol.Should().StartWith("before_exam_sam_submitted,");
			lineProtocol.Should().Contain($"session_id={sessionId}");
			lineProtocol.Should().Contain("aroused_level=2");
			lineProtocol.Should().Contain("pleased_level=3");
		}

		[Fact]
		public void CanDeserializeBeforeExamSAMSubmittedEvent() {
			var sessionId = Guid.NewGuid();
			var fluxRecord = new FluxRecord(1) {
				Values = {
					{ "_measurement", "before_exam_sam_submitted" },
					{ "session_id", sessionId.ToString() },
					{ "aroused_level", 2 },
					{ "pleased_level", 3 }
				}
			};
			var mapper = new DomainObjectMapper();
			var @event = mapper.ConvertToEntity<BeforeExamSAMSubmittedEvent>(fluxRecord);
			@event.SessionId.Should().Be(sessionId);
			@event.SelfAssessmentManikin.ArousedLevel.Should().Be(2);
			@event.SelfAssessmentManikin.PleasedLevel.Should().Be(3);
		}
	}
}
