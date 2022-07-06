using System;
using System.Threading;
using System.Threading.Tasks;
using Grpc.Net.Client;
using Microsoft.Extensions.Options;
using Spectator.Protos.Video;

namespace Spectator.VideoClient {
	public class VideoServices : IDisposable {
		private readonly GrpcChannel _grpcChannel;
		private readonly VideoService.VideoServiceClient _videoClient;

		public VideoServices(
			IOptions<VideoOptions> optionsAccessor
		) {
			_grpcChannel = GrpcChannel.ForAddress(optionsAccessor.Value.Address);
			_videoClient = new(_grpcChannel);
		}

		public async Task PingAsync(CancellationToken cancellationToken) {
			await _videoClient.PingAsync(new EmptyRequest(), cancellationToken: cancellationToken);
		}


		public async Task GetVideoAsync(Guid sessionId, CancellationToken cancellationToken) {
			await _videoClient.GetVideoAsync(
				request: new VideoRequest {
					SessionId = sessionId.ToString()
				},
				cancellationToken: cancellationToken
			);
		}

		public void Dispose(bool disposing) {
			if (disposing) {
				_grpcChannel.Dispose();
			}
		}

		public void Dispose() {
			// Do not change this code. Put cleanup code in 'Dispose(bool disposing)' method
			Dispose(disposing: true);
			GC.SuppressFinalize(this);
		}
	}
}
