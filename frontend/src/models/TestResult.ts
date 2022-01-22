interface PassingTest {
  status: "Passing";
}

interface FailingTest {
  status: "Failing";
  expectedStdout: string;
  actualStdout: string;
}

interface CompileError {
  status: "CompileError";
  stderr: string;
}

interface RuntimeError {
  status: "RuntimeError";
  stderr: string;
}

export type TestResult = {
  testNumber: number;
} & (
    | PassingTest
    | FailingTest
    | CompileError
    | RuntimeError
  );