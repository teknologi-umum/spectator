import { BrowserRouter, Routes, Route } from "react-router-dom";
import Countdown from "@/components/Countdown";
import PersonalInfo from "@/pages/PersonalInfo";
import Instructions from "@/pages/Instructions";
import SAMTest from "@/pages/SAMTest";
import CodingTest from "@/pages/CodingTest";

const DURATION = 90 * 60; // 90 minutes

function App() {
  return (
    <BrowserRouter>
      <Countdown duration={DURATION} />
      <Routes>
        <Route path="/">
          <Route index element={<PersonalInfo />} />
          <Route path="instructions" element={<Instructions />} />
          <Route path="sam-test" element={<SAMTest />} />
          <Route path="coding-test" element={<CodingTest />} />
        </Route>
      </Routes>
    </BrowserRouter>
  );
}

export default App;
