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
	}
}
