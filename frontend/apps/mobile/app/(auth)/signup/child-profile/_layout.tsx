import { QueryClientProvider } from "@tanstack/react-query";
import { Stack } from "expo-router";
import { StatusBar } from "expo-status-bar";

import { queryClient } from "@/constants/query-client";

export default function AddChildLayout() {
	return (
		<QueryClientProvider client={queryClient}>
			<Stack>
				<Stack.Screen name="index" options={{ headerShown: false }} />
				<Stack.Screen name="add-child" options={{ headerShown: false }} />
				<Stack.Screen name="edit-pic" options={{ headerShown: false }} />
				<Stack.Screen name="avatar-picker" options={{ headerShown: false }} />
			</Stack>
			<StatusBar style="auto" />
		</QueryClientProvider>
	);
}
