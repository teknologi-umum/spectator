import personalInfoReducer from "@/store/reducers/personalInfoReducer";
import questionReducer from "@/store/reducers/questionReducer";
import { combineReducers } from "redux";

const rootReducer = combineReducers({
  personalInfo: personalInfoReducer,
  question: questionReducer
});

export default rootReducer;
