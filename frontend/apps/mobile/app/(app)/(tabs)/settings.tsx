import React from "react";
import { View, TouchableOpacity, useColorScheme } from "react-native";
import { useRouter } from "expo-router";
import { useSafeAreaInsets } from "react-native-safe-area-context";
import { ThemedText } from "@/components/themed-text";
import { ThemedView } from "@/components/themed-view";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { Colors } from "@/constants/theme";
import { useTranslation } from "react-i18next";
import { useAuthContext } from "@/hooks/use-auth-context";

export default function SettingsScreen() {
  const router = useRouter();
  const insets = useSafeAreaInsets();
  const colorScheme = useColorScheme();
  const theme = Colors[colorScheme ?? "light"];
  const { t: translate } = useTranslation();

  const cardBg = colorScheme === "dark" ? "#1c1c1e" : "#EFEFEF";
  const dividerColor = colorScheme === "dark" ? "#3a3a3c" : "#D1D5DB";

  const handleDeleteAccount = () => {};

  const { logout } = useAuthContext();

  const handleLogOut = () => {
    logout();
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
      </View>
    </ThemedView>
  );
}
