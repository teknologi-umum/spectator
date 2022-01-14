using System.Linq;
using System.Threading;
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

		public async Task<Submission> EvaluateSubmissionAsync(int questionNumber, Locale locale, Language language, string solution, string scratchPad, CancellationToken cancellationToken) {
			// load assertion template
			var questionsByLocale = await _questionServices.GetAllAsync(cancellationToken);
			var question = questionsByLocale[locale].Single(q => q.QuestionNumber == questionNumber);
			var assertion = question.AssertionByLanguage[language];

			// insert solution into the placeholder in assertion template
			var testCode = assertion.Replace("_REPLACE_ME_", solution);

			// execute tests
			var testResults = await _pistonClient.ExecuteTestsAsync(
				language: language,
				testCode: testCode,
				cancellationToken: cancellationToken
			);

			return new Submission(
				QuestionNumber: questionNumber,
				Language: language,
				Solution: solution,
				ScratchPad: scratchPad,
				TestResults: testResults,
				Accepted: testResults.All(testResult => testResult is PassingTestResult)
			);
		}
	}
}
