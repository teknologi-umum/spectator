using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Diagnostics.CodeAnalysis;
using System.Linq;
using System.Text.Json;
using System.Threading;
using System.Threading.Tasks;
using Spectator.DomainEvents.SessionDomain;
using Spectator.DomainModels.SessionDomain;
using Spectator.DomainModels.SubmissionDomain;
using Spectator.DomainServices.PistonDomain;
using Spectator.Observables.SessionDomain;
using Spectator.Primitives;
using Spectator.Repositories;

namespace Spectator.DomainServices.SessionDomain {
	public class SessionServices {
		private static readonly TimeSpan EXAM_DURATION = TimeSpan.FromMinutes(90);
		private readonly SessionSilo _sessionSilo;
		private readonly SubmissionServices _submissionServices;
		private readonly ISessionEventRepository _sessionEventRepository;

		public SessionServices(
			SessionSilo sessionSilo,
			SubmissionServices submissionServices,
			ISessionEventRepository sessionEventRepository
		) {
			_sessionSilo = sessionSilo;
			_submissionServices = submissionServices;
			_sessionEventRepository = sessionEventRepository;
		}

		public async Task<AnonymousSession> StartSessionAsync() {
			// Generate random sessionId
			var sessionId = Guid.NewGuid();

			// Create event
			var @event = new SessionStartedEvent(sessionId, DateTimeOffset.UtcNow);

			// Dispatch event
			var session = AnonymousSession.From(@event);
			if (!_sessionSilo.TryAdd(session)) throw new InvalidOperationException("Failed to start session");

			// Raise event
			await _sessionEventRepository.AddEventAsync(@event);

			// Return state
			return session;
		}

		public async Task SubmitPersonalInfoAsync(Guid sessionId, string studentNumber, int yearsOfExperience, int hoursOfPractice, string familiarLanguages) {
			// Get store
			var sessionStore = await GetSessionStoreAsync(sessionId, CancellationToken.None);

			// Create event
			var @event = new PersonalInfoSubmittedEvent(
				SessionId: sessionId,
				Timestamp: DateTimeOffset.UtcNow,
				StudentNumber: studentNumber,
				YearsOfExperience: yearsOfExperience,
				HoursOfPractice: hoursOfPractice,
				FamiliarLanguages: familiarLanguages
			);

			// Dispatch event
			sessionStore.Dispatch(@event);

			// Raise event
			await _sessionEventRepository.AddEventAsync(@event);
		}

		public async Task SubmitBeforeExamSAMAsync(Guid sessionId, int arousedLevel, int pleasedLevel) {
			// Get store
			var sessionStore = await GetSessionStoreAsync(sessionId, CancellationToken.None);

			// Create event
			var @event = new BeforeExamSAMSubmittedEvent(
				SessionId: sessionId,
				Timestamp: DateTimeOffset.UtcNow,
				SelfAssessmentManikin: new SelfAssessmentManikin(arousedLevel, pleasedLevel)
			);

			// Dispatch event
			sessionStore.Dispatch(@event);

			// Raise event
			await _sessionEventRepository.AddEventAsync(@event);
		}

		public async Task<RegisteredSession> StartExamAsync(Guid sessionId) {
			// Get store
			var sessionStore = await GetSessionStoreAsync(sessionId, CancellationToken.None);

			// Create event
			var utcNow = DateTimeOffset.UtcNow;
			var @event = new ExamStartedEvent(
				SessionId: sessionId,
				Timestamp: utcNow,
				QuestionNumbers: Enumerable.Range(1, 6).ToImmutableArray(),
				Deadline: utcNow.Add(EXAM_DURATION)
			);

			// Dispatch event
			sessionStore.Dispatch(@event);

			// Raise event
			await _sessionEventRepository.AddEventAsync(@event);

			// Return state
			return (RegisteredSession)sessionStore.State;
		}

		public async Task<RegisteredSession> ResumeExamAsync(Guid sessionId, CancellationToken cancellationToken) {
			// Get store
			var sessionStore = await GetSessionStoreAsync(sessionId, cancellationToken);

			// Create event
			var @event = new ExamIDEReloadedEvent(
				SessionId: sessionId,
				Timestamp: DateTimeOffset.UtcNow
			);

			// Dispatch event
			sessionStore.Dispatch(@event);

			// Raise event
			await _sessionEventRepository.AddEventAsync(@event);

			// Return state
			return (RegisteredSession)sessionStore.State;
		}

