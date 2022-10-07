using System;

namespace Spectator.DomainEvents.SessionDomain {
	public record PersonalInfoSubmittedEvent(
		Guid SessionId,
		DateTimeOffset Timestamp,
		string Email,
		int Age,
		string Gender,
		string Nationality,
		string StudentNumber,
		int YearsOfExperience,
		int HoursOfPractice,
		string FamiliarLanguages,
		string WalletNumber,
		string WalletType
	) : SessionEventBase(SessionId, Timestamp);
}
