using System.Collections.Immutable;
using Spectator.DomainEvents.SessionDomain;
using Spectator.Primitives;

namespace Spectator.DomainModels.SubmissionDomain {
	public record Submission(
		int QuestionNumber,
		Language Language,
		string Solution,
		string ScratchPad,
		ImmutableArray<TestResultBase> TestResults,
		SelfAssessmentManikin? SAMTestResult,
		bool Accepted
	) {
		public Submission Apply(SolutionSAMSubmittedEvent @event) => this with {
			SAMTestResult = @event.SelfAssessmentManikin
		};
	}
}
