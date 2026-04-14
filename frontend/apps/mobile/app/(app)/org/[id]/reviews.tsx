import { useLocalSearchParams, useRouter } from "expo-router";
import { useQueries } from "@tanstack/react-query";
import {
  useGetReviewAggregateOrganization,
  useGetEventOccurrencesByOrganizationId,
  getGetReviewAggregateQueryOptions,
  type ReviewAggregate,
  type Event,
} from "@skillspark/api-client";
import { useMemo, useState } from "react";
import {
  ActivityIndicator,
  ScrollView,
  TouchableOpacity,
  View,
} from "react-native";
import { useSafeAreaInsets } from "react-native-safe-area-context";
import { useTranslation } from "react-i18next";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { ThemedText } from "@/components/themed-text";
import { ThemedView } from "@/components/themed-view";
import { RatingAggregateCard } from "@/components/ReviewAggregate";
import { FilterTabs } from "@/components/SortingButtons";
import { EventRatingCard } from "@/components/EventRatingCard";
import { Colors } from "@/constants/theme";

type SortValue = "most_rated" | "highest" | "lowest";

export default function OrgReviewsPage() {
  const { id } = useLocalSearchParams<{ id: string }>();
  const router = useRouter();
  const insets = useSafeAreaInsets();
  const { t: translate } = useTranslation();
  const theme = Colors.light;
  const [sortBy, setSortBy] = useState<SortValue>("most_rated");

  const { data: orgAggregateResp, isLoading: orgAggLoading } =
    useGetReviewAggregateOrganization(id ?? "", {
      query: { enabled: !!id },
    });

  const { data: occurrencesResp, isLoading: occurrencesLoading } =
    useGetEventOccurrencesByOrganizationId(id ?? "", {
      query: { enabled: !!id },
    });

  const uniqueEvents = useMemo(() => {
    if (occurrencesResp?.status !== 200) return [];
    const seen = new Set<string>();
    const events: Event[] = [];
    for (const occ of occurrencesResp.data) {
      if (!seen.has(occ.event.id)) {
        seen.add(occ.event.id);
        events.push(occ.event);
      }
    }
    return events;
  }, [occurrencesResp]);

  const eventAggregateQueries = useQueries({
    queries: uniqueEvents.map((event) =>
      getGetReviewAggregateQueryOptions(event.id),
    ),
  });

  if (orgAggLoading || occurrencesLoading) {
    return (
      <View className="flex-1 items-center justify-center">
        <ActivityIndicator size="large" />
      </View>
    );
  }

  const orgAggregate =
    orgAggregateResp?.status === 200
      ? (orgAggregateResp.data as ReviewAggregate)
      : null;

  const eventItems = uniqueEvents.map((event, i) => {
    const aggResp = eventAggregateQueries[i]?.data;
    const aggregate =
      aggResp?.status === 200 ? (aggResp.data as ReviewAggregate) : null;
    return { event, aggregate };
  });

  const sorted = [...eventItems].sort((a, b) => {
    if (sortBy === "most_rated") {
      return (
        (b.aggregate?.total_reviews ?? 0) - (a.aggregate?.total_reviews ?? 0)
      );
    }
    if (sortBy === "highest") {
      return (
        (b.aggregate?.average_rating ?? 0) - (a.aggregate?.average_rating ?? 0)
      );
    }
    if (sortBy === "lowest") {
      return (
        (a.aggregate?.average_rating ?? 0) - (b.aggregate?.average_rating ?? 0)
      );
    }
    return 0;
  });

  return (
    <ThemedView className="flex-1" style={{ paddingTop: insets.top }}>
      <View className="flex-row items-center justify-between px-5 py-[14px]">
        <TouchableOpacity
          onPress={() => router.back()}
          className="w-10 justify-center items-start"
          hitSlop={{ top: 10, bottom: 10, left: 10, right: 10 }}
        >
          <IconSymbol name="chevron.left" size={24} color={theme.text} />
        </TouchableOpacity>
        <ThemedText className="text-xl text-center font-nunito-bold">
          {translate("review.title")}
        </ThemedText>
        <View className="w-5" />
      </View>

      <ScrollView
        showsVerticalScrollIndicator={false}
        contentContainerStyle={{ paddingBottom: 32 }}
      >
        {orgAggregate && <RatingAggregateCard aggregate={orgAggregate} />}

        <FilterTabs
          options={[
            { label: translate("review.mostRated"), value: "most_rated" },
            { label: translate("review.highest"), value: "highest" },
            { label: translate("review.lowest"), value: "lowest" },
          ]}
          onChange={(value) => setSortBy(value as SortValue)}
        />

        <View className="px-5 gap-3 mt-2">
          {sorted.map(({ event, aggregate }) => (
            <EventRatingCard
              key={event.id}
              event={event}
              aggregate={aggregate}
              onPress={() =>
                router.push({
                  pathname: "/event/[id]/reviews",
                  params: {
                    id: event.id,
                    eventName: event.title,
                    eventImageUrl: event.presigned_url,
                  },
                })
              }
            />
          ))}
        </View>
      </ScrollView>
    </ThemedView>
  );
}
