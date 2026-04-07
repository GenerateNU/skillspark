import type {
	ManagerLoginOutputBody,
	ManagerSignUpOutputBody,
	loginManagerResponse,
	signupManagerResponse,
} from "@skillspark/api-client";
import { useLoginManager, useSignupManager } from "@skillspark/api-client";
import { useNavigate } from "react-router-dom";
import { createContext, useEffect, useState } from "react";

interface AuthContextType {
	managerId: string | null;
	jwt: string | null;
	isAuthenticated: boolean;
	isLoading: boolean;
	login: (
		email: string,
		password: string,
		onError: (msg: string) => void,
	) => void;
	signup: (
		name: string,
		email: string,
		username: string,
		password: string,
		language_preference: string,
		organization_id: string,
		role: string,
		profile_picture_s3_key: string | undefined,
		onError: (msg: string) => void,
	) => void;
	logout: () => void;
}

export const AuthContext = createContext<AuthContextType | undefined>(
	undefined,
);

export function AuthProvider({ children }: { children: React.ReactNode }) {
	const [managerId, setManagerId] = useState<string | null>(null);
	const [jwt, setJWT] = useState<string | null>(null);
	const [isLoading, setIsLoading] = useState(true);
	const { mutate: loginFunc } = useLoginManager();
	const { mutate: signupFunc } = useSignupManager();
	const navigate = useNavigate();

	useEffect(() => {
		const checkAlreadyAuth = () => {
			const storedManagerId = localStorage.getItem("manager_id");
			const storedJwt = localStorage.getItem("jwt");

			if (storedJwt && storedManagerId) {
				setJWT(storedJwt);
				setManagerId(storedManagerId);
			}
			setIsLoading(false);
		};
		checkAlreadyAuth();
	}, [navigate]);

	const login = (
		email: string,
		password: string,
		onError: (msg: string) => void,
	) => {
		loginFunc(
			{ data: { email, password } },
			{
				onSuccess: (resp: loginManagerResponse) => {
					const success = resp.data as ManagerLoginOutputBody;
					localStorage.setItem("jwt", success.token);
					setJWT(success.token);
					localStorage.setItem("manager_id", success.manager_id);
					setManagerId(success.manager_id);
					navigate("/");
				},
				onError: (err) => {
					const fail = err as unknown as { data?: { message?: string } };
					onError(fail.data?.message ?? "An unexpected error occurred");
				},
			},
		);
	};

	const signup = (
		name: string,
		email: string,
		username: string,
		password: string,
		language_preference: string,
		organization_id: string,
		role: string,
		profile_picture_s3_key: string | undefined,
		onError: (msg: string) => void,
	) => {
		signupFunc(
			{
				data: {
					email,
					password,
					language_preference,
					name,
					organization_id,
					profile_picture_s3_key,
					role,
					username,
				},
			},
			{
				onSuccess: (resp: signupManagerResponse) => {
					const success = resp.data as ManagerSignUpOutputBody;
					localStorage.setItem("jwt", success.token);
					setJWT(success.token);
					localStorage.setItem("manager_id", success.manager_id);
					setManagerId(success.manager_id);
					navigate("/");
				},
				onError: (err) => {
					const fail = err as unknown as { data?: { message?: string } };
					onError(fail.data?.message ?? "An unexpected error occurred");
				},
			},
		);
	};

	const logout = () => {
		localStorage.removeItem("jwt");
		localStorage.removeItem("manager_id");
		setJWT(null);
		setManagerId(null);
		navigate("/login");
	};

	return (
		<AuthContext.Provider
			value={{
				managerId: managerId,
				jwt,
				isAuthenticated: !!(jwt && managerId),
				isLoading,
				login,
				signup,
				logout,
			}}
		>
			{children}
		</AuthContext.Provider>
	);
}
