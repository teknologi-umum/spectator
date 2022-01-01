using Spectator.DomainEvents.SessionDomain;

namespace Spectator.Repositories {
	public interface ISessionEventRepository {
		IAsyncEnumerable<SessionEventBase> GetAllEventsAsync(Guid sessionId, CancellationToken cancellationToken);
		Task AddEventAsync(SessionEventBase @event);
		Task AddEventsAsync(SessionEventBase[] events);
	}
}
