using System;
using FluentAssertions;
using Spectator.DomainEvents.SessionDomain;
using Spectator.DomainModels.SessionDomain;
using Xunit;

namespace Spectator.DomainModels.Tests {
	public class SessionTests {
		[Fact]
		public AnonymousSession CanCreateAnonymousSession() {
			var sessionId = Guid.NewGuid();
			var timestamp = DateTimeOffset.UtcNow;
			var sessionStartedEvent = new SessionStartedEvent(sessionId, timestamp);
			var anonymousSession = AnonymousSession.From(sessionStartedEvent);
			anonymousSession.Id.Should().Be(sessionId);
			anonymousSession.CreatedAt.Should().Be(timestamp);
			anonymousSession.UpdatedAt.Should().Be(timestamp);

			return anonymousSession;
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
		public RegisteredSession CanSubmitBeforeExamSAMEvent() {
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
		public void CannotSubmitBeforeExamSAMEventTwice() {
			var timestamp = DateTimeOffset.UtcNow.AddSeconds(3);
			var sessionWithBeforeExamSAM = CanSubmitBeforeExamSAMEvent();
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
	}
}
