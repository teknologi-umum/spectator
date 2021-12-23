using System;

namespace Spectator.DomainModels.EventDomain.CodingEvents {
	public record KeystrokeEvent(
		Guid SessionId,
		DateTimeOffset Timestamp,
		ConsoleKeyInfo KeystrokeInfo
	) : CodingEventBase(SessionId, Timestamp);
}
