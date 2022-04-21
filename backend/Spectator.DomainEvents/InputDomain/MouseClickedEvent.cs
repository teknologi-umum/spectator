using System;
using Spectator.Primitives;

namespace Spectator.DomainEvents.InputDomain {
	public record MouseClickedEvent(
		Guid SessionId,
		DateTimeOffset Timestamp,
		int X,
		int Y,
		MouseButton Button
	) : InputEventBase(SessionId, Timestamp);
}
