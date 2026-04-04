import React from "react";
import { View, Text, TouchableOpacity, Platform } from "react-native";
import BottomSheet, { BottomSheetScrollView } from "@gorhom/bottom-sheet";
import { useTranslation } from "react-i18next";
import { ThemedText } from "@/components/themed-text";
import { useThemeColor } from "@/hooks/use-theme-color";
import { AppColors } from "@/constants/theme";
import type { LocationPin } from "@/components/SkillSparkMap";
import type { LocationObject } from "expo-location";
import { haversineDistance } from "@/utils/distance";
import { OrgCard } from "@/components/OrgCard";

const SNAP_POINTS = [30, "65%"];

interface OrgListSheetProps {
  locations: LocationPin[];
  userLocation: LocationObject | null;
}

export function OrgListSheet({ locations, userLocation }: OrgListSheetProps) {
  const { t: translate } = useTranslation();
  const bgColor = useThemeColor({}, "background");
  const borderColor = useThemeColor({}, "borderColor");

  return (
    <BottomSheet
      index={0}
      snapPoints={SNAP_POINTS}
      enableDynamicSizing={false}
      backgroundStyle={{ backgroundColor: bgColor }}
      handleIndicatorStyle={{ backgroundColor: AppColors.borderLight }}
    >
      <View
        className="flex-row items-center gap-2 border-b px-4 pb-5 pt-3"
        style={{ borderBottomColor: borderColor }}
      >
        <FilterPill label={`${translate("map.filterTime")} ▾`} />
        <FilterPill label={`${translate("map.filterDistance")} ▾`} />
        <FilterPill label={`${translate("map.sortBy")} ▾`} />
        <View className="flex-1" />
        <TouchableOpacity
          className="rounded-full px-4 py-2"
          style={{ backgroundColor: AppColors.checkboxSelected }}
        >
          <Text className="font-nunito-semibold text-sm text-white">
            {translate("map.filter")}
          </Text>
        </TouchableOpacity>
      </View>
      <BottomSheetScrollView
        showsVerticalScrollIndicator={false}
        contentContainerStyle={{
          paddingHorizontal: 16,
          paddingTop: 16,
          paddingBottom: Platform.OS === "ios" ? 40 : 24,
        }}
      >
        {locations.map((loc) => {
          const distance =
            userLocation != null
              ? haversineDistance(
                  userLocation.coords.latitude,
                  userLocation.coords.longitude,
                  loc.latitude,
                  loc.longitude,
                )
              : null;
          return <OrgCard key={loc.id} pin={loc} distance={distance} />;
        })}
      </BottomSheetScrollView>
    </BottomSheet>
  );
}

function FilterPill({ label }: { label: string }) {
  const borderColor = useThemeColor({}, "borderColor");
  return (
    <TouchableOpacity
      className="rounded-full border px-3 py-2"
      style={{ borderColor }}
    >
      <ThemedText className="font-nunito text-sm">{label}</ThemedText>
    </TouchableOpacity>
  );
}