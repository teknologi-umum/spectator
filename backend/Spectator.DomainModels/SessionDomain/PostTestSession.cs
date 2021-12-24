using System;
using Spectator.DomainModels.SAMDomain;
using Spectator.DomainModels.UserDomain;

namespace Spectator.DomainModels.SessionDomain {
	public record PostTestSession(
		Guid Id,
		DateTimeOffset CreatedAt,
		DateTimeOffset UpdatedAt,
		User User,
		SelfAssessmentManikin BeforeTestSAM,
		SelfAssessmentManikin AfterTestSAM
	) : SessionBase(Id, CreatedAt, UpdatedAt);
}
