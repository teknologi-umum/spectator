using System;

namespace Spectator.DomainEvents.SessionDomain {
	public record ExamIDEReloadedEvent(
		Guid SessionId,
		DateTimeOffset Timestamp
	) : SessionEventBase(SessionId, Timestamp);
}
