using System;

namespace Spectator.DomainEvents.SessionDomain {
	public record ExamForfeitedEvent(
		Guid SessionId,
		DateTimeOffset Timestamp
	) : SessionEventBase(SessionId, Timestamp);
}
