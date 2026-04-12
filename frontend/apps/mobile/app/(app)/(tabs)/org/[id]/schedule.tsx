import {
  Animated,
  ActivityIndicator,
  Pressable,
  ScrollView,
  Text,
  TouchableOpacity,
  View,
  Image,
} from "react-native";
import { useMemo, useRef, useState } from "react";
import { SafeAreaView } from "react-native-safe-area-context";
import { useLocalSearchParams, useRouter } from "expo-router";
import {
  getGetReviewAggregateQueryOptions,
  useGetEventOccurrencesByOrganizationId,
  type EventOccurrence,
} from "@skillspark/api-client";
import { useQueries } from "@tanstack/react-query";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { AppColors, FontFamilies, FontSizes } from "@/constants/theme";
import { useThemeColor } from "@/hooks/use-theme-color";
import { useTranslation } from "react-i18next";
import { ReservationModal } from "@/components/ReservationModal";
import { RatingSmiley } from "@/components/RatingSmiley";

const BUTTON_ROW_HEIGHT = 52;

function formatSectionDate(dateStr: string): string {
  const date = new Date(dateStr);
  const today = new Date();
  if (date.toDateString() === today.toDateString()) return "Today";
  return date.toLocaleDateString("en-US", { weekday: "short", day: "numeric" });
}

function formatTime(dateStr: string): string {
  return new Date(dateStr).toLocaleTimeString("en-US", {
    hour: "numeric",
    minute: "2-digit",
    hour12: true,
  });
}

function formatPrice(cents: number, currency: string): string {
  const amount = cents / 100;
  if (currency?.toUpperCase() === "THB")
    return `฿${amount % 1 === 0 ? amount.toFixed(0) : amount.toFixed(2)}`;
  return `$${amount % 1 === 0 ? amount.toFixed(0) : amount.toFixed(2)}`;
}

function formatAgeRange(min: number, max: number): string {
  if (!min && !max) return "";
  if (min === max) return `Ages ${min}`;
  return `Ages ${min} - ${max}`;
}

function groupOccurrencesByDate(
  occurrences: EventOccurrence[]
): { label: string; items: EventOccurrence[] }[] {
  const sorted = [...occurrences].sort(
    (a, b) =>
      new Date(a.start_time).getTime() - new Date(b.start_time).getTime()
  );

  const groups: Map<string, EventOccurrence[]> = new Map();
  for (const occ of sorted) {
    const key = new Date(occ.start_time).toDateString();
    if (!groups.has(key)) groups.set(key, []);
    groups.get(key)!.push(occ);
  }

  return Array.from(groups.entries()).map(([, items]) => ({
    label: formatSectionDate(items[0].start_time),
    items,
  }));
}

