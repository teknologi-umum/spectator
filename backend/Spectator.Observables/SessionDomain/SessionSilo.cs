using System;
using System.Collections.Generic;
using System.Diagnostics.CodeAnalysis;
using Spectator.DomainModels.SessionDomain;

namespace Spectator.Observables.SessionDomain {
	public class SessionSilo {
		private readonly Dictionary<Guid, SessionStore> _storeById = new();
		private readonly object _gate = new();

		public bool TryGet(Guid sessionId, [NotNullWhen(true)] out SessionStore? sessionStore) {
			lock (_gate) {
				return _storeById.TryGetValue(sessionId, out sessionStore);
			}
		}

		public bool TryAdd(SessionBase session) {
			lock (_gate) {
				if (_storeById.ContainsKey(session.Id)) {
					return false;
				}
				var sessionStore = new SessionStore(session);
				_storeById.Add(session.Id, sessionStore);
				return true;
			}
		}

		public bool TryAdd(SessionStore sessionStore) {
			if (sessionStore.State == null) return false;
			lock (_gate) {
				var sessionId = sessionStore.State.Id;
				if (_storeById.ContainsKey(sessionId)) {
					return false;
				}
				_storeById.Add(sessionId, sessionStore);
				return true;
			}
		}
	}
}
