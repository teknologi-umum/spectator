using System;

namespace Spectator.DomainModels.SAMDomain {
	public record SelfAssessmentManikin(
		int ArousedLevel,
		int PleasedLevel,
		DateTimeOffset CreatedAt,
		DateTimeOffset UpdatedAt
	);
}
