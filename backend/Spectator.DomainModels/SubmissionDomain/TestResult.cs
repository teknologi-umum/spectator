namespace Spectator.DomainModels.SubmissionDomain {
	public record TestResult(
		bool Success,
		string Message
	);
}
