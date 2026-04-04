import {
	DarkTheme,
	DefaultTheme,
	ThemeProvider,
} from "@react-navigation/native";
import { Stack } from "expo-router";
import { StatusBar } from "expo-status-bar";
import "react-native-reanimated";

import { useColorScheme } from "@/hooks/use-color-scheme";
import { queryClient } from "@/constants/query-client";
import { QueryClientProvider } from "@tanstack/react-query";
import { useAuthContext } from "@/hooks/use-auth-context";

export const unstable_settings = {
	anchor: "(tabs)",
};

export default function RootLayout() {
	const colorScheme = useColorScheme();

	const { isAuthenticated, isLoading } = useAuthContext();

	if (isLoading) return null;

	if (!isAuthenticated) {
		return;
	}

	return (
		<ThemeProvider value={colorScheme === "dark" ? DarkTheme : DefaultTheme}>
			<QueryClientProvider client={queryClient}>
				<Stack>
					<Stack.Screen name="(tabs)" options={{ headerShown: false }} />
					<Stack.Screen
						name="modal"
						options={{ presentation: "modal", title: "Modal" }}
					/>
				</Stack>
				<StatusBar style="auto" />
			</QueryClientProvider>
		</ThemeProvider>
	);
}
