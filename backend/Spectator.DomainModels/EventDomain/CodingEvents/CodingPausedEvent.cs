using System;

namespace Spectator.DomainModels.EventDomain.CodingEvents {
	public record CodingPausedEvent(
		Guid SessionId,
		DateTimeOffset Timestamp
	) : CodingEventBase(SessionId, Timestamp);
}
