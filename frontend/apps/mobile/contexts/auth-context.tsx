import {
	GuardianLoginOutputBody,
	GuardianSignUpOutputBody,
	loginGuardianResponse,
	signupGuardianResponse,
	useLoginGuardian,
	useSignupGuardian,
	useGetGuardianById,
	Guardian,
	setCurrentLanguage,
	getGuardianById,
	getGetSavedByGuardianIdQueryKey,
	getUserByUsername,
	getUserByUsernameResponseError,
} from "@skillspark/api-client";
import * as SecureStore from "expo-secure-store";
import { router } from "expo-router";
import { createContext, useState, useEffect, ReactNode } from "react";
import { useQueryClient } from "@tanstack/react-query";
import i18n from "@/i18n";

interface AuthContextType {
	guardianId: string | null;
	jwt: string | null;
	langPref: string | null;
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
		profile_picture_s3_key: string | undefined,
		onError: (msg: string) => void,
	) => void;
	logout: () => void;
	usernameExists: (
		username: string,
		onError: (msg: string) => void,
	) => Promise<boolean>;
}

export const AuthContext = createContext<AuthContextType | undefined>(
	undefined,
);

export function AuthProvider({ children }: { children: ReactNode }) {
	const [guardianId, setGuardianId] = useState<string | null>(null);
	const [jwt, setJWT] = useState<string | null>(null);
	const [langPref, setLangPref] = useState<string | null>(null);
	const [isLoading, setIsLoading] = useState(true);
	const queryClient = useQueryClient();
	const { mutate: loginFunc } = useLoginGuardian();
	const { mutate: signupFunc } = useSignupGuardian();

	const { data: guardianData } = useGetGuardianById(guardianId!, {
		query: {
			enabled: !!guardianId,
		},
	});

	useEffect(() => {
		const checkAlreadyAuth = async () => {
			const storedJWT = await SecureStore.getItemAsync("token");
			const storedGuardianId = await SecureStore.getItemAsync("guardian_id");
			const storedLangPref = await SecureStore.getItemAsync(
				"language_preference",
			);
			if (storedLangPref) {
				await i18n.changeLanguage(storedLangPref);
				setCurrentLanguage(storedLangPref);
				queryClient.invalidateQueries({ refetchType: "all" });
			}
			if (storedJWT && storedGuardianId) {
				setJWT(storedJWT);
				setGuardianId(storedGuardianId);
				setLangPref(storedLangPref);
			}
			setIsLoading(false);
		};
		checkAlreadyAuth();
	}, [queryClient]);

	useEffect(() => {
		const getUpdatedLangPref = async () => {
			if (!guardianData) return;
			// SecureStore is the source of truth — only fall back to server value
			// if there is no locally stored preference (e.g. first login on a new device).
			const stored = await SecureStore.getItemAsync("language_preference");
			if (stored) {
				setLangPref(stored);
				return;
			}
			const guardian = (guardianData as unknown as { data: Guardian })?.data;
			const pref = guardian?.language_preference ?? "en";
			await i18n.changeLanguage(pref);
			setCurrentLanguage(pref);
			await SecureStore.setItemAsync("language_preference", pref);
			setLangPref(pref);
		};
		getUpdatedLangPref();
	}, [guardianData]);

	const login = (
		email: string,
		password: string,
		onError: (msg: string) => void,
	) => {
		loginFunc(
			{ data: { email, password } },
			{
				onSuccess: async (resp: loginGuardianResponse) => {
					const success = resp.data as GuardianLoginOutputBody;
					await SecureStore.setItemAsync("token", success.token);
					setJWT(success.token);
					await SecureStore.setItemAsync("guardian_id", success.guardian_id);
					setGuardianId(success.guardian_id);
					router.replace("/(app)/(tabs)");
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
		profile_picture_s3_key: string | undefined,
		onError: (msg: string) => void,
	) => {
		signupFunc(
			{
				data: {
					name,
					email,
					username,
					password,
					language_preference,
					profile_picture_s3_key,
				},
			},
			{
				onSuccess: async (resp: signupGuardianResponse) => {
					const success = resp.data as GuardianSignUpOutputBody;
					await SecureStore.setItemAsync("token", success.token);
					setJWT(success.token);
					await SecureStore.setItemAsync("guardian_id", success.guardian_id);
					setGuardianId(success.guardian_id);
					router.replace("/(app)/(tabs)");
				},
				onError: (err) => {
					const fail = err as unknown as { data?: { message?: string } };
					onError(fail.data?.message ?? "An unexpected error occurred");
				},
			},
		);
	};

	const logout = async () => {
		router.replace("/(auth)/login");
		await SecureStore.deleteItemAsync("token");
		setJWT(null);
		await SecureStore.deleteItemAsync("guardian_id");
		setGuardianId(null);
		await SecureStore.deleteItemAsync("language_preference");
		setLangPref(null);
	};

	// add update const for changing the guardian

	const usernameExists = async (
		username: string,
		onError: (msg: string) => void,
	) => {
		try {
			console.log(username, "Username");
			await getUserByUsername(username);
			onError("Username is taken.");
			return false;
		} catch (err) {
			console.log(err, "ERROR!!!");
			const typedErr = err as getUserByUsernameResponseError;
			if (typedErr.status === 404) {
				return true;
			} else {
				onError(typedErr.data.detail ?? "An unexpected error occurred.");
				return false;
			}
		}
	};

	return (
		<AuthContext.Provider
			value={{
				guardianId,
				jwt,
				langPref,
				isAuthenticated: !!(jwt && guardianId),
				isLoading,
				login,
				signup,
				logout,
				usernameExists,
			}}
		>
			{children}
		</AuthContext.Provider>
	);
}
