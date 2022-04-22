using System;
using System.Collections.Generic;
using System.Globalization;
using System.Linq;
using System.Runtime.CompilerServices;
using System.Security.Cryptography;
using System.Text;

namespace Spectator.DomainServices.MemoryCache {
	internal class CacheIdentifierBuilder {
		private readonly MD5 _md5;
		private readonly byte[] _buffer;

		private CacheIdentifierBuilder(Type tValue) {
			_md5 = MD5.Create();
			_buffer = new byte[32];

			// [ 0000000000000000 TVALUEGUIDBYTESX ]
			Array.Copy(tValue.GUID.ToByteArray(), 0, _buffer, 16, 16);
		}

		/// <summary>
		/// Allowed key types:
		/// int,
		/// string,
		/// UrlSlug,
		/// Enum,
		/// Guid,
		/// Tuple (which only contains int, string, Enum, HashSet`int, or HashSet`string)
		/// </summary>
		private CacheIdentifierBuilder(Type tKey, Type tValue) {
			if (typeof(ITuple).IsAssignableFrom(tKey)) {
				foreach (var tTupleValue in tKey.GetGenericArguments()) {
					if (tTupleValue != typeof(string)
						&& tTupleValue != typeof(int)
						&& tTupleValue != typeof(bool)
						&& tTupleValue != typeof(HashSet<int>)
						&& tTupleValue != typeof(HashSet<string>)
						&& !tTupleValue.IsEnum) {
						throw new InvalidOperationException($"Invalid value type in cache key tuple: {tTupleValue}, cache key type: {tKey}");
					}
				}
			} else if (tKey != typeof(string)
				&& tKey != typeof(int)
				&& tKey != typeof(Guid)
				&& !tKey.IsEnum) {
				throw new InvalidOperationException($"Invalid cache key type: {tKey}");
			}

			_md5 = MD5.Create();
			_buffer = new byte[32];

			// [ TKEYGUIDBYTESXXX 0000000000000000 ]
			Array.Copy(tKey.GUID.ToByteArray(), _buffer, 16);

			// [ TKEYGUIDBYTESXXX TVALUEGUIDBYTESX ]
			Array.Copy(tValue.GUID.ToByteArray(), 0, _buffer, 16, 16);
		}

		/// <summary>
		/// Allowed key types:
		/// int,
		/// string,
		/// UrlSlug,
		/// Enum,
		/// Guid,
		/// Tuple (which only contains int, string, Enum, HashSet`int, or HashSet`string)
		/// </summary>
		public static string SerializeKey<TKey>(TKey key) where TKey : notnull {
			switch (key) {
				case ITuple tuple:
					var serializedValues = Enumerable.Range(0, tuple.Length).Select(pos => tuple[pos] switch {
						string s => $"\"{s}\"",
						int i => i.ToString(CultureInfo.InvariantCulture),
						bool b => b.ToString(CultureInfo.InvariantCulture),
						HashSet<int> h => $"[{string.Join(',', h.OrderBy(x => x))}]",
						HashSet<string> h => $"[{string.Join(',', h.OrderBy(x => x))}]",
						null => "(null)",
						object value => value.GetType().IsEnum
							? value.ToString() ?? "(null)"
							: throw new InvalidOperationException($"Invalid value in cache key tuple: {value}")
					}).ToList();
					return $"[{string.Join(',', serializedValues)}]";

				case Guid s: return s.ToString();
				case string s: return $"\"{s}\"";
				case int i: return i.ToString(CultureInfo.InvariantCulture);
				case null: return "(null)";
				default:
					if (key.GetType().IsEnum) {
						return key.ToString() ?? "(null)";
					} else {
						throw new InvalidOperationException($"Invalid key: {key}");
					}
			}
		}

		public static CacheIdentifierBuilder ForResultCache<TValue>() => new(typeof(TValue));
		public static CacheIdentifierBuilder ForResultCache<TKey, TValue>() => new(typeof(TKey), typeof(TValue));

		public CacheIdentifierBuilder AppendScope<TScope>() {
			var bufferHash = _md5.ComputeHash(_buffer);

			// [ BUFFERHASHBYTESX 0000000000000000 ]
			Array.Copy(_buffer, bufferHash, 16);

			// [ BUFFERHASHBYTESX TSCOPEGUIDBYTESX ]
			Array.Copy(typeof(TScope).GUID.ToByteArray(), 0, _buffer, 16, 16);

			return this;
		}

		public CacheIdentifierBuilder AppendScope(string scope) {
			var bufferHash = _md5.ComputeHash(_buffer);

			// [ BUFFERHASHBYTESX 0000000000000000 ]
			Array.Copy(_buffer, bufferHash, 16);

			var scopeHash = _md5.ComputeHash(Encoding.UTF8.GetBytes(scope));

			// [ BUFFERHASHBYTESX SCOPEHASHBYTESXX ]
			Array.Copy(scopeHash, 0, _buffer, 16, 16);

			return this;
		}

		public CacheIdentifierBuilder AppendScopes(params string[] scopes) {
			foreach (var scope in scopes) {
				AppendScope(scope);
			}
			return this;
		}

		public byte[] ToBytes() {
			_md5.Dispose();
			return _buffer;
		}

		public string ToBase64String() {
			_md5.Dispose();
			return Convert.ToBase64String(_buffer);
		}
	}
}
