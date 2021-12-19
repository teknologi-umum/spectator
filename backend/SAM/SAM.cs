using System;
using Spectator.Users;

namespace Spectator.SAM; 

public class Sam {
	/// <summary>
	/// SAM test is a self-assessment to the user to measure their current state
	/// of emotion. The class contains the data of SAM test submission for a user.
	/// </summary>
	/// <param name="user">The user class that's doing the SAM test</param>
	/// <param name="condition">Whether this SAM test is before the coding test or after</param>
	/// <param name="arousedLevel">From 1 to 10, how aroused is the user now</param>
	/// <param name="pleasedLevel">From 1 to 10, how pleased is the user now</param>
	/// <param name="dominantLevel">From 1 to 10, how dominant is the user now</param>
	/// <param name="createdAt">When the data is created</param>
	/// <param name="updatedAt">When the data is last updated</param>
	protected Sam(User user,
		ConditionType condition,
		int arousedLevel,
		int pleasedLevel,
		int dominantLevel,
		DateTime createdAt = default,
		DateTime updatedAt = default) {
		User = user;
		Condition = condition;
		ArousedLevel = arousedLevel;
		PleasedLevel = pleasedLevel;
		DominantLevel = dominantLevel;
		CreatedAt = createdAt;
		UpdatedAt = updatedAt;
	}
	
	public User User { get; set; }
	public ConditionType Condition { get; set; }
	public int ArousedLevel { get; set; }
	public int PleasedLevel { get; set; }
	public int DominantLevel { get; set; }
	public DateTime CreatedAt { get; init; }
	public DateTime UpdatedAt { get; set; }
}
