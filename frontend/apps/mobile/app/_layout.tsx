import { Redirect, Stack } from "expo-router";
import "react-native-reanimated";
import "../global.css";
import { useColorScheme } from "@/hooks/use-color-scheme";
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { AuthProvider } from '@/contexts/auth-context';
import { useAuthContext } from "@/hooks/use-auth-context";

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

const LoginOrHome = () => {
  const { isAuthenticated } = useAuthContext();

  return (
    <>
      <Stack>
        <Stack.Screen name="(auth)" options={{ headerShown: false }} />
        {!isAuthenticated && <Redirect href="/(auth)/login" />}
        <Stack.Screen name="(app)" options={{ headerShown: false }} />
        {isAuthenticated && <Redirect href="/(app)/(tabs)" />}
      </Stack>
    </>
  );
}

export default function RootLayout() {  
  return (
      <QueryClientProvider client={queryClient}>
        <AuthProvider>
          <LoginOrHome />
        </AuthProvider>
      </QueryClientProvider>
  );
}