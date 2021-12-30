using System.Threading.Tasks;
using Spectator.Protos.Session;

namespace Spectator.Protos.HubInterfaces {
	public interface ISessionHub {
		Task<SessionReply> StartSessionAsync();
		Task SubmitPersonalInfoAsync(PersonalInfo personalInfo);
		Task SubmitBeforeCodeSAMAsync(SAM sam);
		Task<Exam> StartExamAsync();
		Task<Exam> ResumeExamAsync();
		Task<ExamResult> EndExamAsync();
		Task<SubmissionResult> SubmitSolutionAsync(Solution solution);
		Task SubmitAfterCodeSAM(SAM sam);
	}
}
