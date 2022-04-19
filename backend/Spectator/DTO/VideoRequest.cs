using System;
using Microsoft.AspNetCore.Http;

namespace Spectator.DTO;

public class VideoRequest {
	public DateTimeOffset? StartedAt { get; set; }
	public DateTimeOffset? StoppedAt { get; set; }
	public IFormFile? File { get; set; }
}
