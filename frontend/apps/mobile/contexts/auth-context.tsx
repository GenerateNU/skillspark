import { useGuardian } from "@/hooks/use-guardian";
import i18n from "@/i18n";
import {
	createStripeCustomer,
	GuardianLoginOutputBody,
	GuardianSignUpOutputBody,
	loginGuardianResponse,
	setCurrentLanguage,
	signupGuardianResponse,
	updateGuardianResponse,
	useLoginGuardian,
	useSignupGuardian,
	useUpdateGuardian,
	usernameExists as checkUsernameExists,
	usernameExistsResponseError,
	UsernameExistsOutputBody,
} from "@skillspark/api-client";
import * as SecureStore from "expo-secure-store";
import { createContext, useState, useEffect, ReactNode } from "react";
import { useQueryClient } from "@tanstack/react-query";
import { useGuardian } from "@/hooks/use-guardian";
import i18n from "@/i18n";
import { useTranslation } from "react-i18next";

interface AuthContextType {
	guardianId: string | null;
	jwt: string | null;
	langPref: string | null;
	isAuthenticated: boolean;
	isLoading: boolean;
	hasAccount: boolean;
	inOnboarding: boolean;
	login: (
		email: string,
		password: string,
		onError: (msg: string) => void,
		onSuccess: () => void,
	) => void;
	signup: (
		name: string,
		email: string,
		username: string,
		password: string,
		language_preference: string,
		profile_picture_s3_key: string | undefined,
		onError: (msg: string) => void,
		onSuccess: () => void,
	) => void;
	logout: () => void;
	update: (
		onSuccess: () => void,
		onError: (msg: string) => void,
		id: string,
		email: string,
		language_preference: string,
		name: string,
		username: string,
		profile_picture_s3_key?: string | undefined,
		expo_push_token?: string | undefined,
		push_notifications?: boolean,
		email_notifications?: boolean,
	) => void;
	usernameExists: (
		username: string,
		onError: (msg: string) => void,
	) => Promise<boolean>;
	setLanguage: (language: string) => void;
	setInOnboarding: (value: boolean) => void;
	completeOnboarding: () => Promise<void>;
}

export const AuthContext = createContext<AuthContextType | undefined>(
	undefined,
);

