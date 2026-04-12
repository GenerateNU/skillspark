import { AppColors, FontFamilies, FontSizes } from "@/constants/theme";
import { useFilters } from "@/hooks/use-filters";
import { useRouter } from "expo-router";
import {
  Pressable,
  ScrollView,
  Text,
  TouchableOpacity,
  View,
} from "react-native";
import { useSafeAreaInsets } from "react-native-safe-area-context";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { SliderCard } from "@/components/filters/SliderCard";
import { PillRow } from "@/components/filters/PillRow";
import { SectionLabel, Divider } from "@/components/filters/FilterLayout";

// ─── Constants ────────────────────────────────────────────────────────────────

const CATEGORIES = ["Sport", "Arts", "Music", "Tech", "Activity", "Tutoring"];

const DISTANCE_MAX = 50;
const DURATION_MAX = 180; // minutes
const PRICE_MAX = 2000;
const AGE_MAX = 18;

function todayStart() {
  const d = new Date();
  d.setHours(0, 0, 0, 0);
  return d.toISOString();
}
function todayEnd() {
  const d = new Date();
  d.setHours(23, 59, 59, 999);
  return d.toISOString();
}
function weekEnd() {
  const d = new Date();
  d.setDate(d.getDate() + (6 - d.getDay()));
  d.setHours(23, 59, 59, 999);
  return d.toISOString();
}
function monthEnd() {
  const d = new Date();
  d.setMonth(d.getMonth() + 1, 0);
  d.setHours(23, 59, 59, 999);
  return d.toISOString();
}

type DateOption = { label: string; minDate?: string; maxDate?: string };
const DATE_OPTIONS: DateOption[] = [
  { label: "Any" },
  { label: "Today", minDate: todayStart(), maxDate: todayEnd() },
  { label: "This week", minDate: todayStart(), maxDate: weekEnd() },
  { label: "This month", minDate: todayStart(), maxDate: monthEnd() },
];

// ─── Label helpers ────────────────────────────────────────────────────────────

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
  if (hi >= DURATION_MAX) return `${durationLabel(lo)} +`;
  if (lo === 0) return `Up to ${durationLabel(hi)}`;
  return `${durationLabel(lo)} – ${durationLabel(hi)}`;
}

function priceRangeLabel(lo: number, hi: number) {
  if (lo === 0 && hi >= PRICE_MAX) return "Any";
  if (hi >= PRICE_MAX) return `฿${lo.toLocaleString()} +`;
  if (lo === 0) return `Up to ฿${hi.toLocaleString()}`;
  return `฿${lo.toLocaleString()} – ฿${hi.toLocaleString()}`;
}

function ageRangeLabel(lo: number, hi: number) {
  if (lo === 0 && hi >= AGE_MAX) return "Any";
  if (hi >= AGE_MAX) return `${lo}y +`;
  if (lo === 0) return `Up to ${hi}y`;
  return `${lo} – ${hi}y`;
}

// ─── Main screen ──────────────────────────────────────────────────────────────

export default function FiltersScreen() {
  const router = useRouter();
  const insets = useSafeAreaInsets();
  const { filters, setFilters, clearFilters } = useFilters();

  const categoryIdx = filters.category
    ? CATEGORIES.indexOf(filters.category)
    : -1;
  const distanceKm = filters.radius_km ?? DISTANCE_MAX;
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
  const soldOut = filters.soldout ?? false;
  const dateIdx = Math.max(
    DATE_OPTIONS.findIndex((o) => o.minDate === filters.min_date),
    0,
  );

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
            Back
          </Text>
        </Pressable>

        <Text
          style={{
            fontFamily: FontFamilies.bold,
            fontSize: FontSizes.lg,
            color: AppColors.primaryText,
          }}
        >
          Filters
        </Text>

        <Pressable onPress={clearFilters} hitSlop={12}>
          <Text
            style={{
              fontFamily: FontFamilies.semiBold,
              fontSize: FontSizes.base,
              color: AppColors.primaryBlue,
            }}
          >
            Reset
          </Text>
        </Pressable>
      </View>

      {/* Scrollable content */}
      <ScrollView
        className="flex-1 px-5"
        showsVerticalScrollIndicator={false}
        contentContainerStyle={{ paddingTop: 20, paddingBottom: 32 }}
        keyboardShouldPersistTaps="handled"
      >
        <SectionLabel label="Category" />
        <PillRow
          options={CATEGORIES}
          activeIndex={categoryIdx}
          onSelect={(idx) =>
            setFilters({
              ...filters,
              category: categoryIdx === idx ? undefined : CATEGORIES[idx],
            })
          }
        />

        <Divider />

        <SliderCard
          label="Distance"
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

        <SliderCard
          label="Duration"
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

        <View className="h-4" />

        <SliderCard
          label="Price"
          valueLabel={priceRangeLabel(priceRange[0], priceRange[1])}
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
          step={100}
          minLabel="฿0"
          maxLabel={`฿${PRICE_MAX.toLocaleString()}+`}
        />

        <View className="h-4" />

        <SliderCard
          label="Age Range"
          valueLabel={ageRangeLabel(ageRange[0], ageRange[1])}
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

        <Divider />

        <SectionLabel label="Date" />
        <PillRow
          options={DATE_OPTIONS.map((o) => o.label)}
          activeIndex={dateIdx}
          onSelect={(idx) => {
            const opt = DATE_OPTIONS[idx];
            setFilters({
              ...filters,
              min_date: opt?.minDate,
              max_date: opt?.maxDate,
            });
          }}
        />

        <Divider />

        <TouchableOpacity
          onPress={() =>
            setFilters({ ...filters, soldout: soldOut ? undefined : true })
          }
          activeOpacity={0.7}
          className="flex-row items-center justify-between"
        >
          <Text
            style={{
              fontFamily: FontFamilies.bold,
              fontSize: FontSizes.lg,
              color: AppColors.primaryText,
            }}
          >
            Show sold out
          </Text>
          <View
            style={{
              width: 46,
              height: 26,
              borderRadius: 13,
              backgroundColor: soldOut
                ? AppColors.primaryText
                : AppColors.borderLight,
              justifyContent: "center",
              paddingHorizontal: 3,
            }}
          >
            <View
              style={{
                width: 20,
                height: 20,
                borderRadius: 10,
                backgroundColor: AppColors.white,
                alignSelf: soldOut ? "flex-end" : "flex-start",
              }}
            />
          </View>
        </TouchableOpacity>
      </ScrollView>

      {/* Search button */}
      <View
        className="px-5 pt-4"
        style={{
          paddingBottom: Math.max(insets.bottom, 16) + 8,
          borderTopWidth: 1,
          borderTopColor: AppColors.divider,
        }}
      >
        <TouchableOpacity
          onPress={() => {
            setFilters({ ...filters, search: undefined });
            router.push("/(app)/search");
          }}
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
            Search
          </Text>
        </TouchableOpacity>
      </View>
    </View>
  );
}
