type Languages = "java" | "javascript" | "php" | "python" | "c++" | "c";

interface Solution {
  language: Languages;
  code: string;
}

export interface InitialState {
  currentLanguage: Languages;
  fontSize: number;
  code: string;
  solutions: Solution[];
}
