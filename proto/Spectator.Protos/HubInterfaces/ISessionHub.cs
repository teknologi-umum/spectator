using System.Threading.Tasks;
using Spectator.Protos.Session;

namespace Spectator.Protos.HubInterfaces {
	public interface ISessionHub {
		Task<SessionReply> StartSessionAsync(StartSessionRequest request);
		Task SetLocaleAsync(SetLocaleRequest request);
		Task SubmitPersonalInfoAsync(SubmitPersonalInfoRequest request);
		Task SubmitBeforeExamSAMAsync(SubmitSAMRequest request);
		Task SubmitAfterExamSAMAsync(SubmitSAMRequest request);
		Task<Exam> StartExamAsync(EmptyRequest request);
		Task<Exam> ResumeExamAsync(EmptyRequest request);
		Task<ExamResult> EndExamAsync(EmptyRequest request);
		Task<ExamResult> PassDeadlineAsync(EmptyRequest request);
		Task<ExamResult> ForfeitExamAsync(EmptyRequest request);
		Task<SubmissionResult> SubmitSolutionAsync(SubmissionRequest request);
		Task<SubmissionResult> TestSolutionAsync(SubmissionRequest request);
	}
}
