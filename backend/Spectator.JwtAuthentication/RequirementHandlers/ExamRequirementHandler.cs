using System;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Authorization;
using Spectator.JwtAuthentication.Requirements;

namespace Spectator.JwtAuthentication.RequirementHandlers {
	public class ExamRequirementHandler : AuthorizationHandler<ExamRequirement> {
		protected override Task HandleRequirementAsync(AuthorizationHandlerContext context, ExamRequirement requirement) => throw new NotImplementedException();
	}
}
