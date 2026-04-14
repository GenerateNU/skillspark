import { EventRatingCard } from "@/components/EventRatingCard";
import { RatingAggregateCard } from "@/components/ReviewAggregate";
import { FilterTabs } from "@/components/SortingButtons";
import { ThemedText } from "@/components/themed-text";
import { ThemedView } from "@/components/themed-view";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { Colors } from "@/constants/theme";
import {
  useGetEventReviewsForOrganization,
  useGetReviewAggregateOrganization,
  type ReviewAggregate,
} from "@skillspark/api-client";
import { useLocalSearchParams, useRouter } from "expo-router";
import { useState } from "react";
import { useTranslation } from "react-i18next";
import {
  ActivityIndicator,
  ScrollView,
  TouchableOpacity,
  View,
} from "react-native";
import { useSafeAreaInsets } from "react-native-safe-area-context";

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

  const { data: eventReviewsResp, isLoading: eventReviewsLoading } =
    useGetEventReviewsForOrganization(
      id ?? "",
      { sort_by: sortBy },
      { query: { enabled: !!id } }
    );

  if (orgAggLoading || eventReviewsLoading) {
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

  const items = eventReviewsResp?.status === 200 ? eventReviewsResp.data : [];

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
          value={sortBy}
          options={[
            { label: translate("review.mostRated"), value: "most_rated" },
            { label: translate("review.highest"), value: "highest" },
            { label: translate("review.lowest"), value: "lowest" },
          ]}
          onChange={(value) => setSortBy(value as SortValue)}
        />

        <View className="px-5 gap-3 mt-2">
          {items.map((item) => (
            <EventRatingCard
              key={item.event_id}
              event={item.event}
              aggregate={item}
              onPress={() =>
                router.push({
                  pathname: "/event/[id]/reviews",
                  params: {
                    id: item.event.id,
                    eventName: item.event.title,
                    eventImageUrl: item.event.presigned_url,
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