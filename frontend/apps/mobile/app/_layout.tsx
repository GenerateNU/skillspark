import { DefaultTheme, ThemeProvider } from "@react-navigation/native";
import { useFonts } from "expo-font";
import * as SplashScreen from "expo-splash-screen";
import * as SecureStore from "expo-secure-store";
import { useEffect, useState } from "react";
import "react-native-reanimated";
import { GestureHandlerRootView } from "react-native-gesture-handler";
import { QueryClientProvider } from "@tanstack/react-query";
import { queryClient } from "@/constants/query-client";
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
  const [langReady, setLangReady] = useState(false);
  const [loaded, error] = useFonts({
    NunitoSans_400Regular,
    NunitoSans_500Medium,
    NunitoSans_600SemiBold,
    NunitoSans_700Bold,
    MuseoModerno_700Bold: {
      uri: "https://fonts.gstatic.com/s/museomoderno/v21/zrf30VXsoJQLfl-LiGQLaGoBRhuRKNR8m3k.woff2",
    },
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
    <GestureHandlerRootView>
      <QueryClientProvider client={queryClient}>
        <ThemeProvider value={DefaultTheme}>
          <AuthProvider>
            <LoginRedirect />
          </AuthProvider>
        </ThemeProvider>
      </QueryClientProvider>
    </GestureHandlerRootView>
  );

  const stripePublishableKey: string | undefined = process.env.EXPO_PUBLIC_STRIPE_KEY;

  if (!stripePublishableKey) {
    throw new Error("EXPO_PUBLIC_STRIPE_KEY is not set up properly");
  }

  if (StripeProvider) {
    return (
      <StripeProvider publishableKey={stripePublishableKey}>
        {content}
      </StripeProvider>
    );
  }

  return content;
}
