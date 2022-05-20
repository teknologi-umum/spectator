using System;
using System.Threading;
using System.Threading.Tasks;
using Spectator.WorkerClient;
using Spectator.DomainModels.ExamEndedDomain;

namespace Spectator.DomainServices.ExamResultDomain {
	public class ExamResultServices {
		private readonly WorkerServices _workerServices;

		public ExamResultServices(WorkerServices workerServices) {
			_workerServices = workerServices;
		}

		public async Task<Funfact> GenerateFunfactAsync(Guid sessionID, CancellationToken cancellationToken) {
			// retrieve funfact
			var funfact = await _workerServices.FunFactAsync(sessionID, cancellationToken);

			// generate the files
			await _workerServices.GenerateFilesAsync(sessionID, cancellationToken);

			return new Funfact(
				wordsPerMinute: funfact.WordsPerMinute,
				deletionRate: funfact.DeletionRate,
				submissionAttempts: funfact.SubmissionAttempts
			);
		}
	}
}
