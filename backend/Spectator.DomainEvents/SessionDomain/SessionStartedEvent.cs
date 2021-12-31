namespace Spectator.DomainEvents.SessionDomain {
	public record SessionStartedEvent(
		Guid SessionId,
		DateTimeOffset Timestamp
	) : SessionEventBase(SessionId, Timestamp);
}
