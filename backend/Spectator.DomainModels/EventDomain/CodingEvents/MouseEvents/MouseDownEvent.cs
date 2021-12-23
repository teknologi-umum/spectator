using System;
using Spectator.Primitives;

namespace Spectator.DomainModels.EventDomain.CodingEvents.MouseEvents {
	public record MouseDownEvent(
		Guid SessionId,
		DateTimeOffset Timestamp,
		int X,
		int Y,
		MouseButton Button
	) : MouseEventBase(SessionId, Timestamp, X, Y);
}
