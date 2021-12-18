const questionInit = {
  mark: 0,
  mouseEngagementScore: 0,
  timeUsed: 0
};

const initState = {
  nextQuestion: 0,
  question1: questionInit,
  question2: questionInit,
  question3: questionInit,
  question4: questionInit,
  question5: questionInit
};

// FIXME: proper action interface
// eslint-disable-next-line
const questionReducer = (state = initState, action: any) => {
  switch (action.type) {
    case "SET_NEXT_QUESTION":
      return {
        ...state,
        nextQuestion: action.questionNum || 0
      };
    case "SUBMIT_QUESTION_ONE":
      return {
        ...state,
        question1:
          (action.questionResponse && {
            ...questionInit,
            ...action.questionResponse
          }) ||
          questionInit
      };
    case "SUBMIT_QUESTION_TWO":
      return {
        ...state,
        question2:
          (action.questionResponse && {
            ...questionInit,
            ...action.questionResponse
          }) ||
          questionInit
      };
    case "SUBMIT_QUESTION_THREE":
      return {
        ...state,
        question3:
          (action.questionResponse && {
            ...questionInit,
            ...action.questionResponse
          }) ||
          questionInit
      };
    case "SUBMIT_QUESTION_FOUR":
      return {
        ...state,
        question4:
          (action.questionResponse && {
            ...questionInit,
            ...action.questionResponse
          }) ||
          questionInit
      };
    case "SUBMIT_QUESTION_FIVE":
      return {
        ...state,
        question5:
          (action.questionResponse && {
            ...questionInit,
            ...action.questionResponse
          }) ||
          questionInit
      };
    default:
      return state;
  }
};

export default questionReducer;