		public async Task<Submission> SubmitSolutionAsync(Guid sessionId, int questionNumber, Language language, string solution, string scratchPad) {
			// Get store
			var sessionStore = await GetSessionStoreAsync(sessionId, CancellationToken.None);

			// Execute solution in piston
			var submission = await _submissionServices.EvaluateSubmissionAsync(questionNumber, language, solution, scratchPad);

			// Create event
			SessionEventBase @event = submission.Accepted
				? new SolutionAcceptedEvent(sessionId, DateTimeOffset.UtcNow, questionNumber, language, solution, scratchPad, JsonSerializer.Serialize(submission.TestResults))
				: new SolutionRejectedEvent(sessionId, DateTimeOffset.UtcNow, questionNumber, language, solution, scratchPad, JsonSerializer.Serialize(submission.TestResults));

			// Dispatch event
			sessionStore.Dispatch(@event);

			// Raise event
			await _sessionEventRepository.AddEventAsync(@event);

			// Return submission
			return submission;
		}

		public async Task<RegisteredSession> EndExamAsync(Guid sessionId) {
			// Get store
			var sessionStore = await GetSessionStoreAsync(sessionId, CancellationToken.None);

			// Create event
			var @event = new ExamEndedEvent(sessionId, DateTimeOffset.UtcNow);

			// Dispatch event
			sessionStore.Dispatch(@event);

			// Raise event
			await _sessionEventRepository.AddEventAsync(@event);

			// Return state
			return (RegisteredSession)sessionStore.State;
		}

		public async Task<RegisteredSession> PassDeadlineAsync(Guid sessionId) {
			// Get store
			var sessionStore = await GetSessionStoreAsync(sessionId, CancellationToken.None);

			// Create event
			var @event = new DeadlinePassedEvent(sessionId, DateTimeOffset.UtcNow);

			// Dispatch event
			sessionStore.Dispatch(@event);

			// Raise event
			await _sessionEventRepository.AddEventAsync(@event);

			// Return state
			return (RegisteredSession)sessionStore.State;
		}

		public async Task<RegisteredSession> ForfeitExamAsync(Guid sessionId) {
			// Get store
			var sessionStore = await GetSessionStoreAsync(sessionId, CancellationToken.None);

			// Create event
			var @event = new ExamForfeitedEvent(sessionId, DateTimeOffset.UtcNow);

			// Dispatch event
			sessionStore.Dispatch(@event);

			// Raise event
			await _sessionEventRepository.AddEventAsync(@event);

			// Return state
			return (RegisteredSession)sessionStore.State;
		}

		public async Task SubmitAfterExamSAMAsync(Guid sessionId, int arousedLevel, int pleasedLevel) {
			// Get store
			var sessionStore = await GetSessionStoreAsync(sessionId, CancellationToken.None);

			// Create event
			var @event = new AfterExamSAMSubmittedEvent(
				SessionId: sessionId,
				Timestamp: DateTimeOffset.UtcNow,
				SelfAssessmentManikin: new SelfAssessmentManikin(arousedLevel, pleasedLevel)
			);

			// Dispatch event
			sessionStore.Dispatch(@event);

			// Raise event
			await _sessionEventRepository.AddEventAsync(@event);
		}

		[SuppressMessage("Performance", "RG0001:Do not await inside a loop", Justification = "Async enumerator")]
		private async Task<SessionStore> GetSessionStoreAsync(Guid sessionId, CancellationToken cancellationToken) {
			if (_sessionSilo.TryGet(sessionId, out var sessionStore)) return sessionStore;

			var asyncEnumerator = _sessionEventRepository.GetAllEventsAsync(sessionId, cancellationToken).GetAsyncEnumerator(cancellationToken);
			if (!await asyncEnumerator.MoveNextAsync()) throw new KeyNotFoundException();
			if (asyncEnumerator.Current is not SessionStartedEvent sessionStartedEvent) throw new InvalidProgramException("Session event stream must start with SessionStartedEvent");

			sessionStore = new SessionStore(sessionStartedEvent);
			while (await asyncEnumerator.MoveNextAsync()) {
				sessionStore.Dispatch(asyncEnumerator.Current);
			}

			_sessionSilo.TryAdd(sessionStore);

			return sessionStore;
		}
	}
}
