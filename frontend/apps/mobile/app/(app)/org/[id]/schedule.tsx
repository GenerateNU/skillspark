import {
  ActivityIndicator,
  ScrollView,
  Text,
  TouchableOpacity,
  View,
} from "react-native";
import { useEffect, useMemo } from "react";
import { SafeAreaView } from "react-native-safe-area-context";
import { useLocalSearchParams, useRouter } from "expo-router";
import {
  getGetReviewAggregateQueryOptions,
  useGetEventOccurrencesByOrganizationId,
  type EventOccurrence,
} from "@skillspark/api-client";
import { useQueries, type UseQueryOptions } from "@tanstack/react-query";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { AppColors, FontFamilies } from "@/constants/theme";
import { useThemeColor } from "@/hooks/use-theme-color";
import { useTranslation } from "react-i18next";
import { OccurrenceCard } from "./OccurrenceCard";
import { formatSectionDate, formatSectionMonth } from "@/utils/format";
import { useOrgScheduleFilters } from "@/hooks/use-org-schedule-filters";

export default function OrgScheduleScreen() {
  const { id, filterClass } = useLocalSearchParams<{
    id: string;
    filterClass?: string;
  }>();
  const router = useRouter();
  const { t: translate } = useTranslation();
  const backgroundColor = useThemeColor({}, "background");
  const borderColor = useThemeColor({}, "borderColor");
  const { filters, setFilters, activeCount } = useOrgScheduleFilters(id);

  // Honour the filterClass URL param on first mount
  useEffect(() => {
    if (filterClass && !filters.class_name) {
      setFilters({ ...filters, class_name: filterClass });
    }
  }, [filterClass, filters, setFilters]);

  const { data: occurrencesResp, isLoading: occurrencesLoading } =
    useGetEventOccurrencesByOrganizationId(id);

  const occurrences = useMemo(() => {
    const d = occurrencesResp as unknown as
      | { data: EventOccurrence[] }
      | undefined;
    return Array.isArray(d?.data) ? d!.data : [];
  }, [occurrencesResp]);

  const filteredOccurrences = useMemo(() => {
    const {
      class_name,
      min_start_minutes,
      max_start_minutes,
      min_duration,
      max_duration,
      min_price,
      max_price,
      min_age,
      max_age,
    } = filters;

    return occurrences.filter((o) => {
      if (class_name && o.event.title !== class_name) return false;

      if (min_start_minutes !== undefined || max_start_minutes !== undefined) {
        const d = new Date(o.start_time);
        const mins = d.getHours() * 60 + d.getMinutes();
        if (min_start_minutes !== undefined && mins < min_start_minutes)
          return false;
        if (max_start_minutes !== undefined && mins > max_start_minutes)
          return false;
      }

      if (min_duration !== undefined || max_duration !== undefined) {
        const durationMin =
          (new Date(o.end_time).getTime() - new Date(o.start_time).getTime()) /
          60000;
        if (min_duration !== undefined && durationMin < min_duration)
          return false;
        if (max_duration !== undefined && durationMin > max_duration)
          return false;
      }

      if (min_price !== undefined && o.price < min_price) return false;
      if (max_price !== undefined && o.price > max_price) return false;

      if (min_age !== undefined && o.event.age_range_max < min_age)
        return false;
      if (max_age !== undefined && o.event.age_range_min > max_age)
        return false;

      return true;
    });
  }, [occurrences, filters]);

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

  const grouped = useMemo(() => {
    const sorted = [...filteredOccurrences].sort(
      (a, b) =>
        new Date(a.start_time).getTime() - new Date(b.start_time).getTime()
    );

    const groups = new Map<string, EventOccurrence[]>();
    for (const occ of sorted) {
      const key = new Date(occ.start_time).toDateString();
      if (!groups.has(key)) groups.set(key, []);
      groups.get(key)!.push(occ);
    }

    return Array.from(groups.entries()).map(([, items]) => ({
      month: formatSectionMonth(items[0].start_time),
      label: formatSectionDate(items[0].start_time),
      items,
    }));
  }, [filteredOccurrences]);

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
          className="absolute inset-x-0 text-center text-[16px] font-nunito-bold"
          style={{ color: AppColors.primaryText }}
          numberOfLines={1}
          pointerEvents="none"
        >
          {translate("org.schedule")}
        </Text>
        <View className="flex-1 items-end">
        <TouchableOpacity
          onPress={() => router.push(`/org/${id}/filters`)}
          activeOpacity={0.7}
          className="flex-row items-center gap-1.5 rounded-full px-4 py-1.5"
          style={{ backgroundColor: AppColors.primaryText }}
        >
          <Text
            className="text-[13px] font-nunito-bold"
            style={{ color: AppColors.white }}
          >
            {translate("map.filter")}
          </Text>
          {activeCount > 0 && (
            <View
              className="h-[18px] w-[18px] items-center justify-center rounded-full"
              style={{ backgroundColor: AppColors.white }}
            >
              <Text
                style={{
                  fontFamily: FontFamilies.bold,
                  fontSize: 10,
                  color: AppColors.primaryText,
                }}
              >
                {activeCount}
              </Text>
            </View>
          )}
        </TouchableOpacity>
        </View>
      </View>

      {occurrencesLoading ? (
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
              <View className="px-4 pb-3">
                <Text
                  style={{
                    fontFamily: FontFamilies.regular,
                    fontSize: 11,
                    color: AppColors.primaryText,
                    opacity: 0.6,
                    textTransform: "uppercase",
                    letterSpacing: 0.5,
                  }}
                >
                  {group.month}
                </Text>
                <Text
                  style={{
                    fontFamily: FontFamilies.bold,
                    fontSize: 22,
                    color: AppColors.primaryText,
                  }}
                >
                  {group.label}
                </Text>
              </View>
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
