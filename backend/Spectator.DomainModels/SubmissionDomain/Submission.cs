using System.Collections.Immutable;
using Spectator.Primitives;

namespace Spectator.DomainModels.SubmissionDomain {
	public record Submission(
		int QuestionNumber,
		Language Language,
		string Solution,
		string ScratchPad,
		ImmutableArray<TestResult> TestResults,
		bool Accepted
	);
}
