using System;
using Spectator.Primitives;

namespace Spectator.DomainEvents.SessionDomain {
	public record SolutionSAMSubmittedEvent(
		Guid SessionId,
		int QuestionNumber,
		DateTimeOffset Timestamp,
		SelfAssessmentManikin SelfAssessmentManikin
	) : SessionEventBase(SessionId, Timestamp);
}
