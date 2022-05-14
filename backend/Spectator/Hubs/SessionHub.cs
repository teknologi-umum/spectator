using System;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.SignalR;
using Microsoft.Extensions.DependencyInjection;
using SignalRSwaggerGen.Attributes;
using SignalRSwaggerGen.Enums;
using Spectator.DomainModels.SessionDomain;
using Spectator.DomainModels.SubmissionDomain;
using Spectator.DomainServices.SessionDomain;
using Spectator.JwtAuthentication;
using Spectator.PoormansAuth;
using Spectator.Primitives;
using Spectator.Protos.HubInterfaces;
using Spectator.Protos.Session;

namespace Spectator.Hubs {
	[SignalRHub(autoDiscover: AutoDiscover.MethodsAndArgs)]
	public class SessionHub : Hub<ISessionHub>, ISessionHub {
		private readonly PoormansAuthentication _poormansAuthentication;
		private readonly SessionServices _sessionServices;
		private readonly IServiceProvider _serviceProvider;

		public SessionHub(
			PoormansAuthentication poormansAuthentication,
			SessionServices sessionServices,
			IServiceProvider serviceProvider
		) {
			_poormansAuthentication = poormansAuthentication;
			_sessionServices = sessionServices;
			_serviceProvider = serviceProvider;
		}

		public async Task<SessionReply> StartSessionAsync(StartSessionRequest request) {
			// Create session
			var session = await _sessionServices.StartSessionAsync((Locale)request.Locale);

			// Encode as JWT and map results
			var jwtAuthenticationServices = _serviceProvider.GetRequiredService<JwtAuthenticationServices>();
			var tokenPayload = jwtAuthenticationServices.CreatePayload(session.Id);
			return new SessionReply {
				AccessToken = jwtAuthenticationServices.EncodeToken(tokenPayload)
			};
		}

		public Task SetLocaleAsync(SetLocaleRequest request) {
			// Authenticate
			var session = _poormansAuthentication.Authenticate(request.AccessToken);

			// Set locale and map results
			return _sessionServices.SetLocaleAsync(
				sessionId: session.Id,
				locale: (Locale)request.Locale
			);
		}

		public Task SubmitPersonalInfoAsync(SubmitPersonalInfoRequest request) {
			// Authenticate
			var session = _poormansAuthentication.Authenticate(request.AccessToken);

			// Authorize: session must be anonymous, personal info has not been submitted
			if (session is not AnonymousSession) throw new UnauthorizedAccessException("Personal Info already submitted");

			// Submit personal info and map results
			return _sessionServices.SubmitPersonalInfoAsync(
				sessionId: session.Id,
				studentNumber: request.StudentNumber,
				yearsOfExperience: request.YearsOfExperience,
				hoursOfPractice: request.HoursOfPractice,
				familiarLanguages: request.FamiliarLanguages,
				walletNumber: request.WalletNumber,
				walletType: request.WalletType
			);
		}

		public Task SubmitBeforeExamSAMAsync(SubmitSAMRequest request) {
			// Authenticate
			var session = _poormansAuthentication.Authenticate(request.AccessToken);

			// Authorize: personal info must be already submitted, but SAM not submitted
			if (session is not RegisteredSession registeredSession) throw new UnauthorizedAccessException("Personal Info not yet submitted");
			if (registeredSession.BeforeExamSAM is not null) throw new UnauthorizedAccessException("Before Exam SAM already submitted");

			// Submit SAM and map results
			return _sessionServices.SubmitBeforeExamSAMAsync(
				sessionId: session.Id,
				arousedLevel: request.ArousedLevel,
				pleasedLevel: request.PleasedLevel
			);
		}

		public async Task<Exam> StartExamAsync(EmptyRequest request) {
			// Authenticate
			var session = _poormansAuthentication.Authenticate(request.AccessToken);

			// Authorize: Before exam SAM must be already submitted, but exam has not been started
			if (session is not RegisteredSession registeredSession) throw new UnauthorizedAccessException("Personal Info not yet submitted");
			if (registeredSession.BeforeExamSAM is null) throw new UnauthorizedAccessException("Before Exam SAM not yet submitted");
			if (registeredSession.ExamStartedAt is not null) throw new UnauthorizedAccessException("Exam already started");

			// Start exam
			(registeredSession, var questionByQuestionNumber) = await _sessionServices.StartExamAsync(session.Id);

			// Map results
			return new Exam {
				Deadline = registeredSession.ExamDeadline!.Value.ToUnixTimeMilliseconds(),
				Questions = {
					from questionNumber in registeredSession.QuestionNumbers!.Value
					let question = questionByQuestionNumber[questionNumber]
					select new Question {
						QuestionNumber = questionNumber,
						Title = question.TitleByLocale[registeredSession.Locale],
						Instruction = question.InstructionByLocale[registeredSession.Locale],
						LanguageAndTemplates = {
							from kvp in question.TemplateByLanguageByLocale[registeredSession.Locale]
							select new Question.Types.LanguageAndTemplate {
								Language = (Protos.Enums.Language)kvp.Key,
								Template = kvp.Value
							}
						}
					}
				},
				AnsweredQuestionNumbers = { }
			};
		}

