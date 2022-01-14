namespace Spectator.DomainModels.SubmissionDomain {
	public record FailingTestResult(
		int TestNumber,
		string ExpectedStdout,
		string ActualStdout
	) : TestResultBase(TestNumber);
}
