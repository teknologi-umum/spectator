using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Diagnostics.CodeAnalysis;
using System.Linq;
using System.Text.Json;
using System.Threading;
using System.Threading.Tasks;
using Microsoft.Extensions.DependencyInjection;
using Spectator.DomainEvents.SessionDomain;
using Spectator.DomainModels.QuestionDomain;
using Spectator.DomainModels.SessionDomain;
using Spectator.DomainModels.SubmissionDomain;
using Spectator.DomainServices.PistonDomain;
using Spectator.DomainServices.QuestionDomain;
using Spectator.Observables.SessionDomain;
using Spectator.Primitives;
using Spectator.Repositories;

namespace Spectator.DomainServices.SessionDomain {
	public class SessionServices {
		private static readonly TimeSpan EXAM_DURATION = TimeSpan.FromMinutes(90);
		private readonly SessionSilo _sessionSilo;
		private readonly ISessionEventRepository _sessionEventRepository;
		private readonly IServiceProvider _serviceProvider;

		public SessionServices(
			SessionSilo sessionSilo,
			ISessionEventRepository sessionEventRepository,
			IServiceProvider serviceProvider
		) {
			_sessionSilo = sessionSilo;
			_sessionEventRepository = sessionEventRepository;
			_serviceProvider = serviceProvider;
		}

		public async Task<AnonymousSession> StartSessionAsync(Locale locale) {
			// Generate random sessionId
			var sessionId = Guid.NewGuid();

			// Create event
			var @event = new SessionStartedEvent(sessionId, DateTimeOffset.UtcNow, locale);

			// Dispatch event
			var session = AnonymousSession.From(@event);
			if (!_sessionSilo.TryAdd(session)) throw new InvalidOperationException("Failed to start session");

			// Raise event
			await _sessionEventRepository.AddEventAsync(@event);

			// Return state
			return session;
		}

		public async Task SetLocaleAsync(Guid sessionId, Locale locale) {
			// Get store
			var sessionStore = await GetSessionStoreAsync(sessionId, CancellationToken.None);

			// Create event
			var @event = new LocaleSetEvent(
				SessionId: sessionId,
				Timestamp: DateTimeOffset.UtcNow,
				Locale: locale
			);

			// Dispatch event
			sessionStore.Dispatch(@event);

			// Raise event
			await _sessionEventRepository.AddEventAsync(@event);
		}

