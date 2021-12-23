using System;

namespace Spectator.DomainModels.EventDomain.CodingEvents.MouseEvents {
	public record MouseMovedEvent(
		Guid SessionId,
		DateTimeOffset Timestamp,
		int X,
		int Y
	) : MouseEventBase(SessionId, Timestamp, X, Y);
}
