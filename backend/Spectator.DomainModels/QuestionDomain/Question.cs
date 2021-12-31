using System.Collections.Immutable;
using Spectator.Primitives;

namespace Spectator.DomainModels.QuestionDomain {
	public record Question(
		int QuestionNumber,
		string Title,
		string Instruction,
		ImmutableDictionary<Language, string> BoilerplateByLanguage
	);
}
