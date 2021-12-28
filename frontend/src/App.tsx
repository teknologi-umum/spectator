import { BrowserRouter, Routes, Route } from "react-router-dom";
import PersonalInfo from "@/pages/PersonalInfo";
import Instructions from "@/pages/Instructions";
import SAMTest from "@/pages/SAMTest";
import CodingTest from "@/pages/CodingTest";
import FunFact from "./pages/FunFact";
import PrivateRoute from "./components/PrivateRoute";
import PublicRoute from "./components/PublicRoute";
import FinalRoute from "./components/FinalRoute";

function App() {
  // basically, the rules are:
  //  - if a user has a token and they haven't finished, go to /coding-test
  //  - if a user has a token and they have finished, go to /fun-fact
  //  - if a user doesn't have a token, go to /
  // TODO(elianiva): this could be wrong so, revisit this later
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/">
          <Route
            index
            element={
              <PublicRoute>
                <PersonalInfo />
              </PublicRoute>
            }
          />
          <Route
            path="instructions"
            element={
              <PublicRoute>
                <Instructions />
              </PublicRoute>
            }
          />
          <Route
            path="sam-test"
            element={
              <PublicRoute>
                <SAMTest />
              </PublicRoute>
            }
          />
          <Route
            path="coding-test"
            element={
              <PrivateRoute>
                <CodingTest />
              </PrivateRoute>
            }
          />
          <Route
            path="fun-fact"
            element={
              <FinalRoute>
                <FunFact />
              </FinalRoute>
            }
          />
        </Route>
      </Routes>
    </BrowserRouter>
  );
}

export default App;
