export const setNextQuestion = (questionNum) => ({
  type: "SET_NEXT_QUESTION",
  questionNum
});

export const submitPreTest = (samTestScore) => ({
  type: "SUBMIT_PRETEST",
  samTestScore
});

export const submitQues1Reponse = (questionResponse) => ({
    type: "SUBMIT_QUESTION_ONE",
    questionResponse
  });

export const submitQues2Reponse = (questionResponse) => ({
    type: "SUBMIT_QUESTION_TWO",
    questionResponse
  });

export const submitQues3Reponse = (questionResponse) => ({
    type: "SUBMIT_QUESTION_THREE",
    questionResponse
  });

export const submitQues4Reponse = (questionResponse) => ({
    type: "SUBMIT_QUESTION_FOUR",
    questionResponse
  });

export const submitQues5Reponse = (questionResponse) => ({
    type: "SUBMIT_QUESTION_FIVE",
    questionResponse
  });

export const submitPostTest = (samTestScore) => ({
  type: "SUBMIT_POSTTEST",
  samTestScore
});