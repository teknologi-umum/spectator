const initState = {
  name: "",
  degreeYear: 0,
  gender: 0,
  major: 0,
  race: "",
  programmingGrade: 0,
  programmingExp: 0
};

// FIXME: proper action interface
// eslint-disable-next-line
const personalInfoReducer = (state = initState, action: any) => {
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

