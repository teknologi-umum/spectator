using System;
using Spectator.Primitives;

namespace Spectator.DomainEvents.SessionDomain {
	public record SolutionAcceptedEvent(
		Guid SessionId,
		DateTimeOffset Timestamp,
		int QuestionNumber,
		Language Language,
		string Solution,
		string ScratchPad,
		string? ErrorMessage,
		string ConsoleOutput
	) : SessionEventBase(SessionId, Timestamp);
}
