﻿using System;
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
		private readonly ResultCache<Locale, IReadOnlyCollection<Question>> _questionsByLocaleCache;
		private static readonly TimeSpan CACHE_TTL = TimeSpan.FromHours(1);

		public QuestionServices(
			IQuestionRepository questionRepository,
			IMemoryCache memoryCache
		) {
			_questionRepository = questionRepository;
			ResultCache.Initialize(out _questionsByLocaleCache, memoryCache);
		}

		public async Task<ImmutableArray<Question>> GetAllAsync(Locale locale, CancellationToken cancellationToken) {
			return (ImmutableArray<Question>)await _questionsByLocaleCache.GetOrCreateAsync(locale, async () => {
				return await _questionRepository.GetAllAsync(locale, cancellationToken);
			}, absoluteExpirationRelativeToNow: CACHE_TTL);
		}
	}
}
