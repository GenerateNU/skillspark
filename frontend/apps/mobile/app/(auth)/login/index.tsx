import { ThemedText } from "@/components/themed-text";
import { FontSizes } from "@/constants/theme";
import { router } from "expo-router";
import React, { useState } from "react";
import {
  Alert,
  KeyboardAvoidingView,
  Platform,
  ScrollView,
  TouchableOpacity,
  View,
} from "react-native";
import { AuthBackground } from "@/components/AuthBackground";
import { useAuthContext } from "@/hooks/use-auth-context";
import { useForm } from "react-hook-form";
import { ErrorMessage } from "@/components/ErrorMessage";
import { PageRedirectButton } from "@/components/PageRedirectButton";
import { Button } from "@/components/Button";
import { AuthFormInput } from "@/components/AuthFormInput";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { useTranslation } from "react-i18next";
import { useSafeAreaInsets } from "react-native-safe-area-context";
import { JumpingCharacter } from "@/components/JumpingCharacter";

type LoginFormData = {
  email: string;
  password: string;
};

export default function LoginScreen() {
  const insets = useSafeAreaInsets();
  const [errorText, setErrorText] = useState("");
  const { t: translate } = useTranslation();
  const { login } = useAuthContext();

  const { control, handleSubmit } = useForm<LoginFormData>({
    defaultValues: {
      email: "",
      password: "",
    },
  });

  const onSubmit = (formData: LoginFormData) => {
    if (!formData.email || !formData.password) {
      Alert.alert(
        translate("common.error"),
        translate("childProfile.requiredFieldsError"),
      );
      return;
    }
    login(formData.email, formData.password, setErrorText, () =>
      router.push("/(app)/(tabs)"),
    );
  };

  return (
    <View className="absolute inset-0">
      <AuthBackground />
      <View className="flex-1" style={{ paddingTop: insets.top }}>
        <KeyboardAvoidingView
          behavior={Platform.OS === "ios" ? "padding" : "height"}
          className="flex-1"
        >
          <ScrollView
            contentContainerStyle={{
              flexGrow: 1,
              paddingBottom: insets.bottom + 24,
            }}
            keyboardShouldPersistTaps="handled"
            showsVerticalScrollIndicator={false}
          >
            {/* Back button — space always reserved */}
            <View className="h-11 justify-center">
              {router.canGoBack() && (
                <TouchableOpacity
                  onPress={() => router.back()}
                  className="flex-row items-center px-5 gap-1"
                  hitSlop={{ top: 10, bottom: 10, left: 10, right: 10 }}
                >
                  <IconSymbol name="chevron.left" size={18} color="#11181C" />
                  <ThemedText className="text-base font-nunito">
                    {translate("onboarding.back")}
                  </ThemedText>
                </TouchableOpacity>
              )}
            </View>

            {/* Title */}
            <View className="px-6 pt-2 items-center">
              <ThemedText
                className="font-nunito-bold text-[#111] text-center"
                style={{
                  fontSize: FontSizes.hero,
                  lineHeight: FontSizes.hero + 8,
                  letterSpacing: -0.5,
                }}
              >
                {translate("onboarding.signIn")}
              </ThemedText>
            </View>

            {/* Character image */}
            <View className="items-center py-6">
              <JumpingCharacter />
            </View>

            {/* Form fields */}
            <View className="px-6 gap-6">
              <View className="gap-2">
                <ThemedText className="text-base font-nunito-semibold">
                  {translate("onboarding.emailOrUsername")}
                </ThemedText>
                <AuthFormInput
                  control={control}
                  name="email"
                  keyboardType="email-address"
                  autoCapitalize="none"
                />
              </View>

              <View className="gap-2">
                <ThemedText className="text-base font-nunito-semibold">
                  {translate("onboarding.password")}
                </ThemedText>
                <AuthFormInput
                  control={control}
                  name="password"
                  secureTextEntry
                />
                {/* Forgot Password link */}
                <TouchableOpacity
                  className="self-end"
                  onPress={() => {}}
                  hitSlop={{ top: 8, bottom: 8, left: 8, right: 8 }}
                >
                  <ThemedText className="text-base font-nunito underline text-[#111]">
                    {translate("onboarding.forgotPassword")}
                  </ThemedText>
                </TouchableOpacity>
              </View>
            </View>

            {/* Buttons */}
            <View className="px-6 pt-8 items-center">
              <Button
                label={translate("onboarding.signIn")}
                onPress={handleSubmit(onSubmit)}
                disabled={false}
              />
              <PageRedirectButton
                label={translate("onboarding.dontHaveAccount")}
                onPress={() => router.navigate("/(auth)/signup")}
              />
              <ErrorMessage message={errorText} />
            </View>
          </ScrollView>
        </KeyboardAvoidingView>
      </View>
    </View>
  );
}
