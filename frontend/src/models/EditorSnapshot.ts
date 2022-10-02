import type { Language } from "./Language";
import type { SAMTestResult } from "./SAMTestResult";
import type { TestResult } from "./TestResult";

export interface EditorSnapshot {
  questionNumber: number;
  language: Language;
  solutionByLanguage: Record<Language, string>;
  scratchPad: string;
  submissionSubmitted: boolean;
  submissionAccepted: boolean;
  submissionRefactored: boolean;
  testResults: TestResult[] | null;
  samTestResult: SAMTestResult | null;
}
