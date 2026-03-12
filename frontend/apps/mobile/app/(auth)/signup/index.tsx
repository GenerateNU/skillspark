import { ThemedText } from "@/components/themed-text";
import { ThemedView } from "@/components/themed-view";
import { router } from "expo-router";
import { useState } from "react";
import { View } from "react-native";
import { useAuthContext } from "@/hooks/use-auth-context";
import { useForm } from "react-hook-form";
import { ErrorMessage } from "@/components/ErrorMessage";
import { PageRedirectButton } from "@/components/PageRedirectButton";
import { SubmitButton } from "@/components/SubmitButton";
import { AuthFormInput } from "@/components/AuthFormInput";

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
    <ThemedView
      style={{ flex: 1, alignItems: "center", justifyContent: "center" }}
    >
      <ThemedText
        type="title"
        style={{ fontSize: 30, fontWeight: "bold", marginBottom: 30 }}
      >
        Sign Up
      </ThemedText>
      <View
        style={{
          width: "100%",
          paddingHorizontal: 24,
          gap: 16,
          alignItems: "center",
        }}
      >
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
        <AuthFormInput
          control={control}
          name="language_preference"
          placeholder="Language Preference"
          autoCapitalize="none"
        />
        <SubmitButton label="Sign Up" onPress={handleSubmit(onSubmit)} />
        <PageRedirectButton
          label="Already have an account? Log in"
          onPress={handleGoToLogIn}
        />
        <ErrorMessage message={errorText} />
      </View>
    </ThemedView>
  );
}
