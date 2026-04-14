import React from "react";
import { View, TouchableOpacity } from "react-native";
import { ThemedText } from "@/components/themed-text";
import { Colors } from "@/constants/theme";

type SectionHeaderProps = {
  title: string;
  actionLabel?: string;
  onAction?: () => void;
};

export function SectionHeader({
  title,
  actionLabel,
  onAction,
}: SectionHeaderProps) {
  const theme = Colors.light;

  return (
    <View className="flex-row justify-between items-center py-[14px]">
      <ThemedText className="text-[17px] font-nunito-bold">{title}</ThemedText>
      {actionLabel && (
        <TouchableOpacity onPress={onAction}>
          <ThemedText
            className="text-sm font-nunito-medium"
            style={{ color: theme.tint }}
          >
            {actionLabel}
          </ThemedText>
        </TouchableOpacity>
      )}
    </View>
  );
}
