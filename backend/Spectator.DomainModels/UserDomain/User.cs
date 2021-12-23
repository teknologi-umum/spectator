using System;
using Spectator.DomainModels.EventDomain;
using Spectator.DomainModels.EventDomain.CodingEvents;

namespace Spectator.DomainModels.UserDomain {
	public record User {
		public string StudentNumber { get; }
		public int YearsOfExperience { get; }
		public int HoursOfPractice { get; }
		public string FamiliarLanguages { get; }
		public bool FinishesTest { get; private init; }
		public DateTimeOffset CreatedAt { get; }
		public DateTimeOffset UpdatedAt { get; private init; }

		private User(
			string studentNumber,
			int yearsOfExperience,
			int hoursOfPractice,
			string familiarLanguages,
			bool finishesTest,
			DateTimeOffset createdAt,
			DateTimeOffset updatedAt
		) {
			StudentNumber = studentNumber;
			YearsOfExperience = yearsOfExperience;
			HoursOfPractice = hoursOfPractice;
			FamiliarLanguages = familiarLanguages;
			FinishesTest = finishesTest;
			CreatedAt = createdAt;
			UpdatedAt = updatedAt;
		}

		public static User From(PersonalInfoSubmittedEvent @event) => new(
			studentNumber: @event.StudentNumber,
			yearsOfExperience: @event.YearsOfExperience,
			hoursOfPractice: @event.HoursOfPractice,
			familiarLanguages: @event.FamiliarLanguages,
			finishesTest: false,
			createdAt: @event.Timestamp,
			updatedAt: @event.Timestamp
		);

		public User Apply(CodingCompletedEvent @event) => this with {
			FinishesTest = true,
			UpdatedAt = @event.Timestamp
		};
	}
}