export function AuthProvider({ children }: { children: ReactNode }) {
	const [guardianId, setGuardianId] = useState<string | null>(null);
	const [jwt, setJWT] = useState<string | null>(null);
	const [langPref, setLangPref] = useState<string | null>(null);
	const [isLoading, setIsLoading] = useState(false);
	const [hasAccount, setHasAccount] = useState(false);
	const [inOnboarding, setInOnboarding] = useState(false);
	const queryClient = useQueryClient();
	const { mutate: loginFunc } = useLoginGuardian();
	const { mutate: signupFunc } = useSignupGuardian();
	const { mutate: updateFunc } = useUpdateGuardian();
	const { t: translate } = useTranslation();

	const { guardian, hasError } = useGuardian(guardianId);

	const logout = async () => {
		await SecureStore.deleteItemAsync("token");
		setJWT(null);
		await SecureStore.deleteItemAsync("guardian_id");
		setGuardianId(null);
		await SecureStore.deleteItemAsync("language_preference");
		setLangPref(null);
		console.log(hasAccount);
	};

	useEffect(() => {
		const checkAlreadyAuth = async () => {
			if (hasError) {
				// invalid JWT, log out old session
				logout();
				return;
			}
			const storedLangPref = await SecureStore.getItemAsync(
				"language_preference",
			);
			if (storedLangPref) {
				await i18n.changeLanguage(storedLangPref);
				setCurrentLanguage(storedLangPref);
				queryClient.invalidateQueries({ refetchType: "all" });
			}
			const storedHasAccount = await SecureStore.getItemAsync("has_account");
			setHasAccount(storedHasAccount === "true");
			console.log("hasAccount: ", storedHasAccount);
			if (storedHasAccount) {
				// values set once onboarding is finished
				const storedJWT = await SecureStore.getItemAsync("token");
				const storedGuardianId = await SecureStore.getItemAsync("guardian_id");
				if (storedJWT && storedGuardianId) {
					setJWT(storedJWT);
					setGuardianId(storedGuardianId);
					setLangPref(storedLangPref);
					console.log(
						"everything set: ",
						storedJWT,
						storedGuardianId,
						storedLangPref,
					);
				}
			}
			setIsLoading(false);
		};
		checkAlreadyAuth();
	}, [queryClient, hasError]);

	useEffect(() => {
		const getUpdatedLangPref = async () => {
			if (!guardian) return;
			const pref = guardian.language_preference ?? "en";
			await i18n.changeLanguage(pref);
			setCurrentLanguage(pref);
			await SecureStore.setItemAsync("language_preference", pref);
			setLangPref(pref);
			queryClient.invalidateQueries({ refetchType: "all" });
		};
		getUpdatedLangPref();
	}, [guardian, queryClient]);

	const login = (
		email: string,
		password: string,
		onError: (msg: string) => void,
		onSuccess: () => void,
	) => {
		console.log(hasAccount);
		loginFunc(
			{ data: { email, password } },
			{
				onSuccess: async (resp: loginGuardianResponse) => {
					const success = resp.data as GuardianLoginOutputBody;
					await SecureStore.setItemAsync("token", success.token);
					setJWT(success.token);
					await SecureStore.setItemAsync("guardian_id", success.guardian_id);
					setGuardianId(success.guardian_id);
					onSuccess();
				},
				onError: (err) => {
					const fail = err as unknown as { data?: { message?: string } };
					onError(fail.data?.message ?? translate("common.errorOccurred"));
				},
			},
		);
	};

	const completeOnboarding = async () => {
		await SecureStore.setItemAsync("has_account", "true");
		setHasAccount(true);
		setInOnboarding(false);
	};

	const signup = (
		name: string,
		email: string,
		username: string,
		password: string,
		language_preference: string,
		profile_picture_s3_key: string | undefined,
		onError: (msg: string) => void,
		onSuccess: () => void,
	) => {
		// entering onboarding -> no redirects until onboarding is complete
		setInOnboarding(true);
		// reset for each new account during onboarding
		setHasAccount(false);
		SecureStore.setItemAsync("has_account", "false");
		console.log(
			"values: {",
			name,
			email,
			username,
			password,
			language_preference,
			profile_picture_s3_key,
			"}",
		);
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
					console.log("response: ", JSON.stringify(resp));
					await SecureStore.setItemAsync("token", success.token);
					setJWT(success.token);
					await SecureStore.setItemAsync("guardian_id", success.guardian_id);
					setGuardianId(success.guardian_id);
					await createStripeCustomer(success.guardian_id);
					onSuccess();
				},
				onError: (err) => {
					console.log(JSON.stringify(err));
					const fail = err as unknown as { data?: { message?: string } };
					onError(fail.data?.message ?? translate("common.errorOccurred"));
				},
			},
		);
	};

	const update = (
		onSuccess: () => void,
		onError: (msg: string) => void,
		id: string,
		email: string,
		language_preference: string,
		name: string,
		username: string,
		profile_picture_s3_key?: string | undefined,
		expo_push_token?: string | undefined,
		push_notifications?: boolean,
		email_notifications?: boolean,
	) => {
		updateFunc(
			{
				id: id,
				data: {
					email,
					language_preference,
					name,
					profile_picture_s3_key,
					username,
					expo_push_token,
					push_notifications,
					email_notifications,
				},
			},
			{
				onSuccess: async (_resp: updateGuardianResponse) => {
					// refetch all getGuardian queries to show changes
					queryClient.invalidateQueries({
						queryKey: [`/api/v1/guardians/${id}`],
					});
					onSuccess();
				},
				onError: (err) => {
					const fail = err as unknown as { data?: { message?: string } };
					onError(fail.data?.message ?? "An unexpected error occurred");
				},
			},
		);
	};

	const usernameExists = async (
		username: string,
		onError: (msg: string) => void,
	) => {
		try {
			const resp = await checkUsernameExists(username);
			const data = resp.data as UsernameExistsOutputBody;
			if (data.exists) {
				onError("Username is taken.");
				return false;
			}
			return true;
		} catch (err) {
			const typedErr = err as usernameExistsResponseError;
			onError(typedErr.data?.detail ?? "An unexpected error occurred.");
			return false;
		}
	};

	return (
		<AuthContext.Provider
			value={{
				guardianId,
				jwt,
				langPref,
				isAuthenticated: !!jwt && !!guardianId,
				isLoading,
				hasAccount,
				inOnboarding,
				login,
				signup,
				logout,
				update,
				usernameExists,
				setLanguage,
				setInOnboarding,
				completeOnboarding,
			}}
		>
			{children}
		</AuthContext.Provider>
	);
}
