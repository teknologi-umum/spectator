namespace Spectator.DomainEvents.SessionDomain {
	public record ExamEndedEvent(
		Guid SessionId,
		DateTimeOffset Timestamp
	) : SessionEventBase(SessionId, Timestamp);
}
