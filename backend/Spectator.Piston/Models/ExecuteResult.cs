namespace Spectator.Piston.Models {
	public record ExecuteResult(
		string Language,
		string Version,
		RunResult Run
	);
}
