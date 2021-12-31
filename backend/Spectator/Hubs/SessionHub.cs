using System.Threading.Tasks;
using Microsoft.AspNetCore.SignalR;
using Spectator.DomainServices;
using Spectator.Protos.HubInterfaces;
using Spectator.Protos.Session;

namespace Spectator.Hubs {
	public class SessionHub : Hub<ISessionHub>, ISessionHub {
		private readonly SessionServices _sessionServices;

		public SessionHub(
			SessionServices sessionServices
		) {
			_sessionServices = sessionServices;
		}

		public Task<SessionReply> StartSessionAsync() => throw new System.NotImplementedException();
		public Task<ExamResult> EndExamAsync() => throw new System.NotImplementedException();
		public Task<Exam> ResumeExamAsync() => throw new System.NotImplementedException();
		public Task<Exam> StartExamAsync() => throw new System.NotImplementedException();
		public Task SubmitAfterCodeSAM(SAM sam) => throw new System.NotImplementedException();
		public Task SubmitBeforeCodeSAMAsync(SAM sam) => throw new System.NotImplementedException();
		public Task SubmitPersonalInfoAsync(PersonalInfo personalInfo) => throw new System.NotImplementedException();
		public Task<SubmissionResult> SubmitSolutionAsync(Solution solution) => throw new System.NotImplementedException();
	}
}
