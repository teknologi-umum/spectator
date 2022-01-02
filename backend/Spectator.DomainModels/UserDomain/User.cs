using System;
using Spectator.DomainEvents.SessionDomain;

namespace Spectator.DomainModels.UserDomain {
	public record User {
		public string StudentNumber { get; }
		public int YearsOfExperience { get; }
		public int HoursOfPractice { get; }
		public string FamiliarLanguages { get; }
		public DateTimeOffset CreatedAt { get; }
		public DateTimeOffset UpdatedAt { get; private init; }

		private User(
			string studentNumber,
			int yearsOfExperience,
			int hoursOfPractice,
			string familiarLanguages,
			DateTimeOffset createdAt,
			DateTimeOffset updatedAt
		) {
			StudentNumber = studentNumber;
			YearsOfExperience = yearsOfExperience;
			HoursOfPractice = hoursOfPractice;
			FamiliarLanguages = familiarLanguages;
			CreatedAt = createdAt;
			UpdatedAt = updatedAt;
		}

		public static User From(PersonalInfoSubmittedEvent @event) => new(
			studentNumber: @event.StudentNumber,
			yearsOfExperience: @event.YearsOfExperience,
			hoursOfPractice: @event.HoursOfPractice,
			familiarLanguages: @event.FamiliarLanguages,
			createdAt: @event.Timestamp,
			updatedAt: @event.Timestamp
		);
	}
}
