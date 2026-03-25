import { ThemedText } from "@/components/themed-text";
import { ThemedView } from "@/components/themed-view";
import { router } from "expo-router";
import { useState } from "react";
import { KeyboardAvoidingView, Platform, ScrollView, View } from "react-native";
import { useAuthContext } from "@/hooks/use-auth-context";
import { Controller, useForm } from "react-hook-form";
import { ErrorMessage } from "@/components/ErrorMessage";
import { PageRedirectButton } from "@/components/PageRedirectButton";
import { Button } from "@/components/Button";
import { AuthFormInput } from "@/components/AuthFormInput";
import { Dropdown } from "@/components/Dropdown";

type SignupFormData = {
  name: string;
  email: string;
  username: string;
  password: string;
  language_preference: string;
  profile_picture_s3_key: string | undefined;
};

export default function SignupScreen() {
  const [errorText, setErrorText] = useState("");
  const { signup } = useAuthContext();
  const { control, handleSubmit } = useForm<SignupFormData>({
    defaultValues: {
      name: "",
      email: "",
      username: "",
      password: "",
      language_preference: "",
      profile_picture_s3_key: undefined,
    },
  });

  const onSubmit = (formData: SignupFormData) => {
    if (
      formData.name === "" ||
      formData.email === "" ||
      formData.username === "" ||
      formData.password === "" ||
      formData.language_preference === ""
    ) {
      setErrorText("Missing a required field");
    } else {
      signup(
        formData.name,
        formData.email,
        formData.username,
        formData.password,
        formData.language_preference,
        formData.profile_picture_s3_key,
        setErrorText,
      );
    }
  };

  const handleGoToLogIn = () => {
    router.push("/(auth)/login");
  };

  return (
    <ThemedView className="flex-1">
      <KeyboardAvoidingView
        behavior={Platform.OS === "ios" ? "padding" : "height"}
        className="flex-1"
      >
        <ScrollView
          contentContainerStyle={{ flexGrow: 1 }}
          keyboardShouldPersistTaps="handled"
        >
          <View className="flex-1 items-center justify-center px-6 gap-4">
            <ThemedText type="title" className="text-3xl font-bold mb-8">
              Sign Up
            </ThemedText>
            <AuthFormInput
              control={control}
              name="name"
              placeholder="Full Name"
              autoCapitalize="none"
            />
            <AuthFormInput
              control={control}
              name="email"
              placeholder="Email"
              keyboardType="email-address"
              autoCapitalize="none"
            />
            <AuthFormInput
              control={control}
              name="username"
              placeholder="Username"
              autoCapitalize="none"
            />
            <AuthFormInput
              control={control}
              name="password"
              placeholder="Password"
              secureTextEntry={true}
            />
            <Controller
              control={control}
              name="language_preference"
              render={({ field: { onChange, value } }) => (
                <Dropdown
                  value={value}
                  onChange={onChange}
                  options={[
                    { label: "English", value: "en" },
                    { label: "Thai", value: "th" },
                  ]}
                  placeholder="Select a language..."
                />
              )}
            />
            <Button label="Sign Up" onPress={handleSubmit(onSubmit)} />
            <PageRedirectButton
              label="Already have an account? Log in"
              onPress={handleGoToLogIn}
            />
            <ErrorMessage message={errorText} />
          </View>
        </ScrollView>
      </KeyboardAvoidingView>
    </ThemedView>
  );
}