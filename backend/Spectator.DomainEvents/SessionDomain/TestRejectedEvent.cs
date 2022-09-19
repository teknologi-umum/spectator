using System;
using Spectator.Primitives;

namespace Spectator.DomainEvents.SessionDomain {
	public record TestRejectedEvent(
		Guid SessionId,
		DateTimeOffset Timestamp,
		int QuestionNumber,
		Language Language,
		string Solution,
		string ScratchPad,
		string SerializedTestResults
	) : SessionEventBase(SessionId, Timestamp);
}
