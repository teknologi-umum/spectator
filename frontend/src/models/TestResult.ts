// these are not arbitrary numbers, refer to protos/Spectator.Protos/session.proto#L121-L125
// this enum is used to map the result of `oneofKind` from the protobuf file
export enum ResultCase {
  Passing = 2,
  Failing = 3,
  CompileError = 4,
  RuntimeError = 5,
  InvalidInput = 6,
}

interface PassingTest {
  resultCase: ResultCase.Passing;
  passingTest: {
    expectedStdout: string;
    actualStdout: string;
    argumentsStdout: string;
  };
}

interface FailingTest {
  resultCase: ResultCase.Failing;
  failingTest: {
    expectedStdout: string;
    actualStdout: string;
    argumentsStdout: string;
  };
}

interface CompileError {
  resultCase: ResultCase.CompileError;
  compileError: {
    stderr: string;
  };
}

interface RuntimeError {
  resultCase: ResultCase.RuntimeError;
  runtimeError: {
    stderr: string;
  };
}

interface InvalidInput {
  resultCase: ResultCase.InvalidInput;
  invalidInput: {
    stderr: string;
  };
}

export type TestResult = {
  testNumber: number;
} & (PassingTest | FailingTest | CompileError | RuntimeError | InvalidInput);
