import {
  ActivityIndicator,
  ScrollView,
  Text,
  TouchableOpacity,
  View,
} from "react-native";
import { useMemo } from "react";
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
import { groupOccurrencesByDate } from "./utils";

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
