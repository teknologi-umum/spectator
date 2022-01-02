using System;

namespace Spectator.DomainModels.SessionDomain {
	public abstract record SessionBase {
		public Guid Id { get; }
		public DateTimeOffset CreatedAt { get; }
		public DateTimeOffset UpdatedAt { get; protected init; }

		public SessionBase(
			Guid id,
			DateTimeOffset createdAt,
			DateTimeOffset updatedAt
		) {
			if (updatedAt < createdAt) throw new ArgumentException("UpdatedAt cannot happen before CreatedAt", nameof(updatedAt));

			Id = id;
			CreatedAt = createdAt;
			UpdatedAt = updatedAt;
		}
	}
}
