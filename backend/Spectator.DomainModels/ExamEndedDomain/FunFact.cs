namespace Spectator.DomainModels.ExamEndedDomain {
	public record Funfact {
		public long WordsPerMinute { get; }
		public double DeletionRate { get; }
		public long SubmissionAttempts { get; }

		public Funfact(long wordsPerMinute, double deletionRate, long submissionAttempts) {
			WordsPerMinute = wordsPerMinute;
			DeletionRate = deletionRate;
			SubmissionAttempts = submissionAttempts;
		}
	}
}
