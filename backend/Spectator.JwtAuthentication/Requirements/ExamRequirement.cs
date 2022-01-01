using System;
using Microsoft.AspNetCore.Authorization;

namespace Spectator.JwtAuthentication.Requirements {
	public class ExamRequirement : IAuthorizationRequirement {
		public bool? Started { get; }
		public bool? Ended { get; }

		public ExamRequirement(
			bool? started = null,
			bool? ended = null
		) {
			if (started == null && ended == null) throw new ArgumentException("started and ended cannot be both null");
			Started = started;
			Ended = ended;
		}
	}
}
