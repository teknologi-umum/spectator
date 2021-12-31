namespace Spectator.DomainEvents.InputDomain {
	public record MouseScrolledEvent(
		Guid SessionId,
		DateTimeOffset Timestamp,
		int X,
		int Y,
		int Delta
	) : InputEventBase(SessionId, Timestamp);
}
