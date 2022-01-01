using System.Threading.Tasks;
using Spectator.Protos.Session;

namespace Spectator.Protos.HubInterfaces {
	public interface ISessionHub {
		Task<SessionReply> StartSessionAsync();
		Task SubmitPersonalInfoAsync(PersonalInfo personalInfo);
		Task SubmitBeforeExamSAMAsync(SAM sam);
		Task<Exam> StartExamAsync();
		Task<Exam> ResumeExamAsync();
		Task<ExamResult> EndExamAsync();
		Task<ExamResult> PassDeadlineAsync();
		Task<ExamResult> ForfeitExamAsync();
		Task<SubmissionResult> SubmitSolutionAsync(SubmissionRequest submissionRequest);
		Task SubmitAfterExamSAM(SAM sam);
	}
}
