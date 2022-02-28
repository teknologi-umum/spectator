using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Security.Authentication;
using System.Threading;
using System.Threading.Tasks;
using Microsoft.Extensions.Caching.Memory;
using Microsoft.Extensions.Options;
using Spectator.DomainEvents.ExamReportDomain;
using Spectator.DomainModels.ExamReportDoman;
using Spectator.DomainServices.MemoryCache;

namespace Spectator.DomainServices.ExamReportDomain {
	public class ExamReportServices {
		private readonly ResultCache<Guid, AdministratorSession> _adminSessionById;
		private readonly ExamReportOptions _examReportOptions;
		private static readonly TimeSpan ADMIN_SESSION_TTL = TimeSpan.FromMinutes(60);

		public ExamReportServices(
			IMemoryCache memoryCache,
			IOptions<ExamReportOptions> optionsAccessor
		) {
			ResultCache.Initialize(out _adminSessionById, memoryCache);
			_examReportOptions = optionsAccessor.Value;
		}

		public AdministratorSession Login(string username, string password) {
			if (string.IsNullOrEmpty(username) || string.IsNullOrEmpty(password)) {
				throw new ArgumentNullException("username and/or password should not be empty");
			}

			if (username != _examReportOptions.Username || password != _examReportOptions.Password) {
				throw new AuthenticationException("username and/or password do not match");
			}

			var sessionId = Guid.NewGuid();

			// Create event
			var timestamp = DateTimeOffset.UtcNow;
			var @event = new AdministratorSessionCreatedEvent(
				SessionId: sessionId,
				CreatedAt: timestamp,
				ExpiresAt: timestamp.Add(ADMIN_SESSION_TTL)
			);

			// Dispatch event
			var adminSession = AdministratorSession.From(@event);

			// Store session
			_adminSessionById.Set(sessionId, adminSession, absoluteExpirationRelativeToNow: ADMIN_SESSION_TTL);

			// Return session
			return adminSession;
		}

		public void Logout(Guid sessionId) {
			// Create event (let's skip this ceremony karena belum dipake)
			// var adminSessionDeletedEvent = ....

			// Remove session
			_adminSessionById.Remove(sessionId);
		}

		public async Task<ImmutableList<ReportFile>> GetFilesAsync(Guid sessionId, CancellationToken cancellationToken) {
			if (!_adminSessionById.TryGetValue(sessionId, out var adminSession)
				|| adminSession.ExpiresAt <= DateTimeOffset.UtcNow) {
				throw new UnauthorizedAccessException();
			}
			// TODO: acquire the list of session id from InfluxDB directly
			// TODO: make a gRPC client call to the worker service to acquire the files list
		}
	}
}
