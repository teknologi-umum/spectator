using System;

namespace Spectator.DomainModels.EventDomain.CodingEvents.MouseEvents {
	public record MouseUpEvent(
		Guid SessionId,
		DateTimeOffset Timestamp,
		int X,
		int Y
	) : MouseEventBase(SessionId, Timestamp, X, Y);
}
