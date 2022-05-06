import type { Language } from "./Language";
import type { TestResult } from "@/stub/session";

export interface EditorSnapshot {
  questionNumber: number;
  language: Language;
  directivesByLanguage: Record<Language, string>;
  solutionByLanguage: Record<Language, string>;
  scratchPad: string;
  submissionSubmitted: boolean;
  submissionAccepted: boolean;
  submissionRefactored: boolean;
  testResults: TestResult[] | null;
}
