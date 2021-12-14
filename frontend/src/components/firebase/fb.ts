// Import the functions you need from the SDKs you need
import { initializeApp } from "firebase/app";
import { getAnalytics } from "firebase/analytics";
// TODO: Add SDKs for Firebase products that you want to use
// https://firebase.google.com/docs/web/setup#available-libraries

// Your web app's Firebase configuration
// For Firebase JS SDK v7.20.0 and later, measurementId is optional
const firebaseConfig = {
  apiKey: "AIzaSyCxwCA2DmUk6geit3fyyNwxbdTJG3oVeWA",
  authDomain: "codingtest-1234.firebaseapp.com",
  projectId: "codingtest-1234",
  storageBucket: "codingtest-1234.appspot.com",
  messagingSenderId: "412829956011",
  appId: "1:412829956011:web:66e38a639c5fbac1a4c4ee",
  measurementId: "G-C0N352CM3K"
};

// Initialize Firebase
const app = initializeApp(firebaseConfig);
const analytics = getAnalytics(app);