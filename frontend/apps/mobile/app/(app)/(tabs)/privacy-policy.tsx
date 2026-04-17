import React from "react";
import { View, TouchableOpacity, ScrollView } from "react-native";
import { useRouter } from "expo-router";
import { useSafeAreaInsets } from "react-native-safe-area-context";
import { ThemedText } from "@/components/themed-text";
import { ThemedView } from "@/components/themed-view";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { Colors } from "@/constants/theme";
import { useTranslation } from "react-i18next";
import { FLOATING_TAB_BAR_SCROLL_PADDING } from "@/components/floating-tab-bar";

export default function PrivacyPolicyScreen() {
  const router = useRouter();
  const insets = useSafeAreaInsets();
  const theme = Colors.light;
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
          {translate("settings.privacyPolicy")}
        </ThemedText>
        <View className="w-10" />
      </View>
      <ScrollView
        showsVerticalScrollIndicator={false}
        className="px-5 pt-2"
        contentContainerStyle={{
          paddingBottom: FLOATING_TAB_BAR_SCROLL_PADDING,
        }}
      >
        <ThemedText
          className="text-[15px] font-nunito italic mb-4"
          style={{ color: theme.icon }}
        >
          {translate("privacyPolicy.lastUpdated")}
        </ThemedText>

        <ThemedText className="text-[15px] font-nunito mb-4 leading-[22px]">
          {translate("privacyPolicy.intro1")}
        </ThemedText>
        <ThemedText className="text-[15px] font-nunito mb-4 leading-[22px]">
          {translate("privacyPolicy.intro2")}
        </ThemedText>

        <ThemedText className="text-[15px] font-nunito-bold mb-2">
          {translate("privacyPolicy.collectingDataTitle")}
        </ThemedText>
        <ThemedText className="text-[15px] font-nunito-bold mb-2">
          {translate("privacyPolicy.typesOfDataTitle")}
        </ThemedText>
        <ThemedText className="text-[15px] font-nunito-bold mb-2">
          {translate("privacyPolicy.personalDataTitle")}
        </ThemedText>
        <ThemedText className="text-[15px] font-nunito mb-3 leading-[22px]">
          {translate("privacyPolicy.personalDataDesc")}
        </ThemedText>
        <ThemedText className="text-[15px] font-nunito mb-1 ml-4 leading-[22px]">
          {`\u2022 ${translate("privacyPolicy.emailAddress")}`}
        </ThemedText>
        <ThemedText className="text-[15px] font-nunito mb-1 ml-4 leading-[22px]">
          {`\u2022 ${translate("privacyPolicy.name")}`}
        </ThemedText>
        <ThemedText className="text-[15px] font-nunito mb-1 ml-4 leading-[22px]">
          {`\u2022 ${translate("privacyPolicy.phoneNumber")}`}
        </ThemedText>
        <ThemedText className="text-[15px] font-nunito mb-4 ml-4 leading-[22px]">
          {`\u2022 ${translate("privacyPolicy.usageDataItem")}`}
        </ThemedText>

        <ThemedText className="text-[15px] font-nunito-bold mb-2">
          {translate("privacyPolicy.usageDataTitle")}
        </ThemedText>
        <ThemedText className="text-[15px] font-nunito mb-4 leading-[22px]">
          {translate("privacyPolicy.usageDataDesc")}
        </ThemedText>

        <ThemedText className="text-[15px] font-nunito-bold mb-2">
          {translate("privacyPolicy.useOfDataTitle")}
        </ThemedText>
        <ThemedText className="text-[15px] font-nunito mb-4 leading-[22px]">
          {translate("privacyPolicy.useOfDataDesc")}
        </ThemedText>

        <ThemedText className="text-[15px] font-nunito-bold mb-2">
          {translate("privacyPolicy.securityTitle")}
        </ThemedText>
        <ThemedText className="text-[15px] font-nunito mb-4 leading-[22px]">
          {translate("privacyPolicy.securityDesc")}
        </ThemedText>

        <ThemedText className="text-[15px] font-nunito-bold mb-2">
          {translate("privacyPolicy.childrenTitle")}
        </ThemedText>
        <ThemedText className="text-[15px] font-nunito mb-4 leading-[22px]">
          {translate("privacyPolicy.childrenDesc")}
        </ThemedText>

        <ThemedText className="text-[15px] font-nunito-bold mb-2">
          {translate("privacyPolicy.changesTitle")}
        </ThemedText>
        <ThemedText className="text-[15px] font-nunito mb-4 leading-[22px]">
          {translate("privacyPolicy.changesDesc")}
        </ThemedText>

        <ThemedText className="text-[15px] font-nunito-bold mb-2">
          {translate("privacyPolicy.contactTitle")}
        </ThemedText>
        <ThemedText className="text-[15px] font-nunito mb-4 leading-[22px]">
          {translate("privacyPolicy.contactDesc")}
        </ThemedText>
      </ScrollView>
    </ThemedView>
  );
}
