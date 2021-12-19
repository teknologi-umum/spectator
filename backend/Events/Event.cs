using System;
using Spectator.Users;

namespace Spectator.Events;

public abstract class EventBase {
	/// <summary>
	/// This event constructor takes 3 things:
	/// 1. type with a string value
	/// 2. value with a string value
	/// 3. unixTime with an integer value (must be in UNIX millisecond format)
	/// </summary>
	protected EventBase(User user, string type, string value, int unixTime) {
		// TODO: Refactor this user from a User class.
		User = user;
		Type = type switch {
			"keystroke" => EventType.Keystroke,
			"mouse" => EventType.Mouse,
			// I don't know if we should just throw an exception
			// or if there's a better way to handle this.
			_ => throw new Exception("Type is not identified")
		};
		Value = value;
		Date = DateTimeOffset.FromUnixTimeMilliseconds(unixTime).UtcDateTime;
	}

	public EventType Type { get; set; }
	public string Value { get; set; }
	public DateTime Date { get; init; }
	public User User { get; set; }
}
