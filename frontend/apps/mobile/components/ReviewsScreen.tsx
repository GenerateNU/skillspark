import { ErrorScreen } from "@/components/ErrorScreen";
import { LeaveReviewModal } from "@/components/LeaveReviewModal";
import { RatingSmileys } from "@/components/RatingSmileys";
import { RatingAggregateCard } from "@/components/ReviewAggregate";
import { ReviewCard } from "@/components/ReviewCard";
import { FilterTabs } from "@/components/SortingButtons";
import { ThemedText } from "@/components/themed-text";
import { ThemedView } from "@/components/themed-view";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { AppColors, Colors } from "@/constants/theme";
import { useAuthContext } from "@/hooks/use-auth-context";
import {
  Review,
  ReviewAggregate,
  useGetRegistrationsByGuardianId,
} from "@skillspark/api-client";
import { useLocalSearchParams, useRouter } from "expo-router";
import React, { useMemo, useState } from "react";
import { useTranslation } from "react-i18next";
import {
  ActivityIndicator,
  ScrollView,
  Text,
  TouchableOpacity,
  useColorScheme,
  View,
} from "react-native";
import { useSafeAreaInsets } from "react-native-safe-area-context";

interface QueryResult<T> {
  data: T | undefined;
  isLoading: boolean;
  error: unknown;
}

interface ReviewsScreenProps<
  TAggregate extends { status: number; data: unknown },
  TReviews extends { status: number; data: unknown },
> {
  useGetAggregate: (id: string) => QueryResult<TAggregate>;
  useGetReviews: (id: string) => QueryResult<TReviews>;
  /** When true, shows the rating smileys and write-a-review CTA */
  canReview?: boolean;
  /** Event occurrence ID, used to look up the guardian's registration */
  occurrenceId?: string;
  eventName?: string;
  eventLocation?: string;
  eventImageUrl?: string;
}

export default function ReviewsScreen<
  TAggregate extends { status: number; data: unknown },
  TReviews extends { status: number; data: unknown },
>({
  useGetAggregate,
  useGetReviews,
  canReview = false,
  occurrenceId,
  eventName = "",
  eventLocation = "",
  eventImageUrl,
}: ReviewsScreenProps<TAggregate, TReviews>) {
  const { id } = useLocalSearchParams<{ id: string }>();

  const colorScheme = useColorScheme();
  const scheme = (colorScheme ?? "light") as "light" | "dark";
  const theme = Colors[scheme];

  const router = useRouter();
  const insets = useSafeAreaInsets();
  const { t: translate } = useTranslation();
  const { guardianId } = useAuthContext();

  const [modalVisible, setModalVisible] = useState(false);
  const [initialRating, setInitialRating] = useState<number | undefined>(
    undefined,
  );

  const {
    data: aggregateResponse,
    isLoading: aggregateIsLoading,
    error: aggregateError,
  } = useGetAggregate(id ?? "");

  const {
    data: reviewsResponse,
    isLoading: reviewsIsLoading,
    error: reviewsError,
  } = useGetReviews(id ?? "");

  const { data: registrationsResp } = useGetRegistrationsByGuardianId(
    guardianId ?? "",
    { query: { enabled: canReview && !!guardianId && !!occurrenceId } },
  );
  const registration = useMemo(() => {
    const list =
      registrationsResp?.status === 200
        ? registrationsResp.data.registrations
        : [];
    return list.find(
      (r) => r.event_occurrence_id === occurrenceId && r.status === "registered",
    );
  }, [registrationsResp, occurrenceId]);

  const showReviewCTA = canReview && !!registration;

  if (!id) {
    return <ErrorScreen message={translate("common.noEventId")} />;
  }

  if (aggregateIsLoading || reviewsIsLoading) {
    return (
      <View className="flex-1 items-center justify-center gap-2">
        <ActivityIndicator size="large" />
        <ThemedText>{translate("common.loadingEvents")}</ThemedText>
      </View>
    );
  }

  if (aggregateError || reviewsError) {
    return (
      <View className="flex-1 items-center justify-center p-4">
        <ThemedText className="text-red-500 font-semibold">
          {translate("common.errorLoadingEvents")}
        </ThemedText>
        <ThemedText>{translate("common.errorOccurred")}</ThemedText>
      </View>
    );
  }

  if (!aggregateResponse || aggregateResponse.status !== 200) {
    return (
      <View className="flex-1 items-center justify-center p-4">
        <ThemedText>{translate("common.noReviewsAvailable")}</ThemedText>
      </View>
    );
  }

  if (!reviewsResponse || reviewsResponse.status !== 200) {
    return (
      <View className="flex-1 items-center justify-center p-4">
        <ThemedText>{translate("common.noReviewsAvailable")}</ThemedText>
      </View>
    );
  }

  const aggregate = aggregateResponse.data as ReviewAggregate;
  const reviews = reviewsResponse.data as Review[];

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
        <RatingAggregateCard aggregate={aggregate} />

        {showReviewCTA && (
          <View
            className="mx-5 mt-4 p-4 rounded-2xl border"
            style={{ borderColor: AppColors.borderLight }}
          >
            <ThemedText className="text-lg mb-4 text-center">
              {aggregate.total_reviews > 0
                ? translate("review.tapToReview")
                : translate("review.firstReview")}
            </ThemedText>
            <RatingSmileys
              onSelect={(r) => {
                setInitialRating(r);
                setModalVisible(true);
              }}
            />
            <TouchableOpacity
              className="mt-4 py-4 rounded-2xl items-center"
              style={{ backgroundColor: AppColors.primaryText }}
              onPress={() => {
                setInitialRating(undefined);
                setModalVisible(true);
              }}
            >
              <Text className="text-white text-base">
                {translate("review.writeReview")}
              </Text>
            </TouchableOpacity>
          </View>
        )}

        {showReviewCTA && guardianId && (
          <LeaveReviewModal
            visible={modalVisible}
            onClose={() => setModalVisible(false)}
            eventId={id}
            eventName={eventName}
            eventLocation={eventLocation}
            eventImageUrl={eventImageUrl}
            registrationId={registration!.id}
            guardianId={guardianId}
            initialRating={initialRating}
          />
        )}

        {aggregate.total_reviews > 0 && (
          <FilterTabs
            options={[
              { label: translate("review.mostRecent"), value: "most_recent" },
              { label: translate("review.highest"), value: "highest" },
              { label: translate("review.lowest"), value: "lowest" },
            ]}
            onChange={(value) => {}}
          />
        )}

        <View className="gap-0 px-5 mt-4">
          {reviews.map((review, index) => (
            <View key={review.id}>
              <ReviewCard review={review} />
              {index < reviews.length - 1 && (
                <View
                  style={{ backgroundColor: AppColors.borderLight }}
                  className="my-3 h-px"
                />
              )}
            </View>
          ))}
        </View>
      </ScrollView>
    </ThemedView>
  );
}