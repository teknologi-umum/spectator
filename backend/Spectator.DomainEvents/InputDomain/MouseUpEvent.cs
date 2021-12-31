namespace Spectator.DomainEvents.InputDomain {
	public record MouseUpEvent(
		Guid SessionId,
		DateTimeOffset Timestamp,
		int X,
		int Y
	) : InputEventBase(SessionId, Timestamp);
}
