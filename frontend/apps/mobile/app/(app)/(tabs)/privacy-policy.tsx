import React from "react";
import { View, TouchableOpacity, ScrollView, useColorScheme } from "react-native";
import { useRouter } from "expo-router";
import { useSafeAreaInsets } from "react-native-safe-area-context";
import { ThemedText } from "@/components/themed-text";
import { ThemedView } from "@/components/themed-view";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { Colors } from "@/constants/theme";
import { useTranslation } from "react-i18next";

export default function PrivacyPolicyScreen() {
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
          {translate("settings.privacyPolicy")}
        </ThemedText>
        <View className="w-10" />
      </View>
      <ScrollView
        showsVerticalScrollIndicator={false}
        contentContainerStyle={{ paddingHorizontal: 20, paddingTop: 8, paddingBottom: 40 }}
      >
        <ThemedText className="text-[15px] font-nunito italic mb-4" style={{ color: theme.icon }}>
          Last Updated: April 4, 2026
        </ThemedText>

        <ThemedText className="text-[15px] font-nunito mb-4" style={{ lineHeight: 22 }}>
          This Privacy Policy describes Our policies and procedures on the collection, use and disclosure of Your information when You use the Service and tells You about Your privacy rights and how the law protects You.
        </ThemedText>
        <ThemedText className="text-[15px] font-nunito mb-4" style={{ lineHeight: 22 }}>
          We use Your Personal data to provide and improve the Service. By using the Service, You agree to the collection and use of information in accordance with this Privacy Policy.
        </ThemedText>

        <ThemedText className="text-[15px] font-nunito-bold mb-2">
          Collecting and Using Your Personal Data
        </ThemedText>
        <ThemedText className="text-[15px] font-nunito-bold mb-2">
          Types of Data Collected
        </ThemedText>
        <ThemedText className="text-[15px] font-nunito-bold mb-2">
          Personal Data
        </ThemedText>
        <ThemedText className="text-[15px] font-nunito mb-3" style={{ lineHeight: 22 }}>
          While using Our Service, We may ask You to provide Us with certain personally identifiable information that can be used to contact or identify You. Personally identifiable information may include, but is not limited to:
        </ThemedText>
        <ThemedText className="text-[15px] font-nunito mb-1 ml-4" style={{ lineHeight: 22 }}>
          • Email address
        </ThemedText>
        <ThemedText className="text-[15px] font-nunito mb-1 ml-4" style={{ lineHeight: 22 }}>
          • Name
        </ThemedText>
        <ThemedText className="text-[15px] font-nunito mb-1 ml-4" style={{ lineHeight: 22 }}>
          • Phone number
        </ThemedText>
        <ThemedText className="text-[15px] font-nunito mb-4 ml-4" style={{ lineHeight: 22 }}>
          • Usage Data
        </ThemedText>

        <ThemedText className="text-[15px] font-nunito-bold mb-2">
          Usage Data
        </ThemedText>
        <ThemedText className="text-[15px] font-nunito mb-4" style={{ lineHeight: 22 }}>
          Usage Data is collected automatically when using the Service. Usage Data may include information such as Your Device's Internet Protocol address (e.g. IP address), browser type, browser version, the pages of our Service that You visit, the time and date of Your visit, the time spent on those pages, unique device identifiers and other diagnostic data.
        </ThemedText>

        <ThemedText className="text-[15px] font-nunito-bold mb-2">
          Use of Your Personal Data
        </ThemedText>
        <ThemedText className="text-[15px] font-nunito mb-4" style={{ lineHeight: 22 }}>
          The Company may use Personal Data for the following purposes: to provide and maintain our Service, to manage Your Account, to contact You, to provide You with news and special offers, and for other purposes such as data analysis and improving our Service.
        </ThemedText>

        <ThemedText className="text-[15px] font-nunito-bold mb-2">
          Security of Your Personal Data
        </ThemedText>
        <ThemedText className="text-[15px] font-nunito mb-4" style={{ lineHeight: 22 }}>
          The security of Your Personal Data is important to Us, but remember that no method of transmission over the Internet, or method of electronic storage is 100% secure. While We strive to use commercially acceptable means to protect Your Personal Data, We cannot guarantee its absolute security.
        </ThemedText>

        <ThemedText className="text-[15px] font-nunito-bold mb-2">
          Children's Privacy
        </ThemedText>
        <ThemedText className="text-[15px] font-nunito mb-4" style={{ lineHeight: 22 }}>
          Our Service does not address anyone under the age of 13. We do not knowingly collect personally identifiable information from anyone under the age of 13. If You are a parent or guardian and You are aware that Your child has provided Us with Personal Data, please contact Us.
        </ThemedText>

        <ThemedText className="text-[15px] font-nunito-bold mb-2">
          Changes to This Privacy Policy
        </ThemedText>
        <ThemedText className="text-[15px] font-nunito mb-4" style={{ lineHeight: 22 }}>
          We may update Our Privacy Policy from time to time. We will notify You of any changes by posting the new Privacy Policy on this page. You are advised to review this Privacy Policy periodically for any changes.
        </ThemedText>

        <ThemedText className="text-[15px] font-nunito-bold mb-2">
          Contact Us
        </ThemedText>
        <ThemedText className="text-[15px] font-nunito mb-4" style={{ lineHeight: 22 }}>
          If you have any questions about this Privacy Policy, You can contact us by email.
        </ThemedText>
      </ScrollView>
    </ThemedView>
  );
}
