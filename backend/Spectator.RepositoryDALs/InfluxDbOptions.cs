namespace Spectator.RepositoryDALs {
	public class InfluxDbOptions {
		public string Url { get; set; } = "http://localhost:8086";
		public string Token { get; set; } = "H76G7mEgcyeV2ffM%E#Vd8U^eA6ZY8GH";
		public string SessionEventsBucket { get; set; } = "session_events";
		public string InputEventsBucket { get; set; } = "input_events";
		public string Org { get; set; } = "teknum_spectator";
	}
}
