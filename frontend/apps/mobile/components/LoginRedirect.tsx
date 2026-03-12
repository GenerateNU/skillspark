import { useAuthContext } from "@/hooks/use-auth-context";
import { Redirect, Stack } from "expo-router";

export const LoginRedirect = () => {
  const { isAuthenticated, isLoading } = useAuthContext();

  if (isLoading) {
    return (
      <Stack />
    );
  }

  return (
    <>
      <Stack>
        <Stack.Screen name="(auth)" options={{ headerShown: false }} />
        <Stack.Screen name="(app)" options={{ headerShown: false }} />
      </Stack>
      {!isAuthenticated && <Redirect href="/(auth)/login" />}
      {isAuthenticated && <Redirect href="/(app)/(tabs)" />}
    </>
  );
}