export type Languages = "java" | "javascript" | "php" | "python" | "cpp" | "c";

export interface Solution {
  questionNo: number;
  language: Languages;
  code: string;
}

export interface ScratchPad {
  questionNo: number;
  value: string;
}

export interface InitialState {
  currentLanguage: Languages;
  fontSize: number;
  solutions: Solution[];
  scratchPads: ScratchPad[];
}
