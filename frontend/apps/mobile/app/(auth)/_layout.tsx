import { DarkTheme, DefaultTheme, ThemeProvider } from "@react-navigation/native";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { Stack } from "expo-router";
import { useColorScheme } from "react-native";
import { StatusBar } from "expo-status-bar";


export const unstable_settings = {
  anchor: "login",
};

export default function AuthStackLayout() {
  const queryClient = new QueryClient({
    defaultOptions: {
      queries: {
        staleTime: 1000 * 60 * 5,
        retry: 1,
      },
    },
  });

  return (
      <QueryClientProvider client={queryClient}>
        <Stack>
          <Stack.Screen name="login/index" options={{ headerShown: false }} />
          <Stack.Screen name="signup/index" options={{ headerShown: false }} />
        </Stack>
        <StatusBar style="auto" />
      </QueryClientProvider>
  );
}