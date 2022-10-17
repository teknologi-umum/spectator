using System;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.SignalR;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Logging;
using SignalRSwaggerGen.Attributes;
using SignalRSwaggerGen.Enums;
using Spectator.DomainModels.SessionDomain;
using Spectator.DomainModels.SubmissionDomain;
using Spectator.DomainServices.ExamResultDomain;
using Spectator.DomainServices.SessionDomain;
using Spectator.JwtAuthentication;
using Spectator.PoormansAuth;
using Spectator.Primitives;
using Spectator.Protos.HubInterfaces;
using Spectator.Protos.Session;

namespace Spectator.Hubs {
	[SignalRHub(autoDiscover: AutoDiscover.MethodsAndParams)]
	public class SessionHub : Hub<ISessionHub>, ISessionHub {
		private readonly PoormansAuthentication _poormansAuthentication;
		private readonly SessionServices _sessionServices;
		private readonly ExamResultServices _examResultServices;
		private readonly IServiceProvider _serviceProvider;
		private readonly ILogger<SessionHub> _logger;

		public SessionHub(
			PoormansAuthentication poormansAuthentication,
			SessionServices sessionServices,
			ExamResultServices examResultServices,
			IServiceProvider serviceProvider,
			ILogger<SessionHub> logger
		) {
			_poormansAuthentication = poormansAuthentication;
			_sessionServices = sessionServices;
			_examResultServices = examResultServices;
			_serviceProvider = serviceProvider;
			_logger = logger;
		}

		public async Task<SessionReply> StartSessionAsync(StartSessionRequest request) {
			try {
				// Create session
				var session = await _sessionServices.StartSessionAsync((Locale)request.Locale);

				// Encode as JWT and map results
				var jwtAuthenticationServices = _serviceProvider.GetRequiredService<JwtAuthenticationServices>();
				var tokenPayload = jwtAuthenticationServices.CreatePayload(session.Id);
				var sessionReply = new SessionReply {
					AccessToken = jwtAuthenticationServices.EncodeToken(tokenPayload)
				};

				_logger.LogInformation($"Session started: {session.Id}");

				return sessionReply;
			} catch (Exception exc) {
				_logger.LogError(exc, $"Error starting session. Request: {request}");
				throw;
			}
		}

		public async Task SetLocaleAsync(SetLocaleRequest request) {
			try {
				// Authenticate
				var session = _poormansAuthentication.Authenticate(request.AccessToken);

				// Set locale and map results
				await _sessionServices.SetLocaleAsync(
					sessionId: session.Id,
					locale: (Locale)request.Locale
				);

				_logger.LogInformation($"Set locale. Session: {session.Id}, Locale: {request.Locale}");
			} catch (Exception exc) {
				_logger.LogError(exc, $"Error setting locale. Request: {request}");
				throw;
			}
		}

		public async Task SubmitPersonalInfoAsync(SubmitPersonalInfoRequest request) {
			try {
				// Authenticate
				var session = _poormansAuthentication.Authenticate(request.AccessToken);

				// Authorize: session must be anonymous, personal info has not been submitted
				if (session is not AnonymousSession) throw new UnauthorizedAccessException("Personal Info already submitted");

				// Submit personal info
				await _sessionServices.SubmitPersonalInfoAsync(
					sessionId: session.Id,
					email: request.Email,
					age: request.Age,
					gender: request.Gender,
					nationality: request.Nationality,
					studentNumber: request.StudentNumber,
					yearsOfExperience: request.YearsOfExperience,
					hoursOfPractice: request.HoursOfPractice,
					familiarLanguages: request.FamiliarLanguages,
					walletNumber: request.WalletNumber,
					walletType: request.WalletType
				);

				_logger.LogInformation($"Submitted personal info. Session: {session.Id}");
			} catch (Exception exc) {
				_logger.LogError(exc, $"Error submitting personal info. Request: {request}");
				throw;
			}
		}

