namespace Spectator.DomainModels.SubmissionDomain {
	public record CompileErrorResult(
		string Stderr
	) : TestResultBase(0);
}
