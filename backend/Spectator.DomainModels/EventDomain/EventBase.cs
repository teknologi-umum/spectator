using System;

namespace Spectator.DomainModels.EventDomain {
	public abstract record EventBase(
		Guid SessionId,
		DateTimeOffset Timestamp
	);
}
