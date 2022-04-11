using System;
using System.Collections.Immutable;
using FluentAssertions;
using Spectator.DomainEvents.SessionDomain;
using Spectator.DomainModels.SessionDomain;
using Spectator.Primitives;
using Xunit;

namespace Spectator.DomainModels.Tests {
	public class SessionTests {
		[Fact]
		public AnonymousSession CanCreateAnonymousSession() {
			var sessionId = Guid.NewGuid();
			var timestamp = DateTimeOffset.UtcNow;
			var sessionStartedEvent = new SessionStartedEvent(sessionId, timestamp, Locale.ID);
			var anonymousSession = AnonymousSession.From(sessionStartedEvent);
			anonymousSession.Id.Should().Be(sessionId);
			anonymousSession.CreatedAt.Should().Be(timestamp);
			anonymousSession.UpdatedAt.Should().Be(timestamp);
			anonymousSession.Locale.Should().Be(Locale.ID);

			return anonymousSession;
		}

		[Fact]
		public void CanChangeLocaleOfAnonymousSession() {
			var timestamp = DateTimeOffset.UtcNow.AddSeconds(1);
			var anonymousSession = CanCreateAnonymousSession();
			var localeSetEvent = new LocaleSetEvent(
				SessionId: anonymousSession.Id,
				Timestamp: timestamp,
				Locale: Locale.EN
			);
			var resultingSession = anonymousSession.Apply(localeSetEvent);
			resultingSession.Id.Should().Be(anonymousSession.Id);
			resultingSession.CreatedAt.Should().Be(anonymousSession.CreatedAt);
			resultingSession.UpdatedAt.Should().Be(timestamp);
			resultingSession.Locale.Should().Be(Locale.EN);
		}

		[Fact]
		public RegisteredSession CanCreateRegisteredSession() {
			var timestamp = DateTimeOffset.UtcNow.AddSeconds(1);
			var anonymousSession = CanCreateAnonymousSession();
			var personalInfoSubmittedEvent = new PersonalInfoSubmittedEvent(
				SessionId: anonymousSession.Id,
				Timestamp: timestamp,
				StudentNumber: "1234567890",
				YearsOfExperience: 1,
				HoursOfPractice: 2,
				FamiliarLanguages: "C"
			);
			var registeredSession = anonymousSession.Apply(personalInfoSubmittedEvent);
			registeredSession.Id.Should().Be(anonymousSession.Id);
			registeredSession.CreatedAt.Should().Be(anonymousSession.CreatedAt);
			registeredSession.UpdatedAt.Should().Be(timestamp);
			registeredSession.Locale.Should().Be(Locale.ID);
			registeredSession.User.Should().NotBeNull();
			registeredSession.BeforeExamSAM.Should().BeNull();
			registeredSession.AfterExamSAM.Should().BeNull();
			registeredSession.QuestionNumbers.Should().BeNull();
			registeredSession.SubmissionByQuestionNumber.Should().BeNull();
			registeredSession.ExamStartedAt.Should().BeNull();
			registeredSession.ExamEndedAt.Should().BeNull();
			registeredSession.ExamDeadline.Should().BeNull();

			registeredSession.User.CreatedAt.Should().Be(timestamp);
			registeredSession.User.UpdatedAt.Should().Be(timestamp);
			registeredSession.User.StudentNumber.Should().Be("1234567890");
			registeredSession.User.YearsOfExperience.Should().Be(1);
			registeredSession.User.HoursOfPractice.Should().Be(2);
			registeredSession.User.FamiliarLanguages.Should().Be("C");

			return registeredSession;
		}

		[Fact]
		public void CannotCreateRegisteredSessionUsingInvalidEvent() {
			var timestamp = DateTimeOffset.UtcNow.AddSeconds(1);
			var anonymousSession = CanCreateAnonymousSession();
			var personalInfoSubmittedEvent = new PersonalInfoSubmittedEvent(
				SessionId: Guid.NewGuid(),
				Timestamp: timestamp,
				StudentNumber: "1234567890",
				YearsOfExperience: 1,
				HoursOfPractice: 2,
				FamiliarLanguages: "C"
			);
			new Action(() => anonymousSession.Apply(personalInfoSubmittedEvent)).Should().Throw<ArgumentException>()
				.And.Message.Should().Be("Applied event has different SessionId (Parameter 'event')");
		}

