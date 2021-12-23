using System;

namespace Spectator.DomainModels.EventDomain.CodingEvents {
	public abstract record MouseEventBase(
		Guid SessionId,
		DateTimeOffset Timestamp,
		int X,
		int Y
	) : CodingEventBase(SessionId, Timestamp);
}
