using System.Collections.Immutable;

namespace Spectator.RepositoryDALs.JsonModels {
	internal record QuestionSet(
		ImmutableArray<Question> Questions
	);
}
