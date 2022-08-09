namespace Spectator.DomainModels.SubmissionDomain {
	public record PassingTestResult(
		int TestNumber,
		string ExpectedStdout,
		string ActualStdout,
		string ArgumentsStdout
	) : TestResultBase(TestNumber);
}
