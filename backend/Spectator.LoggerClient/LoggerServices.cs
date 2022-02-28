using System;
using System.Collections.Immutable;
using System.Threading;
using System.Threading.Tasks;
using Grpc.Net.Client;
using Microsoft.Extensions.Options;
using Spectator.Protos.Logger;

namespace Spectator.LoggerClient {
	public class LoggerServices {
		private readonly GrpcChannel _grpcChannel;
		private readonly Logger.LoggerClient _loggerClient;

		public LoggerServices(
			IOptions<LoggerOptions> optionsAccessor
		) {
			_grpcChannel = GrpcChannel.ForAddress(optionsAccessor.Value.Address);
			_loggerClient = new(_grpcChannel);
		}

		public async Task CreateLogAsync(string accessToken, LogData logData, CancellationToken cancellationToken) {
			await _loggerClient.CreateLogAsync(new LogRequest {
				AccessToken = accessToken,
				Data = logData
			}, cancellationToken: cancellationToken);
		}

		public async Task<ImmutableList<LogData>> ReadLogAsync(
			Level level,
			string requestId,
			string application,
			DateTimeOffset timestampFrom,
			DateTimeOffset timestampTo,
			CancellationToken cancellationToken
		) {
			var logResponse = await _loggerClient.ReadLogAsync(new ReadLogRequest {
				Level = level,
				RequestId = requestId,
				Application = application,
				TimestampFrom = timestampFrom.ToUnixTimeMilliseconds(),
				TimestampTo = timestampTo.ToUnixTimeMilliseconds(),
			}, cancellationToken: cancellationToken);

			return logResponse.Data.ToImmutableList();
		}

		public async Task PingAsync(CancellationToken cancellationToken) {
			await _loggerClient.PingAsync(new EmptyRequest(), cancellationToken: cancellationToken);
		}
	}
}
