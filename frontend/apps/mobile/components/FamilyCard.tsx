import React from "react";
import { View } from "react-native";
import { ThemedText } from "@/components/themed-text";
import { Colors, AppColors } from "@/constants/theme";
import { ChildAvatar } from "@/components/ChildAvatar";

type FamilyCardProps = {
  name: string;
  date: string;
  avatarFace?: string | null;
  avatarBackground?: string | null;
};

export function FamilyCard({
  name,
  date,
  avatarFace,
  avatarBackground,
}: FamilyCardProps) {
  const theme = Colors.light;

  return (
    <View
      className="w-[48%] rounded-xl p-[10px] flex-row items-center"
      style={{ backgroundColor: theme.inputBg }}
    >
      <View className="mr-2">
        <ChildAvatar
          name={name}
          avatarFace={avatarFace}
          avatarBackground={avatarBackground}
          size={32}
        />
      </View>
      <View className="flex-1">
        <ThemedText className="text-sm font-nunito-medium" numberOfLines={1}>
          {name}
        </ThemedText>
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
