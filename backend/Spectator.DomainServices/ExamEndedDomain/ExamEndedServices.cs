using System;
using System.Threading;
using System.Threading.Tasks;
using Spectator.WorkerClient;
using Spectator.DomainModels.ExamEndedDomain;

namespace Spectator.DomainServices.ExamEndedDomain {
	public class ExamEndedServices {
		private readonly WorkerServices _workerServices;

		public ExamEndedServices(WorkerServices workerServices) {
			_workerServices = workerServices;
		}

		public async Task<Funfact> GenerateFunfact(Guid sessionID, CancellationToken cancellationToken) {
			var funfact = await _workerServices.FunFactAsync(sessionID, cancellationToken);

			await _workerServices.GenerateFilesAsync(sessionID, cancellationToken);

			return new Funfact(funfact.WordsPerMinute, funfact.DeletionRate, funfact.SubmissionAttempts);
		}
	}
}
