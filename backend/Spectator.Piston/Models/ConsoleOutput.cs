namespace Spectator.Piston.Models {
	public record ConsoleOutput(
		string Stdout,
		string Stderr,
		string Output,
		int Code,
		string? Signal
	);
}
