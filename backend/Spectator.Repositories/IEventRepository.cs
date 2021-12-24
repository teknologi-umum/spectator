using Spectator.DomainModels.EventDomain;

namespace Spectator.Repositories {
	public interface IEventRepository {
		IAsyncEnumerable<EventBase> GetAllEventsAsync(Guid sessionId, CancellationToken cancellationToken);
		Task AddEventAsync(EventBase @event);
		Task AddEventsAsync(EventBase[] events);
	}
}
