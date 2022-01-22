import type { Language } from "./Language";
import type { TestResult } from "./TestResult";

export interface EditorSnapshot {
  questionNumber: number;
  language: Language;
  solutionByLanguage: Record<Language, string>;
  scratchPad: string;
  submissionAccepted: boolean | null;
  testResults: TestResult[] | null;
}