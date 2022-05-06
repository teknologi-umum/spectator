using System;
using Microsoft.AspNetCore.Http;

namespace Spectator.DTO;

public class VideoRequest {
	public long? StartedAt { get; set; }
	public long? StoppedAt { get; set; }
	public IFormFile? File { get; set; }
}
