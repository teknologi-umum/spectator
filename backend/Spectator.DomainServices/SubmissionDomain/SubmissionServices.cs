using System.Collections.Immutable;
using System.Threading.Tasks;
using Spectator.DomainModels.SubmissionDomain;
using Spectator.DomainServices.QuestionDomain;
using Spectator.Piston;
using Spectator.Primitives;

namespace Spectator.DomainServices.PistonDomain {
	public class SubmissionServices {
		private readonly PistonClient _pistonClient;
		private readonly QuestionServices _questionServices;

		public SubmissionServices(
			PistonClient pistonClient,
			QuestionServices questionServices
		) {
			_pistonClient = pistonClient;
			_questionServices = questionServices;
		}

		public async Task<Submission> EvaluateSubmissionAsync(int questionNumber, Language language, string solution, string scratchPad) {
			// TODO: implement this method properly
			// HACK: dummy implementation
			return new Submission(
				QuestionNumber: questionNumber,
				Language: language,
				Solution: solution,
				ScratchPad: scratchPad,
				TestResults: ImmutableArray<TestResult>.Empty,
				Accepted: true
			);
		}
	}
}
