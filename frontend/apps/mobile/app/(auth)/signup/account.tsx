import { AuthFormInput } from "@/components/AuthFormInput";
import { FontSizes } from "@/constants/theme";
import { Button } from "@/components/Button";
import { ThemedText } from "@/components/themed-text";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { useRouter } from "expo-router";
import { useFormContext, useWatch } from "react-hook-form";
import { useTranslation } from "react-i18next";
import { SignupFormData } from "@/constants/signup-types";
import {
  Alert,
  KeyboardAvoidingView,
  Platform,
  ScrollView,
  TouchableOpacity,
  View,
} from "react-native";
import { useSafeAreaInsets } from "react-native-safe-area-context";
import { AuthBackground } from "@/components/AuthBackground";
import { JumpingCharacter } from "@/components/JumpingCharacter";
import { PageRedirectButton } from "@/components/PageRedirectButton";

// Matches the backend validatePasswordStrength rules exactly
const PASSWORD_RULES = [
  { key: "length",    check: (p: string) => p.length >= 8,                          i18nKey: "onboarding.passwordReqLength" },
  { key: "upper",     check: (p: string) => /[A-Z]/.test(p),                        i18nKey: "onboarding.passwordReqUppercase" },
  { key: "lower",     check: (p: string) => /[a-z]/.test(p),                        i18nKey: "onboarding.passwordReqLowercase" },
  { key: "number",    check: (p: string) => /[0-9]/.test(p),                        i18nKey: "onboarding.passwordReqNumber" },
  { key: "special",   check: (p: string) => /[!@#~$%^&*()+|_.,;<>?/{}\\-]/.test(p), i18nKey: "onboarding.passwordReqSpecial" },
];

export default function AccountScreen() {
  const router = useRouter();
  const { t: translate } = useTranslation();
  const insets = useSafeAreaInsets();
  const { getValues, control } = useFormContext<SignupFormData>();
  const password = useWatch({ control, name: "password" }) ?? "";

  const handleCreateAccount = () => {
    const email = getValues("email");
    const pw = getValues("password");
    const confirmPassword = getValues("confirm_password");

    if (!email || !pw || !confirmPassword) {
      Alert.alert(
        translate("common.error"),
        translate("childProfile.requiredFieldsError"),
      );
      return;
    }
    const failed = PASSWORD_RULES.filter((r) => !r.check(pw));
    if (failed.length > 0) {
      Alert.alert(
        translate("onboarding.passwordInvalid"),
        failed.map((r) => `• ${translate(r.i18nKey)}`).join("\n"),
      );
      return;
    }
    if (pw !== confirmPassword) {
      Alert.alert(
        translate("common.error"),
        translate("onboarding.passwordMismatch"),
      );
      return;
    }
    router.push("/(auth)/signup/phone");
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
            contentContainerStyle={{ flexGrow: 1 }}
            keyboardShouldPersistTaps="handled"
            showsVerticalScrollIndicator={false}
          >
            {/* Back button */}
            <TouchableOpacity
              onPress={() => router.back()}
              className="flex-row items-center px-5 py-3 gap-1"
              hitSlop={{ top: 10, bottom: 10, left: 10, right: 10 }}
            >
              <IconSymbol name="chevron.left" size={18} color="#11181C" />
              <ThemedText className="text-base font-nunito">
                {translate("onboarding.back")}
              </ThemedText>
            </TouchableOpacity>

            {/* Title */}
            <View className="px-6 pt-4 pb-2 items-center">
              <ThemedText
                className="font-nunito-bold text-[#111] text-center"
                style={{
                  fontSize: FontSizes.hero,
                  lineHeight: FontSizes.hero + 8,
                  letterSpacing: -0.5,
                }}
              >
                {translate("onboarding.makeAccount")}
              </ThemedText>
            </View>

            {/* Character mascot */}
            <View className="items-center py-4">
              <JumpingCharacter width={180} height={140} />
            </View>

            {/* Form fields */}
            <View className="flex-1 justify-center px-6">
              <View className="gap-6">
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

                <View className="gap-2">
                  <ThemedText className="text-base font-nunito-semibold">
                    {translate("onboarding.password")}
                  </ThemedText>
                  <AuthFormInput
                    control={control}
                    name="password"
                    secureTextEntry
                  />
                  {password.length > 0 && (
                    <View className="gap-1 mt-1">
                      {PASSWORD_RULES.map((rule) => {
                        const met = rule.check(password);
                        return (
                          <View key={rule.key} className="flex-row items-center gap-1.5">
                            <IconSymbol
                              name={met ? "checkmark.circle.fill" : "xmark.circle"}
                              size={14}
                              color={met ? "#22C55E" : "#9CA3AF"}
                            />
                            <ThemedText
                              className="text-xs font-nunito"
                              style={{ color: met ? "#22C55E" : "#9CA3AF" }}
                            >
                              {translate(rule.i18nKey)}
                            </ThemedText>
                          </View>
                        );
                      })}
                    </View>
                  )}
                </View>

                <View className="gap-2">
                  <ThemedText className="text-base font-nunito-semibold">
                    {translate("onboarding.confirmPassword")}
                  </ThemedText>
                  <AuthFormInput
                    control={control}
                    name="confirm_password"
                    secureTextEntry
                  />
                </View>
              </View>
            </View>
          </ScrollView>
        </KeyboardAvoidingView>

        {/* Buttons pinned to bottom */}
        <View
          className="items-center px-6 pt-4"
          style={{ paddingBottom: insets.bottom + 16 }}
        >
          <Button
            label={translate("onboarding.createAccount")}
            onPress={handleCreateAccount}
            disabled={false}
          />
          <View className="items-center justify-center" style={{ height: 48 }}>
            <PageRedirectButton
              label={translate("onboarding.alreadyHaveAccount")}
              onPress={() => router.navigate("/(auth)/login")}
            />
          </View>
        </View>
      </View>
    </View>
  );
}
