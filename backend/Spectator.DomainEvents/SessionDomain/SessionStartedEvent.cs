using System;
using Spectator.Primitives;

namespace Spectator.DomainEvents.SessionDomain {
	public record SessionStartedEvent(
		Guid SessionId,
		DateTimeOffset Timestamp,
		Locale Locale
	) : SessionEventBase(SessionId, Timestamp);
}
