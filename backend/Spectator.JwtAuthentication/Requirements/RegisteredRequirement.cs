using Microsoft.AspNetCore.Authorization;

namespace Spectator.JwtAuthentication.Requirements {
	public class RegisteredRequirement : IAuthorizationRequirement {
		public bool IsRegistered { get; }

		public RegisteredRequirement(
			bool isRegistered
		) {
			IsRegistered = isRegistered;
		}
	}
}
