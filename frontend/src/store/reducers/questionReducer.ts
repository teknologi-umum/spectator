const questionInit = {
    mark: 0,
    mouseEngagementScore: 0,
    timeUsed: 0
};

const initState = {
    nextQuestion: 0,
    question1: questionInit,
    panas1: 0,
    question2: questionInit,
    panas2: 0,
    question3: questionInit,
    panas3: 0,
    question4: questionInit,
    panas4: 0,
    question5: questionInit,
    panas5: 0
};

const questionReducer = (state = initState, action) => {
    switch (action.type) {
        case "SET_NEXT_QUESTION":
            return {
                ...state,
                nextQuestion: action.questionNum || 0
            };
        case "SUBMIT_QUESTION_ONE":
            return {
                ...state,
                question1: action.questionResponse && { ...questionInit, ...action.questionResponse }
                || questionInit
            };
        case "SUBMIT_PANAS_ONE":
            return {
                ...state,
                panas1: action.panasScore || 0
            };
        case "SUBMIT_QUESTION_TWO":
            return {
                ...state,
                question2: action.questionResponse && { ...questionInit, ...action.questionResponse }
                || questionInit
            };
        case "SUBMIT_PANAS_TWO":
            return {
                ...state,
                panas2: action.panasScore || 0
            };
        case "SUBMIT_QUESTION_THREE":
            return {
                ...state,
                question3: action.questionResponse && { ...questionInit, ...action.questionResponse }
                || questionInit
            };
        case "SUBMIT_PANAS_THREE":
            return {
                ...state,
                panas3: action.panasScore || 0
            };
        case "SUBMIT_QUESTION_FOUR":
            return {
                ...state,
                question4: action.questionResponse && { ...questionInit, ...action.questionResponse }
                || questionInit
            };
        case "SUBMIT_PANAS_FOUR":
            return {
                ...state,
                panas4: action.panasScore || 0
            };
        case "SUBMIT_QUESTION_FIVE":
            return {
                ...state,
                question5: action.questionResponse && { ...questionInit, ...action.questionResponse }
                || questionInit
            };
        case "SUBMIT_PANAS_FIVE":
            return {
                ...state,
                panas5: action.panasScore || 0
            };
        default:
            return state;
    }
};

export default questionReducer;