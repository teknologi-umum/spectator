export type Languages = "java" | "javascript" | "php" | "python" | "cpp" | "c";

export interface Submission {
  questionNo: number;
  language: Languages;
  code: string;
  isSubmitted: boolean;
  isRefactored: boolean;
}

export interface InitialState {
  currentQuestion: number;
  submissions: Submission[];
}
