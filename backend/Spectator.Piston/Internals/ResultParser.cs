using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Text.RegularExpressions;
using Spectator.DomainModels.SubmissionDomain;

namespace Spectator.Piston.Internals {
	internal static class ResultParser {
		public static ImmutableArray<TestResultBase> ParseTestResults(string stdout) {
			var lines = stdout.Split(new[] { '\r', '\n' }, System.StringSplitOptions.RemoveEmptyEntries);
			var testResults = new List<TestResultBase>();

			for (var i = 0; i < lines.Length; i++) {
				if (Regex.Match(lines[i], "# ([0-9]+) PASSING").Groups is { Count: 2 } passingGroups
						&& int.TryParse(passingGroups[1].Value, out var passingTestNumber)) {
					testResults.Add(new PassingTestResult(
						TestNumber: passingTestNumber
					));
					continue;
				}

				if (Regex.Match(lines[i], "# ([0-9]+) FAILED").Groups is { Count: 2 } failingGroups
					&& int.TryParse(failingGroups[1].Value, out var failingTestNumber)
					&& Regex.Match(lines[++i], "> EXPECTED (.*)").Groups is { Count: 2 } expectedGroups
					&& Regex.Match(lines[++i], "> GOT (.*)").Groups is { Count: 2 } actualGroups) {
					testResults.Add(new FailingTestResult(
						TestNumber: failingTestNumber,
						ExpectedStdout: expectedGroups[1].Value,
						ActualStdout: actualGroups[1].Value
					));
					continue;
				}

				throw new ArgumentException("Cannot parse stdout", nameof(stdout));
			}

			return testResults.ToImmutableArray();
		}
	}
}
