import React from "react";
import { View, TouchableOpacity, useColorScheme } from "react-native";
import { useRouter } from "expo-router";
import { useSafeAreaInsets } from "react-native-safe-area-context";
import { ThemedText } from "@/components/themed-text";
import { ThemedView } from "@/components/themed-view";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { Colors, AppColors } from "@/constants/theme";
import { useTranslation } from "react-i18next";

export default function PaymentScreen() {
  const router = useRouter();
  const insets = useSafeAreaInsets();
  const colorScheme = useColorScheme();
  const theme = Colors[colorScheme ?? "light"];
  const { t: translate } = useTranslation();

  const handleUpdateBilling = () => {};
  const handleDelete = () => {};

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
          {translate("payment.title")}
        </ThemedText>
        <View className="w-10" />
      </View>
      <View className="px-5 pt-5">
        <ThemedText className="text-[22px] font-nunito-bold mb-5">
          {translate("payment.manageBilling")}
        </ThemedText>
        {/* TODO: Replace with real payment method data from billing API */}
        <ThemedText className="text-base font-nunito mb-[6px]">
          {translate("payment.creditCard")}
        </ThemedText>
        <ThemedText className="text-base font-nunito mb-[6px]">
          {translate("payment.name")}
        </ThemedText>
        <ThemedText className="text-base font-nunito mb-8 tracking-widest">
          **** **** **** XXXX
        </ThemedText>
        <View className="flex-row gap-4">
          <TouchableOpacity
            className="flex-1 py-[14px] rounded-lg items-center justify-center"
            style={{ backgroundColor: AppColors.primaryBlue }}
            onPress={handleUpdateBilling}
            activeOpacity={0.8}
          >
            <ThemedText className="text-white text-[15px] font-nunito-semibold">
              {translate("payment.updateBilling")}
            </ThemedText>
          </TouchableOpacity>

          <TouchableOpacity
            className="flex-1 py-[14px] rounded-lg border-[1.5px] items-center justify-center"
            style={{ borderColor: theme.text }}
            onPress={handleDelete}
            activeOpacity={0.8}
          >
            <ThemedText
              className="text-[15px] font-nunito"
              style={{ color: theme.text }}
            >
              {translate("payment.delete")}
            </ThemedText>
          </TouchableOpacity>
        </View>
      </View>
    </ThemedView>
  );
}
