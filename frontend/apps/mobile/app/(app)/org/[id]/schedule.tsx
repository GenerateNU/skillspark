import {
  ActivityIndicator,
  ScrollView,
  Text,
  TouchableOpacity,
  View,
} from "react-native";
import { useCallback, useEffect, useMemo } from "react";
import { SafeAreaView } from "react-native-safe-area-context";
import { useLocalSearchParams, useRouter } from "expo-router";
import {
  getGetReviewAggregateQueryOptions,
  useGetEventOccurrencesByOrganizationId,
  type EventOccurrence,
} from "@skillspark/api-client";
import { useQueries } from "@tanstack/react-query";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { AppColors, FontFamilies } from "@/constants/theme";
import { useThemeColor } from "@/hooks/use-theme-color";
import { useTranslation } from "react-i18next";
import { OccurrenceCard } from "./OccurrenceCard";
import { formatSectionDate } from "@/utils/format";
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
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [filterClass]);

  const { data: occurrencesResp, isLoading: occurrencesLoading } =
    useGetEventOccurrencesByOrganizationId(id);

  const occurrences = useMemo(() => {
    const d = occurrencesResp as unknown as
      | { data: EventOccurrence[] }
      | undefined;
    return Array.isArray(d?.data) ? d!.data : [];
  }, [occurrencesResp]);

  const filteredOccurrences = useMemo(() => {
    let result = occurrences;

    if (filters.class_name) {
      result = result.filter((o) => o.event.title === filters.class_name);
    }
    if (filters.min_start_minutes !== undefined) {
      result = result.filter((o) => {
        const d = new Date(o.start_time);
        const mins = d.getHours() * 60 + d.getMinutes();
        return mins >= filters.min_start_minutes!;
      });
    }
    if (filters.max_start_minutes !== undefined) {
      result = result.filter((o) => {
        const d = new Date(o.start_time);
        const mins = d.getHours() * 60 + d.getMinutes();
        return mins <= filters.max_start_minutes!;
      });
    }
    if (
      filters.min_duration !== undefined ||
      filters.max_duration !== undefined
    ) {
      result = result.filter((o) => {
        const durationMin =
          (new Date(o.end_time).getTime() -
            new Date(o.start_time).getTime()) /
          60000;
        if (
          filters.min_duration !== undefined &&
          durationMin < filters.min_duration
        )
          return false;
        if (
          filters.max_duration !== undefined &&
          durationMin > filters.max_duration
        )
          return false;
        return true;
      });
    }
    if (filters.min_price !== undefined || filters.max_price !== undefined) {
      result = result.filter((o) => {
        if (filters.min_price !== undefined && o.price < filters.min_price)
          return false;
        if (filters.max_price !== undefined && o.price > filters.max_price)
          return false;
        return true;
      });
    }
    if (filters.min_age !== undefined) {
      result = result.filter(
        (o) => o.event.age_range_max >= filters.min_age!,
      );
    }
    if (filters.max_age !== undefined) {
      result = result.filter(
        (o) => o.event.age_range_min <= filters.max_age!,
      );
    }

    return result;
  }, [occurrences, filters]);

  const uniqueEventIds = useMemo(
    () => [...new Set(occurrences.map((o) => o.event.id))],
    [occurrences],
  );

  const reviewResults = useQueries({
    queries: uniqueEventIds.map((eventId) =>
      getGetReviewAggregateQueryOptions(eventId),
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

  const groupOccurrencesByDate = useCallback(
    (
      occs: EventOccurrence[],
    ): { label: string; items: EventOccurrence[] }[] => {
      const sorted = [...occs].sort(
        (a, b) =>
          new Date(a.start_time).getTime() - new Date(b.start_time).getTime(),
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
    },
    [],
  );

  const grouped = useMemo(
    () => groupOccurrencesByDate(filteredOccurrences),
    [filteredOccurrences, groupOccurrencesByDate],
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
        <TouchableOpacity
          onPress={() => router.push(`/org/${id}/filters` as never)}
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
              className="h-4 w-4 items-center justify-center rounded-full"
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
