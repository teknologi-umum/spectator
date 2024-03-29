﻿using System;
using Spectator.Primitives;

namespace Spectator.DomainEvents.InputDomain {
	public record MouseMovedEvent(
		Guid SessionId,
		DateTimeOffset Timestamp,
		int QuestionNumber,
		int X,
		int Y,
		MouseDirection Direction
	) : InputEventBase(SessionId, Timestamp);
}
