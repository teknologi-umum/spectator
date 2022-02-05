using System.Security.Claims;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Authentication.JwtBearer;
using Microsoft.Extensions.Options;

namespace Spectator.JwtAuthentication {
	public class PostConfigureJwtBearerOptions : IPostConfigureOptions<JwtBearerOptions> {
		private readonly JwtAuthenticationSettings _jwtAuthenticationSettings;

		public PostConfigureJwtBearerOptions(
			JwtAuthenticationSettings jwtAuthenticationSettings
		) {
			_jwtAuthenticationSettings = jwtAuthenticationSettings;
		}

		public void PostConfigure(string name, JwtBearerOptions options) {
			options.SaveToken = false;
			options.IncludeErrorDetails = true;
			options.RequireHttpsMetadata = false;
			options.TokenValidationParameters = new() {
				ValidateActor = false,
				ValidateIssuer = false,
				ValidateAudience = false,
				ValidateLifetime = true,
				ValidateIssuerSigningKey = true,
				IssuerSigningKey = _jwtAuthenticationSettings.SymmetricSecurityKey,
				NameClaimType = ClaimTypes.NameIdentifier
			};
			options.Events = new JwtBearerEvents {
				OnMessageReceived = context => {
					context.Token = context.Request.Query["access_token"];
					return Task.CompletedTask;
				}
			};
		}
	}
}