		[Fact]
		public void CanChangeLocaleOfRegisteredSession() {
			var timestamp = DateTimeOffset.UtcNow.AddSeconds(2);
			var registeredSession = CanCreateRegisteredSession();
			var localeSetEvent = new LocaleSetEvent(
				SessionId: registeredSession.Id,
				Timestamp: timestamp,
				Locale: Locale.EN
			);
			var resultingSession = registeredSession.Apply(localeSetEvent);

			resultingSession.Id.Should().Be(registeredSession.Id);
			resultingSession.CreatedAt.Should().Be(registeredSession.CreatedAt);
			resultingSession.UpdatedAt.Should().Be(timestamp);
			resultingSession.Locale.Should().Be(Locale.EN);
			resultingSession.User.Should().NotBeNull();
			resultingSession.BeforeExamSAM.Should().BeNull();
			resultingSession.AfterExamSAM.Should().BeNull();
			resultingSession.QuestionNumbers.Should().BeNull();
			resultingSession.SubmissionByQuestionNumber.Should().BeNull();
			resultingSession.ExamStartedAt.Should().BeNull();
			resultingSession.ExamEndedAt.Should().BeNull();
			resultingSession.ExamDeadline.Should().BeNull();

			resultingSession.User.CreatedAt.Should().Be(registeredSession.User.CreatedAt);
			resultingSession.User.UpdatedAt.Should().Be(registeredSession.User.UpdatedAt);
			resultingSession.User.StudentNumber.Should().Be("1234567890");
			resultingSession.User.YearsOfExperience.Should().Be(1);
			resultingSession.User.HoursOfPractice.Should().Be(2);
			resultingSession.User.FamiliarLanguages.Should().Be("C");
		}

		[Fact]
		public void CannotChangeLocaleUsingInvalidEvent() {
			var timestamp = DateTimeOffset.UtcNow.AddSeconds(2);
			var registeredSession = CanCreateRegisteredSession();
			var localeSetEvent = new LocaleSetEvent(
				SessionId: Guid.NewGuid(),
				Timestamp: timestamp,
				Locale: Locale.EN
			);
			new Action(() => registeredSession.Apply(localeSetEvent)).Should().Throw<ArgumentException>()
				.And.Message.Should().Be("Applied event has different SessionId (Parameter 'event')");
		}

		[Fact]
		public RegisteredSession CanSubmitBeforeExamSAM() {
			var timestamp = DateTimeOffset.UtcNow.AddSeconds(2);
			var registeredSession = CanCreateRegisteredSession();
			var beforeExamSAMSubmittedEvent = new BeforeExamSAMSubmittedEvent(
				SessionId: registeredSession.Id,
				Timestamp: timestamp,
				SelfAssessmentManikin: new(2, 3)
			);
			var sessionWithBeforeExamSAM = registeredSession.Apply(beforeExamSAMSubmittedEvent);
			sessionWithBeforeExamSAM.Id.Should().Be(registeredSession.Id);
			sessionWithBeforeExamSAM.CreatedAt.Should().Be(registeredSession.CreatedAt);
			sessionWithBeforeExamSAM.UpdatedAt.Should().Be(timestamp);
			sessionWithBeforeExamSAM.Locale.Should().Be(Locale.ID);
			sessionWithBeforeExamSAM.User.Should().NotBeNull();
			sessionWithBeforeExamSAM.BeforeExamSAM.Should().NotBeNull();
			sessionWithBeforeExamSAM.AfterExamSAM.Should().BeNull();
			sessionWithBeforeExamSAM.QuestionNumbers.Should().BeNull();
			sessionWithBeforeExamSAM.SubmissionByQuestionNumber.Should().BeNull();
			sessionWithBeforeExamSAM.ExamStartedAt.Should().BeNull();
			sessionWithBeforeExamSAM.ExamEndedAt.Should().BeNull();
			sessionWithBeforeExamSAM.ExamDeadline.Should().BeNull();

			sessionWithBeforeExamSAM.BeforeExamSAM!.ArousedLevel.Should().Be(2);
			sessionWithBeforeExamSAM.BeforeExamSAM.PleasedLevel.Should().Be(3);

			return sessionWithBeforeExamSAM;
		}

		[Fact]
		public void CannotSubmitBeforeExamSAMTwice() {
			var timestamp = DateTimeOffset.UtcNow.AddSeconds(3);
			var sessionWithBeforeExamSAM = CanSubmitBeforeExamSAM();
			var beforeExamSAMSubmittedEvent = new BeforeExamSAMSubmittedEvent(
				SessionId: sessionWithBeforeExamSAM.Id,
				Timestamp: timestamp,
				SelfAssessmentManikin: new(2, 3)
			);
			new Action(() => sessionWithBeforeExamSAM.Apply(beforeExamSAMSubmittedEvent)).Should().Throw<InvalidOperationException>()
				.And.Message.Should().Be("SAM already submitted");
		}

