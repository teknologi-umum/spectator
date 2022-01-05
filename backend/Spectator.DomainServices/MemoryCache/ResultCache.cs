using System;
using System.Collections;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Diagnostics;
using System.Diagnostics.CodeAnalysis;
using System.Threading.Tasks;
using Microsoft.Extensions.Caching.Memory;

namespace Spectator.DomainServices.MemoryCache {
	[DebuggerNonUserCode]
	internal class ResultCache<TKey, TValue> where TKey : notnull where TValue : class {
		private readonly IMemoryCache _memoryCache;
		private readonly string _prefix;

		public ResultCache(
			IMemoryCache memoryCache,
			byte[] identifier
		) {
			_memoryCache = memoryCache;
			_prefix = Convert.ToBase64String(identifier);
		}

		private string SerializeKey(TKey key) {
			return _prefix + CacheIdentifierBuilder.SerializeKey(key);
		}

		public bool TryGetValue(TKey key, [NotNullWhen(true)] out TValue? value) {
			if (_memoryCache.TryGetValue($"{_prefix}{key}", out TValue cachedValue)) {
				value = cachedValue;
				return true;
			} else {
				value = null;
				return false;
			}
		}

		public bool TryGetNullableValue(TKey key, out TValue? value) {
			return _memoryCache.TryGetValue($"{_prefix}{key}", out value);
		}

		public Task<TValue> GetOrCreateAsync(
			TKey key,
			Func<Task<TValue>> valueFactory,
			TimeSpan? absoluteExpirationRelativeToNow = null,
			TimeSpan? slidingExpiration = null
		) {
			return _memoryCache.GetOrCreateAsync(SerializeKey(key), async cacheEntry => {
				var value = await valueFactory.Invoke().ConfigureAwait(false);
				cacheEntry.AbsoluteExpirationRelativeToNow = absoluteExpirationRelativeToNow;
				cacheEntry.SlidingExpiration = slidingExpiration;
				cacheEntry.Value = value;
				return value;
			});
		}

		public Task<TValue?> GetOrCreateNullableValueAsync(
			TKey key,
			Func<Task<TValue?>> valueFactory,
			TimeSpan? absoluteExpirationRelativeToNow = null,
			TimeSpan? slidingExpiration = null
		) {
			return _memoryCache.GetOrCreateAsync(SerializeKey(key), async cacheEntry => {
				var value = await valueFactory.Invoke().ConfigureAwait(false);
				cacheEntry.AbsoluteExpirationRelativeToNow = absoluteExpirationRelativeToNow;
				cacheEntry.SlidingExpiration = slidingExpiration;
				cacheEntry.Value = value;
				return value;
			});
		}

		public async Task<ImmutableDictionary<TKey, TValue?>> GetOrCreateManyAsync(
			ImmutableHashSet<TKey> keys,
			Func<ImmutableHashSet<TKey>, Task<ImmutableDictionary<TKey, TValue?>>> missingValuesFactory,
			TimeSpan? absoluteExpirationRelativeToNow = null,
			TimeSpan? slidingExpiration = null
		) {
			Dictionary<TKey, TValue?> valueDictionary = new();
			foreach (var key in keys) {
				if (_memoryCache.TryGetValue(SerializeKey(key), out TValue? value)) {
					valueDictionary.Add(key, value);
				}
			}
			if (valueDictionary.Count < keys.Count) {
				var missingValueKeys = keys.Except(valueDictionary.Keys);
				var missingValueDictionary = await missingValuesFactory.Invoke(missingValueKeys).ConfigureAwait(false);
				Debug.Assert(missingValueKeys.Count == missingValueDictionary.Count, "MissingValuesFactory didn't return same number of elements as number of missing value keys.");
				foreach ((var key, var value) in missingValueDictionary) {
					valueDictionary.Add(key, value);
					_memoryCache.Set(SerializeKey(key), value, new MemoryCacheEntryOptions {
						AbsoluteExpirationRelativeToNow = absoluteExpirationRelativeToNow,
						SlidingExpiration = slidingExpiration
					});
				}
			}
			return valueDictionary.ToImmutableDictionary();
		}

		[return: NotNullIfNotNull("value")]
		public TValue? Set(
			TKey key,
			TValue? value,
			TimeSpan? absoluteExpirationRelativeToNow = null,
			TimeSpan? slidingExpiration = null
		) {
			return _memoryCache.Set(SerializeKey(key), value, new MemoryCacheEntryOptions {
				AbsoluteExpirationRelativeToNow = absoluteExpirationRelativeToNow,
				SlidingExpiration = slidingExpiration
			});
		}

