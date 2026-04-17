import { AppColors, FontFamilies, FontSizes } from "@/constants/theme";
import { useFilters } from "@/hooks/use-filters";
import {
  useOrgMapFilters,
  type OrgSortOption,
} from "@/hooks/use-org-map-filters";
import { useRouter } from "expo-router";
import { Pressable, ScrollView, Text, TouchableOpacity, View } from "react-native";
import { useSafeAreaInsets } from "react-native-safe-area-context";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { SliderCard } from "@/components/filters/SliderCard";
import { useTranslation } from "react-i18next";

const DISTANCE_MAX = 50;
const DURATION_MAX = 180;

const SORT_OPTIONS: { key: OrgSortOption; label: string }[] = [
  { key: "distance", label: "Distance" },
  { key: "rating", label: "Rating" },
];

function distanceLabel(km: number) {
  return km >= DISTANCE_MAX ? `${DISTANCE_MAX}+ km` : `${km} km`;
}

function durationLabel(min: number) {
  if (min >= DURATION_MAX) return `${DURATION_MAX / 60}h+`;
  if (min < 60) return `${min}m`;
  const h = Math.floor(min / 60);
  const m = min % 60;
  return m === 0 ? `${h}h` : `${h}h ${m}m`;
}

function durationRangeLabel(lo: number, hi: number) {
  if (lo === 0 && hi >= DURATION_MAX) return "Any";
  if (hi >= DURATION_MAX) return `${durationLabel(lo)}+`;
  if (lo === 0) return `≤ ${durationLabel(hi)}`;
  return `${durationLabel(lo)} – ${durationLabel(hi)}`;
}

export default function MapFiltersScreen() {
  const router = useRouter();
  const insets = useSafeAreaInsets();
  const { t } = useTranslation();
  const { filters, setFilters, clearFilters } = useFilters();
  const {
    filters: mapFilters,
    setFilters: setMapFilters,
    clearFilters: clearMapFilters,
  } = useOrgMapFilters();

  const distanceKm = filters.radius_km ?? DISTANCE_MAX;
  const durationRange: [number, number] = [
    filters.min_duration ?? 0,
    filters.max_duration ?? DURATION_MAX,
  ];
  const sortBy = mapFilters.sort_by ?? null;

  function handleReset() {
    clearFilters();
    clearMapFilters();
  }

  return (
    <View className="flex-1 bg-white" style={{ paddingTop: insets.top }}>
      {/* Header */}
      <View
        className="flex-row items-center justify-between px-5 py-4"
        style={{ borderBottomWidth: 1, borderBottomColor: AppColors.divider }}
      >
        <Pressable
          onPress={() => router.back()}
          hitSlop={12}
          className="flex-row items-center gap-1"
        >
          <IconSymbol
            name="chevron.left"
            size={20}
            color={AppColors.primaryText}
          />
          <Text
            style={{
              fontFamily: FontFamilies.semiBold,
              fontSize: FontSizes.base,
              color: AppColors.primaryText,
            }}
          >
            {t("filters.back")}
          </Text>
        </Pressable>

        <Text
          style={{
            fontFamily: FontFamilies.bold,
            fontSize: FontSizes.lg,
            color: AppColors.primaryText,
          }}
        >
          {t("filters.title")}
        </Text>

        <Pressable onPress={handleReset} hitSlop={12}>
          <Text
            style={{
              fontFamily: FontFamilies.semiBold,
              fontSize: FontSizes.base,
              color: AppColors.primaryBlue,
            }}
          >
            {t("filters.reset")}
          </Text>
        </Pressable>
      </View>

      <ScrollView
        className="flex-1 px-5"
        showsVerticalScrollIndicator={false}
        contentContainerStyle={{ paddingTop: 20, paddingBottom: 32 }}
        keyboardShouldPersistTaps="handled"
      >
        {/* Distance */}
        <SliderCard
          label={t("filters.distance")}
          valueLabel={distanceLabel(distanceKm)}
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

        <View className="h-4" />

        {/* Duration */}
        <SliderCard
          label={t("filters.duration")}
          valueLabel={durationRangeLabel(durationRange[0], durationRange[1])}
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

        <View
          className="my-5 h-px"
          style={{ backgroundColor: AppColors.divider }}
        />

        {/* Sort By */}
        <Text
          style={{
            fontFamily: FontFamilies.bold,
            fontSize: FontSizes.lg,
            color: AppColors.primaryText,
            marginBottom: 12,
          }}
        >
          {t("map.sortBy")}
        </Text>
        <View className="flex-row flex-wrap gap-2">
          {SORT_OPTIONS.map((opt) => (
            <TouchableOpacity
              key={opt.key}
              onPress={() =>
                setMapFilters({
                  sort_by: sortBy === opt.key ? undefined : opt.key,
                })
              }
              className="rounded-full px-5 py-2"
              style={{
                backgroundColor:
                  sortBy === opt.key ? AppColors.primaryText : "#F0F0F0",
              }}
            >
              <Text
                style={{
                  fontFamily: FontFamilies.semiBold,
                  fontSize: FontSizes.sm,
                  color:
                    sortBy === opt.key ? "#fff" : AppColors.primaryText,
                }}
              >
                {t(`map.sort_${opt.key}`)}
              </Text>
            </TouchableOpacity>
          ))}
        </View>
      </ScrollView>

      {/* Show results button — goes back to the map, filters apply automatically */}
      <View
        className="px-5 pt-4"
        style={{
          paddingBottom: Math.max(insets.bottom, 16) + 8,
          borderTopWidth: 1,
          borderTopColor: AppColors.divider,
        }}
      >
        <TouchableOpacity
          onPress={() => router.back()}
          className="rounded-2xl py-4 items-center"
          style={{ backgroundColor: AppColors.primaryText }}
          activeOpacity={0.8}
        >
          <Text
            style={{
              fontFamily: FontFamilies.bold,
              fontSize: FontSizes.base,
              color: AppColors.white,
            }}
          >
            {t("filters.showResults")}
          </Text>
        </TouchableOpacity>
      </View>
    </View>
  );
}