		public async Task SubmitPersonalInfoAsync(
			Guid sessionId,
			string studentNumber,
			int yearsOfExperience,
			int hoursOfPractice,
			string familiarLanguages,
			string walletNumber,
			string walletType
		) {
			// Get store
			var sessionStore = await GetSessionStoreAsync(sessionId, CancellationToken.None);

			// Create event
			var @event = new PersonalInfoSubmittedEvent(
				SessionId: sessionId,
				Timestamp: DateTimeOffset.UtcNow,
				StudentNumber: studentNumber,
				YearsOfExperience: yearsOfExperience,
				HoursOfPractice: hoursOfPractice,
				FamiliarLanguages: familiarLanguages,
				WalletNumber: walletNumber,
				WalletType: walletType
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

		public async Task<(RegisteredSession Session, ImmutableDictionary<int, Question> QuestionByQuestionNumber)> StartExamAsync(Guid sessionId) {
			// Get store
			var sessionStore = await GetSessionStoreAsync(sessionId, CancellationToken.None);

			// Get locale from state
			var locale = sessionStore.State switch {
				AnonymousSession a => a.Locale,
				RegisteredSession r => r.Locale,
				_ => throw new InvalidProgramException("Unhandled session type")
			};

			// Get questions
			var questionsByLocale = await _serviceProvider.GetRequiredService<QuestionServices>().GetAllAsync(CancellationToken.None);
			var questions = questionsByLocale[locale];

			// Create event
			var utcNow = DateTimeOffset.UtcNow;
			var @event = new ExamStartedEvent(
				SessionId: sessionId,
				Timestamp: utcNow,
				QuestionNumbers: questions.Select(q => q.QuestionNumber).ToImmutableArray(),
				Deadline: utcNow.Add(EXAM_DURATION)
			);

			// Dispatch event
			sessionStore.Dispatch(@event);

			// Raise event
			await _sessionEventRepository.AddEventAsync(@event);

			// Return state
			return (
				Session: (RegisteredSession)sessionStore.State,
				QuestionByQuestionNumber: questions.ToImmutableDictionary(q => q.QuestionNumber)
			);
		}

		public async Task<(RegisteredSession Session, ImmutableDictionary<int, Question> QuestionByQuestionNumber)> ResumeExamAsync(Guid sessionId, CancellationToken cancellationToken) {
			// Get store
			var sessionStore = await GetSessionStoreAsync(sessionId, cancellationToken);

			// Get locale from state
			var locale = sessionStore.State switch {
				AnonymousSession a => a.Locale,
				RegisteredSession r => r.Locale,
				_ => throw new InvalidProgramException("Unhandled session type")
			};

			// Get questions
			var questionsByLocale = await _serviceProvider.GetRequiredService<QuestionServices>().GetAllAsync(CancellationToken.None);
			var questions = questionsByLocale[locale];

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
			return (
				Session: (RegisteredSession)sessionStore.State,
				QuestionByQuestionNumber: questions.ToImmutableDictionary(q => q.QuestionNumber)
			);
		}

		public async Task<Submission> SubmitSolutionAsync(Guid sessionId, int questionNumber, Language language, string directives, string solution, string scratchPad, CancellationToken cancellationToken) {
			// Get store
			var sessionStore = await GetSessionStoreAsync(sessionId, cancellationToken);

			// Get locale from state
			var locale = sessionStore.State switch {
				AnonymousSession a => a.Locale,
				RegisteredSession r => r.Locale,
				_ => throw new InvalidProgramException("Unhandled session type")
			};

			// Execute solution in piston
			var submission = await _serviceProvider.GetRequiredService<SubmissionServices>().EvaluateSubmissionAsync(
				questionNumber: questionNumber,
				locale: locale,
				language: language,
				directives: directives,
				solution: solution,
				scratchPad: scratchPad,
				cancellationToken: cancellationToken
			);

			// ----- NOTE: DO NOT USE cancellationToken BELOW THE FOLD!! -----

			// Create event
			SessionEventBase @event = submission.Accepted
				? new SolutionAcceptedEvent(sessionId, DateTimeOffset.UtcNow, questionNumber, language, solution, scratchPad, JsonSerializer.Serialize(submission.TestResults, TestResultBase.JSON_SERIALIZER_OPTIONS))
				: new SolutionRejectedEvent(sessionId, DateTimeOffset.UtcNow, questionNumber, language, solution, scratchPad, JsonSerializer.Serialize(submission.TestResults, TestResultBase.JSON_SERIALIZER_OPTIONS));

			// Dispatch event
			sessionStore.Dispatch(@event);

			// Raise event
			await _sessionEventRepository.AddEventAsync(@event);

			// Return submission
			return submission;
		}

		// Test the solution without submitting it
		public async Task<Submission> TestSolutionAsync(Guid sessionId, int questionNumber, Language language, string directives, string solution, string scratchPad, CancellationToken cancellationToken) {
			// Get store
			var sessionStore = await GetSessionStoreAsync(sessionId, cancellationToken);

			// Get locale from state
			var locale = sessionStore.State switch {
				AnonymousSession a => a.Locale,
				RegisteredSession r => r.Locale,
				_ => throw new InvalidProgramException("Unhandled session type")
			};

			// Execute solution in piston
			var submission = await _serviceProvider.GetRequiredService<SubmissionServices>().EvaluateSubmissionAsync(
				questionNumber: questionNumber,
				locale: locale,
				language: language,
				directives: directives,
				solution: solution,
				scratchPad: scratchPad,
				cancellationToken: cancellationToken
			);

			// Create event
			SessionEventBase @event = submission.Accepted 
				? new TestAcceptedEvent(sessionId, DateTimeOffset.UtcNow, questionNumber, language, solution, scratchPad, JsonSerializer.Serialize(submission.TestResults, TestResultBase.JSON_SERIALIZER_OPTIONS))
				: new TestRejectedEvent(sessionId, DateTimeOffset.UtcNow, questionNumber, language, solution, scratchPad, JsonSerializer.Serialize(submission.TestResults, TestResultBase.JSON_SERIALIZER_OPTIONS));

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
