using System;
using System.Globalization;
using System.IdentityModel.Tokens.Jwt;
using System.Linq;
using System.Security.Claims;

namespace Spectator.JwtAuthentication {
	public record TokenPayload(
		Guid SessionId,
		DateTimeOffset Issued,
		DateTimeOffset Expires
	) {
		public static TokenPayload FromClaimsPrincipal(ClaimsPrincipal claimsPrincipal) {
			return new TokenPayload(
				SessionId: Guid.Parse(claimsPrincipal.Claims.Single(claim => claim.Type is JwtRegisteredClaimNames.Jti).Value),
				Issued: DateTimeOffset.FromUnixTimeSeconds(int.Parse(claimsPrincipal.Claims.Single(claim => claim.Type is JwtRegisteredClaimNames.Iat or JwtRegisteredClaimNames.Nbf).Value, CultureInfo.InvariantCulture)),
				Expires: DateTimeOffset.FromUnixTimeSeconds(int.Parse(claimsPrincipal.Claims.Single(claim => claim.Type == JwtRegisteredClaimNames.Exp).Value, CultureInfo.InvariantCulture))
			);
		}
	}
}
