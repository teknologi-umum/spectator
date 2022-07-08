using System.Collections.Generic;

namespace Spectator.DomainModels.ExamReportDomain {
	public class Report {
		public IReadOnlyList<ReportFile> ReportFiles { get; private init; }

		public Report(IReadOnlyList<ReportFile> reportFiles, string videoUrl) {
			ReportFiles = reportFiles;
		}
	}
}
