import "./App.css";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import StartingPage from "./StartingPage";
import LandingPage from "./LandingPage";

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<StartingPage />} />
        <Route path="/main" element={<LandingPage />} />
      </Routes>
    </Router>
  );
}

export default App;
