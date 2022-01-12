namespace Spectator.DomainModels.SubmissionDomain {
	public record PassingTestResult(
		int TestNumber
	) : TestResultBase(TestNumber);
}
