import "./App.css";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import StartingPage from "./StartingPage";

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<StartingPage />} />
      </Routes>
    </Router>
  );
}

export default App;
