export const setNextQuestion = (questionNum) => ({
    type: 'SET_NEXT_QUESTION',
    questionNum
});

export const submitPretest = (SAMtest) => ({
    type: 'SUBMIT_PRETEST',
    SAMtest
});

export const submitQ1 = (questionResponse) => ({
    type: 'SUBMIT_Q1',
    questionResponse
});

export const submitQ2 = (questionResponse) => ({
    type: 'SUBMIT_Q2',
    questionResponse
});

export const submitQ3= (questionResponse) => ({
    type: 'SUBMIT_Q3',
    questionResponse
});

export const submitQ4 = (questionResponse) => ({
    type: 'SUBMIT_Q4',
    questionResponse
});

export const submitQ5 = (questionResponse) => ({
    type: 'SUBMIT_Q5',
    questionResponse
});

export const submitQ6 = (questionResponse) => ({
    type: 'SUBMIT_Q6',
    questionResponse
});

export const submitPosttest = (SAMtest) => ({
    type: 'SUBMIT_POSTTEST',
    SAMtest
});