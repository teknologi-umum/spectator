using System;
using System.Threading;
using System.Threading.Tasks;
using Spectator.WorkerClient;
using Spectator.DomainModels.ExamEndedDomain;
using Spectator.VideoClient;

namespace Spectator.DomainServices.ExamResultDomain {
	public class ExamResultServices {
		private readonly WorkerServices _workerServices;
		private readonly VideoServices _videoService;

		public ExamResultServices(WorkerServices workerServices, VideoServices videoService) {
			_workerServices = workerServices;
			_videoService = videoService;
		}

		public async Task<Funfact> GenerateFunfactAsync(Guid sessionID, CancellationToken cancellationToken) {
			// retrieve funfact
			var funfact = await _workerServices.FunFactAsync(sessionID, cancellationToken);

			// generate the files
			await _workerServices.GenerateFilesAsync(sessionID, cancellationToken);

			// concat video chunks but just forget about them so we don't force the user to wait
			// for ~60 seconds before seeing their funfact
			_videoService.GetVideoAsync(sessionID, CancellationToken.None).Ignore();

			return new Funfact(
				wordsPerMinute: funfact.WordsPerMinute,
				deletionRate: funfact.DeletionRate,
				submissionAttempts: funfact.SubmissionAttempts
			);
		}
	}
}
