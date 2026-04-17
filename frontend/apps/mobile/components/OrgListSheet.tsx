import React, { useRef, useState, useMemo } from "react";
import { View, Text, TouchableOpacity, StyleSheet } from "react-native";
import BottomSheet, { BottomSheetScrollView } from "@gorhom/bottom-sheet";
import { useTranslation } from "react-i18next";
import { ThemedText } from "@/components/themed-text";
import { useThemeColor } from "@/hooks/use-theme-color";
import { AppColors, FontFamilies, FontSizes } from "@/constants/theme";
import type { LocationPin } from "@/components/SkillSparkMap";
import type { LocationObject } from "expo-location";
import { haversineDistance } from "@/utils/distance";
import { OrgCard } from "@/components/OrgCard";
import { FLOATING_TAB_BAR_SCROLL_PADDING } from "@/components/floating-tab-bar";
import { SliderCard } from "@/components/filters/SliderCard";
import { useFilters } from "@/hooks/use-filters";
import { useOrgMapFilters, type OrgSortOption } from "@/hooks/use-org-map-filters";
import { useRouter } from "expo-router";

// Index 0 = handle strip above the floating nav bar (~120 px tall)
// Index 1 = full list view with filters
const SNAP_POINTS = [140, "75%"];

const DISTANCE_MAX = 50;
const DURATION_MAX = 180; // minutes

type ActivePill = "time" | "distance" | "sortBy" | null;

interface OrgListSheetProps {
  locations: LocationPin[];
  userLocation: LocationObject | null;
}

