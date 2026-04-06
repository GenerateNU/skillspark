import { createRoot } from "react-dom/client";
import "./index.css";
import App from "./App.tsx";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import Login from "./login/login.tsx";
import { AuthProvider } from "./contexts/auth-context.tsx";
import { QueryClientProvider } from "@tanstack/react-query";
import { queryClient } from "./constants/query-client.ts";
import SignUp from "./signup/signup.tsx";

createRoot(document.getElementById("root")!).render(
	<QueryClientProvider client={queryClient}>
		<BrowserRouter>
			<AuthProvider>
				<Routes>
					<Route path="/" element={<App />} />
					<Route path="/login" element={<Login />} />
					<Route path="/signup" element={<SignUp />} />
				</Routes>
			</AuthProvider>
		</BrowserRouter>
	</QueryClientProvider>,
);
