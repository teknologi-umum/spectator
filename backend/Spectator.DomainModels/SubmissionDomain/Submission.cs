namespace Spectator.DomainModels.SubmissionDomain {
	public record Submission(
		int QuestionNumber,
		string Solution,
		string ScratchPad,
		string ConsoleOutput,
		bool Accepted
	);
}
