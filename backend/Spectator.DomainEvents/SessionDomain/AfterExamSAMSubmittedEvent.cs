using Spectator.Primitives;

namespace Spectator.DomainEvents.SessionDomain {
	public record AfterExamSAMSubmittedEvent(
		Guid SessionId,
		DateTimeOffset Timestamp,
		SelfAssessmentManikin SelfAssessmentManikin
	) : SessionEventBase(SessionId, Timestamp);
}
