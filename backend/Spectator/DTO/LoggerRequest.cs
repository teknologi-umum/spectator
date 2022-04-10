using System;

namespace Spectator.DTO; 

public class LoggerRequest {
	public string? Message { get; set; }
	public string? Level { get; set; }
	public DateTimeOffset? Timestamp { get; set; }
}
