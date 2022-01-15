using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading;
using System.Threading.Tasks;
using Microsoft.Extensions.Caching.Memory;
using Spectator.DomainModels.QuestionDomain;
using Spectator.DomainServices.MemoryCache;
using Spectator.Primitives;
using Spectator.Repositories;

namespace Spectator.DomainServices.QuestionDomain {
	public class QuestionServices {
		private readonly IQuestionRepository _questionRepository;
		private readonly ResultCache<IReadOnlyDictionary<Locale, ImmutableArray<Question>>> _cache;
		private static readonly TimeSpan CACHE_TTL = TimeSpan.FromHours(1);

		public QuestionServices(
			IQuestionRepository questionRepository,
			IMemoryCache memoryCache
		) {
			_questionRepository = questionRepository;
			ResultCache.Initialize(out _cache, memoryCache);
		}

		public async Task<ImmutableDictionary<Locale, ImmutableArray<Question>>> GetAllAsync(CancellationToken cancellationToken) {
			return (ImmutableDictionary<Locale, ImmutableArray<Question>>)await _cache.GetOrCreateAsync(async () => {
				return await _questionRepository.GetAllAsync(cancellationToken);
			}, absoluteExpirationRelativeToNow: CACHE_TTL);
		}
	}
}