		public async Task<Exam> ResumeExamAsync(EmptyRequest request) {
			// Authenticate
			var session = _poormansAuthentication.Authenticate(request.AccessToken);

			// Authorize: Exam must be in progress
			if (session is not RegisteredSession registeredSession) throw new UnauthorizedAccessException("Personal Info not yet submitted");
			if (registeredSession.ExamStartedAt is null) throw new UnauthorizedAccessException("Exam not yet started");
			if (registeredSession.ExamEndedAt is not null) throw new UnauthorizedAccessException("Exam already ended");
			if (registeredSession.ExamDeadline is null) throw new InvalidProgramException("Exam deadline not set");
			if (registeredSession.ExamDeadline.Value < DateTimeOffset.UtcNow) throw new UnauthorizedAccessException("Exam deadline exceeded");

			// Resume exam
			(registeredSession, var questionByQuestionNumber) = await _sessionServices.ResumeExamAsync(session.Id, Context.ConnectionAborted);

			// Map results
			return new Exam {
				Deadline = registeredSession.ExamDeadline!.Value.ToUnixTimeMilliseconds(),
				Questions = {
					from questionNumber in registeredSession.QuestionNumbers!.Value
					let question = questionByQuestionNumber[questionNumber]
					select new Question {
						QuestionNumber = questionNumber,
						Title = question.TitleByLocale[registeredSession.Locale],
						Instruction = question.InstructionByLocale[registeredSession.Locale],
						LanguageAndTemplates = {
							from kvp in question.TemplateByLanguageByLocale[registeredSession.Locale]
							select new Question.Types.LanguageAndTemplate {
								Language = (Protos.Enums.Language)kvp.Key,
								Template = kvp.Value
							}
						}
					}
				},
				AnsweredQuestionNumbers = { }
			};
		}

		public async Task<SubmissionResult> SubmitSolutionAsync(SubmissionRequest request) {
			// Authenticate
			var session = _poormansAuthentication.Authenticate(request.AccessToken);

			// Authorize: Exam must be in progress
			if (session is not RegisteredSession registeredSession) throw new UnauthorizedAccessException("Personal Info not yet submitted");
			if (registeredSession.ExamStartedAt is null) throw new UnauthorizedAccessException("Exam not yet started");
			if (registeredSession.ExamEndedAt is not null) throw new UnauthorizedAccessException("Exam already ended");
			if (registeredSession.ExamDeadline is null) throw new InvalidProgramException("Exam deadline not set");
			if (registeredSession.ExamDeadline.Value < DateTimeOffset.UtcNow) throw new UnauthorizedAccessException("Exam deadline exceeded");

			// Submit solution
			var submission = await _sessionServices.SubmitSolutionAsync(
				sessionId: session.Id,
				questionNumber: request.QuestionNumber,
				language: (Language)request.Language,
				directives: request.Directives,
				solution: request.Solution,
				scratchPad: request.ScratchPad,
				cancellationToken: Context.ConnectionAborted
			);

			// Map results
			return new SubmissionResult {
				Accepted = submission.Accepted,
				TestResults = {
					from testResult in submission.TestResults
					select testResult switch {
						PassingTestResult passing => new TestResult {
							TestNumber = passing.TestNumber,
							PassingTest = new TestResult.Types.PassingTest()
						},
						FailingTestResult failing => new TestResult {
							TestNumber = failing.TestNumber,
							FailingTest = new TestResult.Types.FailingTest {
								ExpectedStdout = failing.ExpectedStdout,
								ActualStdout = failing.ActualStdout
							}
						},
						CompileErrorResult compileError => new TestResult {
							TestNumber = compileError.TestNumber,
							CompileError = new TestResult.Types.CompileError {
								Stderr = compileError.Stderr
							}
						},
						RuntimeErrorResult runtimeError => new TestResult {
							TestNumber = runtimeError.TestNumber,
							RuntimeError = new TestResult.Types.RuntimeError {
								Stderr = runtimeError.Stderr
							}
						},
						_ => throw new InvalidProgramException("Unhandled TestResult type")
					}
				}
			};
		}

