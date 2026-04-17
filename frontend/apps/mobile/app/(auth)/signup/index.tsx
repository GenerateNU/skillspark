import React, { useState } from "react";
import { FontSizes } from "@/constants/theme";
import { View, TouchableOpacity, Text } from "react-native";
import { JumpingCharacter } from "@/components/JumpingCharacter";
import { useRouter } from "expo-router";
import { useSafeAreaInsets } from "react-native-safe-area-context";
import { ThemedText } from "@/components/themed-text";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { useTranslation } from "react-i18next";
import { setCurrentLanguage } from "@skillspark/api-client";
import { useAuthContext } from "@/hooks/use-auth-context";
import { Button } from "@/components/Button";
import { useFormContext } from "react-hook-form";
import { SignupFormData } from "@/constants/signup-types";
import { AuthBackground } from "@/components/AuthBackground";

const LANGUAGES = [
  { code: "en", label: "English", flag: "🇺🇸" },
  { code: "th", label: "Thai", flag: "🇹🇭" },
];

export default function WelcomeScreen() {
  const router = useRouter();
  const insets = useSafeAreaInsets();
  const { t: translate, i18n } = useTranslation();

  const { setValue } = useFormContext<SignupFormData>();
  const [selected, setSelected] = useState(i18n.language ?? "en");
  setValue("language_preference", selected);
  const { setLanguage } = useAuthContext();

  const updateLanguageData = async (langCode: string) => {
    setSelected(langCode);
    await i18n.changeLanguage(langCode);
    setCurrentLanguage(langCode);
    setLanguage(langCode);
    setValue("language_preference", langCode);
  };

  return (
    <View className="absolute inset-0">
      <AuthBackground />
      <View className="flex-1" style={{ paddingTop: insets.top }}>
        <View className="h-11" />
        {/* Title */}
        <View className="items-center px-6 pt-10 pb-5">
          <ThemedText
            className="font-nunito-bold text-[#111] text-center"
            style={{
              fontSize: FontSizes.hero,
              lineHeight: FontSizes.hero + 8,
              letterSpacing: -0.5,
            }}
            numberOfLines={1}
            adjustsFontSizeToFit
          >
            {translate("onboarding.welcome")}
          </ThemedText>
        </View>

        {/* Character image */}
        <View className="items-center justify-center pt-2 pb-4">
          <JumpingCharacter width={210} height={160} />
          <ThemedText className="font-nunito-semibold text-xl text-[#374151] mt-6 mb-2 text-center">
            {translate("onboarding.chooseLanguage")}
          </ThemedText>
        </View>

        {/* Language section */}
        <View className="px-6 flex-1 justify-center">
          <View>
            {LANGUAGES.map((lang) => {
              const isSelected = selected === lang.code;
              return (
                <React.Fragment key={lang.code}>
                  <TouchableOpacity
                    onPress={() => updateLanguageData(lang.code)}
                    activeOpacity={0.6}
                    className="flex-row items-center gap-[14px] py-6"
                  >
                    <Text className="text-[30px] leading-[38px]">
                      {lang.flag}
                    </Text>
                    <Text className="flex-1 font-nunito text-xl text-[#111]">
                      {translate(`settings.languages.${lang.code}`)}
                    </Text>
                    <IconSymbol
                      name={isSelected ? "checkmark.circle.fill" : "circle"}
                      size={24}
                      color="#1C1C1E"
                    />
                  </TouchableOpacity>
                </React.Fragment>
              );
            })}
          </View>
        </View>

        {/* Submit button */}
        <View
          className="items-center px-6 pt-4"
          style={{ paddingBottom: insets.bottom + 48 }}
        >
          <Button
            label={translate("common.submit")}
            onPress={() => router.push("/(auth)/signup/account")}
            disabled={false}
          />
        </View>
      </View>
    </View>
  );
}
