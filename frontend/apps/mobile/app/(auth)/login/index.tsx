import { ThemedText } from "@/components/themed-text";
import { ThemedView } from "@/components/themed-view";
import { router } from "expo-router";
import React, { useState } from "react";
import { View } from "react-native";
import { useAuthContext } from "@/hooks/use-auth-context";
import { useForm } from "react-hook-form";
import { ErrorMessage } from "@/components/ErrorMessage";
import { PageRedirectButton } from "@/components/PageRedirectButton";
import { Button } from "@/components/Button";
import { AuthFormInput } from "@/components/AuthFormInput";

type LoginFormData = {
  email: string;
  password: string;
};

export default function LoginScreen() {
  const [errorText, setErrorText] = useState("");
  const { login } = useAuthContext();

  const { control, handleSubmit } = useForm<LoginFormData>({
    defaultValues: {
      email: "",
      password: "",
    },
  });

  const onSubmit = (formData: LoginFormData) => {
    if (formData.email === "" || formData.password === "") {
      setErrorText("Missing email or password");
    } else {
      login(formData.email, formData.password, setErrorText);
    }
  };

  const handleGoToSignUp = () => {
    router.push("/(auth)/signup");
  };

  return (
    <ThemedView className="flex-1 items-center justify-center">
      <ThemedText type="title" className="text-3xl font-bold mb-8">
        Log In
      </ThemedText>

      <View className="w-full px-6 gap-4 items-center">
        <AuthFormInput
          control={control}
          name="email"
          placeholder="Email"
          keyboardType="email-address"
          autoCapitalize="none"
        />
        <AuthFormInput
          control={control}
          name="password"
          placeholder="Password"
          secureTextEntry={true}
        />
        <Button label="Log In" onPress={handleSubmit(onSubmit)} />
        <PageRedirectButton
          label="Don't have an account? Sign up"
          onPress={handleGoToSignUp}
        />
        <ErrorMessage message={errorText} />
      </View>
    </ThemedView>
  );
}