		[return: NotNullIfNotNull("value")]
		public TActualValue? Set<TActualValue>(
			TKey key,
			TActualValue? value,
			TimeSpan? absoluteExpirationRelativeToNow = null,
			TimeSpan? slidingExpiration = null
		) where TActualValue : class, TValue {
			return _memoryCache.Set(SerializeKey(key), value, new MemoryCacheEntryOptions {
				AbsoluteExpirationRelativeToNow = absoluteExpirationRelativeToNow,
				SlidingExpiration = slidingExpiration
			});
		}

		public void Remove(TKey key) {
			_memoryCache.Remove(SerializeKey(key));
		}
	}

	[DebuggerNonUserCode]
	public class ResultCache<TValue> where TValue : class {
		private readonly IMemoryCache _memoryCache;
		private readonly string _key;

		public ResultCache(
			IMemoryCache memoryCache,
			byte[] identifier
		) {
			_memoryCache = memoryCache;
			_key = Convert.ToBase64String(identifier);
		}

		public Task<TValue> GetOrCreateAsync(
			Func<Task<TValue>> valueFactory,
			TimeSpan? absoluteExpirationRelativeToNow = null,
			TimeSpan? slidingExpiration = null
		) {
			return _memoryCache.GetOrCreateAsync(_key, async cacheEntry => {
				var value = await valueFactory.Invoke().ConfigureAwait(false);
				cacheEntry.AbsoluteExpirationRelativeToNow = absoluteExpirationRelativeToNow;
				cacheEntry.SlidingExpiration = slidingExpiration;
				cacheEntry.Value = value;
				return value;
			});
		}

		public Task<TValue?> GetOrCreateNullableValueAsync(
			Func<Task<TValue?>> valueFactory,
			TimeSpan? absoluteExpirationRelativeToNow = null,
			TimeSpan? slidingExpiration = null
		) {
			return _memoryCache.GetOrCreateAsync(_key, async cacheEntry => {
				var value = await valueFactory.Invoke().ConfigureAwait(false);
				cacheEntry.AbsoluteExpirationRelativeToNow = absoluteExpirationRelativeToNow;
				cacheEntry.SlidingExpiration = slidingExpiration;
				cacheEntry.Value = value;
				return value;
			});
		}

		[return: NotNullIfNotNull("value")]
		public TValue? Set(
			TValue? value,
			TimeSpan? absoluteExpirationRelativeToNow = null,
			TimeSpan? slidingExpiration = null
		) {
			return _memoryCache.Set(_key, value, new MemoryCacheEntryOptions {
				AbsoluteExpirationRelativeToNow = absoluteExpirationRelativeToNow,
				SlidingExpiration = slidingExpiration
			});
		}

		public void Remove() {
			_memoryCache.Remove(_key);
		}
	}

	internal static class ResultCache {
		private static readonly Dictionary<byte[], object> CACHE_DICTIONARY = new(new ByteArrayComparer());
		private static readonly object GATE = new();

		public static void Initialize<TValue>(out ResultCache<TValue> resultCache, IMemoryCache memoryCache, params string[] scopes) where TValue : class {
			var identifier = CacheIdentifierBuilder.ForResultCache<TValue>()
				.AppendScopes(scopes)
				.ToBytes();
			lock (GATE) {
				if (CACHE_DICTIONARY.TryGetValue(identifier, out var cache)) {
					resultCache = (ResultCache<TValue>)cache;
				} else {
					cache = new ResultCache<TValue>(memoryCache, identifier);
					CACHE_DICTIONARY.Add(identifier, cache);
					resultCache = (ResultCache<TValue>)cache;
				}
			}
		}

		public static void Initialize<TKey, TValue>(out ResultCache<TKey, TValue> resultCache, IMemoryCache memoryCache, params string[] scopes) where TKey : notnull where TValue : class {
			var identifier = CacheIdentifierBuilder.ForResultCache<TKey, TValue>()
				.AppendScopes(scopes)
				.ToBytes();
			lock (GATE) {
				if (CACHE_DICTIONARY.TryGetValue(identifier, out var cache)) {
					resultCache = (ResultCache<TKey, TValue>)cache;
				} else {
					cache = new ResultCache<TKey, TValue>(memoryCache, identifier);
					CACHE_DICTIONARY.Add(identifier, cache);
					resultCache = (ResultCache<TKey, TValue>)cache;
				}
			}
		}

		private class ByteArrayComparer : IEqualityComparer<byte[]> {
			public bool Equals(byte[]? x, byte[]? y) => StructuralComparisons.StructuralEqualityComparer.Equals(x, y);
			public int GetHashCode([DisallowNull] byte[] obj) => StructuralComparisons.StructuralEqualityComparer.GetHashCode(obj);
		}
	}
}
