import React from "react";
import { BrowserRouter, Routes, Route } from "react-router-dom";

const Instructions = React.lazy(() => import("@/pages/InstructionsPage"));
const PersonalInfo = React.lazy(() => import("@/pages/PersonalInfoPage"));
const SAMTest = React.lazy(() => import("@/pages/SAMTestPage"));
const CodingTest = React.lazy(() => import("@/pages/CodingTestPage"));
const FunFact = React.lazy(() => import("@/pages/FunFactPage"));
const VideoTest = React.lazy(() => import("@/pages/VideoTestPage"));
const Login = React.lazy(() => import("@/pages/secret/LoginPage"));
const Download = React.lazy(() => import("@/pages/secret/DownloadPage"));
const CoercedRoute = React.lazy(() => import("@/hoc/CoercedRoute"));
const SecretRoute = React.lazy(() => import("@/hoc/SecretRoute"));

function App() {
  return (
    <BrowserRouter>
      <Routes>
          <Route path="sam-test" element={<SAMTest />} />
          <Route path="video-test" element={<VideoTest />} />
          <Route path="coding-test" element={<CodingTest />} />
          <Route path="fun-fact" element={<FunFact />} />
        <Route path="/" element={<CoercedRoute />}>
          <Route index element={<PersonalInfo />} />
          <Route path="instructions" element={<Instructions />} />
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
