import React from "react";
import { Text, View, TouchableOpacity } from "react-native";
import { ThemedText } from "@/components/themed-text";
import { AppColors, Colors, FontSizes } from "@/constants/theme";

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

/** Inline section title used between scroll sections (px-5 mb-3 layout). */
export function HomeSectionHeader({ title }: { title: string }) {
  return (
    <Text
      className="font-nunito-bold px-5 mb-3"
      style={{ fontSize: FontSizes.lg, color: AppColors.primaryText }}
    >
      {title}
    </Text>
  );
}
