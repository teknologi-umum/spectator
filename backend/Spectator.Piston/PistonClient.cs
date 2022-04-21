using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Linq;
using System.Net;
using System.Net.Http;
using System.Net.Http.Json;
using System.Text.Json;
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
					Language.C => "c",
					Language.CPP => "c++",
					Language.PHP => "php",
					Language.Javascript => "javascript",
					Language.Java => "java",
					Language.Python => "python",
					_ => throw new InvalidProgramException("Unhandled language")
				},
				code: testCode,
				cancellationToken: cancellationToken
			);

			// TODO: report runtime error together with passing and failing tests
			if (executeResult.ExitCode != 0) return ImmutableArray.Create<TestResultBase>(new RuntimeErrorResult(executeResult.Stderr));

			return ResultParser.ParseTestResults(executeResult.Stdout);
		}

		internal async Task<CodeResponse> ExecuteAsync(string language, string code, CancellationToken cancellationToken) {
			await _semaphore!.WaitAsync(cancellationToken);
			try {
				var runtime = await GetRuntimeAsync(language, cancellationToken) ?? throw new KeyNotFoundException($"Runtime for {language} not found.");
				return await _rceClient.ExecuteAsync(
					new CodeRequest {
						Code = code,
						Language = language,
						CompileTimeout = _pistonOptions.CompileTimeout,
						RunTimeout = _pistonOptions.RunTimeout,
						MemoryLimit = _pistonOptions.MemoryLimit,
					},
					cancellationToken: cancellationToken
				);
			} finally {
				_semaphore!.Release();
			}
		}
	}
}
