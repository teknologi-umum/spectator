using System;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Authorization;
using Spectator.DomainModels.SessionDomain;
using Spectator.JwtAuthentication.Requirements;
using Spectator.Observables.SessionDomain;

namespace Spectator.JwtAuthentication.RequirementHandlers {
	public class RegisteredRequirementHandler : AuthorizationHandler<RegisteredRequirement> {
		private readonly SessionSilo _sessionSilo;

		public RegisteredRequirementHandler(
			SessionSilo sessionSilo
		) {
			_sessionSilo = sessionSilo;
		}

		protected override Task HandleRequirementAsync(AuthorizationHandlerContext context, RegisteredRequirement requirement) {
			try {
				var tokenPayload = TokenPayload.FromClaimsPrincipal(context.User);
				if (!_sessionSilo.TryGet(tokenPayload.SessionId, out var sessionStore)) {
					context.Fail(new AuthorizationFailureReason(this, "Session not found"));
					return Task.CompletedTask;
				}

				if (requirement.IsRegistered == (sessionStore.State is RegisteredSession)) {
					context.Succeed(requirement);
				} else {
					context.Fail(new AuthorizationFailureReason(this, requirement.IsRegistered ? "User not registered" : "User already registered"));
				}
				return Task.CompletedTask;
			} catch (Exception exc) {
#if DEBUG
				context.Fail(new AuthorizationFailureReason(this, exc.Message));
#else
				context.Fail();
#endif
				return Task.CompletedTask;
			}
		}
	}
}
