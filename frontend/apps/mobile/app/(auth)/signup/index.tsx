import { ThemedText } from "@/components/themed-text";
import { ThemedView } from "@/components/themed-view";
import { router } from "expo-router";
import { useState } from "react";
import {
  BlurEvent,
  KeyboardAvoidingView,
  Platform,
  ScrollView,
  TextInputChangeEvent,
  View,
} from "react-native";
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
  const [cannotSubmit, setCannotSubmit] = useState(false);
  const { signup, usernameExists } = useAuthContext();
  const {
    control,
    handleSubmit,
    formState: { errors },
    getValues,
    setError,
    clearErrors,
  } = useForm<SignupFormData>({
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

  const onClickOut = async () => {
    let username = getValues("username");
    if (!username) {
      setCannotSubmit(false);
      setErrorText("");
      return;
    }
    const result = await usernameExists(username, setErrorText);
    if (!result) {
      setError("username", {
        type: "manual",
        message: "Username is taken.",
      });
      setCannotSubmit(true);
    } else {
      clearErrors("username");
      setErrorText("");
      setCannotSubmit(false);
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
              error={errors.name}
              placeholder="Full Name"
              autoCapitalize="none"
            />
            <AuthFormInput
              control={control}
              name="email"
              error={errors.email}
              placeholder="Email"
              keyboardType="email-address"
              autoCapitalize="none"
            />
            <AuthFormInput
              control={control}
              name="username"
              error={errors.username}
              placeholder="Username"
              autoCapitalize="none"
              onBlur={(e) => onClickOut()}
            />
            <AuthFormInput
              control={control}
              name="password"
              error={errors.password}
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
            <Button
              label="Sign Up"
              onPress={handleSubmit(onSubmit)}
              disabled={cannotSubmit}
            />
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
