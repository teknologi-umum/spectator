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
using Microsoft.Extensions.Logging;
using Spectator.Piston.Exceptions;

namespace Spectator.Piston {
	public class PistonClient {
		private static ImmutableList<Runtime>? _runtimes;
		private static SemaphoreSlim? _semaphore;

		private readonly GrpcChannel _grpcChannel;
		private readonly CodeExecutionEngineService.CodeExecutionEngineServiceClient _rceClient;
		private readonly PistonOptions _pistonOptions;
		private readonly int _maxConcurrentExecutions;
		private readonly ILogger<PistonClient> _logger;

		public PistonClient(
			IOptions<PistonOptions> pistonOptionsAccessor,
			ILogger<PistonClient> logger
		) {
			_pistonOptions = pistonOptionsAccessor.Value;
			_maxConcurrentExecutions = _pistonOptions.MaxConcurrentExecutions;
			_semaphore ??= new SemaphoreSlim(_maxConcurrentExecutions, _maxConcurrentExecutions);
			_grpcChannel = GrpcChannel.ForAddress(_pistonOptions.Address);
			_rceClient = new(_grpcChannel);
			_logger = logger;
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
					Language.Java => "17",
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

			try {
				var stdout = executeResult.Runtime.Output_ ?? executeResult.Runtime.Stdout ?? "";
				return ResultParser.ParseTestResults(stdout, executeResult.Runtime.Stderr);
			} catch (CannotParseStdoutException e) {
				return ImmutableArray.Create<TestResultBase>(
					new InvalidInputResult(e.Stderr)
				);
			}
		}

		// HACK: Hard coded check for zeroth question
		public async Task<ImmutableArray<TestResultBase>> ExecuteHelloWorldTestAsync(Language language, string testCode, CancellationToken cancellationToken) {
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
					Language.Java => "17",
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

			if (executeResult.Runtime.Stdout.Trim() != "Hello world") {
				return ImmutableArray.Create<TestResultBase>(
					new FailingTestResult(
						TestNumber: 0,
						ExpectedStdout: "Hello world",
						ActualStdout: executeResult.Runtime.Stdout,
						ArgumentsStdout: ""
					)
				);
			}

			return ImmutableArray.Create<TestResultBase>(
				new PassingTestResult(
					TestNumber: 0,
					ExpectedStdout: "Hello world",
					ActualStdout: executeResult.Runtime.Stdout,
					ArgumentsStdout: ""
				)
			);
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
					Language.Java => "17",
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
						ActualStdout: executeResult.Runtime.Stdout,
						ArgumentsStdout: ""
					)
				);
			}

			return ImmutableArray.Create<TestResultBase>(
				new PassingTestResult(
					TestNumber: 1,
					ExpectedStdout: "Twinkle twinkle little star\nHow I wonder what you are\nUp above the world so high\nLike a diamond in the sky\nTwinkle twinkle little star\nHow I wonder what you are",
					ActualStdout: executeResult.Runtime.Stdout,
					ArgumentsStdout: ""
				)
			);
		}

		internal async Task<CodeResponse> ExecuteAsync(string language, string version, string code, CancellationToken cancellationToken) {
			await _semaphore!.WaitAsync(cancellationToken);
			_logger.LogInformation($"Concurrent executions: {_maxConcurrentExecutions - _semaphore.CurrentCount}");
			try {
				// HACK: java is expensive, they need more RAM than anything else
				var memoryLimit = language.ToLowerInvariant() switch {
					"java" or "javascript" => 512 * 1024 * 1024,
					_ => _pistonOptions.MemoryLimit,
				};

				var runTimeout = language.ToLowerInvariant() switch {
					"java" => 15_000,
					_ => _pistonOptions.RunTimeout,
				};

				var compileTimeout = language.ToLowerInvariant() switch {
					"java" => 15_000,
					_ => _pistonOptions.CompileTimeout,
				};

				var runtime = await GetRuntimeAsync(language, cancellationToken) ?? throw new KeyNotFoundException($"Runtime for {language} not found.");
				return await _rceClient.ExecuteAsync(
					new CodeRequest {
						Code = code,
						Version = version,
						Language = language,
						CompileTimeout = compileTimeout,
						RunTimeout = runTimeout,
						MemoryLimit = memoryLimit,
					},
					cancellationToken: cancellationToken
				);
			} finally {
				_semaphore!.Release();
				_logger.LogInformation($"Concurrent executions: {_maxConcurrentExecutions - _semaphore.CurrentCount}");
			}
		}

		// HACK: Hard coded check for first question
		private static bool IsTwinkleTwinkleLittleStarLyrics(string text) {
			var lines = text.Split(new[] { '\r', '\n' }, StringSplitOptions.RemoveEmptyEntries);
			if (lines.Length != 6) return false;
			if (lines[0].Trim().ToLowerInvariant() != "twinkle twinkle little star") return false;
			if (lines[1].Trim().ToLowerInvariant() != "how i wonder what you are") return false;
			if (lines[2].Trim().ToLowerInvariant() != "up above the world so high") return false;
			if (lines[3].Trim().ToLowerInvariant() != "like a diamond in the sky") return false;
			if (lines[4].Trim().ToLowerInvariant() != "twinkle twinkle little star") return false;
			if (lines[5].Trim().ToLowerInvariant() != "how i wonder what you are") return false;
			return true;
		}
	}
}
