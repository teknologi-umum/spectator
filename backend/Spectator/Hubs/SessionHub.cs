using System;
using System.Collections.Immutable;
using System.Linq;
using System.Threading;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.SignalR;
using Microsoft.Extensions.DependencyInjection;
using SignalRSwaggerGen.Attributes;
using SignalRSwaggerGen.Enums;
using Spectator.DomainServices.QuestionDomain;
using Spectator.DomainServices.SessionDomain;
using Spectator.JwtAuthentication;
using Spectator.Primitives;
using Spectator.Protos.HubInterfaces;
using Spectator.Protos.Session;

namespace Spectator.Hubs {
	[SignalRHub(autoDiscover: AutoDiscover.MethodsAndArgs)]
	public class SessionHub : Hub<ISessionHub>, ISessionHub {
		private readonly SessionServices _sessionServices;
		private readonly IServiceProvider _serviceProvider;

		public SessionHub(
			SessionServices sessionServices,
			IServiceProvider serviceProvider
		) {
			_sessionServices = sessionServices;
			_serviceProvider = serviceProvider;
		}

		public async Task<SessionReply> StartSessionAsync() {
			var session = await _sessionServices.StartSessionAsync();
			var jwtAuthenticationServices = _serviceProvider.GetRequiredService<JwtAuthenticationServices>();
			var tokenPayload = jwtAuthenticationServices.CreatePayload(session.Id);
			return new SessionReply {
				AccessToken = jwtAuthenticationServices.EncodeToken(tokenPayload)
			};
		}

		[Authorize(AuthPolicy.ANONYMOUS)]
		public async Task SubmitPersonalInfoAsync(PersonalInfo personalInfo) {
			if (Context.User == null) throw new UnauthorizedAccessException();
			var tokenPayload = TokenPayload.FromClaimsPrincipal(Context.User);
			await _sessionServices.SubmitPersonalInfoAsync(
				sessionId: tokenPayload.SessionId,
				studentNumber: personalInfo.StudentNumber,
				yearsOfExperience: personalInfo.YearsOfExperience,
				hoursOfPractice: personalInfo.HoursOfPractice,
				familiarLanguages: personalInfo.FamiliarLanguages
			);
		}

		[Authorize(AuthPolicy.REGISTERED)]
		public async Task SubmitBeforeExamSAMAsync(SAM sam) {
			if (Context.User == null) throw new UnauthorizedAccessException();
			var tokenPayload = TokenPayload.FromClaimsPrincipal(Context.User);
			await _sessionServices.SubmitBeforeExamSAMAsync(
				sessionId: tokenPayload.SessionId,
				arousedLevel: sam.ArousedLevel,
				pleasedLevel: sam.PleasedLevel
			);
		}

