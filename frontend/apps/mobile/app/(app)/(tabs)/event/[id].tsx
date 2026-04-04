import { Image } from "expo-image";
import {
  ActivityIndicator,
  ScrollView,
  Text,
  TouchableOpacity,
  View,
} from "react-native";
import { useSafeAreaInsets } from "react-native-safe-area-context";
import { useLocalSearchParams, useRouter } from "expo-router";
import {
  useGetEventOccurrencesById,
  useGetReviewByEventId,
  useGetReviewAggregate,
} from "@skillspark/api-client";
import type { EventOccurrence } from "@skillspark/api-client";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { AppColors, FontFamilies, FontSizes } from "@/constants/theme";
import { StarRating } from "@/components/StarRating";
import { useTranslation } from "react-i18next";

function formatRelativeTime(dateStr: string): string {
  const diff = Date.now() - new Date(dateStr).getTime();
  const days = Math.floor(diff / 86400000);
  if (days < 1) return "today";
  if (days < 7) return `${days}d`;
  if (days < 30) return `${Math.floor(days / 7)}w`;
  return `${Math.floor(days / 30)}mo`;
}

function formatLocation(occurrence: EventOccurrence) {
  const loc = occurrence.location;
  const parts = [loc.address_line1, loc.district].filter(Boolean);
  return parts.join(", ") || "Location";
}

