import { useAuthContext } from "@/hooks/use-auth-context";
import { Stack, router } from "expo-router";
import { useEffect } from "react";

export const LoginRedirect = () => {
	const { isAuthenticated, isLoading, hasAccount } = useAuthContext();

	useEffect(() => {
		if (isLoading) return;
		if (isAuthenticated) {
			router.replace("/(app)/(tabs)");
		} else if (hasAccount) {
			router.replace("/(auth)/login");
		} else {
			router.replace("/(auth)/signup");
		}
	}, [isAuthenticated, isLoading, hasAccount]);

	return (
		<Stack>
			<Stack.Screen name="(auth)" options={{ headerShown: false }} />
			<Stack.Screen name="(app)" options={{ headerShown: false }} />
		</Stack>
	);
};
