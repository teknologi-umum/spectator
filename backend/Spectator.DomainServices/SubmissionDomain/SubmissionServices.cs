using System.Threading.Tasks;
using Spectator.DomainModels.SubmissionDomain;
using Spectator.Piston;
using Spectator.Primitives;

namespace Spectator.DomainServices.PistonDomain {
	public class SubmissionServices {
		private readonly PistonClient _pistonClient;

		public SubmissionServices(
			PistonClient pistonClient
		) {
			_pistonClient = pistonClient;
		}

		public async Task<Submission> EvaluateSubmissionAsync(int questionNumber, Language language, string solution, string scratchPad) {
			// TODO: implement this method properly
			// HACK: dummy implementation
			return new Submission(
				QuestionNumber: questionNumber,
				Language: language,
				Solution: solution,
				ScratchPad: scratchPad,
				ErrorMessage: null,
				ConsoleOutput: "",
				Accepted: true
			);
		}
	}
}
