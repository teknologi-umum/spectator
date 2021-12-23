using System;

namespace Spectator.DomainModels.EventDomain {
	public record SessionStartedEvent(
		Guid SessionId,
		DateTimeOffset Timestamp
	) : EventBase(SessionId, Timestamp);
}
