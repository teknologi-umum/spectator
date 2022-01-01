using System;
using Spectator.Primitives;

namespace Spectator.DomainEvents.SessionDomain {
	public record BeforeExamSAMSubmittedEvent(
		Guid SessionId,
		DateTimeOffset Timestamp,
		SelfAssessmentManikin SelfAssessmentManikin
	) : SessionEventBase(SessionId, Timestamp);
}
