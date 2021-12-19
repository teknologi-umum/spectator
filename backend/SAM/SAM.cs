using System;
using Spectator.Users;

namespace Spectator.SAM; 

public class Sam {
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