		[Authorize(AuthPolicy.READY_TO_TAKE_EXAM)]
		public async Task<Exam> StartExamAsync() {
			if (Context.User == null) throw new UnauthorizedAccessException();
			var tokenPayload = TokenPayload.FromClaimsPrincipal(Context.User);
			// TODO: accept locale parameter
			// HACK: temporary dummy value Locale.ID
			var session = await _sessionServices.StartExamAsync(tokenPayload.SessionId, Locale.ID);
			var questions = await _serviceProvider.GetRequiredService<QuestionServices>().GetAllAsync(Locale.ID, CancellationToken.None);
			var questionById = questions.ToImmutableDictionary(q => q.QuestionNumber);
			return new Exam {
				Deadline = session.ExamDeadline!.Value.ToUnixTimeMilliseconds(),
				Questions = {
					from questionNumber in session.QuestionNumbers!.Value
					let question = questionById[questionNumber]
					select new Question {
						QuestionNumber = questionNumber,
						Title = question.Title,
						Instruction = question.Instruction,
						LanguageAndTemplates = {
							from kvp in question.TemplateByLanguage
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

		[Authorize(AuthPolicy.TAKING_EXAM)]
		public async Task<Exam> ResumeExamAsync() {
			if (Context.User == null) throw new UnauthorizedAccessException();
			var tokenPayload = TokenPayload.FromClaimsPrincipal(Context.User);
			var session = await _sessionServices.ResumeExamAsync(tokenPayload.SessionId, Context.ConnectionAborted);
			// TODO: accept locale parameter
			// HACK: temporary dummy value Locale.ID
			var questions = await _serviceProvider.GetRequiredService<QuestionServices>().GetAllAsync(Locale.ID, CancellationToken.None);
			var questionById = questions.ToImmutableDictionary(q => q.QuestionNumber);
			return new Exam {
				Deadline = session.ExamDeadline!.Value.ToUnixTimeMilliseconds(),
				Questions = {
					from questionNumber in session.QuestionNumbers!.Value
					let question = questionById[questionNumber]
					select new Question {
						QuestionNumber = questionNumber,
						Title = question.Title,
						Instruction = question.Instruction,
						LanguageAndTemplates = {
							from kvp in question.TemplateByLanguage
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

		[Authorize(AuthPolicy.TAKING_EXAM)]
		public async Task<SubmissionResult> SubmitSolutionAsync(SubmissionRequest submissionRequest) {
			if (Context.User == null) throw new UnauthorizedAccessException();
			var tokenPayload = TokenPayload.FromClaimsPrincipal(Context.User);
			var submission = await _sessionServices.SubmitSolutionAsync(
				sessionId: tokenPayload.SessionId,
				questionNumber: submissionRequest.QuestionNumber,
				language: (Language)submissionRequest.Language,
				solution: submissionRequest.Solution,
				scratchPad: submissionRequest.ScratchPad
			);
			return new SubmissionResult {
				Accepted = submission.Accepted,
				TestResults = {
					from testResult in submission.TestResults
					select new SubmissionResult.Types.TestResult {
						Success = testResult.Success,
						Message = testResult.Message
					}
				}
			};
		}

		[Authorize(AuthPolicy.TAKING_EXAM)]
		public async Task<ExamResult> EndExamAsync() {
			if (Context.User == null) throw new UnauthorizedAccessException();
			var tokenPayload = TokenPayload.FromClaimsPrincipal(Context.User);
			var session = await _sessionServices.EndExamAsync(tokenPayload.SessionId);
			return new ExamResult {
				Duration = session.ExamEndedAt!.Value.ToUnixTimeMilliseconds() - session.ExamStartedAt!.Value.ToUnixTimeMilliseconds(),
				AnsweredQuestionNumbers = {
					from kvp in session.SubmissionByQuestionNumber
					where kvp.Value.Accepted
					select kvp.Key
				}
			};
		}

		[Authorize(AuthPolicy.TAKING_EXAM)]
		public async Task<ExamResult> PassDeadlineAsync() {
			if (Context.User == null) throw new UnauthorizedAccessException();
			var tokenPayload = TokenPayload.FromClaimsPrincipal(Context.User);
			var session = await _sessionServices.PassDeadlineAsync(tokenPayload.SessionId);
			return new ExamResult {
				Duration = session.ExamEndedAt!.Value.ToUnixTimeMilliseconds() - session.ExamStartedAt!.Value.ToUnixTimeMilliseconds(),
				AnsweredQuestionNumbers = {
					from kvp in session.SubmissionByQuestionNumber
					where kvp.Value.Accepted
					select kvp.Key
				}
			};
		}

		[Authorize(AuthPolicy.TAKING_EXAM)]
		public async Task<ExamResult> ForfeitExamAsync() {
			if (Context.User == null) throw new UnauthorizedAccessException();
			var tokenPayload = TokenPayload.FromClaimsPrincipal(Context.User);
			var session = await _sessionServices.ForfeitExamAsync(tokenPayload.SessionId);
			return new ExamResult {
				Duration = session.ExamEndedAt!.Value.ToUnixTimeMilliseconds() - session.ExamStartedAt!.Value.ToUnixTimeMilliseconds(),
				AnsweredQuestionNumbers = {
					from kvp in session.SubmissionByQuestionNumber
					where kvp.Value.Accepted
					select kvp.Key
				}
			};
		}

		[Authorize(AuthPolicy.HAS_TAKEN_EXAM)]
		public async Task SubmitAfterExamSAM(SAM sam) {
			if (Context.User == null) throw new UnauthorizedAccessException();
			var tokenPayload = TokenPayload.FromClaimsPrincipal(Context.User);
			await _sessionServices.SubmitAfterExamSAMAsync(
				sessionId: tokenPayload.SessionId,
				arousedLevel: sam.ArousedLevel,
				pleasedLevel: sam.PleasedLevel
			);
		}
	}
}
