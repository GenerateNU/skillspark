import React from "react";
import {
  View,
  TouchableOpacity,
  ScrollView,
  useColorScheme,
} from "react-native";
import { useRouter } from "expo-router";
import { useSafeAreaInsets } from "react-native-safe-area-context";
import { ThemedText } from "@/components/themed-text";
import { ThemedView } from "@/components/themed-view";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { Colors } from "@/constants/theme";
import { useTranslation } from "react-i18next";

export default function TermsAndConditionsScreen() {
  const router = useRouter();
  const insets = useSafeAreaInsets();
  const colorScheme = useColorScheme();
  const theme = Colors[colorScheme ?? "light"];
  const { t: translate } = useTranslation();

  return (
    <ThemedView className="flex-1" style={{ paddingTop: insets.top }}>
      <View className="flex-row items-center justify-between px-5 py-[14px]">
        <TouchableOpacity
          onPress={() => router.navigate("/settings")}
          className="w-10 justify-center items-start"
          hitSlop={{ top: 10, bottom: 10, left: 10, right: 10 }}
        >
          <IconSymbol name="chevron.left" size={24} color={theme.text} />
        </TouchableOpacity>
        <ThemedText className="text-xl text-center font-nunito-bold">
          {translate("settings.termsAndConditions")}
        </ThemedText>
        <View className="w-10" />
      </View>
      <ScrollView
        showsVerticalScrollIndicator={false}
        className="px-5 pt-2 pb-10"
      >
        <ThemedText
          className="text-[15px] font-nunito italic mb-4"
          style={{ color: theme.icon }}
        >
          {translate("termsAndConditions.lastUpdated")}
        </ThemedText>

        <ThemedText className="text-[15px] font-nunito-bold mb-2">
          {translate("termsAndConditions.acknowledgmentTitle")}
        </ThemedText>
        <ThemedText className="text-[15px] font-nunito mb-4 leading-[22px]">
          {translate("termsAndConditions.acknowledgment1")}
        </ThemedText>
        <ThemedText className="text-[15px] font-nunito mb-4 leading-[22px]">
          {translate("termsAndConditions.acknowledgment2")}
        </ThemedText>
        <ThemedText className="text-[15px] font-nunito mb-4 leading-[22px]">
          {translate("termsAndConditions.acknowledgment3")}
        </ThemedText>
        <ThemedText className="text-[15px] font-nunito mb-4 leading-[22px]">
          {translate("termsAndConditions.acknowledgment4")}
        </ThemedText>
        <ThemedText className="text-[15px] font-nunito mb-4 leading-[22px]">
          {translate("termsAndConditions.acknowledgment5")}
        </ThemedText>

        <ThemedText className="text-[15px] font-nunito-bold mb-2">
          {translate("termsAndConditions.userAccountsTitle")}
        </ThemedText>
        <ThemedText className="text-[15px] font-nunito mb-4 leading-[22px]">
          {translate("termsAndConditions.userAccounts1")}
        </ThemedText>
        <ThemedText className="text-[15px] font-nunito mb-4 leading-[22px]">
          {translate("termsAndConditions.userAccounts2")}
        </ThemedText>

        <ThemedText className="text-[15px] font-nunito-bold mb-2">
          {translate("termsAndConditions.intellectualPropertyTitle")}
        </ThemedText>
        <ThemedText className="text-[15px] font-nunito mb-4 leading-[22px]">
          {translate("termsAndConditions.intellectualPropertyDesc")}
        </ThemedText>

        <ThemedText className="text-[15px] font-nunito-bold mb-2">
          {translate("termsAndConditions.limitationTitle")}
        </ThemedText>
        <ThemedText className="text-[15px] font-nunito mb-4 leading-[22px]">
          {translate("termsAndConditions.limitationDesc")}
        </ThemedText>

        <ThemedText className="text-[15px] font-nunito-bold mb-2">
          {translate("termsAndConditions.governingLawTitle")}
        </ThemedText>
        <ThemedText className="text-[15px] font-nunito mb-4 leading-[22px]">
          {translate("termsAndConditions.governingLawDesc")}
        </ThemedText>

        <ThemedText className="text-[15px] font-nunito-bold mb-2">
          {translate("termsAndConditions.changesTitle")}
        </ThemedText>
        <ThemedText className="text-[15px] font-nunito mb-4 leading-[22px]">
          {translate("termsAndConditions.changesDesc")}
        </ThemedText>

        <ThemedText className="text-[15px] font-nunito-bold mb-2">
          {translate("termsAndConditions.contactTitle")}
        </ThemedText>
        <ThemedText className="text-[15px] font-nunito mb-4 leading-[22px]">
          {translate("termsAndConditions.contactDesc")}
        </ThemedText>
      </ScrollView>
    </ThemedView>
  );
}
