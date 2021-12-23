using System;
using Spectator.DomainModels.SAMDomain;

namespace Spectator.DomainModels.EventDomain {
	public record BeforeTestSAMSubmittedEvent(
		Guid SessionId,
		DateTimeOffset Timestamp,
		SelfAssessmentManikin SelfAssessmentManikin
	) : EventBase(SessionId, Timestamp);
}
