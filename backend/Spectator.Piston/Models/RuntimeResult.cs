using System.Collections.Immutable;

namespace Spectator.Piston.Models {
	public record RuntimeResult(
		string Language,
		string Version,
		ImmutableHashSet<string> Aliases
	);
}
