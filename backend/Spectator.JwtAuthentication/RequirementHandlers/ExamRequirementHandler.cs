using System;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Authorization;
using Spectator.DomainModels.SessionDomain;
using Spectator.JwtAuthentication.Requirements;
using Spectator.Observables.SessionDomain;

namespace Spectator.JwtAuthentication.RequirementHandlers {
	public class ExamRequirementHandler : AuthorizationHandler<ExamRequirement> {
		private readonly SessionSilo _sessionSilo;

		public ExamRequirementHandler(
			SessionSilo sessionSilo
		) {
			_sessionSilo = sessionSilo;
		}

		protected override Task HandleRequirementAsync(AuthorizationHandlerContext context, ExamRequirement requirement) {
			try {
				var tokenPayload = TokenPayload.FromClaimsPrincipal(context.User);
				if (!_sessionSilo.TryGet(tokenPayload.SessionId, out var sessionStore)
					|| sessionStore.State is not RegisteredSession registeredSession) {
					context.Fail(new AuthorizationFailureReason(this, "Session not found"));
					return Task.CompletedTask;
				}

				if (requirement.Started.HasValue
					&& requirement.Started.Value != registeredSession.ExamStartedAt.HasValue) {
					context.Fail(new AuthorizationFailureReason(this, requirement.Started.Value ? "Exam has not started" : "Exam already started"));
					return Task.CompletedTask;
				}

				if (requirement.Ended.HasValue
					&& requirement.Ended.Value != registeredSession.ExamEndedAt.HasValue) {
					context.Fail(new AuthorizationFailureReason(this, requirement.Ended.Value ? "Exam has not ended" : "Exam already ended"));
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
