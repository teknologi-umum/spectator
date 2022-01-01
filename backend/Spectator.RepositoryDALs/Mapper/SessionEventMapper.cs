using System.Collections.Immutable;
using System.Reflection;
using InfluxDB.Client.Api.Domain;
using InfluxDB.Client.Core.Flux.Domain;
using InfluxDB.Client.Writes;
using RG.Ninja;
using Spectator.DomainEvents.SessionDomain;
using Spectator.Primitives;

namespace Spectator.RepositoryDALs.Mapper {
	internal record SessionEventMapper(
		Type Type,
		string FluxTypeName,
		ConstructorInfo Constructor,
		ImmutableList<FluxPropertyInfo> FluxProperties
	) {
		public static SessionEventMapper For<T>() {
			if (typeof(T) is { IsClass: false } or { IsAbstract: true }) {
				throw new InvalidOperationException("Event type must be a concrete class");
			}

			if (!typeof(SessionEventBase).IsAssignableFrom(typeof(T))) {
				throw new InvalidOperationException("Event type must be derived from SessionEventBase");
			}

			var typeName = typeof(T).Name;
			if (!typeName.EndsWith("Event", out var pascalCaseName)) {
				throw new InvalidOperationException("Event type must end with 'Event'");
			}

			return new(
				Type: typeof(T),
				FluxTypeName: ToSnakeCase(pascalCaseName),
				Constructor: typeof(T)
					.GetConstructors()
					.Single(),
				FluxProperties: typeof(T)
					.GetProperties(BindingFlags.Public | BindingFlags.NonPublic | BindingFlags.DeclaredOnly | BindingFlags.Instance)
					.Where(prop => prop.Name is not nameof(SessionEventBase.SessionId) and not nameof(SessionEventBase.Timestamp))
					.Select(prop => new FluxPropertyInfo(
						FluxFieldName: ToSnakeCase(prop.Name),
						PropertyInfo: prop
					))
					.ToImmutableList()
			);
		}

		public T ConvertToEntity<T>(FluxRecord fluxRecord) {
			if (fluxRecord.GetMeasurement() != "event") {
				throw new InvalidOperationException("This mapper only converts event measurement");
			}

			var parameters = Constructor.GetParameters();
			var arguments = new object[parameters.Length];

			arguments[0] = fluxRecord.GetValueByKey("type");
			arguments[1] = Guid.Parse((string)fluxRecord.GetValueByKey("session_id"));

			for (var i = 2; i < parameters.Length; i++) {
				if (FluxProperties.SingleOrDefault(fluxProp => fluxProp.PropertyInfo.Name == parameters[i].Name) is not FluxPropertyInfo fluxProp) {
					throw new InvalidOperationException("This exception should never be thrown");
				}

				var parameterType = parameters[i].ParameterType;
				if (parameterType == typeof(string)
					|| parameterType == typeof(int)
					|| parameterType == typeof(DateTimeOffset)) {
					arguments[i] = fluxRecord.GetValueByKey(fluxProp.FluxFieldName);
				} else if (parameterType == typeof(MouseButton)) {
					arguments[i] = Enum.Parse<MouseButton>((string)fluxRecord.GetValueByKey(fluxProp.FluxFieldName));
				} else if (parameterType == typeof(SelfAssessmentManikin)) {
					arguments[i] = new SelfAssessmentManikin(
						ArousedLevel: (int)fluxRecord.GetValueByKey("aroused_level"),
						PleasedLevel: (int)fluxRecord.GetValueByKey("pleased_level")
					);
				} else if (parameterType == typeof(ConsoleKeyInfo)) {
					arguments[i] = new ConsoleKeyInfo(
						keyChar: (char)fluxRecord.GetValueByKey("key_char"),
						key: (ConsoleKey)fluxRecord.GetValueByKey("key_code"),
						shift: (bool)fluxRecord.GetValueByKey("shift"),
						alt: (bool)fluxRecord.GetValueByKey("alt"),
						control: (bool)fluxRecord.GetValueByKey("ctrl")
					);
				} else {
					throw new InvalidProgramException($"Unhandled parameter type {parameterType}");
				}
			}

			return (T)Constructor.Invoke(arguments);
		}

		public PointData ConvertToPointData<T>(T @event) where T : SessionEventBase {
			if (typeof(T) != Type) {
				throw new InvalidOperationException($"This mapper only converts {Type.Name} and cannot be used to convert {@event.GetType().Name}");
			}

			var pointData = PointData
				.Measurement("event")
				.Tag("type", FluxTypeName)
				.Tag("session_id", @event.SessionId.ToString())
				.Timestamp(@event.Timestamp, WritePrecision.Ns);

			foreach (var fluxProperty in FluxProperties) {
				var value = fluxProperty.PropertyInfo.GetValue(@event);
				pointData = value switch {
					null => throw new ArgumentException($"event.{fluxProperty.PropertyInfo.Name} is null", nameof(@event)),
					string s => pointData.Field(fluxProperty.FluxFieldName, s),
					int i => pointData.Field(fluxProperty.FluxFieldName, i),
					DateTimeOffset dto => pointData.Field(fluxProperty.FluxFieldName, dto.ToUnixTimeMilliseconds() * 1_000_000),
					MouseButton mb => pointData.Field(fluxProperty.FluxFieldName, mb.ToString()),
					SelfAssessmentManikin sam => pointData
						.Field("aroused_level", sam.ArousedLevel)
						.Field("pleased_level", sam.PleasedLevel),
					ConsoleKeyInfo cki => pointData
						.Field("key_char", cki.KeyChar)
						.Field("key_code", (int)cki.Key)
						.Field("shift", cki.Modifiers.HasFlag(ConsoleModifiers.Shift))
						.Field("alt", cki.Modifiers.HasFlag(ConsoleModifiers.Alt))
						.Field("ctrl", cki.Modifiers.HasFlag(ConsoleModifiers.Control)),
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

			return string.Concat(
				pascalCaseName
					.Select((c, i) => char.IsUpper(c)
						? i > 0
							? $"_{char.ToLower(c)}"
							: char.ToLower(c).ToString()
						: c.ToString()
					)
			);
		}
	}
}
