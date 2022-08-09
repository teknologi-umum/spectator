using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Text.RegularExpressions;
using Spectator.DomainModels.SubmissionDomain;

namespace Spectator.Piston.Internals {
	internal static class ResultParser {
		public static ImmutableArray<TestResultBase> ParseTestResults(string stdout) {
			var lines = stdout.Split(new[] { '\r', '\n' }, StringSplitOptions.RemoveEmptyEntries);
			var testResults = new List<TestResultBase>();

			for (var i = 0; i < lines.Length; i++) {
				if (Regex.Match(lines[i], "# ([0-9]+) PASSING").Groups is { Count: 2 } passingGroups
					&& int.TryParse(passingGroups[1].Value, out var passingTestNumber)
					&& Regex.Match(lines[++i], "> ARGUMENTS (.*)").Groups is { Count: 2 } passingArgumentsGroups
					&& Regex.Match(lines[++i], "> EXPECTED (.*)").Groups is { Count: 2 } passingExpectedGroups
					&& Regex.Match(lines[++i], "> GOT (.*)").Groups is { Count: 2 } passingActualGroups) {
					testResults.Add(new PassingTestResult(
						TestNumber: passingTestNumber,
						ExpectedStdout: passingExpectedGroups[1].Value,
						ActualStdout: passingActualGroups[1].Value,
						ArgumentsStdout: passingArgumentsGroups[1].Value

					));
					continue;
				}

				if (Regex.Match(lines[i], "# ([0-9]+) FAILED").Groups is { Count: 2 } failingGroups
					&& int.TryParse(failingGroups[1].Value, out var failingTestNumber)
					&& Regex.Match(lines[++i], "> ARGUMENTS (.*)").Groups is { Count: 2 } failingArgumentsGroups
					&& Regex.Match(lines[++i], "> EXPECTED (.*)").Groups is { Count: 2 } failingExpectedGroups
					&& Regex.Match(lines[++i], "> GOT (.*)").Groups is { Count: 2 } failingActualGroups) {
					testResults.Add(new FailingTestResult(
						TestNumber: failingTestNumber,
						ExpectedStdout: failingExpectedGroups[1].Value,
						ActualStdout: failingActualGroups[1].Value,
						ArgumentsStdout: failingArgumentsGroups[1].Value
					));
					continue;
				}

				throw new ArgumentException("Cannot parse stdout", nameof(stdout));
			}

			return testResults.ToImmutableArray();
		}
	}
}
