import {
  Pressable,
  ScrollView,
  Text,
  TouchableOpacity,
  View,
} from "react-native";
import { useSafeAreaInsets } from "react-native-safe-area-context";
import { useLocalSearchParams, useRouter } from "expo-router";
import { useMemo } from "react";
import {
  useGetEventOccurrencesByOrganizationId,
  type EventOccurrence,
} from "@skillspark/api-client";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { AppColors, FontFamilies, FontSizes } from "@/constants/theme";
import { SliderCard } from "@/components/filters/SliderCard";
import { useOrgScheduleFilters } from "@/hooks/use-org-schedule-filters";
import { useTranslation } from "react-i18next";
import { TFunction } from "i18next";

const START_TIME_MAX = 1380; // 11:00 PM in minutes from midnight
const DURATION_MAX = 180; // minutes
const PRICE_MAX = 200000; // cents ($2000)
const AGE_MAX = 18;

function formatMinutes(totalMinutes: number): string {
  const h = Math.floor(totalMinutes / 60) % 24;
  const m = totalMinutes % 60;
  const ampm = h < 12 ? "AM" : "PM";
  const displayH = h === 0 ? 12 : h > 12 ? h - 12 : h;
  return `${displayH}:${m.toString().padStart(2, "0")} ${ampm}`;
}

function timeRangeLabel(lo: number, hi: number, t: TFunction): string {
  if (lo === 0 && hi >= START_TIME_MAX) return t("filters.any");
  if (lo === 0) return t("filters.upTo", { value: formatMinutes(hi) });
  if (hi >= START_TIME_MAX)
    return t("filters.orMore", { value: formatMinutes(lo) });
  return `${formatMinutes(lo)} – ${formatMinutes(hi)}`;
}

function durationLabel(min: number) {
  if (min >= DURATION_MAX) return `${DURATION_MAX / 60}h+`;
  if (min < 60) return `${min}m`;
  const h = Math.floor(min / 60);
  const m = min % 60;
  return m === 0 ? `${h}h` : `${h}h ${m}m`;
}

function durationRangeLabel(lo: number, hi: number, t: TFunction) {
  if (lo === 0 && hi >= DURATION_MAX) return t("filters.any");
  if (hi >= DURATION_MAX)
    return t("filters.orMore", { value: durationLabel(lo) });
  if (lo === 0) return t("filters.upTo", { value: durationLabel(hi) });
  return `${durationLabel(lo)} – ${durationLabel(hi)}`;
}

function priceLabel(cents: number): string {
  const amount = cents / 100;
  return `$${amount % 1 === 0 ? amount.toFixed(0) : amount.toFixed(2)}`;
}

function priceRangeLabel(lo: number, hi: number, t: TFunction) {
  if (lo === 0 && hi >= PRICE_MAX) return t("filters.any");
  if (hi >= PRICE_MAX) return t("filters.orMore", { value: priceLabel(lo) });
  if (lo === 0) return t("filters.upTo", { value: priceLabel(hi) });
  return `${priceLabel(lo)} – ${priceLabel(hi)}`;
}

function ageRangeLabel(lo: number, hi: number, t: TFunction) {
  if (lo === 0 && hi >= AGE_MAX) return t("filters.any");
  if (hi >= AGE_MAX) return t("filters.orMore", { value: `${lo}y` });
  if (lo === 0) return t("filters.upTo", { value: `${hi}y` });
  return `${lo} – ${hi}y`;
}

