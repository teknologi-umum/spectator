using System;
using System.Collections.Generic;
using System.Reactive.Subjects;
using Spectator.DomainEvents;

namespace Spectator.Observables.Redux {
	public record Store<TState, TEvent> : IStore<TState, TEvent>, IDisposable where TEvent : IEvent {
		private readonly BehaviorSubject<TState> _subject;
		private readonly Reducer<TState, TEvent> _reducer;
		private bool _disposedValue;

		public TState State => _subject.Value;

		public Store(Reducer<TState, TEvent> reducer, TState initialValue, IEnumerable<TEvent>? initialEvents = null) {
			if (initialEvents != null) {
				foreach (var @event in initialEvents) {
					initialValue = reducer.Invoke(initialValue, @event);
				}
			}
			_subject = new BehaviorSubject<TState>(initialValue);
			_reducer = reducer;
		}

		public IDisposable Subscribe(IObserver<TState> observer) => _subject.Subscribe(observer);

		public virtual TEvent Dispatch(TEvent @event) {
			var newValue = _reducer.Invoke(_subject.Value, @event);
			_subject.OnNext(newValue);
			return @event;
		}

		protected virtual void Dispose(bool disposing) {
			if (!_disposedValue) {
				if (disposing) {
					// dispose managed state (managed objects)
					_subject.Dispose();
				}

				_disposedValue = true;
			}
		}

		public void Dispose() {
			// Do not change this code. Put cleanup code in 'Dispose(bool disposing)' method
			Dispose(disposing: true);
			GC.SuppressFinalize(this);
		}
	}

	public record Store<TState> : Store<TState, IEvent> {
		public Store(Reducer<TState, IEvent> reducer, TState initialValue) : base(reducer, initialValue) { }
	}
}
