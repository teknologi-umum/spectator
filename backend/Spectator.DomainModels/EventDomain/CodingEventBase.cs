using System;

namespace Spectator.DomainModels.EventDomain {
	public abstract record CodingEventBase(
		Guid SessionId,
		DateTimeOffset Timestamp
	) : EventBase(SessionId, Timestamp);
}
