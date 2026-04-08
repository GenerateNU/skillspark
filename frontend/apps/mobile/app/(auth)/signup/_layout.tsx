import { QueryClientProvider } from "@tanstack/react-query";
import { Stack } from "expo-router";
import { StatusBar } from "expo-status-bar";

import { queryClient } from "@/constants/query-client";

export default function OnboardingStackLayout() {
	return (
		<QueryClientProvider client={queryClient}>
			<Stack>
				<Stack.Screen name="index" options={{ headerShown: false }} />
				<Stack.Screen name="name" options={{ headerShown: false }} />
				<Stack.Screen name="photo" options={{ headerShown: false }} />
				<Stack.Screen
					name="child-profile/index"
					options={{ headerShown: false }}
				/>
				<Stack.Screen
					name="emergency-contact"
					options={{ headerShown: false }}
				/>
			</Stack>
			<StatusBar style="auto" />
		</QueryClientProvider>
	);
}
