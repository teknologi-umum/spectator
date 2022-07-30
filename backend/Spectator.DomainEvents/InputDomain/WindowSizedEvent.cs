using System;

namespace Spectator.DomainEvents.InputDomain {
	public record WindowSizedEvent(
		Guid SessionId,
		DateTimeOffset Timestamp,
		int QuestionNumber,
		int Width,
		int Height
	) : InputEventBase(SessionId, Timestamp);
}
