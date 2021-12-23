using System;

namespace Spectator.DomainModels.EventDomain.CodingEvents {
	public record CodingStartedEvent(
		Guid SessionId,
		DateTimeOffset Timestamp,
		int WindowWidth,
		int WindowHeight
	) : CodingEventBase(SessionId, Timestamp);
}
