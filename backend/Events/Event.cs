using System;

namespace Spectator.Events; 

public enum EventType {
	Keystroke,
	Mouse
}

public class Event {
	/**
	 * This event constructor takes 3 things:
	 * 1. type with a string value
	 * 2. value with a string value
	 * 3. date with an integer value (must be in UNIX millisecond format)
	 */
	protected Event(string user, string type, string value, int date) {
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
		Date = (new DateTime(1970, 1, 1)).AddMilliseconds(Convert.ToDouble(date));
	}

	public EventType Type { get; set; }
	public string Value { get; set; }
	public DateTime Date { get; init; }
	public string User { get; set; }
}
