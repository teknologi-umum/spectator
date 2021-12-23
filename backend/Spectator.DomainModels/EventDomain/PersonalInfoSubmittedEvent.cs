using System;

namespace Spectator.DomainModels.EventDomain {
	public record PersonalInfoSubmittedEvent(
		Guid SessionId,
		DateTimeOffset Timestamp,
		string StudentNumber,
		int YearsOfExperience,
		int HoursOfPractice,
		string FamiliarLanguages
	) : EventBase(SessionId, Timestamp);
}
