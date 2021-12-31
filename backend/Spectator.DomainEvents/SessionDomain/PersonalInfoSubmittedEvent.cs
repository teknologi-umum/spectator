namespace Spectator.DomainEvents.SessionDomain {
	public record PersonalInfoSubmittedEvent(
		Guid SessionId,
		DateTimeOffset Timestamp,
		string StudentNumber,
		int YearsOfExperience,
		int HoursOfPractice,
		string FamiliarLanguages
	) : SessionEventBase(SessionId, Timestamp);
}
