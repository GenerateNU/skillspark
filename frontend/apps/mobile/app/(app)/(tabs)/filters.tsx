import { CategoryPicker } from "@/components/filters/CategoryPicker";
import { DateRangePicker } from "@/components/filters/DateRangePicker";
import { SliderCard } from "@/components/filters/SliderCard";
import { SoldOutToggle } from "@/components/filters/SoldOutToggle";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { CATEGORY_KEYS } from "@/constants/eventCategories";
import { AppColors, FontFamilies, FontSizes } from "@/constants/theme";
import { useFilters } from "@/hooks/use-filters";
import { useRouter } from "expo-router";
import { TFunction } from "i18next";
import { useTranslation } from "react-i18next";
import {
  Pressable,
  ScrollView,
  Text,
  TouchableOpacity,
  View,
} from "react-native";
import { useSafeAreaInsets } from "react-native-safe-area-context";

const DISTANCE_MAX = 50;
const DURATION_MAX = 180; // minutes
const PRICE_MAX = 200000; // 2000 in cents, since backend is in cents
const AGE_MAX = 18;

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

export default function FiltersScreen() {
  const router = useRouter();
  const insets = useSafeAreaInsets();
  const { filters, setFilters, clearFilters } = useFilters();
  const { t } = useTranslation();

  const categoryLabels: string[] = t("filters.categories", {
    returnObjects: true,
  }) as string[];

  const categoryIdx = filters.category
    ? CATEGORY_KEYS.indexOf(filters.category)
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
  const startDate = filters.min_date ? new Date(filters.min_date) : undefined;
  const endDate = filters.max_date ? new Date(filters.max_date) : undefined;

  return (
    <View className="flex-1 bg-white" style={{ paddingTop: insets.top }}>
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
        <CategoryPicker
          label={t("filters.category")}
          categoryLabels={categoryLabels}
          activeIndex={categoryIdx}
          onSelect={(idx) =>
            setFilters({
              ...filters,
              category: categoryIdx === idx ? undefined : CATEGORY_KEYS[idx],
            })
          }
        />

        <View
          className="my-5 h-px"
          style={{ backgroundColor: AppColors.divider }}
        />

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
          step={100}
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

        <DateRangePicker
          startDate={startDate}
          endDate={endDate}
          onChange={(start, end) =>
            setFilters({
              ...filters,
              min_date: start?.toISOString(),
              max_date: end?.toISOString(),
            })
          }
        />

        <View
          className="my-5 h-px"
          style={{ backgroundColor: AppColors.divider }}
        />

        <SoldOutToggle
          label={t("filters.showSoldOut")}
          value={soldOut}
          onToggle={() =>
            setFilters({ ...filters, soldout: soldOut ? undefined : true })
          }
        />
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
            {t("filters.search")}
          </Text>
        </TouchableOpacity>
      </View>
    </View>
  );
}
