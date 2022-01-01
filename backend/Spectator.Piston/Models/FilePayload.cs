namespace Spectator.Piston.Models {
	public record FilePayload(
		string Content,
		string? Name = null,
		string? Encoding = "utf8"
	);
}