		[Fact]
		public void CannotSubmitBeforeExamSAMUsingInvalidEvent() {
			var timestamp = DateTimeOffset.UtcNow.AddSeconds(2);
			var registeredSession = CanCreateRegisteredSession();
			var beforeExamSAMSubmittedEvent = new BeforeExamSAMSubmittedEvent(
				SessionId: Guid.NewGuid(),
				Timestamp: timestamp,
				SelfAssessmentManikin: new(2, 3)
			);
			new Action(() => registeredSession.Apply(beforeExamSAMSubmittedEvent)).Should().Throw<ArgumentException>()
				.And.Message.Should().Be("Applied event has different SessionId (Parameter 'event')");
		}

		[Fact]
		public RegisteredSession CanStartExam() {
			var timestamp = DateTimeOffset.UtcNow.AddSeconds(3);
			var sessionWithBeforeExamSAM = CanSubmitBeforeExamSAM();
			var examStartedEvent = new ExamStartedEvent(
				SessionId: sessionWithBeforeExamSAM.Id,
				Timestamp: timestamp,
				QuestionNumbers: ImmutableArray.Create(1, 2, 3, 4, 5, 6),
				Deadline: timestamp.AddMinutes(71)
			);
			var examSession = sessionWithBeforeExamSAM.Apply(examStartedEvent);
			examSession.Id.Should().Be(sessionWithBeforeExamSAM.Id);
			examSession.CreatedAt.Should().Be(sessionWithBeforeExamSAM.CreatedAt);
			examSession.UpdatedAt.Should().Be(timestamp);
			examSession.Locale.Should().Be(Locale.ID);
			examSession.User.Should().NotBeNull();
			examSession.BeforeExamSAM.Should().NotBeNull();
			examSession.AfterExamSAM.Should().BeNull();
			examSession.QuestionNumbers.Should().NotBeNull();
			examSession.SubmissionByQuestionNumber.Should().NotBeNull();
			examSession.ExamStartedAt.Should().NotBeNull();
			examSession.ExamEndedAt.Should().BeNull();
			examSession.ExamDeadline.Should().NotBeNull();

			examSession.QuestionNumbers!.Value.Should().ContainInOrder(1, 2, 3, 4, 5, 6);
			examSession.SubmissionByQuestionNumber.Should().BeEmpty();
			examSession.ExamStartedAt!.Value.Should().Be(timestamp);
			examSession.ExamDeadline!.Value.Should().Be(timestamp.AddMinutes(71));

			return examSession;
		}

		[Fact]
		public void CannotStartExamBeforeSubmittingSAM() {
			var timestamp = DateTimeOffset.UtcNow.AddSeconds(2);
			var registeredSession = CanCreateRegisteredSession();
			var examStartedEvent = new ExamStartedEvent(
				SessionId: registeredSession.Id,
				Timestamp: timestamp,
				QuestionNumbers: ImmutableArray.Create(1, 2, 3, 4, 5, 6),
				Deadline: timestamp.AddMinutes(71)
			);
			new Action(() => registeredSession.Apply(examStartedEvent)).Should().Throw<InvalidOperationException>()
				.And.Message.Should().Be("SAM hasn't been submitted");
		}

		[Fact]
		public void CannotStartExamTwice() {
			var timestamp = DateTimeOffset.UtcNow.AddSeconds(4);
			var examSession = CanStartExam();
			var examStartedEvent = new ExamStartedEvent(
				SessionId: examSession.Id,
				Timestamp: timestamp,
				QuestionNumbers: ImmutableArray.Create(1, 2, 3, 4, 5, 6),
				Deadline: timestamp.AddMinutes(75)
			);
			new Action(() => examSession.Apply(examStartedEvent)).Should().Throw<InvalidOperationException>()
				.And.Message.Should().Be("Exam already started");
		}

		[Fact]
		public void CannotStartExamUsingInvalidEvent() {
			var timestamp = DateTimeOffset.UtcNow.AddSeconds(3);
			var sessionWithBeforeExamSAM = CanSubmitBeforeExamSAM();
			var examStartedEvent = new ExamStartedEvent(
				SessionId: Guid.NewGuid(),
				Timestamp: timestamp,
				QuestionNumbers: ImmutableArray.Create(1, 2, 3, 4, 5, 6),
				Deadline: timestamp.AddMinutes(71)
			);
			new Action(() => sessionWithBeforeExamSAM.Apply(examStartedEvent)).Should().Throw<ArgumentException>()
				.And.Message.Should().Be("Applied event has different SessionId (Parameter 'event')");
		}

