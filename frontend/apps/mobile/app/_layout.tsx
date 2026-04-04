import {
  DarkTheme,
  DefaultTheme,
  ThemeProvider,
} from "@react-navigation/native";
import { useFonts } from "expo-font";
import * as SplashScreen from "expo-splash-screen";
import * as SecureStore from "expo-secure-store";
import { useEffect, useState } from "react";
import "react-native-reanimated";
import { QueryClientProvider } from "@tanstack/react-query";
import { queryClient } from "@/constants/query-client";
import { useColorScheme } from "@/hooks/use-color-scheme";
import "../global.css";
import {
  NunitoSans_400Regular,
  NunitoSans_500Medium,
  NunitoSans_600SemiBold,
  NunitoSans_700Bold,
} from "@expo-google-fonts/nunito-sans";
import { AuthProvider } from "@/contexts/auth-context";
import { LoginRedirect } from "@/components/LoginRedirect";
import { setCurrentLanguage } from "@skillspark/api-client";

let StripeProvider: React.ComponentType<{
  publishableKey: string;
  children: React.ReactNode;
}> | null = null;
try {
  StripeProvider = require("@stripe/stripe-react-native").StripeProvider;
} catch {
  // Native module unavailable (e.g. Expo Go). Skip Stripe
}

SplashScreen.preventAutoHideAsync();

export default function RootLayout() {
  const colorScheme = useColorScheme();
  const [langReady, setLangReady] = useState(false);
  const [loaded, error] = useFonts({
    NunitoSans_400Regular,
    NunitoSans_500Medium,
    NunitoSans_600SemiBold,
    NunitoSans_700Bold,
  });

  useEffect(() => {
    SecureStore.getItemAsync("language_preference").then((lang) => {
      if (lang) setCurrentLanguage(lang);
      setLangReady(true);
    });
  }, []);

  useEffect(() => {
    if ((loaded || error) && langReady) {
      SplashScreen.hideAsync();
    }
  }, [loaded, error, langReady]);

  if ((!loaded && !error) || !langReady) {
    return null;
  }

  const content = (
    <QueryClientProvider client={queryClient}>
      <ThemeProvider value={colorScheme === "dark" ? DarkTheme : DefaultTheme}>
        <AuthProvider>
          <LoginRedirect />
        </AuthProvider>
      </ThemeProvider>
    </QueryClientProvider>
  );

  if (StripeProvider) {
    return (
      <StripeProvider publishableKey={process.env.EXPO_PUBLIC_STRIPE_KEY ?? ""}>
        {content}
      </StripeProvider>
    );
  }

  return content;
}
