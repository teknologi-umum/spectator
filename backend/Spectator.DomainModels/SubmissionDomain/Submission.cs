using Spectator.Primitives;

namespace Spectator.DomainModels.SubmissionDomain {
	public record Submission(
		int QuestionNumber,
		Language Language,
		string Solution,
		string ScratchPad,
		string? ErrorMessage,
		string ConsoleOutput,
		bool Accepted
	);
}
