using System;
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
					context.Fail(new AuthorizationFailureReason(this, "Session not found"));
					return Task.CompletedTask;
				}

				if (requirement.BeforeExam.HasValue
					&& requirement.BeforeExam.Value != (registeredSession.BeforeExamSAM != null)) {
					context.Fail(new AuthorizationFailureReason(this, requirement.BeforeExam.Value ? "Before exam SAM not submitted" : "Before exam SAM already submitted"));
					return Task.CompletedTask;
				}

				if (requirement.AfterExam.HasValue
					&& requirement.AfterExam.Value != (registeredSession.AfterExamSAM != null)) {
					context.Fail(new AuthorizationFailureReason(this, requirement.AfterExam.Value ? "After exam SAM not submitted" : "After exam SAM already submitted"));
					return Task.CompletedTask;
				}

				context.Succeed(requirement);
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
