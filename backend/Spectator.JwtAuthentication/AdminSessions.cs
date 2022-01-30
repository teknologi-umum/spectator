﻿using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Spectator.JwtAuthentication {
	public class AdminSessions {
		private readonly HashSet<Guid> _ADMIN_IDS = new();

		public void Add(Guid sessionId) => _ADMIN_IDS.Add(sessionId);
		public bool Contains(Guid sessionId) => _ADMIN_IDS.Contains(sessionId);
	}
}
