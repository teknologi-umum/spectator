export type Languages = "java" | "javascript" | "php" | "python" | "cpp" | "c";

export interface Solution {
  language: Languages;
  code: string;
}

export interface InitialState {
  currentLanguage: Languages;
  fontSize: number;
  code: string;
  solutions: Solution[];
}
