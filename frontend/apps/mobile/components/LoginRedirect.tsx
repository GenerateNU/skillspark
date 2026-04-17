import { useAuthContext } from "@/hooks/use-auth-context";
import { Redirect, Stack } from "expo-router";
import { useEffect, useState } from "react";
import * as Linking from "expo-linking";

export const LoginRedirect = () => {
	const { isAuthenticated, isLoading } = useAuthContext();
	const [initialUrl, setInitialUrl] = useState<string | null | undefined>(
		undefined,
	);

	useEffect(() => {
		Linking.getInitialURL()
			.then((url) => setInitialUrl(url ?? null))
			.catch(() => setInitialUrl(null));
	}, []);

	if (isLoading || initialUrl === undefined) {
		return <Stack />;
	}

	return (
		<>
			<Stack>
				<Stack.Screen name="(auth)" options={{ headerShown: false }} />
				<Stack.Screen name="(app)" options={{ headerShown: false }} />
			</Stack>
			{!isAuthenticated && <Redirect href="/(auth)/login" />}
			{isAuthenticated && !initialUrl && <Redirect href="/(app)/(tabs)" />}
		</>
	);
};
