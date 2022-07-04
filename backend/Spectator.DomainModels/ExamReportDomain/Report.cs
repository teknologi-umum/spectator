using System.Collections.Generic;

namespace Spectator.DomainModels.ExamReportDomain {
	public class Report {
		public IEnumerable<ReportFile> ReportFiles { get; private init; }

		public Report(IEnumerable<ReportFile> reportFiles, string videoUrl) {
			ReportFiles = reportFiles;
		}
	}
}
