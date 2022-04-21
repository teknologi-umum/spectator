using System;

namespace Spectator.DomainModels.ExamReportDoman {
	// I'm going to create a POCO because to be honest,
	// I don't know what I'm doing. This is all written
	// as I think it's probably right.
	public class ReportFile {
		public Guid SessionId { get; private init; }
		public string StudentNumber { get; private init; }
		public string JSONFileUrl { get; private init; }
		public string CSVFileUrl { get; private init; }

		public ReportFile(Guid sessionId, string studentNumber, string jsonFileUrl, string csvFileUrl) {
			SessionId = sessionId;
			StudentNumber = studentNumber;
			JSONFileUrl = jsonFileUrl;
			CSVFileUrl = csvFileUrl;
		}
	}
}
