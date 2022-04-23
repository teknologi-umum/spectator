using System;
using Spectator.Primitives;

namespace Spectator.DomainEvents.InputDomain {
	public record MouseUpEvent (
		Guid SessionId,
		DateTimeOffset Timestamp,
		int X,
		int Y,
		MouseButton Button
	) : InputEventBase(SessionId, Timestamp);
}
