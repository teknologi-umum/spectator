using System;

namespace Spectator.Users; 

public abstract class User {
	protected User(string studentNumber,
		int yearsOfExperience,
		int hoursOfPractice,
		string familiarLanguage,
		DateTime createdAt = default,
		DateTime updatedAt = default,
		bool finishedTest = false) {
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
	public int HoursOfPractice { get; set;  }
	public string FamiliarLanguage { get; set; }
	public DateTime CreatedAt { get; init; }
	public DateTime UpdatedAt { get; set; }
	public bool FinishedTest { get; set; }
}