		public async Task<ExamResult> EndExamAsync(EmptyRequest request) {
			// Authenticate
			var session = _poormansAuthentication.Authenticate(request.AccessToken);

			// Authorize: Exam must be in progress
			if (session is not RegisteredSession registeredSession) throw new UnauthorizedAccessException("Personal Info not yet submitted");
			if (registeredSession.ExamStartedAt is null) throw new UnauthorizedAccessException("Exam not yet started");
			if (registeredSession.ExamEndedAt is not null) throw new UnauthorizedAccessException("Exam already ended");
			if (registeredSession.ExamDeadline is null) throw new InvalidProgramException("Exam deadline not set");
			if (registeredSession.ExamDeadline.Value < DateTimeOffset.UtcNow) throw new UnauthorizedAccessException("Exam deadline exceeded");

			// End exam
			registeredSession = await _sessionServices.EndExamAsync(session.Id);


			// Map results
			return new ExamResult {
				Duration = registeredSession.ExamEndedAt!.Value.ToUnixTimeMilliseconds() - registeredSession.ExamStartedAt!.Value.ToUnixTimeMilliseconds(),
				AnsweredQuestionNumbers = {
					from kvp in registeredSession.SubmissionByQuestionNumber
					where kvp.Value.Accepted
					select kvp.Key
				}
			};
		}

		public async Task<ExamResult> PassDeadlineAsync(EmptyRequest request) {
			// Authenticate
			var session = _poormansAuthentication.Authenticate(request.AccessToken);

			// Authorize: Exam must be in progress but deadline is exceeded
			if (session is not RegisteredSession registeredSession) throw new UnauthorizedAccessException("Personal Info not yet submitted");
			if (registeredSession.ExamStartedAt is null) throw new UnauthorizedAccessException("Exam not yet started");
			if (registeredSession.ExamEndedAt is not null) throw new UnauthorizedAccessException("Exam already ended");
			if (registeredSession.ExamDeadline is null) throw new InvalidProgramException("Exam deadline not set");
			if (registeredSession.ExamDeadline.Value >= DateTimeOffset.UtcNow) throw new UnauthorizedAccessException("Exam deadline not yet exceeded");

			// Mark deadline passed
			registeredSession = await _sessionServices.PassDeadlineAsync(session.Id);

			// Map results
			return new ExamResult {
				Duration = registeredSession.ExamEndedAt!.Value.ToUnixTimeMilliseconds() - registeredSession.ExamStartedAt!.Value.ToUnixTimeMilliseconds(),
				AnsweredQuestionNumbers = {
					from kvp in registeredSession.SubmissionByQuestionNumber
					where kvp.Value.Accepted
					select kvp.Key
				}
			};
		}

		public async Task<ExamResult> ForfeitExamAsync(EmptyRequest request) {
			// Authenticate
			var session = _poormansAuthentication.Authenticate(request.AccessToken);

			// Authorize: Exam must be in progress
			if (session is not RegisteredSession registeredSession) throw new UnauthorizedAccessException("Personal Info not yet submitted");
			if (registeredSession.ExamStartedAt is null) throw new UnauthorizedAccessException("Exam not yet started");
			if (registeredSession.ExamEndedAt is not null) throw new UnauthorizedAccessException("Exam already ended");
			if (registeredSession.ExamDeadline is null) throw new InvalidProgramException("Exam deadline not set");
			if (registeredSession.ExamDeadline.Value < DateTimeOffset.UtcNow) throw new UnauthorizedAccessException("Exam deadline exceeded");

			// Forfeit exam
			registeredSession = await _sessionServices.ForfeitExamAsync(session.Id);

			// Map results
			return new ExamResult {
				Duration = registeredSession.ExamEndedAt!.Value.ToUnixTimeMilliseconds() - registeredSession.ExamStartedAt!.Value.ToUnixTimeMilliseconds(),
				AnsweredQuestionNumbers = {
					from kvp in registeredSession.SubmissionByQuestionNumber
					where kvp.Value.Accepted
					select kvp.Key
				}
			};
		}

		public Task SubmitAfterExamSAM(SubmitSAMRequest request) {
			// Authenticate
			var session = _poormansAuthentication.Authenticate(request.AccessToken);

			// Authorize: Exam must already been ended and second SAM not submitted
			if (session is not RegisteredSession registeredSession) throw new UnauthorizedAccessException("Personal Info not yet submitted");
			if (registeredSession.ExamEndedAt is null) throw new UnauthorizedAccessException("Exam not yet ended");
			if (registeredSession.AfterExamSAM is not null) throw new InvalidProgramException("After exam SAM already submitted");

			// Submit SAM and map results
			return _sessionServices.SubmitAfterExamSAMAsync(
				sessionId: session.Id,
				arousedLevel: request.ArousedLevel,
				pleasedLevel: request.PleasedLevel
			);
		}
	}
}
