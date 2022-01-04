using System;
using System.Collections.Generic;
using System.IdentityModel.Tokens.Jwt;
using System.Security.Claims;
using Microsoft.IdentityModel.Tokens;

namespace Spectator.JwtAuthentication {
	public class JwtAuthenticationServices {
		private readonly JwtSecurityTokenHandler _jwtSecurityTokenHandler;
		private readonly JwtAuthenticationSettings _jwtAuthenticationSettings;

		public JwtAuthenticationServices(
			JwtSecurityTokenHandler jwtSecurityTokenHandler,
			JwtAuthenticationSettings jwtAuthenticationSettings
		) {
			_jwtSecurityTokenHandler = jwtSecurityTokenHandler;
			_jwtAuthenticationSettings = jwtAuthenticationSettings;
		}

		public TokenPayload CreatePayload(Guid sessionId) {
			var now = DateTimeOffset.UtcNow;
			var expires = now.Add(_jwtAuthenticationSettings.Lifetime);
			return new TokenPayload(
				SessionId: sessionId,
				Issued: now,
				Expires: expires
			);
		}

		public string EncodeToken(TokenPayload tokenPayload) {
			List<Claim> claims = new() {
				new Claim(JwtRegisteredClaimNames.Jti, tokenPayload.SessionId.ToString()),
				new Claim(JwtRegisteredClaimNames.Nonce, Guid.NewGuid().ToString())
			};
			JwtSecurityToken jwtSecurityToken = new(
				issuer: _jwtAuthenticationSettings.Issuer,
				audience: _jwtAuthenticationSettings.Audience,
				claims: claims,
				notBefore: tokenPayload.Issued.UtcDateTime,
				expires: tokenPayload.Expires.UtcDateTime,
				signingCredentials: _jwtAuthenticationSettings.SigningCredentials
			);
			return _jwtSecurityTokenHandler.WriteToken(jwtSecurityToken);
		}

		public TokenPayload DecodeToken(string jwtToken) {
			var _ = _jwtSecurityTokenHandler.ValidateToken(
				token: jwtToken.Trim(),
				validationParameters: new TokenValidationParameters {
					ValidIssuer = _jwtAuthenticationSettings.Issuer,
					ValidAudience = _jwtAuthenticationSettings.Audience,
					IssuerSigningKey = _jwtAuthenticationSettings.SymmetricSecurityKey
				},
				validatedToken: out var validatedToken
			);
			var jwtSecurityToken = (JwtSecurityToken)validatedToken;
			return new TokenPayload(
				SessionId: Guid.Parse(jwtSecurityToken.Payload[JwtRegisteredClaimNames.Jti].ToString()!),
				Issued: validatedToken.ValidFrom,
				Expires: validatedToken.ValidTo
			);
		}
	}
}
