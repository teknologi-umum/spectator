namespace Spectator.DomainModels.ExamReportDoman {
	// I'm going to create a POCO because to be honest,
	// I don't know what I'm doing. This is all written
	// as I think it's probably right.
	public class File {
		public string SessionId { get; set; }
		public string StudentNumber { get; set; }
		public string JSONFileUrl { get; set; }
		public string CSVFileUrl { get; set; }
	}
}
