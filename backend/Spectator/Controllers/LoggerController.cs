using System;
using System.Threading;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Options;
using Spectator.DTO;
using Spectator.LoggerClient;
using Spectator.Protos.Logger;
using Environment = Spectator.Protos.Logger.Environment;

namespace Spectator.Controllers;

[Route("log")]
public class Logger : Controller {
	private readonly LoggerServices _loggerClient;
	private readonly string _loggerAccessToken;

	public Logger(LoggerServices loggerClient, IOptions<LoggerOptions> loggerOptionsAssessor) {
		_loggerClient = loggerClient;
		_loggerAccessToken = loggerOptionsAssessor.Value.AccessToken ??
							 throw new InvalidOperationException("LoggerOptions:AccessToken is required");
	}

	[HttpPost]
	public async Task<IActionResult> LogAsync([FromBody] LoggerRequest request, CancellationToken cancellationToken) {
		if (request == null) throw new ArgumentNullException(nameof(request));
		if (request.Timestamp == null) throw new ArgumentNullException(nameof(request.Timestamp));
		var logData = new LogData {
			RequestId = Guid.NewGuid().ToString(),
			// TODO: implement correct way of injecting environment variable to reflect the enum provided in the GRPC stub
			Environment = Environment.Unset,
			Application = "frontend",
			Language = "Javascript",
			Level = request.Level switch {
				"error" => Level.Error,
				"warn" => Level.Warning,
				"info" => Level.Info,
				_ => Level.Debug
			},
			Message = request.Message,
			Timestamp = request.Timestamp.Value.ToUnixTimeMilliseconds()
		};
		await _loggerClient.CreateLogAsync(_loggerAccessToken, logData, cancellationToken);
		return Ok();
	}
}
