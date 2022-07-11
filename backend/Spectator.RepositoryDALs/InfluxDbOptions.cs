using System.Diagnostics.CodeAnalysis;

namespace Spectator.RepositoryDALs {
	public class InfluxDbOptions {
		[SuppressMessage("Design", "CA1056:URI-like properties should not be strings", Justification = "Options object")]
		public string Url { get; set; } = "http://localhost:8086";
		public string Token { get; set; } = "nMfrRYVcTyqFwDARAdqB92Ywj6GNMgPEd";
		public string SessionEventsBucket { get; set; } = "session_events";
		public string InputEventsBucket { get; set; } = "input_events";
		public string Org { get; set; } = "teknum_spectator";
	}
}
