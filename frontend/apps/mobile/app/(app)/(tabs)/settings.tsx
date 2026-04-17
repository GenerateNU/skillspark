import { ErrorScreen } from "@/components/ErrorScreen";
import { ThemedText } from "@/components/themed-text";
import { ThemedView } from "@/components/themed-view";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { AppColors, Colors } from "@/constants/theme";
import { useAuthContext } from "@/hooks/use-auth-context";
import { useGuardian } from "@/hooks/use-guardian";
import { useRouter } from "expo-router";
import React, { useEffect, useState } from "react";
import { useTranslation } from "react-i18next";
import { Switch, TouchableOpacity, View } from "react-native";
import { useSafeAreaInsets } from "react-native-safe-area-context";

export default function SettingsScreen() {
  const router = useRouter();
  const insets = useSafeAreaInsets();
  const theme = Colors.light;
  const { t: translate } = useTranslation();

  const cardBg = "#EFEFEF";
  const dividerColor = "#D1D5DB";

  const handleDeleteAccount = () => {};

  const { logout } = useAuthContext();

  const handleLogOut = () => {
    logout();
  };

  const { guardianId, update } = useAuthContext();

  const { guardian } = useGuardian(guardianId);

  const [pushEnabled, setPushEnabled] = useState(false);
  const [emailEnabled, setEmailEnabled] = useState(false);

  useEffect(() => {
    if (!guardian) return;
    setPushEnabled(guardian.push_notifications);
    setEmailEnabled(guardian.email_notifications);
  }, [guardian]);

  const handlePushToggle = (value: boolean) => {
    if (!guardian) return;
    setPushEnabled(value);
    update(
      () => {},
      () => setPushEnabled(!value),
      guardianId!,
      guardian.email,
      guardian.language_preference,
      guardian.name,
      guardian.username,
      guardian.profile_picture_s3_key,
      guardian.expo_push_token,
      value,
      emailEnabled,
    );
  };

  const handleEmailToggle = (value: boolean) => {
    if (!guardian) return;
    setEmailEnabled(value);
    update(
      () => {},
      () => setEmailEnabled(!value),
      guardianId!,
      guardian.email,
      guardian.language_preference,
      guardian.name,
      guardian.username,
      guardian.profile_picture_s3_key,
      guardian.expo_push_token,
      pushEnabled,
      value,
    );
  };

  return (
    <ThemedView className="flex-1" style={{ paddingTop: insets.top }}>
      <View className="flex-row items-center justify-between px-5 py-[14px]">
        <TouchableOpacity
          onPress={() => router.navigate("/profile")}
          className="w-10 justify-center items-start"
          hitSlop={{ top: 10, bottom: 10, left: 10, right: 10 }}
        >
          <IconSymbol name="chevron.left" size={24} color={theme.text} />
        </TouchableOpacity>
        <ThemedText className="text-xl text-center font-nunito-bold">
          {translate("settings.title")}
        </ThemedText>
        <View className="w-10" />
      </View>

      <View className="px-4 pt-5">
        <View
          className="rounded-2xl overflow-hidden"
          style={{ backgroundColor: cardBg }}
        >
          <TouchableOpacity
            className="flex-row items-center justify-between px-4 py-[18px]"
            activeOpacity={0.6}
            onPress={() => router.push("/language")}
          >
            <ThemedText className="text-[17px] font-nunito">
              {translate("settings.language")}
            </ThemedText>
            <IconSymbol name="chevron.right" size={16} color="#9CA3AF" />
          </TouchableOpacity>
          <View className="h-px" style={{ backgroundColor: dividerColor }} />
          <TouchableOpacity
            className="flex-row items-center justify-between px-4 py-[18px]"
            activeOpacity={0.6}
            onPress={() => router.push("/terms-and-conditions")}
          >
            <ThemedText className="text-[17px] font-nunito">
              {translate("settings.termsAndConditions")}
            </ThemedText>
            <IconSymbol name="chevron.right" size={16} color="#9CA3AF" />
          </TouchableOpacity>
          <View className="h-px" style={{ backgroundColor: dividerColor }} />
          <TouchableOpacity
            className="flex-row items-center justify-between px-4 py-[18px]"
            activeOpacity={0.6}
            onPress={() => router.push("/privacy-policy")}
          >
            <ThemedText className="text-[17px] font-nunito">
              {translate("settings.privacyPolicy")}
            </ThemedText>
            <IconSymbol name="chevron.right" size={16} color="#9CA3AF" />
          </TouchableOpacity>
          <View className="h-px" style={{ backgroundColor: dividerColor }} />
          <TouchableOpacity
            className="flex-row items-center justify-between px-4 py-[18px]"
            activeOpacity={0.6}
            onPress={handleLogOut}
          >
            <ThemedText className="text-[17px] font-nunito">
              {translate("settings.logOut")}
            </ThemedText>
          </TouchableOpacity>
          <View className="h-px" style={{ backgroundColor: dividerColor }} />
          <TouchableOpacity
            className="flex-row items-center justify-between px-4 py-[18px]"
            activeOpacity={0.6}
            onPress={handleDeleteAccount}
          >
            <ThemedText className="text-[17px] font-nunito">
              {translate("settings.deleteAccount")}
            </ThemedText>
          </TouchableOpacity>
        </View>

        <View
          className="rounded-2xl overflow-hidden mt-5"
          style={{ backgroundColor: cardBg }}
        >
          <View className="flex-row items-center justify-between px-4 py-[18px]">
            <ThemedText className="text-[17px] font-nunito">
              {translate("settings.pushNotifications")}
            </ThemedText>
            <Switch
              value={pushEnabled}
              onValueChange={handlePushToggle}
              trackColor={{
                false: AppColors.borderLight,
                true: AppColors.checkboxSelected,
              }}
              thumbColor={Colors.light.dropdownBg}
            />
          </View>
          <View className="h-px" style={{ backgroundColor: dividerColor }} />
          <View className="flex-row items-center justify-between px-4 py-[18px]">
            <ThemedText className="text-[17px] font-nunito">
              {translate("settings.emailNotifications")}
            </ThemedText>
            <Switch
              value={emailEnabled}
              onValueChange={handleEmailToggle}
              trackColor={{
                false: AppColors.borderLight,
                true: AppColors.checkboxSelected,
              }}
              thumbColor={Colors.light.dropdownBg}
            />
          </View>
        </View>
      </View>
    </ThemedView>
  );
}
