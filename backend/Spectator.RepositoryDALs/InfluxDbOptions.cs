namespace Spectator.RepositoryDALs {
	public class InfluxDbOptions {
		public string? Url { get; set; }
		public string? Token { get; set; }
		public string? SessionEventsBucket { get; set; }
		public string? InputEventsBucket { get; set; }
		public string? Org { get; set; }
	}
}
