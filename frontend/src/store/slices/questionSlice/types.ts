export interface Result {
  questionNo: number;
  isSubmitted: boolean;
  isRefactored: boolean;
}

export interface InitialState {
  currentQuestion: number;
  results: Result[];
}
