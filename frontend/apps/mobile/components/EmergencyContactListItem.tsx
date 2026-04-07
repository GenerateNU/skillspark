import React from "react";
import { View, TouchableOpacity } from "react-native";
import { ThemedText } from "@/components/themed-text";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { InterestTags } from "./InterestTags";
import { EmergencyContact } from "@skillspark/api-client";
import { AppColors } from "@/constants/theme";
import { useTranslation } from "react-i18next";

type EmergencyContactListItemProps = {
  emergencyContact: EmergencyContact;
  onPress?: () => void;
};

export function EmergencyContactListItem({
  emergencyContact,
  onPress,
}: EmergencyContactListItemProps) {
  const { t: translate } = useTranslation();
  return (
    <TouchableOpacity
      className="flex-row items-start py-3 gap-3"
      onPress={onPress}
      activeOpacity={0.7}
    >
      <View
        className="w-11 h-11 rounded-[22px] border-[1.5px] items-center justify-center"
        style={{ borderColor: AppColors.subtleText }}
      >
        <ThemedText className="text-[15px] font-nunito-semibold">
          {emergencyContact.name?.slice(0, 2).toUpperCase() || "??"}
        </ThemedText>
      </View>
      <View className="flex-1 gap-[2px]">
        <ThemedText className="text-base font-nunito-semibold">
          {emergencyContact.name}
        </ThemedText>
        <ThemedText
          className="text-[13px] font-nunito"
          style={{ color: AppColors.mutedText }}
        >
          {emergencyContact.phone_number}
        </ThemedText>
      </View>
      <IconSymbol name="chevron.right" size={18} color={AppColors.subtleText} />
    </TouchableOpacity>
  );
}
