using System.Collections.Immutable;
using System.Reflection;
using System.Text.Json;
using Spectator.DomainModels.QuestionDomain;
using Spectator.Primitives;
using Spectator.Repositories;

namespace Spectator.RepositoryDALs.Internals {
	internal class QuestionRepositoryDAL : IQuestionRepository {
		public async Task<ImmutableArray<Question>> GetAllAsync(Locale locale, CancellationToken cancellationToken) {
			var resourceName = locale switch {
				Locale.EN => "Spectator.RepositoryDALs.Imported.FrontendData.en.questions.json",
				Locale.ID => "Spectator.RepositoryDALs.Imported.FrontendData.id.questions.json",
				_ => throw new ArgumentException("Invalid locale", nameof(locale))
			};
			using var stream = Assembly.GetAssembly(typeof(QuestionRepositoryDAL))!.GetManifestResourceStream(resourceName)!;
			using var reader = new StreamReader(stream);
			var json = await reader.ReadToEndAsync();
			var questionSet = JsonSerializer.Deserialize<JsonModels.QuestionSet>(json);
			return questionSet!.Questions.Select((q, i) => new Question(
				QuestionNumber: i + 1,
				Title: q.Title,
				Instruction: q.Instruction,
				TemplateByLanguage: q.TemplateByLanguage.ToImmutableDictionary(
					keySelector: kvp => Enum.Parse<Language>(kvp.Key, ignoreCase: true),
					elementSelector: kvp => kvp.Value
				)
			)).ToImmutableArray();
		}
	}
}
