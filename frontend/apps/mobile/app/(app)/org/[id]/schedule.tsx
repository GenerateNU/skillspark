import {
  ActivityIndicator,
  Modal,
  Pressable,
  ScrollView,
  Text,
  TouchableOpacity,
  View,
} from "react-native";
import { useCallback, useMemo, useState } from "react";
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
import { formatSectionDate, extractResponseData } from "@/utils/format";

export default function OrgScheduleScreen() {
  const { id, filterClass } = useLocalSearchParams<{
    id: string;
    filterClass?: string;
  }>();
  const router = useRouter();
  const { t: translate } = useTranslation();
  const backgroundColor = useThemeColor({}, "background");
  const borderColor = useThemeColor({}, "borderColor");

  const [selectedClass, setSelectedClass] = useState<string | null>(
    filterClass ?? null,
  );
  const [filterVisible, setFilterVisible] = useState(false);

  const { data: occurrencesResp, isLoading: occurrencesLoading } =
    useGetEventOccurrencesByOrganizationId(id);

  const occurrences = useMemo(
    () => extractResponseData<EventOccurrence>(occurrencesResp),
    [occurrencesResp],
  );

  const classNames = useMemo(
    () => [...new Set(occurrences.map((o) => o.event.title))].sort(),
    [occurrences],
  );

  const filteredOccurrences = useMemo(
    () =>
      selectedClass
        ? occurrences.filter((o) => o.event.title === selectedClass)
        : occurrences,
    [occurrences, selectedClass],
  );

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
      occurrences: EventOccurrence[],
    ): { label: string; items: EventOccurrence[] }[] => {
      const sorted = [...occurrences].sort(
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
          onPress={() => setFilterVisible(true)}
          activeOpacity={0.7}
          className="rounded-full px-4 py-1.5"
          style={{ backgroundColor: AppColors.primaryText }}
        >
          <Text
            className="text-[13px] font-nunito-bold"
            style={{ color: AppColors.white }}
          >
            {translate("map.filter")}
          </Text>
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

      {/* Class filter modal */}
      <Modal
        visible={filterVisible}
        transparent
        animationType="fade"
        onRequestClose={() => setFilterVisible(false)}
      >
        <Pressable
          className="flex-1 items-center justify-center"
          style={{ backgroundColor: "rgba(0,0,0,0.3)" }}
          onPress={() => setFilterVisible(false)}
        >
          <Pressable
            className="mx-6 w-full rounded-2xl bg-white p-6"
            style={{ maxWidth: 360 }}
            onPress={(e) => e.stopPropagation()}
          >
            <Text
              className="mb-4 text-[22px] font-nunito-bold"
              style={{ color: AppColors.primaryText }}
            >
              {translate("org.schedule")}
            </Text>

            {/* All Classes option */}
            <TouchableOpacity
              onPress={() => {
                setSelectedClass(null);
                setFilterVisible(false);
              }}
              activeOpacity={0.7}
              className="flex-row items-center justify-between py-3 border-b"
              style={{ borderBottomColor: AppColors.divider }}
            >
              <Text
                className="text-[16px] font-nunito"
                style={{ color: AppColors.primaryText }}
              >
                {translate("org.allClasses")}
              </Text>
              <View
                className="w-6 h-6 rounded-full border-2 items-center justify-center"
                style={{
                  borderColor:
                    selectedClass === null
                      ? AppColors.primaryText
                      : AppColors.borderLight,
                }}
              >
                {selectedClass === null && (
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
                onPress={() => {
                  setSelectedClass(name);
                  setFilterVisible(false);
                }}
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
                  className="text-[16px] font-nunito flex-1 mr-3"
                  style={{ color: AppColors.primaryText }}
                >
                  {name}
                </Text>
                <View
                  className="w-6 h-6 rounded-full border-2 items-center justify-center"
                  style={{
                    borderColor:
                      selectedClass === name
                        ? AppColors.primaryText
                        : AppColors.borderLight,
                  }}
                >
                  {selectedClass === name && (
                    <View
                      className="w-3 h-3 rounded-full"
                      style={{ backgroundColor: AppColors.primaryText }}
                    />
                  )}
                </View>
              </TouchableOpacity>
            ))}
          </Pressable>
        </Pressable>
      </Modal>
    </SafeAreaView>
  );
}
