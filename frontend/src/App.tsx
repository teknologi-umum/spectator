import React from "react";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import PersonalInfo from "@/pages/PersonalInfoPage";
import Instructions from "@/pages/Instructions";
import SAMTest from "@/pages/SAMTest";
import CodingTest from "@/pages/CodingTest";
import FunFact from "@/pages/FunFact";
import { CoercedRoute } from "@/hoc/CoercedRoute";

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<CoercedRoute />}>
          <Route index element={<PersonalInfo />} />
          <Route path="sam-test" element={<SAMTest />} />
          <Route path="fun-fact" element={<FunFact />} />
        </Route>
        <Route path="instructions" element={<Instructions />} />
        <Route path="coding-test" element={<CodingTest />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
