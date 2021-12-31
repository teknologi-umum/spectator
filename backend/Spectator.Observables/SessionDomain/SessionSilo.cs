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

		public bool TryAdd(SessionBase session, [NotNullWhen(true)] out SessionStore? sessionStore) {
			lock (_gate) {
				if (_storeById.ContainsKey(session.Id)) {
					sessionStore = null;
					return false;
				}
				sessionStore = new SessionStore(session);
				_storeById.Add(session.Id, sessionStore);
				return true;
			}
		}
	}
}
