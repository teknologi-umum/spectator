interface PassingTest {
  status: "Passing";
  result: {
    testNumber: number;
  };
}

interface FailingTest {
  status: "Failing";
  result: {
    testNumber: number;
    expectedStdout: string;
    actualStdout: string;
  };
}

interface CompileError {
  status: "CompileError";
  result: {
    testNumber: number;
    stderr: string;
  };
}

interface RuntimeError {
  status: "RuntimeError";
  result: {
    testNumber: number;
    stderr: string;
  };
}

export type TestResult =
  | PassingTest
  | FailingTest
  | CompileError
  | RuntimeError;
