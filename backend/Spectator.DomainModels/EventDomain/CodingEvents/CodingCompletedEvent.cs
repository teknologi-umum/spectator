using System;

namespace Spectator.DomainModels.EventDomain.CodingEvents {
	public record CodingCompletedEvent(
		Guid SessionId,
		DateTimeOffset Timestamp
	) : CodingEventBase(SessionId, Timestamp);
}
