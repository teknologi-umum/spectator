using Spectator.DomainEvents.InputDomain;

namespace Spectator.Repositories {
	public interface IInputEventRepository {
		IAsyncEnumerable<InputEventBase> GetAllEventsAsync(Guid sessionId, CancellationToken cancellationToken);
		Task AddEventAsync(InputEventBase @event);
		Task AddEventsAsync(InputEventBase[] events);
	}
}
