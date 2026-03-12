import { BrowserRouter } from "react-router-dom";
import { AppShell } from "./appShell";

export default function App() {
  return (
    <BrowserRouter>
      <AppShell />
    </BrowserRouter>
  );
}