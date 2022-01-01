using System;

namespace Spectator.DomainEvents.InputDomain {
	public record MouseMovedEvent(
		Guid SessionId,
		DateTimeOffset Timestamp,
		int X,
		int Y
	) : InputEventBase(SessionId, Timestamp);
}
