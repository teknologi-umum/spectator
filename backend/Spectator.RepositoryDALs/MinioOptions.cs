using System.Diagnostics.CodeAnalysis;

namespace Spectator.RepositoryDALs {
	public class MinioOptions {
		[SuppressMessage("Design", "CA1056:URI-like properties should not be strings", Justification = "Options object")]
		public string Url { get; set; } = "localhost:9000";
		public string AccessKey { get; set; } = "teknum";
		public string SecretKey { get; set; } = "c2N9Xz8bzHPkgNcgDtKgwGPTdb76GjD48";
	}
}
