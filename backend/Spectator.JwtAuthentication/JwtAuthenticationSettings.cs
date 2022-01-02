using System;
using System.Globalization;
using Microsoft.Extensions.Configuration;
using Microsoft.IdentityModel.Tokens;

namespace Spectator.JwtAuthentication {
	public class JwtAuthenticationSettings {
		public string Secret { get; }
		public string Issuer { get; }
		public string Audience { get; }
		public TimeSpan Lifetime { get; }

		public JwtAuthenticationSettings(IConfiguration configuration) {
			Secret = configuration["JwtAuthentication:Secret"];
			Issuer = configuration["JwtAuthentication:Issuer"];
			Audience = configuration["JwtAuthentication:Audience"];
			Lifetime = TimeSpan.FromSeconds(int.Parse(configuration["JwtAuthentication:Lifetime"], CultureInfo.InvariantCulture));
		}

		public SymmetricSecurityKey SymmetricSecurityKey => new(Convert.FromBase64String(Secret));
		public SigningCredentials SigningCredentials => new(SymmetricSecurityKey, SecurityAlgorithms.HmacSha512);
	}
}
