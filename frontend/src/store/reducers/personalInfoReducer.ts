const initState = {
    // projects: [
    //     {id: '1', title: 'help me find peach', content: 'blah blah blah'},
    //     {id: '2', title: 'collect all the stars', content: 'blah blah blah'},
    //     {id: '3', title: 'egg hunt with yoshi', content: 'blah blah blah'},
    // ]
    name: "",
    degreeYear: 0,
    gender: 0,
    major: 0,
    race: "",
    programmingGrade: 0,
    programmingExp: 0
};

const personalInfoReducer = (state = initState, action) => {
    switch (action.type) {
        case "RECORD_PERSONAL_INFO":
            return {
                ...state,
                name: action.personalInfo.name || "",
                degreeYear: action.personalInfo.degreeYear || 0,
                gender: action.personalInfo.gender || 0,
                major: action.personalInfo.major || 0,
                race: action.personalInfo.race || "",
                programmingGrade: action.personalInfo.programmingGrade || 0,
                programmingExp: action.personalInfo.programmingExp || 0
            };
        default:
            return state;
    }
};

export default personalInfoReducer;