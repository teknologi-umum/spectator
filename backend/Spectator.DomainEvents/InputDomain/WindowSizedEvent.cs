using System;

namespace Spectator.DomainEvents.InputDomain {
	public record WindowSizedEvent(
		Guid SessionId,
		DateTimeOffset Timestamp,
		int Width,
		int Height
	) : InputEventBase(SessionId, Timestamp);
}
