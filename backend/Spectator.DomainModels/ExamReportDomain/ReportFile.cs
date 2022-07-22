using System;

namespace Spectator.DomainModels.ExamReportDomain {
	public class ReportFile {
		public Guid SessionId { get; }
		public string StudentNumber { get; }
		public Uri? JSONFileUrl { get; }
		public Uri? CSVFileUrl { get; }

		public ReportFile(Guid sessionId, string studentNumber, Uri? jsonFileUrl, Uri? csvFileUrl) {
			SessionId = sessionId;
			StudentNumber = studentNumber;
			JSONFileUrl = jsonFileUrl;
			CSVFileUrl = csvFileUrl;
		}
	}
}
