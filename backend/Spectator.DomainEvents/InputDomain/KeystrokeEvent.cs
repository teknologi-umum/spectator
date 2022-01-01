using System;

namespace Spectator.DomainEvents.InputDomain {
	public record KeystrokeEvent(
		Guid SessionId,
		DateTimeOffset Timestamp,
		ConsoleKeyInfo KeystrokeInfo
	) : InputEventBase(SessionId, Timestamp);
}
