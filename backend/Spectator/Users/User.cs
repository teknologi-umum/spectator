using System;

namespace Spectator.Users;

public abstract class User {
	/// <summary>
	/// User contains the object of a user's personal
	/// information.
	/// </summary>
	/// <param name="studentNumber">Student number that will become a unique identity to a user</param>
	/// <param name="yearsOfExperience">Number of years they have been programming</param>
	/// <param name="hoursOfPractice">Number of hours that they'd practice in a day</param>
	/// <param name="familiarLanguage">List of the programming language that they are familiar in</param>
	/// <param name="finishedTest">Whether they finished the test</param>
	/// <param name="createdAt">When the user is created</param>
	/// <param name="updatedAt">When the data is last updated</param>
	protected User(string studentNumber,
		int yearsOfExperience,
		int hoursOfPractice,
		string familiarLanguage,
		bool finishedTest = false,
		DateTime createdAt = default,
		DateTime updatedAt = default
	) {
		StudentNumber = studentNumber;
		YearsOfExperience = yearsOfExperience;
		HoursOfPractice = hoursOfPractice;
		FamiliarLanguage = familiarLanguage;
		CreatedAt = createdAt;
		UpdatedAt = updatedAt;
		FinishedTest = finishedTest;
	}

	public string StudentNumber { get; init; }
	public int YearsOfExperience { get; set; }
	public int HoursOfPractice { get; set; }
	public string FamiliarLanguage { get; set; }
	public DateTime CreatedAt { get; init; }
	public DateTime UpdatedAt { get; set; }
	public bool FinishedTest { get; set; }
}
