import { createRoot } from "react-dom/client";
import "./index.css";
import App from "./App.tsx";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import Login from "./login/login.tsx";
import Admin from "./admin/admin.tsx";
import SignUp from "./signup/signup.tsx";

createRoot(document.getElementById("root")!).render(
  <BrowserRouter>
    <Routes>
      <Route path="/" element={<App />} />
      <Route path="/login" element={<Login />} />
      <Route path="/admin/*" element={<Admin />} />
      <Route path="/signup" element={<SignUp />} />
    </Routes>
  </BrowserRouter>,
);
