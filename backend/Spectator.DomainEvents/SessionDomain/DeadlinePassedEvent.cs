using System;

namespace Spectator.DomainEvents.SessionDomain {
	public record DeadlinePassedEvent(
		Guid SessionId,
		DateTimeOffset Timestamp
	) : SessionEventBase(SessionId, Timestamp);
}
