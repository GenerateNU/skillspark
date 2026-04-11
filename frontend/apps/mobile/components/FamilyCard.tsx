import React from "react";
import { View } from "react-native";
import { ThemedText } from "@/components/themed-text";
import { Colors, AppColors } from "@/constants/theme";

type FamilyCardProps = {
  initials: string;
  name: string;
  date: string;
};

export function FamilyCard({ initials, name, date }: FamilyCardProps) {
  const theme = Colors.light;

  return (
    <View
      className="w-[48%] rounded-xl p-[10px] flex-row items-center"
      style={{ backgroundColor: theme.inputBg }}
    >
      <View
        className="w-8 h-8 rounded-2xl border items-center justify-center mr-2"
        style={{ borderColor: theme.borderColor }}
      >
        <ThemedText className="text-xs font-nunito-semibold">
          {initials}
        </ThemedText>
      </View>
      <View>
        <ThemedText className="text-sm font-nunito-medium">{name}</ThemedText>
        <ThemedText
          className="text-[10px] font-nunito"
          style={{ color: AppColors.mutedText }}
        >
          {date}
        </ThemedText>
      </View>
    </View>
  );
}
