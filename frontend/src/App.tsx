import React from "react";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import PersonalInfo from "@/pages/PersonalInfoPage";
import Instructions from "@/pages/Instructions";
import SAMTest from "@/pages/SAMTest";
import CodingTest from "@/pages/CodingTest";
import FunFact from "@/pages/FunFact";
import { CoercedRoute } from "@/hoc/CoercedRoute";

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
          <CoercedRoute index element={<PersonalInfo />} />
          <CoercedRoute path="instructions" element={<Instructions />} />
          <CoercedRoute path="sam-test" element={<SAMTest />} />
          <CoercedRoute path="coding-test" element={<CodingTest />} />
          <CoercedRoute path="fun-fact" element={<FunFact />} />
        </Route>
      </Routes>
    </BrowserRouter>
  );
}

export default App;
