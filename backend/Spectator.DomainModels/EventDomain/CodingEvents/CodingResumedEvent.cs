using System;

namespace Spectator.DomainModels.EventDomain.CodingEvents {
	public record CodingResumedEvent(
		Guid SessionId,
		DateTimeOffset Timestamp
	) : CodingEventBase(SessionId, Timestamp);
}
