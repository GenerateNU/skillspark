import "react-native-reanimated";
import "../global.css";
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { AuthProvider } from '@/contexts/auth-context';
import { LoginRedirect } from "@/components/LoginRedirect";

export const unstable_settings = {
  anchor: "(auth)",
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
  return (
      <QueryClientProvider client={queryClient}>
        <AuthProvider>
          <LoginRedirect />
        </AuthProvider>
      </QueryClientProvider>
  );
}