using System.Collections.Immutable;
using Spectator.Primitives;

namespace Spectator.DomainModels.QuestionDomain {
	public record Question(
		int QuestionNumber,
		ImmutableDictionary<Locale, string> TitleByLocale,
		ImmutableDictionary<Locale, string> InstructionByLocale,
		ImmutableDictionary<Locale, ImmutableDictionary<Language, string>> TemplateByLanguageByLocale,
		ImmutableDictionary<Language, string> AssertionByLanguage
	);
}
