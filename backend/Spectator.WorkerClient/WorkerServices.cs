using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Linq;
using System.Threading;
using System.Threading.Tasks;
using Grpc.Net.Client;
using Microsoft.Extensions.Options;
using Spectator.Protos.Worker;

namespace Spectator.WorkerClient {
	public class WorkerServices {
		private readonly GrpcChannel _grpcChannel;
		private readonly Worker.WorkerClient _workerClient;

		public WorkerServices(
			IOptions<WorkerOptions> optionsAccessor
		) {
			_grpcChannel = GrpcChannel.ForAddress(optionsAccessor.Value.Address);
			_workerClient = new(_grpcChannel);
		}

		public async Task PingAsync(CancellationToken cancellationToken) {
			await _workerClient.PingAsync(new EmptyRequest(), cancellationToken: cancellationToken);
		}

		public async Task<ImmutableList<FilesList>> GetFilesListsBySessionIdsAsync(IReadOnlySet<Guid> sessionIds, CancellationToken cancellationToken) {
			var reply = await _workerClient.ListMultipleFilesAsync(
				request: new() {
					RequestId = Guid.NewGuid().ToString(),
					SessionId = {
						from sessionId in sessionIds
						select sessionId.ToString()
					}
				},
				cancellationToken: cancellationToken
			);
			return reply.FilesList.ToImmutableList();
		}

		public Task GenerateFilesAsync(Guid sessionId, CancellationToken cancellationToken) {
			return _workerClient.GenerateFilesAsync(
				request: new Member {
					RequestId = Guid.NewGuid().ToString(),
					SessionId = sessionId.ToString()
				},
				cancellationToken: cancellationToken
			).ResponseAsync;
		}

		public Task<FunFactResponse> FunFactAsync(Guid sessionId, CancellationToken cancellationToken) {
			return _workerClient.FunFactAsync(
				request: new Member {
					RequestId = Guid.NewGuid().ToString(),
					SessionId = sessionId.ToString()
				},
				cancellationToken: cancellationToken
			).ResponseAsync;
		}
	}
}
