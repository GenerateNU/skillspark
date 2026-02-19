import { Redirect, router, Stack } from "expo-router";
import "react-native-reanimated";
import { useColorScheme } from "@/hooks/use-color-scheme";
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';

export const unstable_settings = {
  anchor: "(tabs)",
};

// Create QueryClient outside the component
const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      staleTime: 1000 * 60 * 5,
      retry: 1,
    },
  },
});

export default function RootLayout() {
  const colorScheme = useColorScheme();
  
  return (
      <QueryClientProvider client={queryClient}>
        <StatusBar />
        <LoginOrHome />
      </QueryClientProvider>
  );
}