import { createRoot } from "react-dom/client";
import "./index.css";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import Login from "./login/login.tsx";
import { AppShell } from "./App.tsx";

createRoot(document.getElementById("root")!).render(
  <BrowserRouter>
    <Routes>
      <Route path="/*" element={<AppShell />} />
      <Route path="/login" element={<Login />} />
    </Routes>
  </BrowserRouter>,
);
