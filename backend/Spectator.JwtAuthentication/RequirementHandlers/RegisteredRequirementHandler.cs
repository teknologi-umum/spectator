using System;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Authorization;
using Spectator.JwtAuthentication.Requirements;

namespace Spectator.JwtAuthentication.RequirementHandlers {
	public class RegisteredRequirementHandler : AuthorizationHandler<RegisteredRequirement> {
		protected override Task HandleRequirementAsync(AuthorizationHandlerContext context, RegisteredRequirement requirement) => throw new NotImplementedException();
	}
}
