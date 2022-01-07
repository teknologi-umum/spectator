using System.Collections.Immutable;
using System.Text.Json.Serialization;

namespace Spectator.RepositoryDALs.JsonModels {
	internal record Question(
		int Id,
		string Title,
		[property:JsonPropertyName("question")] string Instruction,
		[property:JsonPropertyName("templates")] ImmutableDictionary<string, string> TemplateByLanguage
	);
}
