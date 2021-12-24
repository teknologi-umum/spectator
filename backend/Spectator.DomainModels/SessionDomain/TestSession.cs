using System;
using Spectator.DomainModels.EventDomain;
using Spectator.DomainModels.SAMDomain;
using Spectator.DomainModels.UserDomain;

namespace Spectator.DomainModels.SessionDomain {
	public record TestSession(
		Guid Id,
		DateTimeOffset CreatedAt,
		DateTimeOffset UpdatedAt,
		User User,
		SelfAssessmentManikin BeforeTestSAM
	) : SessionBase(Id, CreatedAt, UpdatedAt) {
		public PostTestSession Apply(AfterTestSAMSubmittedEvent @event) {
			if (@event.SessionId != Id) throw new ArgumentException("Applied event has different SessionId", nameof(@event));

			return new(
				Id: Id,
				CreatedAt: CreatedAt,
				UpdatedAt: @event.Timestamp,
				User: User,
				BeforeTestSAM: BeforeTestSAM,
				AfterTestSAM: @event.SelfAssessmentManikin
			);
		}
	}
}