function EventOccurrenceDetail({
  occurrence,
}: {
  occurrence: EventOccurrence;
}) {
  const router = useRouter();
  const insets = useSafeAreaInsets();
  const { t: translate } = useTranslation();
  const location = formatLocation(occurrence);
  const categories = occurrence.event.category?.join(" / ") ?? "";

  const { data: aggregateResp } = useGetReviewAggregate(occurrence.event.id);
  const aggregate =
    aggregateResp?.status === 200 ? aggregateResp.data : null;

  const { data: reviewsResp } = useGetReviewByEventId(occurrence.event.id, {
    page: 1,
    page_size: 1,
  });
  const firstReview =
    reviewsResp?.status === 200 ? reviewsResp.data[0] ?? null : null;

  return (
    <View className="flex-1 bg-white">
      <View
        className="flex-row items-center bg-white px-4 border-b"
        style={{
          paddingTop: insets.top + 6,
          paddingBottom: 10,
          borderBottomColor: AppColors.divider,
        }}
      >
        <TouchableOpacity
          onPress={() => router.back()}
          activeOpacity={0.7}
          className="w-8 h-8 items-center justify-center"
        >
          <IconSymbol
            name="chevron.left"
            size={28}
            color={AppColors.primaryText}
          />
        </TouchableOpacity>
        <Text
          className="flex-1 text-center"
          style={{
            fontSize: 17,
            fontFamily: FontFamilies.bold,
            color: AppColors.primaryText,
          }}
          numberOfLines={1}
        >
          {occurrence.event.title}
        </Text>
        <View style={{ width: 32 }} />
      </View>
      <ScrollView
        showsVerticalScrollIndicator={false}
        contentContainerStyle={{ paddingBottom: 24 }}
      >
        <View
          style={{
            height: 220,
            backgroundColor: AppColors.imagePlaceholder,
          }}
        >
          {occurrence.event.presigned_url ? (
            <Image
              source={{ uri: occurrence.event.presigned_url }}
              className="w-full h-full"
              contentFit="cover"
            />
          ) : (
            <View className="flex-1 bg-[#C5C5C5]" />
          )}
        </View>
        <View className="px-4 pt-4 pb-2">
          <View className="flex-row items-start justify-between">
            <View className="flex-1 mr-3">
              <Text
                style={{
                  fontSize: 26,
                  fontFamily: FontFamilies.museoModerno,
                  color: AppColors.primaryText,
                  marginBottom: 2,
                }}
              >
                {occurrence.event.title}
              </Text>
              <Text
                style={{
                  fontSize: FontSizes.base,
                  color: AppColors.secondaryText,
                  fontFamily: FontFamilies.regular,
                  marginBottom: 5,
                }}
              >
                {location}
              </Text>
              {!!categories && (
                <View className="flex-row items-center gap-1.5 mb-1.5">
                  <Text style={{ fontSize: FontSizes.base }}>⚽</Text>
                  <Text
                    style={{
                      fontSize: FontSizes.base,
                      color: AppColors.secondaryText,
                      fontFamily: FontFamilies.regular,
                    }}
                  >
                    {categories}
                  </Text>
                </View>
              )}
              <View className="flex-row items-center gap-1.5">
                <Text style={{ fontSize: FontSizes.base }}>🔥</Text>
                <Text
                  style={{
                    fontSize: FontSizes.base,
                    color: AppColors.primaryBlue,
                    fontFamily: FontFamilies.semiBold,
                  }}
                >
                  {occurrence.curr_enrolled}+{" "}
                  {translate("event.bookingsThisWeek")}
                </Text>
              </View>
            </View>
            <View className="flex-col items-center gap-2 mt-1">
              <TouchableOpacity
                activeOpacity={0.7}
                className="w-9 h-9 rounded-full border-2 items-center justify-center"
                style={{ borderColor: AppColors.borderLight }}
              >
                <IconSymbol
                  name="bookmark"
                  size={18}
                  color={AppColors.secondaryText}
                />
              </TouchableOpacity>
              <TouchableOpacity
                activeOpacity={0.7}
                className="w-9 h-9 rounded-full border-2 items-center justify-center"
                style={{ borderColor: AppColors.borderLight }}
              >
                <IconSymbol
                  name="square.and.arrow.up"
                  size={18}
                  color={AppColors.secondaryText}
                />
              </TouchableOpacity>
            </View>
          </View>
        </View>
        <View
          className="mt-2 mb-3 p-5"
          style={{
            marginHorizontal: 15,
            backgroundColor: AppColors.white,
            borderRadius: 32,
            shadowColor: "#000",
            shadowOffset: { width: 0, height: 2 },
            shadowOpacity: 0.08,
            shadowRadius: 8,
            elevation: 3,
          }}
        >
          {(() => {
            const hasReviews = !!aggregate && aggregate.total_reviews > 0;
            if (!aggregate) return null;
            return (
              <>
                <View className="flex-row gap-3">
                  <View style={{ width: 90, alignItems: "center" }}>
                    <Text
                      style={{
                        fontSize: 22,
                        fontFamily: FontFamilies.bold,
                        color: AppColors.primaryText,
                        marginBottom: 2,
                        textAlign: "center",
                      }}
                    >
                      {translate("event.reviews")}
                    </Text>
                    <Text
                      className="text-[42px] font-nunito-bold text-[#111] leading-[46px] text-center"
                    >
                      {aggregate.average_rating.toFixed(1)}
                    </Text>
                    <StarRating
                      rating={Math.round(aggregate.average_rating)}
                      size={13}
                      filledColor={AppColors.primaryText}
                    />
                    <Text
                      style={{
                        fontSize: FontSizes.sm,
                        color: AppColors.mutedText,
                        fontFamily: FontFamilies.regular,
                        marginTop: 4,
                        textAlign: "center",
                      }}
                    >
                      ({aggregate.total_reviews})
                    </Text>
                  </View>
                  {hasReviews && firstReview && (
                    <View className="justify-start px-0.5 pt-1">
                      <Image
                        source={require("@/assets/images/faces.png")}
                        className="w-[38px] h-[38px] rounded-full"
                        contentFit="cover"
                      />
                    </View>
                  )}
                  {hasReviews && firstReview && (
                    <View className="flex-1">
                      <Text
                        style={{
                          fontSize: FontSizes.lg,
                          color: AppColors.primaryText,
                          fontFamily: FontFamilies.regular,
                          lineHeight: 24,
                          marginBottom: 12,
                        }}
                      >
                        {firstReview.description}
                      </Text>
                      {firstReview.categories?.length > 0 && (
                        <View className="flex-row gap-2 flex-wrap mb-3">
                          {firstReview.categories.map((cat) => (
                            <View
                              key={cat}
                              className="rounded-xl px-4 py-1.5"
                              style={{ backgroundColor: "#BFCFEA" }}
                            >
                              <Text
                                style={{
                                  fontSize: FontSizes.base,
                                  color: AppColors.secondaryText,
                                  fontFamily: FontFamilies.regular,
                                }}
                              >
                                {cat}
                              </Text>
                            </View>
                          ))}
                        </View>
                      )}
                      <View className="flex-row items-center justify-between">
                        <Text
                          style={{
                            fontSize: FontSizes.base,
                            color: AppColors.primaryText,
                            fontFamily: FontFamilies.regular,
                            fontStyle: "italic",
                          }}
                        >
                          Anonymous
                        </Text>
                        <Text
                          style={{
                            fontSize: FontSizes.base,
                            color: AppColors.primaryText,
                            fontFamily: FontFamilies.regular,
                          }}
                        >
                          {formatRelativeTime(firstReview.created_at)}
                        </Text>
                      </View>
                    </View>
                  )}
                </View>
                {hasReviews && (
                  <TouchableOpacity
                    activeOpacity={0.7}
                    className="items-center mt-5"
                    onPress={() => {}}
                  >
                    <Text
                      style={{
                        fontSize: FontSizes.base,
                        color: AppColors.primaryText,
                        fontFamily: FontFamilies.regular,
                      }}
                    >
                      {translate("event.seeMoreReviews")}
                    </Text>
                  </TouchableOpacity>
                )}
              </>
            );
          })()}
        </View>
        <View
          className="px-4"
          style={{ paddingTop: 4, paddingBottom: insets.bottom + 10 }}
        >
          <TouchableOpacity
            activeOpacity={0.85}
            className="w-full items-center rounded-full py-4"
            style={{ backgroundColor: AppColors.primaryText }}
            onPress={() => {}}
          >
            <Text
              style={{
                color: AppColors.white,
                fontSize: 17,
                fontFamily: FontFamilies.bold,
              }}
            >
              {translate("event.reserve")}
            </Text>
          </TouchableOpacity>
        </View>
      </ScrollView>
    </View>
  );
}

export default function EventOccurrenceScreen() {
  const { id } = useLocalSearchParams<{ id: string }>();
  const { data: response, isLoading, error } = useGetEventOccurrencesById(id);
  const { t: translate } = useTranslation();

  if (isLoading) {
    return (
      <View className="flex-1 items-center justify-center">
        <ActivityIndicator size="large" />
      </View>
    );
  }

  if (error || !response || response.status !== 200) {
    return (
      <View className="flex-1 items-center justify-center p-6">
        <Text
          className="text-base font-semibold"
          style={{ color: AppColors.danger }}
        >
          {translate("event.notFound")}
        </Text>
      </View>
    );
  }

  return <EventOccurrenceDetail occurrence={response.data} />;
}
