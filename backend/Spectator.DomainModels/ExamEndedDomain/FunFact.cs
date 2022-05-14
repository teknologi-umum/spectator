namespace Spectator.DomainModels.ExamEndedDomain {
	public class Funfact {
		public long WordsPerMinute { get; init; }
		public double DeletionRate { get; init; }
		public long SubmissionAttempts { get; init; }

		public Funfact(long wordsPerMinute, double deletionRate, long submissionAttempts) {
			WordsPerMinute = wordsPerMinute;
			DeletionRate = deletionRate;
			SubmissionAttempts = submissionAttempts;
		}
	}
}
