using System;
using Spectator.DomainModels.SessionDomain;
using Spectator.JwtAuthentication;
using Spectator.Observables.SessionDomain;

namespace Spectator.PoormansAuth {
	public class PoormansAuthentication {
		private readonly JwtAuthenticationServices _jwtAuthenticationServices;
		private readonly SessionSilo _sessionSilo;

		public PoormansAuthentication(
			JwtAuthenticationServices jwtAuthenticationServices,
			SessionSilo sessionSilo
		) {
			_jwtAuthenticationServices = jwtAuthenticationServices;
			_sessionSilo = sessionSilo;
		}

		public SessionBase Authenticate(string accessToken) {
			var tokenPayload = _jwtAuthenticationServices.DecodeToken(accessToken);
			if (tokenPayload.Expires < DateTimeOffset.UtcNow) throw new UnauthorizedAccessException("Access token expired");
			if (!_sessionSilo.TryGet(tokenPayload.SessionId, out var sessionStore)) throw new UnauthorizedAccessException("Session not found");
			return sessionStore.State;
		}
	}
}
