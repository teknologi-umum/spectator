using System.Collections.Immutable;
using Spectator.DomainModels.QuestionDomain;
using Spectator.Primitives;

namespace Spectator.Repositories {
	public interface IQuestionRepository {
		Task<ImmutableDictionary<Locale, ImmutableArray<Question>>> GetAllAsync(CancellationToken cancellationToken);
	}
}
