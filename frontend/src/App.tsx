import { BrowserRouter, Routes, Route } from "react-router-dom";
import Countdown from "@/components/Countdown";
import PersonalInfo from "@/pages/PersonalInfo";
import Instructions from "@/pages/Instructions";
import SAMTest from "@/pages/SAMTest";
import CodingTest from "@/pages/CodingTest";
import {useColorModeValue } from "@chakra-ui/react";
import ThemeButton from "./components/ThemeButton";

const DURATION = 90 * 60; // 90 minutes

function App() {
  const BackgroundColor = useColorModeValue('white', 'gray.700')
  const Color = useColorModeValue('gray.800', 'white')

  return (
    <BrowserRouter>
      <Countdown duration={DURATION} />
      <ThemeButton />
      <Routes>
        <Route path="/">
          <Route index element={<PersonalInfo background={BackgroundColor} color={Color}/>} />
          <Route path="instructions" element={<Instructions background={BackgroundColor} color={Color}/>} />
          <Route path="sam-test" element={<SAMTest background={BackgroundColor} color={Color} />} />
          <Route path="coding-test" element={<CodingTest />} />
        </Route>
      </Routes>
    </BrowserRouter>
  );
}

export default App;
