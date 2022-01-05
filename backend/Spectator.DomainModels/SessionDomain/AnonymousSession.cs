using System;
using Spectator.DomainEvents.SessionDomain;
using Spectator.DomainModels.UserDomain;
using Spectator.Primitives;

namespace Spectator.DomainModels.SessionDomain {
	public record AnonymousSession : SessionBase {
		public Locale Locale { get; }

		protected AnonymousSession(
			Guid id,
			DateTimeOffset createdAt,
			DateTimeOffset updatedAt,
			Locale locale
		) : base(
			id: id,
			createdAt: createdAt,
			updatedAt: updatedAt
		) {
			Locale = locale;
		}

		public static AnonymousSession From(SessionStartedEvent @event) => new(
			id: @event.SessionId,
			createdAt: @event.Timestamp,
			updatedAt: @event.Timestamp,
			locale: @event.Locale
		);

		public AnonymousSession Apply(LocaleSetEvent @event) {
			if (@event.SessionId != Id) throw new ArgumentException("Applied event has different SessionId", nameof(@event));

			return new(
				id: Id,
				createdAt: CreatedAt,
				updatedAt: @event.Timestamp,
				locale: @event.Locale
			);
		}

		public RegisteredSession Apply(PersonalInfoSubmittedEvent @event) {
			if (@event.SessionId != Id) throw new ArgumentException("Applied event has different SessionId", nameof(@event));

			return new(
				Id: Id,
				CreatedAt: CreatedAt,
				UpdatedAt: @event.Timestamp,
				Locale: Locale,
				User: User.From(@event)
			);
		}
	}
}
