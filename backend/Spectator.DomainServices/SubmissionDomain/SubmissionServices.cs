using System;
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

		public async Task<Submission> EvaluateSubmissionAsync(int questionNumber, Locale locale, Language language, string directives, string solution, string scratchPad, CancellationToken cancellationToken) {
			// load assertion template
			var questionsByLocale = await _questionServices.GetAllAsync(cancellationToken);
			var question = questionsByLocale[locale].Single(q => q.QuestionNumber == questionNumber);
			var assertion = question.AssertionByLanguage[language];

			// insert directives and solution into the placeholder in assertion template
			var testCode = assertion
				.Replace("_REPLACE_ME_WITH_DIRECTIVES_", directives, StringComparison.Ordinal)
				.Replace("_REPLACE_ME_WITH_SOLUTION_", solution, StringComparison.Ordinal);

			// execute tests
			var testResults = questionNumber switch {
				// HACK: Hard coded check for zeroth and first question
				0 => await _pistonClient.ExecuteHelloWorldTestAsync(
					language: language,
					testCode: testCode,
					cancellationToken: cancellationToken
				),
				1 => await _pistonClient.ExecuteTwinkleTwinkleLittleStarTestAsync(
					language: language,
					testCode: testCode,
					cancellationToken: cancellationToken
				),
				_ => await _pistonClient.ExecuteTestsAsync(
					language: language,
					testCode: testCode,
					cancellationToken: cancellationToken
				)
			};

			return new Submission(
				QuestionNumber: questionNumber,
				Language: language,
				Solution: solution,
				ScratchPad: scratchPad,
				TestResults: testResults,
				SAMTestResult: null,
				// always return false if timeout reached because Enumerable.All will return true
				// if the list is empty
				Accepted: testResults.Length > 0 && testResults.All(testResult => testResult is PassingTestResult)
			);
		}
	}
}
