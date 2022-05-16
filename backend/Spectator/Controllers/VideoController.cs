using System;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Options;
using Spectator.DomainModels.SessionDomain;
using Spectator.LoggerClient;
using Spectator.PoormansAuth;
using Spectator.DTO;
using Spectator.Protos.Logger;
using Environment = Spectator.Protos.Logger.Environment;
using Minio;

namespace Spectator.Controllers;

[ApiController]
public class VideoController : ControllerBase {
	private readonly PoormansAuthentication _poormansAuthentication;
	private readonly MinioClient _minioClient;
	private readonly LoggerServices _loggerClient;
	private readonly string _loggerAccessToken;

	public VideoController(
		LoggerServices logger,
		IOptions<LoggerOptions> loggerOptionsAssessor,
		PoormansAuthentication poormansAuthentication,
		MinioClient minioClient
	) {
		_loggerClient = logger;
		_minioClient = minioClient;
		_poormansAuthentication = poormansAuthentication;
		_loggerAccessToken = loggerOptionsAssessor.Value.AccessToken ??
							 throw new InvalidOperationException("LoggerOptions:AccessToken is required");
	}

	// TODO: Refactor this method so it doesn't get fat
	[HttpPost]
	[Route("/video")]
	public async Task<IActionResult> LoginAsync([FromForm] VideoRequest request, [FromHeader(Name = "Authorization")] string accessToken) {
		if (request.File == null) throw new ArgumentNullException(nameof(request.File));
		if (request.StartedAt == null) throw new ArgumentNullException(nameof(request.StartedAt));
		if (request.StoppedAt == null) throw new ArgumentNullException(nameof(request.StoppedAt));

		// Only registered users can upload videos
		if (string.IsNullOrEmpty(accessToken)) return Unauthorized();

		var session = _poormansAuthentication.Authenticate(accessToken);

		// Authorize: Exam must be in progress
		if (session is not RegisteredSession registeredSession) throw new UnauthorizedAccessException("Personal Info not yet submitted");
		if (registeredSession.ExamStartedAt is null) throw new UnauthorizedAccessException("Exam not yet started");
		if (registeredSession.ExamEndedAt is not null) throw new UnauthorizedAccessException("Exam already ended");
		if (registeredSession.ExamDeadline is null) throw new InvalidProgramException("Exam deadline not set");
		if (registeredSession.ExamDeadline.Value < DateTimeOffset.UtcNow) throw new UnauthorizedAccessException("Exam deadline exceeded");

		try {
			var found = await _minioClient.BucketExistsAsync(
				new BucketExistsArgs().WithBucket(registeredSession.Id.ToString())
			);
			if (!found) {
				await _minioClient.MakeBucketAsync(
					new MakeBucketArgs().WithBucket(registeredSession.Id.ToString())
				);
			}

			var filestream = request.File.OpenReadStream();
			await _minioClient.PutObjectAsync(
				new PutObjectArgs().WithBucket(registeredSession.Id.ToString())
								   .WithObject($"{request.StartedAt}_{request.StoppedAt}.webm")
								   .WithStreamData(filestream)
								   .WithObjectSize(filestream.Length)
			);
		} catch (Exception e) {
			var logData = new LogData {
				RequestId = Guid.NewGuid().ToString(),
				// TODO: implement correct way of injecting environment variable to reflect the enum provided in the GRPC stub
				Environment = Environment.Unset,
				Application = "Spectator.VideoController",
				Language = "C#",
				Level = Level.Error,
				Message = e.Message,
				Timestamp = DateTimeOffset.UtcNow.ToUnixTimeMilliseconds()
			};
			await _loggerClient.CreateLogAsync(_loggerAccessToken, logData, default);
		}

		return Ok();
	}
}
