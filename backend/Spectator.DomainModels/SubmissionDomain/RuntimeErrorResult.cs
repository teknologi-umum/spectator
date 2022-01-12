namespace Spectator.DomainModels.SubmissionDomain {
	public record RuntimeErrorResult(
		string Stderr
	) : TestResultBase(0);
}
