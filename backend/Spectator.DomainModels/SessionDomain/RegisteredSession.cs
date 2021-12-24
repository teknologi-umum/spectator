using System;
using Spectator.DomainModels.EventDomain;
using Spectator.DomainModels.UserDomain;

namespace Spectator.DomainModels.SessionDomain {
	public record RegisteredSession(
		Guid Id,
		DateTimeOffset CreatedAt,
		DateTimeOffset UpdatedAt,
		User User
	) : SessionBase(Id, CreatedAt, UpdatedAt) {
		public TestSession Apply(BeforeTestSAMSubmittedEvent @event) {
			if (@event.SessionId != Id) throw new ArgumentException("Applied event has different SessionId", nameof(@event));

			return new(
				Id: Id,
				CreatedAt: CreatedAt,
				UpdatedAt: @event.Timestamp,
				User: User,
				BeforeTestSAM: @event.SelfAssessmentManikin
			);
		}
	}
}