function OccurrenceCard({
  occurrence,
  avgRating,
}: {
  occurrence: EventOccurrence;
  avgRating: number | null;
}) {
  const router = useRouter();
  const [expanded, setExpanded] = useState(false);
  const [reservationVisible, setReservationVisible] = useState(false);
  const progress = useRef(new Animated.Value(0)).current;

  const toggle = () => {
    const toValue = expanded ? 0 : 1;
    setExpanded((prev) => !prev);
    Animated.timing(progress, {
      toValue,
      duration: 250,
      useNativeDriver: false,
    }).start();
  };

  const buttonsAnimStyle = {
    height: progress.interpolate({
      inputRange: [0, 1],
      outputRange: [0, BUTTON_ROW_HEIGHT],
    }),
    opacity: progress,
    overflow: "hidden" as const,
  };

  const chevronAnimStyle = {
    transform: [
      {
        rotate: progress.interpolate({
          inputRange: [0, 1],
          outputRange: ["0deg", "180deg"],
        }),
      },
    ],
  };

  const ageRange = formatAgeRange(
    occurrence.event.age_range_min,
    occurrence.event.age_range_max
  );
  const locationText = [
    occurrence.location?.district,
    occurrence.location?.province,
  ]
    .filter(Boolean)
    .join(", ");

  return (
    <>
      <Pressable
        onPress={toggle}
        className="mx-4 mb-3 rounded-2xl bg-white"
        style={{
          shadowColor: "#000",
          shadowOpacity: 0.1,
          shadowRadius: 8,
          shadowOffset: { width: 0, height: 10 },
          elevation: 3,
        }}
      >
        <View className="rounded-2xl overflow-hidden">
          <View className="flex-row p-3 gap-3">
            {/* Thumbnail */}
            <View
              className="w-[80px] h-[80px] rounded-xl overflow-hidden flex-shrink-0"
              style={{ backgroundColor: AppColors.imagePlaceholder }}
            >
              {occurrence.event.presigned_url ? (
                <Image
                  source={{ uri: occurrence.event.presigned_url }}
                  style={{ width: "100%", height: "100%" }}
                />
              ) : null}
            </View>

            {/* Info */}
            <View className="flex-1 justify-center gap-[3px]">
              <Text
                numberOfLines={1}
                style={{
                  fontFamily: FontFamilies.bold,
                  fontSize: FontSizes.base,
                  color: AppColors.primaryText,
                }}
              >
                {occurrence.event.title}
              </Text>

              <View className="flex-row items-center gap-1">
                <View className="pr-2">
                  <RatingSmiley rating={avgRating} width={16} height={16} />
                </View>
                <Text
                  style={{
                    fontFamily: FontFamilies.regular,
                    fontSize: FontSizes.sm,
                  }}
                >
                  {avgRating ? avgRating : "~"} / 5 Smiles
                </Text>
              </View>

              {!!ageRange && (
                <Text
                  style={{
                    fontFamily: FontFamilies.regular,
                    fontSize: FontSizes.sm,
                  }}
                >
                  {ageRange}
                </Text>
              )}

              {!!locationText && (
                <Text
                  numberOfLines={1}
                  style={{
                    fontFamily: FontFamilies.regular,
                    fontSize: FontSizes.sm,
                  }}
                >
                  {locationText}
                </Text>
              )}
            </View>

            {/* Time & Price */}
            <View className="items-end justify-center gap-1 flex-shrink-0">
              <Text
                style={{
                  fontFamily: FontFamilies.semiBold,
                  fontSize: FontSizes.sm,
                  color: AppColors.primaryText,
                }}
              >
                {formatTime(occurrence.start_time)}
              </Text>
              <Text
                style={{
                  fontFamily: FontFamilies.semiBold,
                  fontSize: FontSizes.base,
                  color: AppColors.primaryText,
                }}
              >
                {formatPrice(occurrence.price, occurrence.currency)}
              </Text>
            </View>
          </View>

          {/* Action buttons (animated expand/collapse) */}
          <Animated.View style={buttonsAnimStyle}>
            <View className="flex-row gap-3 px-3 pb-3">
              <TouchableOpacity
                onPress={() => router.push(`/event/${occurrence.id}`)}
                activeOpacity={0.7}
                className="flex-1 rounded-full py-2.5 items-center"
                style={{ backgroundColor: "#99C0EE" }}
              >
                <Text
                  style={{
                    fontFamily: FontFamilies.regular,
                    fontSize: FontSizes.base,
                  }}
                >
                  Learn more
                </Text>
              </TouchableOpacity>
              <TouchableOpacity
                onPress={() => setReservationVisible(true)}
                activeOpacity={0.7}
                className="flex-1 rounded-full py-2.5 items-center"
                style={{ backgroundColor: AppColors.checkboxSelected }}
              >
                <Text
                  style={{
                    fontFamily: FontFamilies.semiBold,
                    fontSize: FontSizes.base,
                    color: "#fff",
                  }}
                >
                  Reserve
                </Text>
              </TouchableOpacity>
            </View>
          </Animated.View>

          {/* Expand/collapse chevron (rotates 180° when expanded) */}
          <View className="items-center pb-2 -mt-1">
            <Animated.View style={chevronAnimStyle}>
              <IconSymbol name="chevron.down" size={16} color="#000000" />
            </Animated.View>
          </View>
        </View>
      </Pressable>

      <ReservationModal
        visible={reservationVisible}
        onClose={() => setReservationVisible(false)}
        occurrence={occurrence}
      />
    </>
  );
}

