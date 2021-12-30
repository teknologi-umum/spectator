namespace Spectator.Observables.Redux {
	public delegate TState Reducer<TState, TEvent>(TState state, TEvent @event) where TState : notnull where TEvent : IEvent;
}
