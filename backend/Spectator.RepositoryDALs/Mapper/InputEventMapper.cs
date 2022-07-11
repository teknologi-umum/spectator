using System.Collections.Immutable;
using System.Globalization;
using System.Reflection;
using InfluxDB.Client.Api.Domain;
using InfluxDB.Client.Core.Flux.Domain;
using InfluxDB.Client.Writes;
using RG.Ninja;
using Spectator.DomainEvents.InputDomain;
using Spectator.Primitives;

namespace Spectator.RepositoryDALs.Mapper {
	internal record InputEventMapper(
		Type Type,
		string FluxTypeName,
		ConstructorInfo Constructor,
		ImmutableList<FluxPropertyInfo> FluxProperties
	) {
		private static readonly ImmutableDictionary<string, Type> EVENT_TYPE_BY_MEASUREMENT = (
			from eventType in Assembly.GetAssembly(typeof(InputEventBase))!.GetTypes()
			where eventType.FullName!.StartsWith(typeof(InputEventBase).Namespace!, StringComparison.Ordinal)
			select new {
				Measurement = ToSnakeCase(eventType.Name.EndsWith("Event", out var pascalCaseName) ? pascalCaseName : eventType.Name),
				EventType = eventType
			}
		).ToImmutableDictionary(
			keySelector: e => e.Measurement,
			elementSelector: e => e.EventType
		);

		public static InputEventMapper For(Type eventType) {
			if (eventType is { IsClass: false } or { IsAbstract: true }) {
				throw new InvalidOperationException("Event type must be a concrete class");
			}

			if (!typeof(InputEventBase).IsAssignableFrom(eventType)) {
				throw new InvalidOperationException("Event type must be derived from InputEventBase");
			}

			var typeName = eventType.Name;
			if (!typeName.EndsWith("Event", out var pascalCaseName)) {
				throw new InvalidOperationException("Event type must end with 'Event'");
			}

			return new(
				Type: eventType,
				FluxTypeName: ToSnakeCase(pascalCaseName),
				Constructor: eventType
					.GetConstructors()
					.Single(),
				FluxProperties: eventType
					.GetProperties(BindingFlags.Public | BindingFlags.NonPublic | BindingFlags.DeclaredOnly | BindingFlags.Instance)
					.Where(prop => prop.Name is not nameof(InputEventBase.SessionId) and not nameof(InputEventBase.Timestamp) and not "EqualityContract")
					.Select(prop => new FluxPropertyInfo(
						FluxFieldName: ToSnakeCase(prop.Name),
						PropertyInfo: prop
					))
					.ToImmutableList()
			);
		}

		public static Type EventTypeFor(FluxRecord fluxRecord) {
			if (!EVENT_TYPE_BY_MEASUREMENT.TryGetValue(fluxRecord.GetMeasurement(), out var eventType)) {
				throw new KeyNotFoundException();
			}
			return eventType;
		}

		public InputEventBase ConvertToEntity(FluxRecord fluxRecord) {
			var measurement = fluxRecord.GetMeasurement();
			if (measurement != FluxTypeName) {
				throw new ArgumentException($"fluxRecord doesn't contain {FluxTypeName} measurement; instead, it contains {measurement} measurement", nameof(fluxRecord));
			}

			var parameters = Constructor.GetParameters();
			var arguments = new object[parameters.Length];

			arguments[0] = Guid.Parse((string)fluxRecord.GetValueByKey("session_id"));
			arguments[1] = new DateTimeOffset(fluxRecord.GetTimeInDateTime()!.Value, TimeSpan.Zero);

			for (var i = 2; i < parameters.Length; i++) {
				if (FluxProperties.SingleOrDefault(fluxProp => fluxProp.PropertyInfo.Name == parameters[i].Name) is not FluxPropertyInfo fluxProp) {
					throw new InvalidProgramException("This exception should never be thrown");
				}

				var parameterType = parameters[i].ParameterType;
				if (parameterType == typeof(string)) {
					arguments[i] = (string)fluxRecord.GetValueByKey(fluxProp.FluxFieldName);
				} else if (parameterType == typeof(int)) {
					arguments[i] = Convert.ToInt32(fluxRecord.GetValueByKey(fluxProp.FluxFieldName), CultureInfo.InvariantCulture);
				} else if (parameterType == typeof(bool)) {
					arguments[i] = Convert.ToBoolean(fluxRecord.GetValueByKey(fluxProp.FluxFieldName), CultureInfo.InvariantCulture);
				} else if (parameterType == typeof(DateTimeOffset)) {
					arguments[i] = DateTimeOffset.FromUnixTimeMilliseconds(Convert.ToInt64(fluxRecord.GetValueByKey(fluxProp.FluxFieldName), CultureInfo.InvariantCulture) / 1_000_000);
				} else if (parameterType == typeof(MouseButton)) {
					arguments[i] = Enum.Parse<MouseButton>((string)fluxRecord.GetValueByKey(fluxProp.FluxFieldName));
				} else if (parameterType == typeof(MouseDirection)) {
					arguments[i] = Enum.Parse<MouseDirection>((string)fluxRecord.GetValueByKey(fluxProp.FluxFieldName));
				} else if (parameterType == typeof(Locale)) {
					arguments[i] = Enum.Parse<Locale>((string)fluxRecord.GetValueByKey(fluxProp.FluxFieldName), ignoreCase: true);
				} else if (parameterType == typeof(Language)) {
					arguments[i] = Enum.Parse<Language>((string)fluxRecord.GetValueByKey(fluxProp.FluxFieldName), ignoreCase: true);
				} else if (parameterType == typeof(SelfAssessmentManikin)) {
					arguments[i] = new SelfAssessmentManikin(
						ArousedLevel: Convert.ToInt32(fluxRecord.GetValueByKey("aroused_level"), CultureInfo.InvariantCulture),
						PleasedLevel: Convert.ToInt32(fluxRecord.GetValueByKey("pleased_level"), CultureInfo.InvariantCulture)
					);
				} else if (parameterType == typeof(ConsoleKeyInfo)) {
					arguments[i] = new ConsoleKeyInfo(
						keyChar: (char)fluxRecord.GetValueByKey("key_char"),
						key: (ConsoleKey)fluxRecord.GetValueByKey("key_code"),
						shift: (bool)fluxRecord.GetValueByKey("shift"),
						alt: (bool)fluxRecord.GetValueByKey("alt"),
						control: (bool)fluxRecord.GetValueByKey("ctrl")
					);
				} else if (parameterType == typeof(ImmutableArray<int>)) {
					arguments[i] = ((string)fluxRecord.GetValueByKey(fluxProp.FluxFieldName)).Split(',').Select(s => int.Parse(s, CultureInfo.InvariantCulture)).ToImmutableArray();
				} else {
					throw new InvalidProgramException($"Unhandled parameter type {parameterType}");
				}
			}

			return (InputEventBase)Constructor.Invoke(arguments);
		}

		public PointData ConvertToPointData(InputEventBase @event) {
			if (@event.GetType() != Type) {
				throw new InvalidOperationException($"This mapper only converts {Type.Name} and cannot be used to convert {@event.GetType().Name}");
			}

			var pointData = PointData
				.Measurement(FluxTypeName)
				.Tag("session_id", @event.SessionId.ToString())
				.Timestamp(@event.Timestamp, WritePrecision.Ns);

			foreach (var fluxProperty in FluxProperties) {
				var value = fluxProperty.PropertyInfo.GetValue(@event);
				pointData = value switch {
					null => throw new ArgumentException($"event.{fluxProperty.PropertyInfo.Name} is null", nameof(@event)),
					string s => pointData.Field(fluxProperty.FluxFieldName, s),
					int i => pointData.Field(fluxProperty.FluxFieldName, i),
					bool b => pointData.Field(fluxProperty.FluxFieldName, b),
					DateTimeOffset dto => pointData.Field(fluxProperty.FluxFieldName, dto.ToUnixTimeMilliseconds() * 1_000_000),
					MouseButton mb => pointData.Field(fluxProperty.FluxFieldName, mb.ToString()),
					MouseDirection md => pointData.Field(fluxProperty.FluxFieldName, md.ToString()),
					Locale l => pointData.Field(fluxProperty.FluxFieldName, l.ToString().ToUpperInvariant()),
					Language l => pointData.Field(fluxProperty.FluxFieldName, l.ToString().ToLowerInvariant()),
					SelfAssessmentManikin sam => pointData
						.Field("aroused_level", sam.ArousedLevel)
						.Field("pleased_level", sam.PleasedLevel),
					ConsoleKeyInfo cki => pointData
						.Field("key_char", cki.KeyChar)
						.Field("key_code", (int)cki.Key)
						.Field("shift", cki.Modifiers.HasFlag(ConsoleModifiers.Shift))
						.Field("alt", cki.Modifiers.HasFlag(ConsoleModifiers.Alt))
						.Field("ctrl", cki.Modifiers.HasFlag(ConsoleModifiers.Control)),
					ImmutableArray<int> a => pointData.Field(fluxProperty.FluxFieldName, string.Join(',', a.Select(i => i.ToString(CultureInfo.InvariantCulture)))),
					_ => throw new InvalidProgramException($"Unhandled property type {value.GetType()}")
				};
			}

			return pointData;
		}

		private static bool IsInPascalCase(string identifierName) {
			return identifierName.Length > 0
				&& char.IsUpper(identifierName[0])
				&& identifierName.Skip(1).All(c => char.IsLetterOrDigit(c));
		}

		private static string ToSnakeCase(string pascalCaseName) {
			if (!IsInPascalCase(pascalCaseName)) {
				throw new ArgumentException("Name is not in pascal case", nameof(pascalCaseName));
			}

			// Known abbreviations
			pascalCaseName = pascalCaseName.Replace("SAM", "Sam", StringComparison.Ordinal);
			pascalCaseName = pascalCaseName.Replace("IDE", "Ide", StringComparison.Ordinal);

			return string.Concat(
				pascalCaseName
					.Select((c, i) => char.IsUpper(c)
						? i > 0
							? $"_{char.ToLowerInvariant(c)}"
							: char.ToLowerInvariant(c).ToString()
						: c.ToString()
					)
			);
		}
	}
}
