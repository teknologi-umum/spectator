using Microsoft.AspNetCore.Http;

namespace Spectator.DTO;

public class VideoRequest {
	public string SessionId { get; set; }
	public IFormFile File { get; set; }
}