export default function OrgScheduleScreen() {
  const { id } = useLocalSearchParams<{ id: string }>();
  const router = useRouter();
  const { t: translate } = useTranslation();
  const backgroundColor = useThemeColor({}, "background");
  const borderColor = useThemeColor({}, "borderColor");

  const { data: occurrencesResp, isLoading: occurrencesLoading } =
    useGetEventOccurrencesByOrganizationId(id);

  const occurrences = useMemo(() => {
    const d = occurrencesResp as unknown as
      | { data: EventOccurrence[] }
      | undefined;
    return Array.isArray(d?.data) ? d!.data : [];
  }, [occurrencesResp]);

  const uniqueEventIds = useMemo(
    () => [...new Set(occurrences.map((o) => o.event.id))],
    [occurrences]
  );

  const reviewResults = useQueries({
    queries: uniqueEventIds.map((eventId) =>
      getGetReviewAggregateQueryOptions(eventId)
    ),
  });

  const ratingsMap = useMemo(() => {
    const map = new Map<string, number | null>();
    uniqueEventIds.forEach((eventId, i) => {
      const apiResult = reviewResults[i];
      const mapValue =
        apiResult?.data?.status === 200 &&
        apiResult.data.data.total_reviews !== 0
          ? apiResult.data.data.average_rating
          : null;

      map.set(eventId, mapValue);
    });
    return map;
  }, [uniqueEventIds, reviewResults]);

  const isLoading = occurrencesLoading;

  const grouped = useMemo(
    () => groupOccurrencesByDate(occurrences),
    [occurrences]
  );

  return (
    <SafeAreaView
      className="flex-1"
      style={{ backgroundColor }}
      edges={["top", "bottom"]}
    >
      {/* Header */}
      <View
        className="flex-row items-center border-b px-4 pb-2.5 pt-3"
        style={{ backgroundColor, borderBottomColor: borderColor }}
      >
        <TouchableOpacity
          onPress={() => router.back()}
          activeOpacity={0.7}
          className="h-8 w-8 items-center justify-center"
        >
          <IconSymbol
            name="chevron.left"
            size={28}
            color={AppColors.primaryText}
          />
        </TouchableOpacity>
        <Text
          className="flex-1 text-center text-[16px] font-nunito-bold"
          style={{ color: AppColors.primaryText }}
          numberOfLines={1}
        >
          {translate("org.schedule")}
        </Text>
        <View className="w-8" />
      </View>

      {isLoading ? (
        <View className="flex-1 items-center justify-center">
          <ActivityIndicator size="large" />
        </View>
      ) : grouped.length === 0 ? (
        <View className="flex-1 items-center justify-center px-6">
          <Text
            className="text-base text-center"
            style={{
              fontFamily: FontFamilies.regular,
            }}
          >
            {translate("org.noSchedule")}
          </Text>
        </View>
      ) : (
        <ScrollView
          showsVerticalScrollIndicator={false}
          contentContainerStyle={{ paddingBottom: 32, paddingTop: 16 }}
        >
          {grouped.map((group) => (
            <View key={group.label} className="mb-4">
              <Text
                className="px-4 pb-3"
                style={{
                  fontFamily: FontFamilies.bold,
                  fontSize: 22,
                  color: AppColors.primaryText,
                }}
              >
                {group.label}
              </Text>
              {group.items.map((occ) => (
                <OccurrenceCard
                  key={occ.id}
                  occurrence={occ}
                  avgRating={ratingsMap.get(occ.event.id) ?? null}
                />
              ))}
            </View>
          ))}
        </ScrollView>
      )}
    </SafeAreaView>
  );
}
