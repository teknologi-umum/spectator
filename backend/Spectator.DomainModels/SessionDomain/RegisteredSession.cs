using System;
using System.Collections.Immutable;
using System.Linq;
using System.Text.Json;
using Spectator.DomainEvents.SessionDomain;
using Spectator.DomainModels.SubmissionDomain;
using Spectator.DomainModels.UserDomain;
using Spectator.Primitives;

namespace Spectator.DomainModels.SessionDomain {
	public record RegisteredSession(
		Guid Id,
		DateTimeOffset CreatedAt,
		DateTimeOffset UpdatedAt,
		Locale Locale,
		User User,
		SelfAssessmentManikin? BeforeExamSAM = null,
		SelfAssessmentManikin? AfterExamSAM = null,
		ImmutableArray<int>? QuestionNumbers = null,
		DateTimeOffset? ExamDeadline = null,
		DateTimeOffset? ExamStartedAt = null,
		DateTimeOffset? ExamEndedAt = null,
		ImmutableDictionary<int, Submission>? SubmissionByQuestionNumber = null
	) : SessionBase(Id, CreatedAt, UpdatedAt) {
		public RegisteredSession Apply(LocaleSetEvent @event) {
			if (@event.SessionId != Id) throw new ArgumentException("Applied event has different SessionId", nameof(@event));

			return this with {
				UpdatedAt = @event.Timestamp,
				Locale = @event.Locale
			};
		}

		public RegisteredSession Apply(BeforeExamSAMSubmittedEvent @event) {
			if (BeforeExamSAM != null) throw new InvalidOperationException("SAM already submitted");
			if (@event.SessionId != Id) throw new ArgumentException("Applied event has different SessionId", nameof(@event));

			return this with {
				UpdatedAt = @event.Timestamp,
				BeforeExamSAM = @event.SelfAssessmentManikin
			};
		}

		public RegisteredSession Apply(AfterExamSAMSubmittedEvent @event) {
			if (ExamEndedAt == null) throw new InvalidOperationException("Exam hasn't ended");
			if (AfterExamSAM != null) throw new InvalidOperationException("SAM already submitted");
			if (@event.SessionId != Id) throw new ArgumentException("Applied event has different SessionId", nameof(@event));

			return this with {
				UpdatedAt = @event.Timestamp,
				AfterExamSAM = @event.SelfAssessmentManikin
			};
		}

		public RegisteredSession Apply(ExamStartedEvent @event) {
			if (BeforeExamSAM == null) throw new InvalidOperationException("SAM hasn't been submitted");
			if (ExamStartedAt != null) throw new InvalidOperationException("Exam already started");
			if (@event.SessionId != Id) throw new ArgumentException("Applied event has different SessionId", nameof(@event));

			return this with {
				UpdatedAt = @event.Timestamp,
				QuestionNumbers = @event.QuestionNumbers,
				ExamStartedAt = @event.Timestamp,
				ExamDeadline = @event.Deadline,
				SubmissionByQuestionNumber = ImmutableDictionary<int, Submission>.Empty
			};
		}

		public RegisteredSession Apply(ExamEndedEvent @event) {
			if (ExamStartedAt == null) throw new InvalidOperationException("Exam hasn't been started");
			if (ExamEndedAt != null) throw new InvalidOperationException("Exam already ended");
			if (ExamDeadline == null) throw new InvalidProgramException("Invalid state");
			if (@event.Timestamp >= ExamDeadline) throw new InvalidOperationException("Deadline passed");
			if (QuestionNumbers == null) throw new InvalidProgramException("Invalid state");
			if (SubmissionByQuestionNumber == null) throw new InvalidProgramException("Invalid state");
			if (SubmissionByQuestionNumber.Count(kvp => kvp.Value.Accepted) < QuestionNumbers.Value.Length) throw new InvalidProgramException("Some solutions haven't been accepted");

			return this with {
				UpdatedAt = @event.Timestamp,
				ExamEndedAt = @event.Timestamp
			};
		}

		public RegisteredSession Apply(DeadlinePassedEvent @event) {
			if (ExamStartedAt == null) throw new InvalidOperationException("Exam hasn't been started");
			if (ExamEndedAt != null) throw new InvalidOperationException("Exam already ended");
			if (ExamDeadline == null) throw new InvalidProgramException("Invalid state");
			if (@event.Timestamp < ExamDeadline) throw new InvalidOperationException("Deadline hasn't been passed");

			return this with {
				UpdatedAt = @event.Timestamp,
				ExamEndedAt = @event.Timestamp
			};
		}

		public RegisteredSession Apply(ExamForfeitedEvent @event) {
			if (ExamStartedAt == null) throw new InvalidOperationException("Exam hasn't been started");
			if (ExamEndedAt != null) throw new InvalidOperationException("Exam already ended");
			if (ExamDeadline == null) throw new InvalidProgramException("Invalid state");
			if (@event.Timestamp >= ExamDeadline) throw new InvalidOperationException("Deadline passed");
			if (QuestionNumbers == null) throw new InvalidProgramException("Invalid state");
			if (SubmissionByQuestionNumber == null) throw new InvalidProgramException("Invalid state");
			if (SubmissionByQuestionNumber.Count(kvp => kvp.Value.Accepted) >= QuestionNumbers.Value.Length) throw new InvalidProgramException("All solutions have been accepted");

			return this with {
				UpdatedAt = @event.Timestamp,
				ExamEndedAt = @event.Timestamp
			};
		}

		public RegisteredSession Apply(SolutionAcceptedEvent @event) {
			if (ExamStartedAt == null) throw new InvalidOperationException("Exam hasn't been started");
			if (ExamEndedAt != null) throw new InvalidOperationException("Exam already ended");
			if (ExamDeadline == null) throw new InvalidProgramException("Invalid state");
			if (@event.Timestamp >= ExamDeadline) throw new InvalidOperationException("Deadline passed");
			if (QuestionNumbers == null) throw new InvalidProgramException("Invalid state");
			if (SubmissionByQuestionNumber == null) throw new InvalidProgramException("Invalid state");
			if (!QuestionNumbers.Value.Contains(@event.QuestionNumber)) throw new InvalidOperationException("Invalid question number");

			return this with {
				UpdatedAt = @event.Timestamp,
				SubmissionByQuestionNumber = SubmissionByQuestionNumber.SetItem(@event.QuestionNumber, new Submission(
					QuestionNumber: @event.QuestionNumber,
					Language: @event.Language,
					Solution: @event.Solution,
					ScratchPad: @event.ScratchPad,
					TestResults: JsonSerializer.Deserialize<ImmutableArray<TestResult>>(@event.SerializedTestResults),
					Accepted: true
				))
			};
		}

		public RegisteredSession Apply(SolutionRejectedEvent @event) {
			if (ExamStartedAt == null) throw new InvalidOperationException("Exam hasn't been started");
			if (ExamEndedAt != null) throw new InvalidOperationException("Exam already ended");
			if (ExamDeadline == null) throw new InvalidProgramException("Invalid state");
			if (@event.Timestamp >= ExamDeadline) throw new InvalidOperationException("Deadline passed");
			if (QuestionNumbers == null) throw new InvalidProgramException("Invalid state");
			if (SubmissionByQuestionNumber == null) throw new InvalidProgramException("Invalid state");
			if (!QuestionNumbers.Value.Contains(@event.QuestionNumber)) throw new InvalidOperationException("Invalid question number");
			if (SubmissionByQuestionNumber.TryGetValue(@event.QuestionNumber, out var submission) && submission.Accepted) throw new InvalidOperationException("Another solution already accepted");

			return this with {
				UpdatedAt = @event.Timestamp,
				SubmissionByQuestionNumber = SubmissionByQuestionNumber.SetItem(@event.QuestionNumber, new Submission(
					QuestionNumber: @event.QuestionNumber,
					Language: @event.Language,
					Solution: @event.Solution,
					ScratchPad: @event.ScratchPad,
					TestResults: JsonSerializer.Deserialize<ImmutableArray<TestResult>>(@event.SerializedTestResults),
					Accepted: false
				))
			};
		}
	}
}
