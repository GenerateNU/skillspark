import React from "react";
import { View, TouchableOpacity, ScrollView, useColorScheme } from "react-native";
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
        contentContainerStyle={{ paddingHorizontal: 20, paddingTop: 8, paddingBottom: 40 }}
      >
        <ThemedText className="text-[15px] font-nunito italic mb-4" style={{ color: theme.icon }}>
          Last Updated: April 4, 2026
        </ThemedText>

        <ThemedText className="text-[15px] font-nunito-bold mb-2">
          Acknowledgment
        </ThemedText>
        <ThemedText className="text-[15px] font-nunito mb-4" style={{ lineHeight: 22 }}>
          These are the Terms and Conditions governing the use of this Service and the agreement that operates between You and the Company. These Terms and Conditions set out the rights and obligations of all users regarding the use of the Service.
        </ThemedText>
        <ThemedText className="text-[15px] font-nunito mb-4" style={{ lineHeight: 22 }}>
          Your access to and use of the Service is conditioned on Your acceptance of and compliance with these Terms and Conditions. These Terms and Conditions apply to all visitors, users and others who access or use the Service.
        </ThemedText>
        <ThemedText className="text-[15px] font-nunito mb-4" style={{ lineHeight: 22 }}>
          By accessing or using the Service You agree to be bound by these Terms and Conditions. If You disagree with any part of these Terms and Conditions then You may not access the Service.
        </ThemedText>
        <ThemedText className="text-[15px] font-nunito mb-4" style={{ lineHeight: 22 }}>
          You represent that you are over the age of 18. The Company does not permit those under 18 to use the Service.
        </ThemedText>
        <ThemedText className="text-[15px] font-nunito mb-4" style={{ lineHeight: 22 }}>
          Your access to and use of the Service is also conditioned on Your acceptance of and compliance with the Privacy Policy of the Company. Our Privacy Policy describes Our policies and procedures on the collection, use and disclosure of Your personal information when You use the Application or the Website and tells You about Your privacy rights and how the law protects You. Please read Our Privacy Policy carefully before using Our Service.
        </ThemedText>

        <ThemedText className="text-[15px] font-nunito-bold mb-2">
          User Accounts
        </ThemedText>
        <ThemedText className="text-[15px] font-nunito mb-4" style={{ lineHeight: 22 }}>
          When You create an account with Us, You must provide Us information that is accurate, complete, and current at all times. Failure to do so constitutes a breach of the Terms, which may result in immediate termination of Your account on Our Service.
        </ThemedText>
        <ThemedText className="text-[15px] font-nunito mb-4" style={{ lineHeight: 22 }}>
          You are responsible for safeguarding the password that You use to access the Service and for any activities or actions under Your password. You agree not to disclose Your password to any third party. You must notify Us immediately upon becoming aware of any breach of security or unauthorized use of Your account.
        </ThemedText>

        <ThemedText className="text-[15px] font-nunito-bold mb-2">
          Intellectual Property
        </ThemedText>
        <ThemedText className="text-[15px] font-nunito mb-4" style={{ lineHeight: 22 }}>
          The Service and its original content (excluding Content provided by You or other users), features and functionality are and will remain the exclusive property of the Company and its licensors. The Service is protected by copyright, trademark, and other laws of both the Country and foreign countries.
        </ThemedText>

        <ThemedText className="text-[15px] font-nunito-bold mb-2">
          Limitation of Liability
        </ThemedText>
        <ThemedText className="text-[15px] font-nunito mb-4" style={{ lineHeight: 22 }}>
          To the maximum extent permitted by applicable law, in no event shall the Company or its suppliers be liable for any special, incidental, indirect, or consequential damages whatsoever (including, but not limited to, damages for loss of profits, loss of data or other information, personal injury, loss of privacy) arising out of or in any way related to the use of or inability to use the Service.
        </ThemedText>

        <ThemedText className="text-[15px] font-nunito-bold mb-2">
          Governing Law
        </ThemedText>
        <ThemedText className="text-[15px] font-nunito mb-4" style={{ lineHeight: 22 }}>
          The laws of the Country, excluding its conflicts of law rules, shall govern this Terms and Your use of the Service. Your use of the Application may also be subject to other local, state, national, or international laws.
        </ThemedText>

        <ThemedText className="text-[15px] font-nunito-bold mb-2">
          Changes to These Terms and Conditions
        </ThemedText>
        <ThemedText className="text-[15px] font-nunito mb-4" style={{ lineHeight: 22 }}>
          We reserve the right, at Our sole discretion, to modify or replace these Terms at any time. If a revision is material We will make reasonable efforts to provide at least 30 days' notice prior to any new terms taking effect. What constitutes a material change will be determined at Our sole discretion.
        </ThemedText>

        <ThemedText className="text-[15px] font-nunito-bold mb-2">
          Contact Us
        </ThemedText>
        <ThemedText className="text-[15px] font-nunito mb-4" style={{ lineHeight: 22 }}>
          If you have any questions about these Terms and Conditions, You can contact us by email.
        </ThemedText>
      </ScrollView>
    </ThemedView>
  );
}
