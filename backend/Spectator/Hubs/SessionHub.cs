using System;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.SignalR;
using Microsoft.Extensions.DependencyInjection;
using SignalRSwaggerGen.Attributes;
using SignalRSwaggerGen.Enums;
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

		public async Task<SessionReply> StartSessionAsync(LocaleInfo localeInfo) {
			var session = await _sessionServices.StartSessionAsync((Locale)localeInfo.Locale);
			var jwtAuthenticationServices = _serviceProvider.GetRequiredService<JwtAuthenticationServices>();
			var tokenPayload = jwtAuthenticationServices.CreatePayload(session.Id);
			return new SessionReply {
				AccessToken = jwtAuthenticationServices.EncodeToken(tokenPayload)
			};
		}

		[Authorize]
		public Task SetLocaleAsync(LocaleInfo localeInfo) {
			if (Context.User == null) throw new UnauthorizedAccessException();
			var tokenPayload = TokenPayload.FromClaimsPrincipal(Context.User);
			return _sessionServices.SetLocaleAsync(
				sessionId: tokenPayload.SessionId,
				locale: (Locale)localeInfo.Locale
			);
		}

		[Authorize(AuthPolicy.ANONYMOUS)]
		public Task SubmitPersonalInfoAsync(PersonalInfo personalInfo) {
			if (Context.User == null) throw new UnauthorizedAccessException();
			var tokenPayload = TokenPayload.FromClaimsPrincipal(Context.User);
			return _sessionServices.SubmitPersonalInfoAsync(
				sessionId: tokenPayload.SessionId,
				studentNumber: personalInfo.StudentNumber,
				yearsOfExperience: personalInfo.YearsOfExperience,
				hoursOfPractice: personalInfo.HoursOfPractice,
				familiarLanguages: personalInfo.FamiliarLanguages
			);
		}

		[Authorize(AuthPolicy.REGISTERED)]
		public Task SubmitBeforeExamSAMAsync(SAM sam) {
			if (Context.User == null) throw new UnauthorizedAccessException();
			var tokenPayload = TokenPayload.FromClaimsPrincipal(Context.User);
			return _sessionServices.SubmitBeforeExamSAMAsync(
				sessionId: tokenPayload.SessionId,
				arousedLevel: sam.ArousedLevel,
				pleasedLevel: sam.PleasedLevel
			);
		}

		[Authorize(AuthPolicy.READY_TO_TAKE_EXAM)]
		public async Task<Exam> StartExamAsync() {
			if (Context.User == null) throw new UnauthorizedAccessException();
			var tokenPayload = TokenPayload.FromClaimsPrincipal(Context.User);
			(var session, var questionByQuestionNumber) = await _sessionServices.StartExamAsync(tokenPayload.SessionId);
			return new Exam {
				Deadline = session.ExamDeadline!.Value.ToUnixTimeMilliseconds(),
				Questions = {
					from questionNumber in session.QuestionNumbers!.Value
					let question = questionByQuestionNumber[questionNumber]
					select new Question {
						QuestionNumber = questionNumber,
						Title = question.TitleByLocale[session.Locale],
						Instruction = question.InstructionByLocale[session.Locale],
						LanguageAndTemplates = {
							from kvp in question.TemplateByLanguageByLocale[session.Locale]
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
			(var session, var questionByQuestionNumber) = await _sessionServices.ResumeExamAsync(tokenPayload.SessionId, Context.ConnectionAborted);
			return new Exam {
				Deadline = session.ExamDeadline!.Value.ToUnixTimeMilliseconds(),
				Questions = {
					from questionNumber in session.QuestionNumbers!.Value
					let question = questionByQuestionNumber[questionNumber]
					select new Question {
						QuestionNumber = questionNumber,
						Title = question.TitleByLocale[session.Locale],
						Instruction = question.InstructionByLocale[session.Locale],
						LanguageAndTemplates = {
							from kvp in question.TemplateByLanguageByLocale[session.Locale]
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
		public Task SubmitAfterExamSAM(SAM sam) {
			if (Context.User == null) throw new UnauthorizedAccessException();
			var tokenPayload = TokenPayload.FromClaimsPrincipal(Context.User);
			return _sessionServices.SubmitAfterExamSAMAsync(
				sessionId: tokenPayload.SessionId,
				arousedLevel: sam.ArousedLevel,
				pleasedLevel: sam.PleasedLevel
			);
		}
	}
}