export default function OrgScheduleFiltersScreen() {
  const router = useRouter();
  const insets = useSafeAreaInsets();
  const { id } = useLocalSearchParams<{ id: string }>();
  const { filters, setFilters, clearFilters } = useOrgScheduleFilters(id);
  const { t } = useTranslation();

  const { data: occurrencesResp } = useGetEventOccurrencesByOrganizationId(id);
  const classNames = useMemo(() => {
    const d = occurrencesResp as unknown as
      | { data: EventOccurrence[] }
      | undefined;
    const occs = Array.isArray(d?.data) ? d!.data : [];
    return [...new Set(occs.map((o) => o.event.title))].sort();
  }, [occurrencesResp]);

  const startTimeRange: [number, number] = [
    filters.min_start_minutes ?? 0,
    filters.max_start_minutes ?? START_TIME_MAX,
  ];
  const durationRange: [number, number] = [
    filters.min_duration ?? 0,
    filters.max_duration ?? DURATION_MAX,
  ];
  const priceRange: [number, number] = [
    filters.min_price ?? 0,
    filters.max_price ?? PRICE_MAX,
  ];
  const ageRange: [number, number] = [
    filters.min_age ?? 0,
    filters.max_age ?? AGE_MAX,
  ];

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

        <Pressable onPress={clearFilters} hitSlop={12}>
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
        <SliderCard
          label={t("filters.startTime")}
          valueLabel={timeRangeLabel(startTimeRange[0], startTimeRange[1], t)}
          value={startTimeRange}
          onValueChange={(val) =>
            setFilters({
              ...filters,
              min_start_minutes: val[0] > 0 ? Math.round(val[0]) : undefined,
              max_start_minutes:
                val[1] < START_TIME_MAX ? Math.round(val[1]) : undefined,
            })
          }
          min={0}
          max={START_TIME_MAX}
          step={30}
          minLabel="12:00 AM"
          maxLabel="11:00 PM"
        />

        <View className="h-4" />

        <SliderCard
          label={t("filters.duration")}
          valueLabel={durationRangeLabel(durationRange[0], durationRange[1], t)}
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

        <View className="h-4" />

        <SliderCard
          label={t("filters.price")}
          valueLabel={priceRangeLabel(priceRange[0], priceRange[1], t)}
          value={priceRange}
          onValueChange={(val) =>
            setFilters({
              ...filters,
              min_price: val[0] > 0 ? Math.round(val[0]) : undefined,
              max_price: val[1] < PRICE_MAX ? Math.round(val[1]) : undefined,
            })
          }
          min={0}
          max={PRICE_MAX}
          step={5000}
          minLabel={priceLabel(0)}
          maxLabel={`${priceLabel(PRICE_MAX)}+`}
        />

        <View className="h-4" />

        <SliderCard
          label={t("filters.ageRange")}
          valueLabel={ageRangeLabel(ageRange[0], ageRange[1], t)}
          value={ageRange}
          onValueChange={(val) =>
            setFilters({
              ...filters,
              min_age: val[0] > 0 ? Math.round(val[0]) : undefined,
              max_age: val[1] < AGE_MAX ? Math.round(val[1]) : undefined,
            })
          }
          min={0}
          max={AGE_MAX}
          step={1}
          minLabel="0"
          maxLabel={`${AGE_MAX}+`}
        />

        <View
          className="my-5 h-px"
          style={{ backgroundColor: AppColors.divider }}
        />

        {/* Class filter */}
        <Text
          style={{
            fontFamily: FontFamilies.bold,
            fontSize: FontSizes.lg,
            color: AppColors.primaryText,
            marginBottom: 12,
          }}
        >
          {t("org.class")}
        </Text>

        {/* All Classes option */}
        <TouchableOpacity
          onPress={() => setFilters({ ...filters, class_name: undefined })}
          activeOpacity={0.7}
          className="flex-row items-center justify-between py-3 border-b"
          style={{ borderBottomColor: AppColors.divider }}
        >
          <Text
            style={{
              fontFamily: FontFamilies.regular,
              fontSize: FontSizes.base,
              color: AppColors.primaryText,
            }}
          >
            {t("org.allClasses")}
          </Text>
          <View
            className="w-6 h-6 rounded-full border-2 items-center justify-center"
            style={{
              borderColor: !filters.class_name
                ? AppColors.primaryText
                : AppColors.borderLight,
            }}
          >
            {!filters.class_name && (
              <View
                className="w-3 h-3 rounded-full"
                style={{ backgroundColor: AppColors.primaryText }}
              />
            )}
          </View>
        </TouchableOpacity>

        {/* Individual class options */}
        {classNames.map((name, idx) => (
          <TouchableOpacity
            key={name}
            onPress={() =>
              setFilters({
                ...filters,
                class_name: filters.class_name === name ? undefined : name,
              })
            }
            activeOpacity={0.7}
            className="flex-row items-center justify-between py-3"
            style={
              idx < classNames.length - 1
                ? {
                    borderBottomWidth: 1,
                    borderBottomColor: AppColors.divider,
                  }
                : undefined
            }
          >
            <Text
              style={{
                fontFamily: FontFamilies.regular,
                fontSize: FontSizes.base,
                color: AppColors.primaryText,
                flex: 1,
                marginRight: 12,
              }}
            >
              {name}
            </Text>
            <View
              className="w-6 h-6 rounded-full border-2 items-center justify-center"
              style={{
                borderColor:
                  filters.class_name === name
                    ? AppColors.primaryText
                    : AppColors.borderLight,
              }}
            >
              {filters.class_name === name && (
                <View
                  className="w-3 h-3 rounded-full"
                  style={{ backgroundColor: AppColors.primaryText }}
                />
              )}
            </View>
          </TouchableOpacity>
        ))}
      </ScrollView>

      {/* Show results button */}
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
