import { BrowserRouter, Routes, Route } from "react-router-dom";
import Layout from "@/components/Layout";
import Countdown from "@/components/Countdown";
import PersonalInfo from "@/pages/PersonalInfo";
import Instructions from "@/pages/Instructions";
import SAMTest from "@/pages/SAMTest";

const DURATION = 90 * 60; // 90 minutes

function App() {
  return (
    <BrowserRouter>
      <Countdown duration={DURATION} />
      <Layout>
        <Routes>
          <Route path="/">
            <Route index element={<PersonalInfo />} />
            <Route path="instructions" element={<Instructions />} />
            <Route path="sam-test" element={<SAMTest />} />
          </Route>
        </Routes>
      </Layout>
    </BrowserRouter>
  );
}

export default App;
