namespace Spectator.Piston.Models {
	public record RunResult(
		string Stdout,
		string Stderr,
		string Output,
		int Code,
		string? Signal
	);
}
