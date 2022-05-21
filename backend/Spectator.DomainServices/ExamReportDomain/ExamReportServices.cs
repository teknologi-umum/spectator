using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Linq;
using System.Security.Authentication;
using System.Threading;
using System.Threading.Tasks;
using Microsoft.Extensions.Caching.Memory;
using Microsoft.Extensions.Options;
using Spectator.DomainEvents.ExamReportDomain;
using Spectator.DomainModels.ExamReportDomain;
using Spectator.DomainServices.MemoryCache;
using Spectator.Repositories;
using Spectator.WorkerClient;

namespace Spectator.DomainServices.ExamReportDomain {
	public class ExamReportServices {
		private readonly ResultCache<Guid, AdministratorSession> _adminSessionById;
		private readonly ExamReportOptions _examReportOptions;
		private static readonly TimeSpan ADMIN_SESSION_TTL = TimeSpan.FromMinutes(60);
		private readonly ISessionEventRepository _sessionEventRepository;
		private readonly WorkerServices _workerServices;

		public ExamReportServices(
			IMemoryCache memoryCache,
			IOptions<ExamReportOptions> optionsAccessor,
			ISessionEventRepository sessionEventRepository,
			WorkerServices workerServices
		) {
			ResultCache.Initialize(out _adminSessionById, memoryCache);
			_examReportOptions = optionsAccessor.Value;
			_sessionEventRepository = sessionEventRepository;
			_workerServices = workerServices;
		}

		public AdministratorSession Login(string password) {
			if (string.IsNullOrEmpty(password)) throw new ArgumentNullException("password should not be empty");
			if (password != _examReportOptions.Password) throw new AuthenticationException("password do not match");

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

		public async Task<ImmutableList<ReportFile>> GetFilesAsync(Guid adminSessionId, CancellationToken cancellationToken) {
			// Authorize admin session
			if (!_adminSessionById.TryGetValue(adminSessionId, out var adminSession)
				|| adminSession.ExpiresAt <= DateTimeOffset.UtcNow) {
				throw new UnauthorizedAccessException();
			}

			// Acquire the list of session id from InfluxDB directly
			var sessionIds = new HashSet<Guid>();
			await foreach (var sessionId in _sessionEventRepository.GetAllSessionIdsAsync(adminSession, cancellationToken)) {
				sessionIds.Add(sessionId);
			}

			// Make a gRPC client call to the worker service to acquire the files list
			var filesLists = await _workerServices.GetFilesListsBySessionIdsAsync(sessionIds, cancellationToken);
			return filesLists
				.SelectMany(filesList => filesList.Files.Select(file => new ReportFile(
					sessionId: Guid.Parse(filesList.SessionId),
					studentNumber: file.StudentNumber,
					jsonFileUrl: file.FileUrlJson,
					csvFileUrl: file.FileUrlCsv
				)))
				.ToImmutableList();
		}
	}
}
