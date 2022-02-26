using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Spectator.DomainEvents.ExamReportDomain;

namespace Spectator.DomainModels.ExamReportDoman {
	public record AdministratorSession {
		public Guid SessionId { get; private init; }
		public DateTimeOffset ExpiresAt { get; private init; }
		public DateTimeOffset CreatedAt { get; private init; } 

		protected AdministratorSession(Guid sessionId, DateTimeOffset expiresAt, DateTimeOffset createdAt) {
				SessionId = sessionId;
				ExpiresAt = expiresAt;
				CreatedAt = createdAt;
		}

		public static AdministratorSession From(AdministratorSessionCreatedEvent @event) => new(
			sessionId: @event.SessionId,
			expiresAt: @event.ExpiresAt,
			createdAt: @event.CreatedAt
		);
	}
}