export function OrgListSheet({ locations, userLocation }: OrgListSheetProps) {
  const { t } = useTranslation();
  const borderColor = useThemeColor({}, "borderColor");
  const sheetRef = useRef<BottomSheet>(null);
  const router = useRouter();
  const { filters, setFilters } = useFilters();
  const { filters: mapFilters, setFilters: setMapFilters } = useOrgMapFilters();
  const [sheetIndex, setSheetIndex] = useState(0);
  const [activePill, setActivePill] = useState<ActivePill>(null);

  const sortBy = mapFilters.sort_by ?? null;
  const distanceKm = filters.radius_km ?? DISTANCE_MAX;
  const durationRange: [number, number] = [
    filters.min_duration ?? 0,
    filters.max_duration ?? DURATION_MAX,
  ];

  function handlePillPress(pill: ActivePill) {
    setActivePill(activePill === pill ? null : pill);
  }

  const filteredAndSorted = useMemo(() => {
    let result = [...locations];

    if (filters.radius_km != null && userLocation != null) {
      result = result.filter((loc) => {
        const d = haversineDistance(
          userLocation.coords.latitude,
          userLocation.coords.longitude,
          loc.latitude,
          loc.longitude,
        );
        return d <= filters.radius_km!;
      });
    }

    if (sortBy === "distance" && userLocation != null) {
      result.sort((a, b) => {
        const da = haversineDistance(
          userLocation.coords.latitude,
          userLocation.coords.longitude,
          a.latitude,
          a.longitude,
        );
        const db = haversineDistance(
          userLocation.coords.latitude,
          userLocation.coords.longitude,
          b.latitude,
          b.longitude,
        );
        return da - db;
      });
    } else if (sortBy === "rating") {
      result.sort((a, b) => b.rating - a.rating);
    }

    return result;
  }, [locations, userLocation, filters.radius_km, sortBy]);

  const isDistanceActive = filters.radius_km !== undefined;
  const isTimeActive =
    filters.min_duration !== undefined || filters.max_duration !== undefined;
  const isSortActive = sortBy !== null;
  const isExpanded = sheetIndex === 1;

  return (
    <BottomSheet
      ref={sheetRef}
      index={0}
      snapPoints={SNAP_POINTS}
      enableDynamicSizing={false}
      backgroundStyle={{ backgroundColor: "#fff" }}
      handleIndicatorStyle={{ backgroundColor: AppColors.borderLight }}
      onChange={(index) => {
        setSheetIndex(index);
        if (index === 0) setActivePill(null);
      }}
    >
      {/* White overlay covers card content while collapsed so nothing peeks
          through around the floating nav-bar pill. */}
      {!isExpanded && (
        <View
          pointerEvents="none"
          style={[StyleSheet.absoluteFillObject, { backgroundColor: "#fff" }]}
        />
      )}

      {/* Filter pills + mini-panels — only rendered when fully expanded */}
        <>
          <View
            className="flex-row items-center gap-2 px-4 pb-3 pt-1"
            style={{ borderBottomWidth: 1, borderBottomColor: borderColor }}
          >
            <FilterPill
              label={`${t("map.filterTime")} ▾`}
              active={isTimeActive || activePill === "time"}
              onPress={() => handlePillPress("time")}
            />
            <FilterPill
              label={`${t("map.filterDistance")} ▾`}
              active={isDistanceActive || activePill === "distance"}
              onPress={() => handlePillPress("distance")}
            />
            <FilterPill
              label={`${t("map.sortBy")} ▾`}
              active={isSortActive || activePill === "sortBy"}
              onPress={() => handlePillPress("sortBy")}
            />
            <View className="flex-1" />
            <TouchableOpacity
              className="rounded-full bg-[#1C1C1E] px-4 py-2"
              onPress={() => router.push("/(app)/map-filters")}
            >
              <Text className="font-nunito-semibold text-sm text-white">
                {t("map.filter")}
              </Text>
            </TouchableOpacity>
          </View>

          {/* Mini-panel: Distance */}
          {activePill === "distance" && (
            <View
              className="px-4 py-4"
              style={{ borderBottomWidth: 1, borderBottomColor: borderColor }}
            >
              <SliderCard
                label={t("filters.distance")}
                valueLabel={distanceLabelFn(distanceKm)}
                value={distanceKm}
                onValueChange={(val) =>
                  setFilters({
                    ...filters,
                    radius_km: val[0] < DISTANCE_MAX ? val[0] : undefined,
                  })
                }
                min={0}
                max={DISTANCE_MAX}
                step={5}
                minLabel="0 km"
                maxLabel={`${DISTANCE_MAX}+ km`}
              />
            </View>
          )}

          {/* Mini-panel: Time (Duration) */}
          {activePill === "time" && (
            <View
              className="px-4 py-4"
              style={{ borderBottomWidth: 1, borderBottomColor: borderColor }}
            >
              <SliderCard
                label={t("filters.duration")}
                valueLabel={durationRangeLabelFn(
                  durationRange[0],
                  durationRange[1],
                )}
                value={durationRange}
                onValueChange={(val) =>
                  setFilters({
                    ...filters,
                    min_duration: val[0] > 0 ? Math.round(val[0]) : undefined,
                    max_duration:
                      val[1] < DURATION_MAX ? Math.round(val[1]) : undefined,
                  })
                }
                min={0}
                max={DURATION_MAX}
                step={15}
                minLabel="0"
                maxLabel={`${DURATION_MAX / 60}h+`}
              />
            </View>
          )}

          {/* Mini-panel: Sort By */}
          {activePill === "sortBy" && (
            <View
              className="px-4 pb-4 pt-3"
              style={{ borderBottomWidth: 1, borderBottomColor: borderColor }}
            >
              <Text
                style={{
                  fontFamily: FontFamilies.bold,
                  fontSize: FontSizes.base,
                  color: AppColors.primaryText,
                  marginBottom: 10,
                }}
              >
                {t("map.sortBy")}
              </Text>
              <View className="flex-row gap-2">
                {(["distance", "rating"] as OrgSortOption[]).map(
                  (opt) => (
                    <TouchableOpacity
                      key={opt}
                      onPress={() =>
                        setMapFilters({
                          sort_by: sortBy === opt ? undefined : opt,
                        })
                      }
                      className="rounded-full px-4 py-2"
                      style={{
                        backgroundColor:
                          sortBy === opt ? AppColors.primaryText : "#F0F0F0",
                      }}
                    >
                      <Text
                        style={{
                          fontFamily: FontFamilies.semiBold,
                          fontSize: FontSizes.sm,
                          color:
                            sortBy === opt ? "#fff" : AppColors.primaryText,
                        }}
                      >
                        {t(`map.sort_${opt}`)}
                      </Text>
                    </TouchableOpacity>
                  ),
                )}
              </View>
            </View>
          )}
        </>

      {/* Org list */}
      <BottomSheetScrollView
        showsVerticalScrollIndicator={false}
        contentContainerStyle={{
          paddingHorizontal: 16,
          paddingTop: 16,
          paddingBottom: FLOATING_TAB_BAR_SCROLL_PADDING,
        }}
      >
        {filteredAndSorted.map((loc) => {
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

function distanceLabelFn(km: number) {
  return km >= DISTANCE_MAX ? `${DISTANCE_MAX}+ km` : `${km} km`;
}

function durationLabelFn(min: number) {
  if (min >= DURATION_MAX) return `${DURATION_MAX / 60}h+`;
  if (min < 60) return `${min}m`;
  const h = Math.floor(min / 60);
  const m = min % 60;
  return m === 0 ? `${h}h` : `${h}h ${m}m`;
}

function durationRangeLabelFn(lo: number, hi: number) {
  if (lo === 0 && hi >= DURATION_MAX) return "Any";
  if (hi >= DURATION_MAX) return `${durationLabelFn(lo)}+`;
  if (lo === 0) return `<= ${durationLabelFn(hi)}`;
  return `${durationLabelFn(lo)} - ${durationLabelFn(hi)}`;
}

interface FilterPillProps {
  label: string;
  active?: boolean;
  onPress?: () => void;
}

function FilterPill({ label, active = false, onPress }: FilterPillProps) {
  const borderColor = useThemeColor({}, "borderColor");
  return (
    <TouchableOpacity
      onPress={onPress}
      className="rounded-full border px-3 py-2"
      style={{
        borderColor: active ? AppColors.primaryText : borderColor,
        backgroundColor: active ? AppColors.primaryText : "transparent",
      }}
    >
      <ThemedText
        className="font-nunito text-sm"
        style={{ color: active ? "#fff" : undefined }}
      >
        {label}
      </ThemedText>
    </TouchableOpacity>
  );
}
