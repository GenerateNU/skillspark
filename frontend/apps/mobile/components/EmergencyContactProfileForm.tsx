import React from "react";
import { View, TextInput, TouchableOpacity, ScrollView } from "react-native";
import { ThemedText } from "@/components/themed-text";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { AppColors, TAG_COLORS, Colors } from "@/constants/theme";
import { useColorScheme } from "@/hooks/use-color-scheme";
import { SchoolPicker } from "@/components/SchoolPicker";
import { useTranslation } from "react-i18next";

const capitalize = (s: string) => s.charAt(0).toUpperCase() + s.slice(1);

export type EmergencyContactFormProps = {
  name: string;
  setName: (v: string) => void;
  setPhoneNumber: (v: string) => void;
  phoneNumber: string;
};

export function EmergencyContactForm({
  name,
  setName,
  phoneNumber,
  setPhoneNumber,
}: EmergencyContactFormProps) {
  const colorScheme = useColorScheme();
  const theme = Colors[colorScheme ?? "light"];
  const { t: translate } = useTranslation();

  return (
    <>
      <TextInput
        className="rounded-[10px] px-4 py-[14px] text-base font-nunito mb-3 bg-[#F3F4F6] dark:bg-[#27272a] text-[#11181C] dark:text-[#ECEDEE]"
        value={name}
        onChangeText={setName}
        placeholder={translate("emergencyContact.name")}
        placeholderTextColor={AppColors.placeholderText}
      />
      <TextInput
        className="rounded-[10px] px-4 py-[14px] text-base font-nunito mb-3 bg-[#F3F4F6] dark:bg-[#27272a] text-[#11181C] dark:text-[#ECEDEE]"
        value={phoneNumber}
        onChangeText={setPhoneNumber}
        placeholder={translate("emergencyContact.phoneNumber")}
        placeholderTextColor={AppColors.placeholderText}
      />
    </>
  );
}
