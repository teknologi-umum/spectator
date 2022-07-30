using System;

namespace Spectator.DomainEvents.InputDomain {
	public record KeystrokeEvent(
		Guid SessionId,
		DateTimeOffset Timestamp,
		int QuestionNumber,
		string KeyChar,
		bool Shift,
		bool Alt,
		bool Control,
		bool Meta,
		bool UnrelatedKey
	) : InputEventBase(SessionId, Timestamp);
}
