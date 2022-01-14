using System;
using Spectator.Primitives;

namespace Spectator.DomainEvents.SessionDomain {
	public record LocaleSetEvent(
		Guid SessionId,
		DateTimeOffset Timestamp,
		Locale Locale
	) : SessionEventBase(SessionId, Timestamp);
}
