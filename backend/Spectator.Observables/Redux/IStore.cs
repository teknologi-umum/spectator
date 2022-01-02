using System;
using Spectator.DomainEvents;

namespace Spectator.Observables.Redux {
	public interface IStore<TState, TEvent> : IObservable<TState> where TEvent : IEvent {
		TState State { get; }
		TEvent Dispatch(TEvent @event);
	}

	public interface IStore<TState> : IStore<TState, IEvent> { }
}
