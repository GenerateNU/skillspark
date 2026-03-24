import { useAuthContext } from "@/hooks/use-auth-context";
import { Redirect, Stack } from "expo-router";
import { useEffect } from "react";
import i18n from "@/i18n";

export const LoginRedirect = () => {
  const { isAuthenticated, isLoading, langPref } = useAuthContext();

  useEffect(() => {
    if (langPref) {
      i18n.changeLanguage(langPref);
    }
  }, [langPref]);

  if (isLoading) {
    return <Stack />;
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
};
