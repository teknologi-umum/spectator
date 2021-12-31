namespace Spectator.DomainEvents.SessionDomain {
	public abstract record SessionEventBase(
		Guid SessionId,
		DateTimeOffset Timestamp
	) : IEvent;
}
