using System;
using System.Threading;
using System.Threading.Tasks;
using Grpc.Net.Client;
using Microsoft.Extensions.Options;
using Spectator.Protos.Video;

namespace Spectator.VideoClient {
	public class VideoServices {
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


		public Task GetVideoAsync(Guid sessionId, CancellationToken cancellationToken) {
			return _videoClient.GetVideoAsync(
				new VideoRequest {
					SessionId = sessionId.ToString()
				}
			).ResponseAsync;
		}
	}
}
