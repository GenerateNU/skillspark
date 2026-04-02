import { createRoot } from "react-dom/client";
import "./index.css";
import App from "./App.tsx";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import Login from "./login/login.tsx";
import SignUp from "./signup/signup.tsx";
import { AuthProvider } from "../contexts/auth-context.tsx";

createRoot(document.getElementById("root")!).render(
	<BrowserRouter>
		<Routes>
			<Route path="/" element={<App />} />
			<Route path="/login" element={<Login />} />
			<Route path="/signup" element={<SignUp />} />
		</Routes>
	</BrowserRouter>,
);
