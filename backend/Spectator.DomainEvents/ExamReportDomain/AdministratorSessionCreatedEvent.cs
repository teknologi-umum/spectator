using System;

namespace Spectator.DomainEvents.ExamReportDomain {
	public record AdministratorSessionCreatedEvent(
		Guid SessionId,
		DateTimeOffset CreatedAt,
		DateTimeOffset ExpiresAt
	);
}
