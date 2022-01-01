using System;

namespace Spectator.DomainEvents.InputDomain {
	public abstract record InputEventBase(
		Guid SessionId,
		DateTimeOffset Timestamp
	) : IEvent;
}
