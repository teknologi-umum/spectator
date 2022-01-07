using System.Collections.Immutable;
using Spectator.DomainModels.QuestionDomain;
using Spectator.Primitives;

namespace Spectator.Repositories {
	public interface IQuestionRepository {
		Task<ImmutableArray<Question>> GetAllAsync(Locale locale, CancellationToken cancellationToken);
	}
}
