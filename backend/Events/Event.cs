using System;
using Spectator.Users;

namespace Spectator.Events;

public abstract class EventBase {
	/// <summary>
	/// Specifies a base event that could be implemented into something else.
	/// </summary>
	/// <param name="user">User that's doing the event</param>
	/// <param name="type">The type of event the user's doing</param>
	/// <param name="payload">The payload of the current event</param>
	/// <param name="unixTime">The unix timestamp of the current event</param>
	/// <exception cref="ArgumentException">Will throw an exception if the event type is not valid</exception>
	protected EventBase(User user, string type, string payload, int unixTime) {
		// TODO: Refactor this user from a User class.
		User = user;
		Type = Enum.TryParse<EventType>(value: type, ignoreCase: true, result: out var eventType)
			? eventType
			: throw new ArgumentException($"Unidentified event type: {type}", nameof(type));
		Value = payload;
		Date = DateTimeOffset.FromUnixTimeMilliseconds(unixTime).UtcDateTime;
	}

	public EventType Type { get; set; }
	public string Value { get; set; }
	public DateTime Date { get; init; }
	public User User { get; set; }
}
