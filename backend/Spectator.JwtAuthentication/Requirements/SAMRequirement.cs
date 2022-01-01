using System;
using Microsoft.AspNetCore.Authorization;

namespace Spectator.JwtAuthentication.Requirements {
	public class SAMRequirement : IAuthorizationRequirement {
		public bool? BeforeTest { get; }
		public bool? AfterTest { get; }

		public SAMRequirement(
			bool? beforeTest = null,
			bool? afterTest = null
		) {
			if (beforeTest == null && afterTest == null) throw new ArgumentException("beforeTest and afterTest cannot be both null");
			BeforeTest = beforeTest;
			AfterTest = afterTest;
		}
	}
}
