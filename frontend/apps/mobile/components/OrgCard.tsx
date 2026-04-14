import React from "react";
import { View, Text, TouchableOpacity } from "react-native";
import { useTranslation } from "react-i18next";
import { useRouter } from "expo-router";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { ThemedText } from "@/components/themed-text";
import { useThemeColor } from "@/hooks/use-theme-color";
import { AppColors } from "@/constants/theme";
import type { LocationPin } from "@/components/SkillSparkMap";

interface OrgCardProps {
  pin: LocationPin;
  distance: number | null;
}

export function OrgCard({ pin, distance }: OrgCardProps) {
  const { t: translate } = useTranslation();
  const router = useRouter();
  const borderColor = useThemeColor({}, "borderColor");

  return (
    <View
      className="mb-5 border-b pb-5"
      style={{ borderBottomColor: borderColor }}
    >
      <View className="flex-row items-center justify-between">
        <View className="flex-1 pr-4">
          <ThemedText className="font-nunito-bold mb-2 text-[23px] leading-[26px]">
            {pin.title}
          </ThemedText>
          <ThemedText className="font-nunito text-[15px] leading-[22px]">
            {pin.members} {translate("dashboard.members")}
          </ThemedText>
          {distance !== null && (
            <ThemedText className="font-nunito text-[15px] leading-[22px]">
              {distance.toFixed(1)} {translate("map.km")}
            </ThemedText>
          )}
          <ThemedText className="font-nunito text-[15px] leading-[22px]">
            {pin.rating.toFixed(1)} {translate("map.stars")}
          </ThemedText>
          <TouchableOpacity
            className="mt-4 w-full items-center rounded-full py-[10px]"
            style={{ backgroundColor: AppColors.checkboxSelected }}
            onPress={() => router.push(`/org/${pin.id}`)}
          >
            <Text className="font-nunito-semibold text-sm text-white">
              {translate("dashboard.learnMore")}
            </Text>
          </TouchableOpacity>
        </View>
        <View
          className="h-[120px] w-[120px] items-center justify-center overflow-hidden rounded-2xl"
          style={{ backgroundColor: AppColors.imagePlaceholder }}
        >
          <IconSymbol name="photo" size={32} color={AppColors.mutedText} />
        </View>
      </View>
    </View>
  );
}
