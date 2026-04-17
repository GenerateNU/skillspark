import React from "react";
import { View, Text, TouchableOpacity } from "react-native";
import { Image } from "expo-image";
import { useTranslation } from "react-i18next";
import { useRouter } from "expo-router";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { ThemedText } from "@/components/themed-text";
import { useThemeColor } from "@/hooks/use-theme-color";
import { AppColors } from "@/constants/theme";
import { RatingSmiley } from "@/components/RatingSmiley";
import type { LocationPin } from "@/components/SkillSparkMap";

interface OrgCardProps {
  pin: LocationPin;
  distance: number | null;
}

export function OrgCard({ pin, distance }: OrgCardProps) {
  const { t: translate } = useTranslation();
  const router = useRouter();
  const borderColor = useThemeColor({}, "borderColor");

  const locationLine = [
    pin.district && pin.district.length > 0 ? pin.district : null,
    distance !== null ? `${distance.toFixed(1)} mi` : null,
  ]
    .filter(Boolean)
    .join(" ");

  return (
    <View className="py-4 border-b" style={{ borderBottomColor: borderColor }}>
      {/* Top row: text + image */}
      <View className="flex-row mb-3">
        <View className="flex-1 pr-3 justify-center gap-[3px]">
          <ThemedText className="font-nunito-bold text-[17px] leading-[22px]">
            {pin.title}
          </ThemedText>
          {pin.members > 0 && (
            <ThemedText className="font-nunito text-[13px] leading-[19px] text-[#6B7280]">
              {pin.members} {translate("dashboard.members")}
            </ThemedText>
          )}
          {locationLine.length > 0 && (
            <ThemedText className="font-nunito text-[13px] leading-[19px] text-[#6B7280]">
              {locationLine}
            </ThemedText>
          )}
          {pin.description != null && pin.description.length > 0 && (
            <ThemedText
              className="font-nunito text-[13px] leading-[19px] text-[#6B7280]"
              numberOfLines={1}
            >
              {pin.description}
            </ThemedText>
          )}
          {pin.rating > 0 && (
            <View className="flex-row items-center gap-[5px] mt-[2px]">
              <RatingSmiley rating={pin.rating} width={14} height={14} />
              <ThemedText className="font-nunito text-[13px] text-[#6B7280]">
                {pin.rating.toFixed(1)} / 5 {translate("map.smiles")}
              </ThemedText>
            </View>
          )}
        </View>

        {/* Image */}
        <View
          className="w-[100px] h-[100px] overflow-hidden rounded-lg"
          style={{ backgroundColor: AppColors.imagePlaceholder }}
        >
          {pin.image ? (
            <Image
              source={{ uri: pin.image }}
              className="h-full w-full"
              contentFit="cover"
            />
          ) : (
            <View className="flex-1 items-center justify-center">
              <IconSymbol name="photo" size={24} color={AppColors.mutedText} />
            </View>
          )}
        </View>
      </View>

      {/* Full-width Learn more button */}
      <TouchableOpacity
        className="w-full items-center rounded-xl bg-[#1C1C1E] py-[11px]"
        activeOpacity={0.85}
        onPress={() => router.push(`/org/${pin.id}`)}
      >
        <Text className="font-nunito-semibold text-sm text-white">
          {translate("map.learnMore")}
        </Text>
      </TouchableOpacity>
    </View>
  );
}
