import { BrowserRouter, Routes, Route } from "react-router-dom";
import PersonalInfo from "@/pages/PersonalInfo";
import Instructions from "@/pages/Instructions";
import SAMTest from "@/pages/SAMTest";
import CodingTest from "@/pages/CodingTest";
import ThemeButton from "./components/ThemeButton";

function App() {
  return (
    <BrowserRouter>
      <ThemeButton />
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
