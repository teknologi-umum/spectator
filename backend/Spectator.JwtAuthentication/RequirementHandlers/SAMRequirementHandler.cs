using System;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Authorization;
using Spectator.JwtAuthentication.Requirements;

namespace Spectator.JwtAuthentication.RequirementHandlers {
	public class SAMRequirementHandler : AuthorizationHandler<SAMRequirement> {
		protected override Task HandleRequirementAsync(AuthorizationHandlerContext context, SAMRequirement requirement) => throw new NotImplementedException();
	}
}
