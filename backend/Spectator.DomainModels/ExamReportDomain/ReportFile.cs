using System;

namespace Spectator.DomainModels.ExamReportDomain {
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
