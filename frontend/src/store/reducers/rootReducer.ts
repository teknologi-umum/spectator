import personalInfoReducer from "./personalInfoReducer";
import questionReducer from "./questionReducer";
import { combineReducers } from "redux";
// import { firestoreReducer } from 'redux-firestore'
// import {firebaseReducer} from 'react-redux-firebase'

const rootReducer = combineReducers({
    personalInfo: personalInfoReducer,
    question: questionReducer
    // firestore: firestoreReducer,
    // firebase: firebaseReducer
});

export default rootReducer;