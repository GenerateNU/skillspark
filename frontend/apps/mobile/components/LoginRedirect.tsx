import { useAuthContext } from "@/hooks/use-auth-context";
import { Stack, router } from "expo-router";
import { useEffect } from "react";

export const LoginRedirect = () => {
	const { isAuthenticated, isLoading, hasAccount, inOnboarding, logout } =
		useAuthContext();

	useEffect(() => {
		if (isLoading) return;
		if (inOnboarding) return;

		if (isAuthenticated && hasAccount) {
			router.replace("/(app)/(tabs)");
		} else if (!isAuthenticated && hasAccount) {
			router.replace("/(auth)/login");
		} else if (!isAuthenticated && !hasAccount) {
			router.replace("/(auth)/signup");
		} else {
			logout();
		}
	}, [isAuthenticated, isLoading, hasAccount, inOnboarding]);

	return (
		<Stack>
			<Stack.Screen name="(auth)" options={{ headerShown: false }} />
			<Stack.Screen name="(app)" options={{ headerShown: false }} />
		</Stack>
	);
};
