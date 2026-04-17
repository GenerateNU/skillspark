import { QueryClientProvider } from "@tanstack/react-query";
import { Stack } from "expo-router";
import { StatusBar } from "expo-status-bar";
import { useForm, FormProvider } from "react-hook-form";

import { queryClient } from "@/constants/query-client";
import {
  SignupFormData,
  signupFormDefaultValues,
} from "@/constants/signup-types";

export default function OnboardingStackLayout() {
  const methods = useForm<SignupFormData>({
    defaultValues: signupFormDefaultValues,
  });

  return (
    <QueryClientProvider client={queryClient}>
      <FormProvider {...methods}>
        <Stack>
          <Stack.Screen name="index" options={{ headerShown: false }} />
          <Stack.Screen name="account" options={{ headerShown: false }} />
          <Stack.Screen name="phone" options={{ headerShown: false }} />
          <Stack.Screen name="name" options={{ headerShown: false }} />
          <Stack.Screen name="photo" options={{ headerShown: false }} />
          <Stack.Screen name="child-profile" options={{ headerShown: false }} />
          <Stack.Screen
            name="emergency-contact"
            options={{ headerShown: false }}
          />
          <Stack.Screen name="payment" options={{ headerShown: false }} />
          <Stack.Screen name="all-set" options={{ headerShown: false }} />
        </Stack>
      </FormProvider>
      <StatusBar style="auto" />
    </QueryClientProvider>
  );
}
