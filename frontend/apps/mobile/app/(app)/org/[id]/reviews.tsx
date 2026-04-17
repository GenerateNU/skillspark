import { EventRatingCard } from "@/components/EventRatingCard";
import { RatingAggregateCard } from "@/components/ReviewAggregate";
import { FilterTabs } from "@/components/SortingButtons";
import { ThemedText } from "@/components/themed-text";
import { ThemedView } from "@/components/themed-view";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { Colors } from "@/constants/theme";
import { useInfiniteEventReviewsForOrganization } from "@/hooks/use-infinite-reviews";
import {
  useGetReviewAggregateOrganization,
  type ReviewAggregate,
  type GetEventReviewsForOrganizationSortBy,
} from "@skillspark/api-client";
import { useLocalSearchParams, useRouter } from "expo-router";
import { useMemo, useState } from "react";
import { useTranslation } from "react-i18next";
import {
  ActivityIndicator,
  FlatList,
  TouchableOpacity,
  View,
} from "react-native";
import { useSafeAreaInsets } from "react-native-safe-area-context";

type SortValue = GetEventReviewsForOrganizationSortBy;

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

  const {
    data: eventReviewsData,
    isLoading: eventReviewsLoading,
    fetchNextPage,
    hasNextPage,
    isFetchingNextPage,
  } = useInfiniteEventReviewsForOrganization(id, sortBy);

  const items = useMemo(
    () =>
      eventReviewsData?.pages.flatMap((page) =>
        Array.isArray(page.data) ? page.data : [],
      ) ?? [],
    [eventReviewsData],
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

  const ListHeader = (
    <>
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
    </>
  );

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

      <FlatList
        data={items}
        keyExtractor={(item) => item.event_id}
        showsVerticalScrollIndicator={false}
        contentContainerStyle={{ paddingBottom: 32, paddingHorizontal: 20 }}
        ListHeaderComponent={ListHeader}
        renderItem={({ item }) => (
          <View className="gap-3 mt-2">
            <EventRatingCard
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
          </View>
        )}
        onEndReached={() => {
          if (hasNextPage && !isFetchingNextPage) {
            fetchNextPage();
          }
        }}
        onEndReachedThreshold={0.3}
        ListFooterComponent={
          isFetchingNextPage ? (
            <View className="py-4 items-center">
              <ActivityIndicator size="small" />
            </View>
          ) : null
        }
      />
    </ThemedView>
  );
}
