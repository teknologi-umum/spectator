using System.Collections.Immutable;
using System.Globalization;
using System.Reflection;
using System.Text.Json;
using Spectator.DomainModels.QuestionDomain;
using Spectator.Primitives;
using Spectator.Repositories;

namespace Spectator.RepositoryDALs.Internals {
	internal class QuestionRepositoryDAL : IQuestionRepository {
		private static readonly Assembly ASSEMBLY = Assembly.GetAssembly(typeof(QuestionRepositoryDAL))!;
		private static readonly JsonSerializerOptions JSON_SERIALIZER_OPTIONS = new JsonSerializerOptions {
			PropertyNamingPolicy = JsonNamingPolicy.CamelCase
		};

		public async Task<ImmutableDictionary<Locale, ImmutableArray<Question>>> GetAllAsync(CancellationToken cancellationToken) {
			// Get all assertions
			var builder = ImmutableDictionary<int, ImmutableDictionary<Language, string>>.Empty.ToBuilder();
			foreach (var g in
				from resourceName in ASSEMBLY.GetManifestResourceNames()
				where resourceName.StartsWith("Spectator.RepositoryDALs.Imported.AssertionData.", StringComparison.Ordinal)
				let names = resourceName.Split('.')
				let language = Enum.Parse<Language>(names[4], ignoreCase: true)
				let questionNumber = int.Parse(names[5][8..], CultureInfo.InvariantCulture)
				group (resourceName, language) by questionNumber into g
				select g
			) {
				var innerBuilder = ImmutableDictionary<Language, string>.Empty.ToBuilder();
				foreach ((var resourceName, var language) in g) {
					using var stream = ASSEMBLY.GetManifestResourceStream(resourceName)!;
					using var reader = new StreamReader(stream);
#pragma warning disable RG0001 // Do not await inside a loop
					var assertion = await reader.ReadToEndAsync();
#pragma warning restore RG0001 // Do not await inside a loop
					innerBuilder.Add(language, assertion);
				}
				builder.Add(g.Key, innerBuilder.ToImmutable());
			}
			var assertionByLanguageByQuestionNumber = builder.ToImmutable();

			// Get all questions
			var questionsByLocale = new Dictionary<Locale, ImmutableDictionary<int, JsonModels.Question>>();
			foreach ((var resourceName, var locale) in
				from resourceName in ASSEMBLY.GetManifestResourceNames()
				where resourceName.StartsWith("Spectator.RepositoryDALs.Imported.FrontendData.", StringComparison.Ordinal)
				let names = resourceName.Split('.')
				where names[5] == "questions"
				let locale = Enum.Parse<Locale>(names[4], ignoreCase: true)
				select (resourceName, locale)
			) {
				using var stream = ASSEMBLY.GetManifestResourceStream(resourceName)!;
				using var reader = new StreamReader(stream);
				var json = await reader.ReadToEndAsync();
				var questionSet = JsonSerializer.Deserialize<JsonModels.QuestionSet>(json, JSON_SERIALIZER_OPTIONS)!;
				questionsByLocale.Add(locale, questionSet.Questions.ToImmutableDictionary(q => q.Id));
			}

			// Validate and get array of question numbers
			if (questionsByLocale.Count == 0) throw new InvalidProgramException("Question database is empty");
			var questionCount = questionsByLocale.Values.First().Count;
			if (questionsByLocale.Values.Any(qs => qs.Count != questionCount)) throw new InvalidProgramException("Every questions.json must contain exactly same number of questions");
			var questionNumbers = questionsByLocale.Values.First().Keys;
			if (questionNumbers.Any(qn => questionsByLocale.Values.Any(qs => !qs.ContainsKey(qn)))) throw new InvalidProgramException("Every questions.json must contain questions with the exact same set of question numbers");

			// Map results
			return questionsByLocale.ToImmutableDictionary(
				keySelector: kvp => kvp.Key,
				elementSelector: kvp => questionNumbers.Select(questionNumber => new Question(
					QuestionNumber: questionNumber,
					TitleByLocale: questionsByLocale.ToImmutableDictionary(
						keySelector: kvp => kvp.Key,
						elementSelector: kvp => kvp.Value[questionNumber].Title
					),
					InstructionByLocale: questionsByLocale.ToImmutableDictionary(
						keySelector: kvp => kvp.Key,
						elementSelector: kvp => kvp.Value[questionNumber].Instruction
					),
					TemplateByLanguageByLocale: questionsByLocale.ToImmutableDictionary(
						keySelector: kvp => kvp.Key,
						elementSelector: kvp => kvp.Value[questionNumber].TemplateByLanguage.ToImmutableDictionary(
							keySelector: kvp => Enum.Parse<Language>(kvp.Key, ignoreCase: true),
							elementSelector: kvp => kvp.Value
						)
					),
					AssertionByLanguage: assertionByLanguageByQuestionNumber[questionNumber]
				)).ToImmutableArray()
			);
		}
	}
}
