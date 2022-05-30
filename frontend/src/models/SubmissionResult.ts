import type { TestResult } from "./TestResult";

export interface SubmissionResult {
  accepted: boolean;
  testResults: TestResult[]
}
