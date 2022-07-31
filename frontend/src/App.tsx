import React from "react";
import { BrowserRouter, Routes, Route } from "react-router-dom";

const Instructions = React.lazy(() => import("@/pages/Instructions"));
const PersonalInfo = React.lazy(() => import("@/pages/PersonalInfoPage"));
const SAMTest = React.lazy(() => import("@/pages/SAMTest"));
const CodingTest = React.lazy(() => import("@/pages/CodingTest"));
const FunFact = React.lazy(() => import("@/pages/FunFact"));
const Login = React.lazy(() => import("@/pages/secret/Login"));
const Download = React.lazy(() => import("@/pages/secret/Download"));
const CoercedRoute = React.lazy(() => import("@/hoc/CoercedRoute"));
const SecretRoute = React.lazy(() => import("@/hoc/SecretRoute"));

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<CoercedRoute />}>
          <Route index element={<PersonalInfo />} />
          <Route path="instructions" element={<Instructions />} />
          <Route path="sam-test" element={<SAMTest />} />
          <Route path="coding-test" element={<CodingTest />} />
          <Route path="fun-fact" element={<FunFact />} />
        </Route>
        <Route path="/secret" element={<SecretRoute />}>
          <Route path="login" element={<Login />} />
          <Route path="download" element={<Download />} />
        </Route>
      </Routes>
    </BrowserRouter>
  );
}

export default App;
