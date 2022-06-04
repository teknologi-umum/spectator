using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Linq;
using System.Threading;
using System.Threading.Tasks;
using Microsoft.Extensions.Options;
using Spectator.DomainModels.SubmissionDomain;
using Spectator.Piston.Internals;
using Spectator.Primitives;
using Spectator.Protos.Rce;
using Grpc.Net.Client;

namespace Spectator.Piston {
	public class PistonClient {
		private static ImmutableList<Runtime>? _runtimes;
		private static SemaphoreSlim? _semaphore;

		private readonly GrpcChannel _grpcChannel;
		private readonly CodeExecutionEngineService.CodeExecutionEngineServiceClient _rceClient;
		private readonly PistonOptions _pistonOptions;

		public PistonClient(
			IOptions<PistonOptions> pistonOptionsAccessor
		) {
			_pistonOptions = pistonOptionsAccessor.Value;
			_semaphore ??= new SemaphoreSlim(_pistonOptions.MaxConcurrentExecutions, _pistonOptions.MaxConcurrentExecutions);
			_grpcChannel = GrpcChannel.ForAddress(_pistonOptions.Address);
			_rceClient = new(_grpcChannel);
		}

		private async Task<Runtime?> GetRuntimeAsync(string language, CancellationToken cancellationToken) {
			if (_runtimes is null) {
				var runtimes = await _rceClient.ListRuntimesAsync(new EmptyRequest(), cancellationToken: cancellationToken);
				_runtimes = runtimes.Runtime.ToImmutableList();
			}
			return _runtimes!
				.Where(runtime => runtime.Language == language)
				.OrderByDescending(runtime => runtime.Version)
				.FirstOrDefault();
		}

		public async Task<ImmutableArray<TestResultBase>> ExecuteTestsAsync(Language language, string testCode, CancellationToken cancellationToken) {
			var executeResult = await ExecuteAsync(
				language: language switch {
					Language.C => "C",
					Language.CPP => "C++",
					Language.PHP => "PHP",
					Language.Javascript => "Javascript",
					Language.Java => "Java",
					Language.Python => "Python",
					_ => throw new InvalidProgramException("Unhandled language")
				},
				version: language switch {
					Language.C => "9.3.0",
					Language.CPP => "9.3.0",
					Language.PHP => "8.1",
					Language.Javascript => "16.15.0",
					Language.Java => "11",
					Language.Python => "3.10.2",
					_ => throw new InvalidProgramException("Unhandled language")

				},
				code: testCode,
				cancellationToken: cancellationToken
			);

			if (executeResult.Compile.ExitCode != 0) {
				return ImmutableArray.Create<TestResultBase>(
					new CompileErrorResult(executeResult.Compile.Output_)
				);
			}

			// TODO: report runtime error together with passing and failing tests
			if (executeResult.Runtime.ExitCode != 0) {
				return ImmutableArray.Create<TestResultBase>(
					new RuntimeErrorResult(executeResult.Runtime.Output_)
				);
			}

			return ResultParser.ParseTestResults(executeResult.Runtime.Stdout);
		}

		// HACK: Hard coded check for first question
		public async Task<ImmutableArray<TestResultBase>> ExecuteTwinkleTwinkleLittleStarTestAsync(Language language, string testCode, CancellationToken cancellationToken) {
			var executeResult = await ExecuteAsync(
				language: language switch {
					Language.C => "C",
					Language.CPP => "C++",
					Language.PHP => "PHP",
					Language.Javascript => "Javascript",
					Language.Java => "Java",
					Language.Python => "Python",
					_ => throw new InvalidProgramException("Unhandled language")
				},
				version: language switch {
					Language.C => "9.3.0",
					Language.CPP => "9.3.0",
					Language.PHP => "8.1",
					Language.Javascript => "16.15.0",
					Language.Java => "11",
					Language.Python => "3.10.2",
					_ => throw new InvalidProgramException("Unhandled language")

				},
				code: testCode,
				cancellationToken: cancellationToken
			);

			if (executeResult.Compile.ExitCode != 0) {
				return ImmutableArray.Create<TestResultBase>(
					new CompileErrorResult(executeResult.Compile.Output_)
				);
			}

			// TODO: report runtime error together with passing and failing tests
			if (executeResult.Runtime.ExitCode != 0) {
				return ImmutableArray.Create<TestResultBase>(
					new RuntimeErrorResult(executeResult.Runtime.Output_)
				);
			}

			if (!IsTwinkleTwinkleLittleStarLyrics(executeResult.Runtime.Stdout)) {
				return ImmutableArray.Create<TestResultBase>(
					new FailingTestResult(
						TestNumber: 1,
						ExpectedStdout: "Twinkle twinkle little star\nHow I wonder what you are\nUp above the world so high\nLike a diamond in the sky\nTwinkle twinkle little star\nHow I wonder what you are",
						ActualStdout: executeResult.Runtime.Stdout
					)
				);
			}

			return ImmutableArray.Create<TestResultBase>(
				new PassingTestResult(
					TestNumber: 1
				)
			);
		}

		internal async Task<CodeResponse> ExecuteAsync(string language, string version, string code, CancellationToken cancellationToken) {
			await _semaphore!.WaitAsync(cancellationToken);
			try {
				// HACK: java is expensive, they need more RAM than anything else
				var memoryLimit = language.ToLower() switch {
					"java" => 512 * 1024 * 1024,
					"javascript" => 512 * 1024 * 1024,
					_ => _pistonOptions.MemoryLimit,
				};

				var runtime = await GetRuntimeAsync(language, cancellationToken) ?? throw new KeyNotFoundException($"Runtime for {language} not found.");
				return await _rceClient.ExecuteAsync(
					new CodeRequest {
						Code = code,
						Version = version,
						Language = language,
						CompileTimeout = _pistonOptions.CompileTimeout,
						RunTimeout = _pistonOptions.RunTimeout,
						MemoryLimit = memoryLimit,
					},
					cancellationToken: cancellationToken
				);
			} finally {
				_semaphore!.Release();
			}
		}

		// HACK: Hard coded check for first question
		private static bool IsTwinkleTwinkleLittleStarLyrics(string text) {
			var lines = text.Split(new[] { '\r', '\n' }, StringSplitOptions.RemoveEmptyEntries);
			if (lines.Length != 6) return false;
			if (lines[0] != "Twinkle twinkle little star") return false;
			if (lines[1] != "How I wonder what you are") return false;
			if (lines[2] != "Up above the world so high") return false;
			if (lines[3] != "Like a diamond in the sky") return false;
			if (lines[4] != "Twinkle twinkle little star") return false;
			if (lines[5] != "How I wonder what you are") return false;
			return true;
		}
	}
}
