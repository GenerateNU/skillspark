import { AuthBackground } from "@/components/AuthBackground";
import { AuthFormInput } from "@/components/AuthFormInput";
import { Button } from "@/components/Button";
import { ErrorMessage } from "@/components/ErrorMessage";
import { JumpingCharacter } from "@/components/JumpingCharacter";
import { ThemedText } from "@/components/themed-text";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { FontSizes } from "@/constants/theme";
import { useForgotPassword } from "@skillspark/api-client";
import { router } from "expo-router";
import { useState } from "react";
import { useForm } from "react-hook-form";
import { useTranslation } from "react-i18next";
import {
  KeyboardAvoidingView,
  Platform,
  ScrollView,
  TouchableOpacity,
  View,
} from "react-native";
import { useSafeAreaInsets } from "react-native-safe-area-context";

type ForgotPasswordFormData = {
  email: string;
};

export default function ForgotPasswordScreen() {
  const insets = useSafeAreaInsets();
  const { t: translate } = useTranslation();
  const [errorText, setErrorText] = useState("");
  const [submitted, setSubmitted] = useState(false);
  const { mutate: forgotPasswordFunc, isPending } = useForgotPassword();

  const { control, handleSubmit } = useForm<ForgotPasswordFormData>({
    defaultValues: { email: "" },
  });

  const onSubmit = (formData: ForgotPasswordFormData) => {
    setErrorText("");
    forgotPasswordFunc(
      { data: { email: formData.email } },
      {
        onSuccess: () => setSubmitted(true),
        onError: () => {
          // Show success regardless to avoid leaking email existence
          setSubmitted(true);
        },
      },
    );
  };

  if (submitted) {
    return (
      <View className="absolute inset-0">
        <AuthBackground />
        <View
          className="flex-1 items-center justify-center px-6"
          style={{ paddingTop: insets.top, paddingBottom: insets.bottom + 24 }}
        >
          <JumpingCharacter />
          <ThemedText
            className="font-nunito-bold text-[#111] text-center mt-6"
            style={{
              fontSize: FontSizes.hero,
              lineHeight: FontSizes.hero + 8,
              letterSpacing: -0.5,
            }}
          >
            {translate("onboarding.checkYourEmail")}
          </ThemedText>
          <ThemedText className="text-base font-nunito text-center text-[#555] mt-4 mb-8">
            {translate("onboarding.resetEmailSent")}
          </ThemedText>
          <TouchableOpacity onPress={() => router.replace("/(auth)/login")}>
            <ThemedText className="text-base font-nunito underline text-[#111]">
              {translate("onboarding.backToSignIn")}
            </ThemedText>
          </TouchableOpacity>
        </View>
      </View>
    );
  }

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
            {/* Back button */}
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
                {translate("onboarding.forgotPasswordTitle")}
              </ThemedText>
            </View>

            {/* Character image */}
            <View className="items-center py-6">
              <JumpingCharacter />
            </View>

            {/* Subtitle */}
            <View className="px-6 pb-4">
              <ThemedText className="text-base font-nunito text-center text-[#555]">
                {translate("onboarding.forgotPasswordSubtitle")}
              </ThemedText>
            </View>

            {/* Form */}
            <View className="px-6 gap-6">
              <View className="gap-2">
                <ThemedText className="text-base font-nunito-semibold">
                  {translate("onboarding.email")}
                </ThemedText>
                <AuthFormInput
                  control={control}
                  name="email"
                  keyboardType="email-address"
                  autoCapitalize="none"
                />
              </View>
            </View>

            {/* Buttons */}
            <View className="px-6 pt-8 items-center">
              <Button
                label={
                  isPending
                    ? translate("onboarding.sending")
                    : translate("onboarding.sendResetLink")
                }
                onPress={handleSubmit(onSubmit)}
                disabled={isPending}
              />
              <ErrorMessage message={errorText} />
            </View>
          </ScrollView>
        </KeyboardAvoidingView>
      </View>
    </View>
  );
}
