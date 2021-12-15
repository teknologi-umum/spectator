interface QuestionScore {
  mark: number;
  mouseEngagementScore: number;
  timeUsed: number;
}

export interface InitialState {
  nextQuestion: number;
  question1: QuestionScore;
  panas1: number;
  question2: QuestionScore;
  panas2: number;
  question3: QuestionScore;
  panas3: number;
  question4: QuestionScore;
  panas4: number;
  question5: QuestionScore;
  panas5: number;
}
