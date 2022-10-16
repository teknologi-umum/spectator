namespace Spectator.DomainModels.SubmissionDomain {
	public record InvalidInputResult(
		string Stderr
	) : TestResultBase(0);
}
