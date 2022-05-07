using System;
using System.Linq;
using FluentAssertions;
using Spectator.DomainModels.SubmissionDomain;
using Spectator.Piston.Internals;
using Spectator.Piston.Tests.Utilities;
using Xunit;

namespace Spectator.Piston.Tests {
	public class ResultParserTests {
		[Fact]
		public void CanParsePassingTestResult() {
			const string stdout =
				"# 1 PASSING\n" +
				"# 2 PASSING\n" +
				"# 3 PASSING\n" +
				"# 4 PASSING\n" +
				"# 5 PASSING\n";

			var testResults = ResultParser.ParseTestResults(stdout);
			testResults.Length.Should().Be(5);
			testResults[0].Should().BeOfType<PassingTestResult>().Which.TestNumber.Should().Be(1);
			testResults[1].Should().BeOfType<PassingTestResult>().Which.TestNumber.Should().Be(2);
			testResults[2].Should().BeOfType<PassingTestResult>().Which.TestNumber.Should().Be(3);
			testResults[3].Should().BeOfType<PassingTestResult>().Which.TestNumber.Should().Be(4);
			testResults[4].Should().BeOfType<PassingTestResult>().Which.TestNumber.Should().Be(5);
		}

		[Fact]
		public void CanParseFailingTestResult() {
			const string stdout =
				"# 1 PASSING\n" +
				"# 2 PASSING\n" +
				"# 3 FAILED\n" +
				"> EXPECTED wkwkwk\n" +
				"> GOT wlwlwl\n" +
				"# 4 PASSING\n" +
				"# 5 PASSING\n";

			var testResults = ResultParser.ParseTestResults(stdout);
			testResults.Length.Should().Be(5);
			testResults.Length.Should().Be(5);
			testResults[0].Should().BeOfType<PassingTestResult>().Which.TestNumber.Should().Be(1);
			testResults[1].Should().BeOfType<PassingTestResult>().Which.TestNumber.Should().Be(2);
			testResults[2].Should().BeOfType<FailingTestResult>().Which.TestNumber.Should().Be(3);
			((FailingTestResult)testResults[2]).ExpectedStdout.Should().Be("wkwkwk");
			((FailingTestResult)testResults[2]).ActualStdout.Should().Be("wlwlwl");
			testResults[3].Should().BeOfType<PassingTestResult>().Which.TestNumber.Should().Be(4);
			testResults[4].Should().BeOfType<PassingTestResult>().Which.TestNumber.Should().Be(5);
		}

		[Fact]
		public void CannotParseIncompeteFailingTestResult() {
			const string stdout =
				"# 1 PASSING\n" +
				"# 2 PASSING\n" +
				"# 3 FAILED\n" +
				"> EXPECTED wkwkwk\n" +
				"# 4 PASSING\n" +
				"# 5 PASSING\n";

			new Action(() => ResultParser.ParseTestResults(stdout)).Should().Throw<ArgumentException>()
				.And.Message.Should().Be("Cannot parse stdout (Parameter 'stdout')");
		}

		[Fact]
		public void CannotParseTestResultsWithUnknownMessage() {
			const string stdout =
				"# 1 PASSING\n" +
				"# 2 PASSING\n" +
				"# 3 FAILED\n" +
				"> EXPECTED wkwkwk\n" +
				"> GOT wlwlwl\n" +
				"# 4 PASSING\n" +
				"# 5 PASSING\n" +
				"# 6 PUSING";

			new Action(() => ResultParser.ParseTestResults(stdout)).Should().Throw<ArgumentException>()
				.And.Message.Should().Be("Cannot parse stdout (Parameter 'stdout')");
		}

		[Fact]
		public void CanParseSampleTestResultFromReinaldy() {
			const string stdout = "# 1 PASSING\n# 2 PASSING\n# 3 PASSING\n# 4 PASSING\n# 5 PASSING\n# 6 PASSING\n# 7 PASSING\n# 8 FAILED\n> EXPECTED -153.15\n> GOT -153.0\n# 9 PASSING\n# 10 PASSING\n";

			var testResults = ResultParser.ParseTestResults(stdout).ToArray();
			testResults.Length.Should().Be(10);
			testResults[0..6].Should().AllBeOfType<PassingTestResult>();
			testResults[7].Should().BeOfType<FailingTestResult>().Which.Should(its => {
				its.ExpectedStdout.Should().Be("-153.15");
				its.ActualStdout.Should().Be("-153.0");
			});
			testResults[8..9].Should().AllBeOfType<PassingTestResult>();
		}
	}
}
