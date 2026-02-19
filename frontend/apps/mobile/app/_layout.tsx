import { Redirect, router, Stack } from "expo-router";
import "react-native-reanimated";

import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { Image, StatusBar } from "react-native";

const queryClient = new QueryClient({
    defaultOptions: {
      queries: {
        staleTime: 1000 * 60 * 5,
        retry: 1,
      },
    },
  });

const LoginOrHome = () => {
  const isAuth = false;

  return (
    <>
        <Stack/>
        {!isAuth && <Redirect href="/(auth)/login" />}
        {isAuth && <Redirect href="/(app)/(tabs)" />}
    </>
  );
}

export default function RootToLoginOrHome() {
  return (
      <QueryClientProvider client={queryClient}>
        <StatusBar />
        <LoginOrHome />
      </QueryClientProvider>
  );
}

