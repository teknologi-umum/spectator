import { BrowserRouter, Routes, Route } from "react-router-dom";
import LastPage from "@/pages/LastPage";
import Instructions from "@/pages/Instructions/Instructions";
import SAMTest from "@/pages/SAMTest/SAMTest";
import Countdown from "@/components/Countdown";
import PersonalInfo from "@/pages/PersonalInfo";
// import CodingTest from "@/pages/CodingTest";
import Layout from "@/components/Layout";

const DURATION = 90 * 60; // 90 minutes

export default function App() {
  return (
    <BrowserRouter>
      <Countdown duration={DURATION} />
      <Layout>
        <Routes>
          <Route path="/">
            <Route index element={<PersonalInfo />} />
            <Route path="sam-test" element={<SAMTest />} />
            <Route path="instructions" element={<Instructions />} />
            <Route path="last" element={<LastPage />} />
            {/* <Route path="test" element={<CodingTest />} /> */}
          </Route>
        </Routes>
      </Layout>
    </BrowserRouter>
  );
}
