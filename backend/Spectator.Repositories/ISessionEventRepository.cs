using Spectator.DomainEvents.SessionDomain;
using Spectator.DomainModels.ExamReportDomain;

namespace Spectator.Repositories {
	public interface ISessionEventRepository {
		IAsyncEnumerable<Guid> GetAllSessionIdsAsync(AdministratorSession administratorSession, CancellationToken cancellationToken);
		IAsyncEnumerable<SessionEventBase> GetAllEventsAsync(Guid sessionId, CancellationToken cancellationToken);
		Task AddEventAsync(SessionEventBase @event);
		Task AddEventsAsync(SessionEventBase[] events);
	}
}
