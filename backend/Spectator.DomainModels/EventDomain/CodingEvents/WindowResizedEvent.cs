using System;

namespace Spectator.DomainModels.EventDomain.CodingEvents {
	public record WindowResizedEvent(
		Guid SessionId,
		DateTimeOffset Timestamp,
		int Width,
		int Height
	) : CodingEventBase(SessionId, Timestamp);
}
