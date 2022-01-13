namespace Spectator.Piston.Models {
	public record ExecuteResult(
		string Language,
		string Version,
		ConsoleOutput Compile,
		ConsoleOutput Run
	);
}
