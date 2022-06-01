using System;
using System.Text.Json;
using System.Text.Json.Serialization;

namespace Spectator.DomainModels.SubmissionDomain {
	public abstract record TestResultBase(
		int TestNumber
	) {
		public static readonly JsonSerializerOptions JSON_SERIALIZER_OPTIONS = new() {
			Converters = {
				TestResultJsonConverter.INSTANCE
			},
			PropertyNamingPolicy = JsonNamingPolicy.CamelCase
		};

		private class TestResultJsonConverter : JsonConverter<TestResultBase> {
			public static readonly TestResultJsonConverter INSTANCE = new();

			public override TestResultBase? Read(ref Utf8JsonReader reader, Type typeToConvert, JsonSerializerOptions options) {
				// It has to start with a {
				if (reader.TokenType != JsonTokenType.StartObject) throw new JsonException("Expected start of object");

				// First property has to be "status"
				if (!reader.Read()
					|| reader.TokenType != JsonTokenType.PropertyName
					|| reader.GetString() != "status") {
					throw new JsonException("Expected first property to be 'status'");
				}

				// Read test status
				if (!reader.Read()
					|| reader.TokenType != JsonTokenType.String
					|| reader.GetString() is not string testStatus) {
					throw new JsonException("Expected test status");
				}

				// Second property has to be "result"
				if (!reader.Read()
					|| reader.TokenType != JsonTokenType.PropertyName
					|| reader.GetString() != "result") {
					throw new JsonException("Expected second property to be 'result'");
				}

				// Read test result
				TestResultBase? testResult = testStatus switch {
					"Passing" => JsonSerializer.Deserialize<PassingTestResult>(ref reader),
					"Failing" => JsonSerializer.Deserialize<FailingTestResult>(ref reader),
					"CompileError" => JsonSerializer.Deserialize<CompileErrorResult>(ref reader),
					"RuntimeError" => JsonSerializer.Deserialize<RuntimeErrorResult>(ref reader),
					_ => throw new JsonException($"Unknown test status {testStatus}")
				};

				// It has to end with a }
				if (!reader.Read()
					|| reader.TokenType != JsonTokenType.EndObject) {
					throw new JsonException("Expected end of object");
				}

				return testResult;
			}

			public override void Write(Utf8JsonWriter writer, TestResultBase value, JsonSerializerOptions options) {
				// Write {
				writer.WriteStartObject();

				// Write status property
				writer.WritePropertyName("status");
				writer.WriteStringValue(value switch {
					PassingTestResult => "Passing",
					FailingTestResult => "Failing",
					CompileErrorResult => "CompileError",
					RuntimeErrorResult => "RuntimeError",
					_ => throw new JsonException($"Unsupported type {value.GetType().Name}")
				});

				// Write result property
				writer.WritePropertyName("result");
				JsonSerializer.Serialize(writer, value);

				// Write }
				writer.WriteEndObject();
			}
		}
	}
}
