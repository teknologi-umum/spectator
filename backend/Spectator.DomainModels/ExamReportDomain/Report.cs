using System;
using System.Collections.Generic;

namespace Spectator.DomainModels.ExamReportDomain {
	public class Report {
		public IReadOnlyList<ReportFile> ReportFiles { get; }
		public Uri VideoUrl { get; }

		public Report(IReadOnlyList<ReportFile> reportFiles, Uri videoUrl) {
			ReportFiles = reportFiles;
			VideoUrl = videoUrl;
		}
	}
}