		public async Task SubmitBeforeExamSAMAsync(SubmitSAMRequest request) {
			try {
				// Authenticate
				var session = _poormansAuthentication.Authenticate(request.AccessToken);

				// Authorize: personal info must be already submitted, but SAM not submitted
				if (session is not RegisteredSession registeredSession) throw new UnauthorizedAccessException("Personal Info not yet submitted");
				if (registeredSession.BeforeExamSAM is not null) throw new UnauthorizedAccessException("Before Exam SAM already submitted");

				// Submit SAM and map results
				await _sessionServices.SubmitBeforeExamSAMAsync(
					sessionId: session.Id,
					arousedLevel: request.ArousedLevel,
					pleasedLevel: request.PleasedLevel
				);

				_logger.LogInformation($"Submitted before exam SAM. Session: {session.Id}");
			} catch (Exception exc) {
				_logger.LogError(exc, $"Error submitting before exam SAM. Request: {request}");
				throw;
			}
		}

		public async Task<Exam> StartExamAsync(EmptyRequest request) {
			try {
				// Authenticate
				var session = _poormansAuthentication.Authenticate(request.AccessToken);

				// Authorize: Before exam SAM must be already submitted, but exam has not been started
				if (session is not RegisteredSession registeredSession) throw new UnauthorizedAccessException("Personal Info not yet submitted");
				if (registeredSession.BeforeExamSAM is null) throw new UnauthorizedAccessException("Before Exam SAM not yet submitted");
				if (registeredSession.ExamStartedAt is not null) throw new UnauthorizedAccessException("Exam already started");

				// Start exam
				(registeredSession, var questionByQuestionNumber) = await _sessionServices.StartExamAsync(session.Id);

				// Map results
				var exam = new Exam {
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

				_logger.LogInformation($"Started exam. Session: {session.Id}");

				return exam;
			} catch (Exception exc) {
				_logger.LogError(exc, $"Error starting exam. Request: {request}");
				throw;
			}
		}

		public async Task<Exam> ResumeExamAsync(EmptyRequest request) {
			try {
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
				var exam = new Exam {
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

				_logger.LogInformation($"Resumed exam. Session: {session.Id}");

				return exam;
			} catch (Exception exc) {
				_logger.LogError(exc, $"Error resuming exam. Request: {request}");
				throw;
			}
		}

		public async Task<SubmissionResult> SubmitSolutionAsync(SubmissionRequest request) {
			try {
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
				var submissionResult = new SubmissionResult {
					Accepted = submission.Accepted,
					TestResults = {
						from testResult in submission.TestResults
						select testResult switch {
							PassingTestResult passing => new TestResult {
								TestNumber = passing.TestNumber,
								PassingTest = new TestResult.Types.PassingTest {
									ExpectedStdout = passing.ExpectedStdout,
									ActualStdout = passing.ActualStdout,
									ArgumentsStdout = passing.ArgumentsStdout
								}
							},
							FailingTestResult failing => new TestResult {
								TestNumber = failing.TestNumber,
								FailingTest = new TestResult.Types.FailingTest {
									ExpectedStdout = failing.ExpectedStdout,
									ActualStdout = failing.ActualStdout,
									ArgumentsStdout = failing.ArgumentsStdout
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
							InvalidInputResult invalidInput => new TestResult {
								TestNumber = invalidInput.TestNumber,
								InvalidInput = new TestResult.Types.InvalidInput {
									Stderr = invalidInput.Stderr
								}
							},
							_ => throw new InvalidProgramException("Unhandled TestResult type")
						}
					}
				};

				_logger.LogInformation($"Submitted solution. Session: {session.Id}, Question: {request.QuestionNumber}, Accepted: {submission.Accepted}");

				return submissionResult;
			} catch (Exception exc) {
				_logger.LogError(exc, $"Error submitting solution. Request: {request}");
				throw;
			}
		}

		public async Task<SubmissionResult> TestSolutionAsync(SubmissionRequest request) {
			try {
				// Authenticate
				var session = _poormansAuthentication.Authenticate(request.AccessToken);

				// Authorize: Exam must be in progress
				if (session is not RegisteredSession registeredSession) throw new UnauthorizedAccessException("Personal Info not yet submitted");
				if (registeredSession.ExamStartedAt is null) throw new UnauthorizedAccessException("Exam not yet started");
				if (registeredSession.ExamEndedAt is not null) throw new UnauthorizedAccessException("Exam already ended");
				if (registeredSession.ExamDeadline is null) throw new InvalidProgramException("Exam deadline not set");
				if (registeredSession.ExamDeadline.Value < DateTimeOffset.UtcNow) throw new UnauthorizedAccessException("Exam deadline exceeded");

				// Test solution
				var submission = await _sessionServices.TestSolutionAsync(
					sessionId: session.Id,
					questionNumber: request.QuestionNumber,
					language: (Language)request.Language,
					directives: request.Directives,
					solution: request.Solution,
					scratchPad: request.ScratchPad,
					cancellationToken: Context.ConnectionAborted
				);

				// Map results
				var submissionResult = new SubmissionResult {
					Accepted = submission.Accepted,
					TestResults = {
						from testResult in submission.TestResults
						select testResult switch {
							PassingTestResult passing => new TestResult {
								TestNumber = passing.TestNumber,
								PassingTest = new TestResult.Types.PassingTest {
									ExpectedStdout = passing.ExpectedStdout,
									ActualStdout = passing.ActualStdout,
									ArgumentsStdout = passing.ArgumentsStdout
								}
							},
							FailingTestResult failing => new TestResult {
								TestNumber = failing.TestNumber,
								FailingTest = new TestResult.Types.FailingTest {
									ExpectedStdout = failing.ExpectedStdout,
									ActualStdout = failing.ActualStdout,
									ArgumentsStdout = failing.ArgumentsStdout
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
							InvalidInputResult invalidInput => new TestResult {
								TestNumber = invalidInput.TestNumber,
								InvalidInput = new TestResult.Types.InvalidInput {
									Stderr = invalidInput.Stderr
								}
							},
							_ => throw new InvalidProgramException("Unhandled TestResult type")
						}
					}
				};

				_logger.LogInformation($"Tested solution. Session: {session.Id}, Question: {request.QuestionNumber}, Accepted: {submission.Accepted}");

				return submissionResult;
			} catch (Exception exc) {
				_logger.LogError(exc, $"Error testing solution. Request: {request}");
				throw;
			}
		}

		public async Task<ExamResult> EndExamAsync(EmptyRequest request) {
			try {
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

				// retrieve funfact from worker
				var funFact = await _examResultServices.GenerateFunfactAsync(session.Id, Context.ConnectionAborted);

				// Map results
				var examResult = new ExamResult {
					Duration = registeredSession.ExamEndedAt!.Value.ToUnixTimeMilliseconds() - registeredSession.ExamStartedAt!.Value.ToUnixTimeMilliseconds(),
					AnsweredQuestionNumbers = {
						from kvp in registeredSession.SubmissionByQuestionNumber
						where kvp.Value.Accepted
						select kvp.Key
					},
					FunFact = new ExamResult.Types.FunFact {
						DeletionRate = funFact.DeletionRate,
						WordsPerMinute = funFact.WordsPerMinute,
						SubmissionAttempts = funFact.SubmissionAttempts,
					}
				};

				_logger.LogInformation($"Ended exam. Session: {session.Id}");

				return examResult;
			} catch (Exception exc) {
				_logger.LogError(exc, $"Error ending exam. Request: {request}");
				throw;
			}
		}

		public async Task<ExamResult> PassDeadlineAsync(EmptyRequest request) {
			try {
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

				// retrieve funfact from worker
				var funFact = await _examResultServices.GenerateFunfactAsync(session.Id, Context.ConnectionAborted);

				// Map results
				var examResult = new ExamResult {
					Duration = registeredSession.ExamEndedAt!.Value.ToUnixTimeMilliseconds() - registeredSession.ExamStartedAt!.Value.ToUnixTimeMilliseconds(),
					AnsweredQuestionNumbers = {
						from kvp in registeredSession.SubmissionByQuestionNumber
						where kvp.Value.Accepted
						select kvp.Key
					},
					FunFact = new ExamResult.Types.FunFact {
						DeletionRate = funFact.DeletionRate,
						WordsPerMinute = funFact.WordsPerMinute,
						SubmissionAttempts = funFact.SubmissionAttempts,
					}
				};

				_logger.LogInformation($"Passed deadline. Session: {session.Id}");

				return examResult;
			} catch (Exception exc) {
				_logger.LogError(exc, $"Error passing deadline. Request: {request}");
				throw;
			}
		}

		public async Task<ExamResult> ForfeitExamAsync(EmptyRequest request) {
			try {
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

				// retrieve funfact from worker
				var funFact = await _examResultServices.GenerateFunfactAsync(session.Id, Context.ConnectionAborted);

				// Map results
				var examResult = new ExamResult {
					Duration = registeredSession.ExamEndedAt!.Value.ToUnixTimeMilliseconds() - registeredSession.ExamStartedAt!.Value.ToUnixTimeMilliseconds(),
					AnsweredQuestionNumbers = {
						from kvp in registeredSession.SubmissionByQuestionNumber
						where kvp.Value.Accepted
						select kvp.Key
					},
					FunFact = new ExamResult.Types.FunFact {
						DeletionRate = funFact.DeletionRate,
						WordsPerMinute = funFact.WordsPerMinute,
						SubmissionAttempts = funFact.SubmissionAttempts,
					}
				};

				_logger.LogInformation($"Forfeited exam. Session: {session.Id}");

				return examResult;
			} catch (Exception exc) {
				_logger.LogError(exc, $"Error forfeiting exam. Request: {request}");
				throw;
			}
		}

		public async Task SubmitAfterExamSAMAsync(SubmitSAMRequest request) {
			try {
				// Authenticate
				var session = _poormansAuthentication.Authenticate(request.AccessToken);

				// Authorize: Exam must already been ended and second SAM not submitted
				if (session is not RegisteredSession registeredSession) throw new UnauthorizedAccessException("Personal Info not yet submitted");
				if (registeredSession.ExamEndedAt is null) throw new UnauthorizedAccessException("Exam not yet ended");
				if (registeredSession.AfterExamSAM is not null) throw new InvalidProgramException("After exam SAM already submitted");

				// Submit SAM and map results
				await _sessionServices.SubmitAfterExamSAMAsync(
					sessionId: session.Id,
					arousedLevel: request.ArousedLevel,
					pleasedLevel: request.PleasedLevel
				);

				_logger.LogInformation($"Submitted after exam SAM. Session: {session.Id}");
			} catch (Exception exc) {
				_logger.LogError(exc, $"Error submitting after exam SAM. Request: {request}");
				throw;
			}
		}

		public async Task SubmitSolutionSAMAsync(SubmitSolutionSAMRequest request) {
			try {
				// Authenticate
				var session = _poormansAuthentication.Authenticate(request.AccessToken);

				// Authorize: Exam must be in progress and has a result
				if (session is not RegisteredSession registeredSession) throw new UnauthorizedAccessException("Personal Info not yet submitted");
				if (registeredSession.ExamEndedAt is not null) throw new UnauthorizedAccessException("Exam has ended");

				await _sessionServices.SubmitSolutionSAMAsync(
					sessionId: session.Id,
					questionNumber: request.QuestionNumber,
					arousedLevel: request.ArousedLevel,
					pleasedLevel: request.PleasedLevel
				);

				_logger.LogInformation($"Submitted solution SAM. Session: {session.Id}");
			} catch (Exception exc) {
				_logger.LogError(exc, $"Error submitting solution SAM. Request: {request}");
				throw;
			}
		}
	}
}
