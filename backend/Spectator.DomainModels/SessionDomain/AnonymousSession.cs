using System;
using Spectator.DomainEvents.SessionDomain;
using Spectator.DomainModels.UserDomain;

namespace Spectator.DomainModels.SessionDomain {
	public record AnonymousSession : SessionBase {
		protected AnonymousSession(
			Guid id,
			DateTimeOffset createdAt,
			DateTimeOffset updatedAt
		) : base(
			id: id,
			createdAt: createdAt,
			updatedAt: updatedAt
		) { }

		public static AnonymousSession From(SessionStartedEvent @event) => new(
			id: @event.SessionId,
			createdAt: @event.Timestamp,
			updatedAt: @event.Timestamp
		);

		public RegisteredSession Apply(PersonalInfoSubmittedEvent @event) {
			if (@event.SessionId != Id) throw new ArgumentException("Applied event has different SessionId", nameof(@event));

			return new(
				Id: Id,
				CreatedAt: CreatedAt,
				UpdatedAt: @event.Timestamp,
				User: User.From(@event)
			);
		}
	}
}