		//[Fact]
		//public void CanEndExam() {
		//	var timestamp = DateTimeOffset.UtcNow.AddSeconds(4);
		//	var sessionWithStartedExam = CanStartExam();
		//	var examEndedEvent = new ExamEndedEvent(
		//		SessionId: sessionWithStartedExam.Id,
		//		Timestamp: timestamp
		//	);
		//	var resultingSession = sessionWithStartedExam.Apply(examEndedEvent);
		//	resultingSession.Id.Should().Be(sessionWithStartedExam.Id);
		//	resultingSession.CreatedAt.Should().Be(sessionWithStartedExam.CreatedAt);
		//	resultingSession.UpdatedAt.Should().Be(timestamp);
		//	resultingSession.Locale.Should().Be(Locale.ID);
		//	resultingSession.User.Should().NotBeNull();
		//	resultingSession.BeforeExamSAM.Should().NotBeNull();
		//	resultingSession.AfterExamSAM.Should().BeNull();
		//	resultingSession.QuestionNumbers.Should().NotBeNull();
		//	resultingSession.SubmissionByQuestionNumber.Should().NotBeNull();
		//	resultingSession.ExamStartedAt.Should().NotBeNull();
		//	resultingSession.ExamEndedAt.Should().BeNull();
		//	resultingSession.ExamDeadline.Should().NotBeNull();

		//	resultingSession.QuestionNumbers!.Value.Should().ContainInOrder(1, 2, 3, 4, 5, 6);
		//	resultingSession.SubmissionByQuestionNumber.Should().BeEmpty();
		//	resultingSession.ExamStartedAt!.Value.Should().Be(timestamp);
		//	resultingSession.ExamDeadline!.Value.Should().Be(timestamp.AddMinutes(71));
		//}

		// Wanna help? Install Visual Studio + Fine Code Coverage or Jetbrains Rider, then implement following test methods:

		// TODO: CanEndExam
		// TODO: CannotEndExamBeforeExamStarted
		// TODO: CannotEndExamTwice
		// TODO: CannotEndExamWithoutDeadline (corrupt session manually to test)
		// TODO: CannotEndExamAfterDeadlinePassed
		// TODO: CannotEndExamWithoutQuestionNumbers (corrupt session manually to test)
		// TODO: CannotEndExamWithoutSubmissionDictionary (corrupt session manually to test)
		// TODO: CannotEndExamWithUnansweredQuestions
		// TODO: CanPassDeadline
		// TODO: CannotPassDeadlineBeforeExamStarted
		// TODO: CannotPassDeadlineAfterExamEnded
		// TODO: CannotPassDeadlineBeforeDeadlinePassed
		// TODO: CanForfeitExam
		// TODO: CannotForfeitExamBeforeExamStarted
		// TODO: CannotForfeitExamAfterExamEnded
		// TODO: CannotForfeitExamWithoutDeadline (corrupt session manually to test)
		// TODO: CannotForfeitExamAfterDeadlinePassed
		// TODO: CannotForfeitExamWithoutQuestionNumbers (corrupt session manually to test)
		// TODO: CannotForfeitExamWithoutSubmissionDictionary (corrupt session manually to test)
		// TODO: CannotForfeitExamAfterAllSubmissionsAccepted
		// TODO: CanAcceptSolution
		// TODO: CannotAcceptSolutionBeforeExamStarted
		// TODO: CannotAcceptSolutionAfterExamEnded
		// TODO: CannotAcceptSolutionWithoutDeadline (corrupt session manually to test)
		// TODO: CannotAcceptSolutionAfterDeadlinePassed
		// TODO: CannotAcceptSolutionWithoutQuestionNumbers (corrupt session manually to test)
		// TODO: CannotAcceptSolutionWithoutSubmissionDictionary (corrupt session manually to test)
		// TODO: CannotAcceptSolutionWithInvalidQuestionNumber
		// TODO: CanRejectSolution
		// TODO: CannotRejectSolutionBeforeExamStarted
		// TODO: CannotRejectSolutionAfterExamEnded
		// TODO: CannotRejectSolutionWithoutDeadline (corrupt session manually to test)
		// TODO: CannotRejectSolutionAfterDeadlinePassed
		// TODO: CannotRejectSolutionWithoutQuestionNumbers (corrupt session manually to test)
		// TODO: CannotRejectSolutionWithoutSubmissionDictionary (corrupt session manually to test)
		// TODO: CannotRejectSolutionWithInvalidQuestionNumber
		// TODO: CannotRejectAlreadyAcceptedSolution
	}
}
