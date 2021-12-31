using System.Threading.Tasks;
using Microsoft.AspNetCore.Authorization;
using Spectator.DomainModels.SessionDomain;
using Spectator.JwtAuthentication.Requirements;
using Spectator.Observables.SessionDomain;

namespace Spectator.JwtAuthentication.RequirementHandlers {
	public class SAMRequirementHandler : AuthorizationHandler<SAMRequirement> {
		private readonly SessionSilo _sessionSilo;

		public SAMRequirementHandler(
			SessionSilo sessionSilo
		) {
			_sessionSilo = sessionSilo;
		}

		protected override Task HandleRequirementAsync(AuthorizationHandlerContext context, SAMRequirement requirement) {
			try {
				var tokenPayload = TokenPayload.FromClaimsPrincipal(context.User);
				if (!_sessionSilo.TryGet(tokenPayload.SessionId, out var sessionStore)
					|| sessionStore.State is not RegisteredSession registeredSession) {
					context.Fail();
					return Task.CompletedTask;
				}

				if (requirement.BeforeTest.HasValue
					&& requirement.BeforeTest.Value != (registeredSession.BeforeExamSAM != null)) {
					context.Fail();
					return Task.CompletedTask;
				}

				if (requirement.AfterTest.HasValue
					&& requirement.AfterTest.Value != (registeredSession.AfterExamSAM != null)) {
					context.Fail();
					return Task.CompletedTask;
				}

				context.Succeed(requirement);
				return Task.CompletedTask;
			} catch {
				context.Fail();
				return Task.CompletedTask;
			}
		}
	}
}
