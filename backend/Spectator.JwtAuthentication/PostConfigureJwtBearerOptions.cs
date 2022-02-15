using System.Security.Claims;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Authentication.JwtBearer;
using Microsoft.Extensions.Options;
using RG.Ninja;

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
					if (context.Request.Cookies.TryGetValue("ACCESS_TOKEN", out var accessTokenCookie)) {
						context.Token = accessTokenCookie;
					} else if (context.Request.Headers.TryGetValue("ACCESS_TOKEN", out var accessTokenHeader)) {
						context.Token = accessTokenHeader;
					} else if (context.Request.Headers.Authorization is { Count: > 0 } authorizationHeader
						&& authorizationHeader[0].StartsWith("Bearer ", out var bearerToken)) {
						context.Token = bearerToken;
					} else if (context.Request.Query.TryGetValue("access_token", out var accessTokenQuery)) {
						context.Token = accessTokenCookie;
					}
					return Task.CompletedTask;
				}
			};
		}
	}
}
