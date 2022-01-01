namespace Spectator.JwtAuthentication {
	public static class AuthPolicy {
		public const string ANONYMOUS = nameof(ANONYMOUS);
		public const string REGISTERED = nameof(REGISTERED);
		public const string READY_TO_TAKE_EXAM = nameof(READY_TO_TAKE_EXAM);
		public const string TAKING_EXAM = nameof(TAKING_EXAM);
		public const string HAS_TAKEN_EXAM = nameof(HAS_TAKEN_EXAM);
		public const string COMPLETED = nameof(COMPLETED);
	}
}
