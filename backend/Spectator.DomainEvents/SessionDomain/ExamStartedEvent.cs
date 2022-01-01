using System;
using System.Collections.Immutable;

namespace Spectator.DomainEvents.SessionDomain {
	public record ExamStartedEvent(
		Guid SessionId,
		DateTimeOffset Timestamp,
		ImmutableArray<int> QuestionNumbers,
		DateTimeOffset Deadline
	) : SessionEventBase(SessionId, Timestamp);
}
