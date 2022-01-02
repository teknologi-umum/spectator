using Spectator.DomainEvents;

namespace Spectator.Observables.Redux {
	public delegate TState Reducer<TState, TEvent>(TState state, TEvent @event) where TEvent : IEvent;
}
