using System;

namespace Spectator.JwtAuthentication {
	public record TokenPayload(
		Guid SessionId,
		DateTimeOffset Issued,
		DateTimeOffset Expires
	);
}
