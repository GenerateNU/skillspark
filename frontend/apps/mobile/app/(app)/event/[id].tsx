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
import { ListItem } from "@/components/ListItem";
import { AboutPage } from "@/components/AboutPage";
import { useRouteInfo } from "expo-router/build/hooks";

function formatAddress(occurrence: EventOccurrence) {
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
          marginTop: insets.top + 8,
          marginBottom: insets.bottom - 20,
          shadowColor: "#000",
          shadowOpacity: 0.12,
          shadowRadius: 20,
        }}
      >
        <ScrollView
          showsVerticalScrollIndicator={false}
          bounces={false}
          contentContainerStyle={{ flexGrow: 1 }}
        >
          {/* Hero image */}
          <View className="h-[250px] bg-[#1a1a1a]">
            {occurrence.event.presigned_url ? (
              <Image
                source={{ uri: occurrence.event.presigned_url }}
                style={{ width: "100%", height: "100%" }}
                contentFit="cover"
              />
            ) : (
              <View className="flex-1 bg-[#C5C5C5]" />
            )}
            <TouchableOpacity
              onPress={() => router.back()}
              activeOpacity={0.7}
              className="absolute top-4 left-4 z-10 flex-row items-center bg-white rounded-full px-4 py-2.5 elevation-10"
            >
              <MaterialIcons
                name="chevron-left"
                size={20}
                color={AppColors.primaryText}
              />
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
            </TouchableOpacity>
          </View>

          {/* Content card */}
          <View className="bg-white rounded-t-[28px] -mt-7 px-[22px] pb-6 elevation-2">
            {/* Drag handle */}
            <View
              className="w-[38px] h-1 rounded-sm self-center mt-3 mb-3.5"
              style={{ backgroundColor: AppColors.borderLight }}
            />
            <View className="absolute -top-[5px] left-2.5">
              <BookmarkButton eventId={occurrence.event.id} />
            </View>
            <Text
              className="text-[28px] font-bold tracking-tight mb-2"
              style={{ color: AppColors.primaryText }}
            >
              {occurrence.event.title}
            </Text>
            <View className="flex-row items-center gap-1.5 mb-3.5">
              <MaterialIcons
                name="location-on"
                size={16}
                color={AppColors.primaryText}
              />
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
              <StarRating size={17} />
              <ListItem
                label={translate("review.title")}
                isLast
                onPress={() =>
                  router.push({
                    pathname: `/event/[id]/reviews`,
                    params: {
                      id: occurrence.event.id,
                      occurrenceId: occurrence.id,
                      canReview: "true",
                      eventName: occurrence.event.title,
                      eventLocation: address,
                      eventImageUrl: occurrence.event.presigned_url ?? "",
                    },
                  })
                }
              />
            </View>
            <AboutPage
              description={occurrence.event.description}
              links={occurrence.org_links ?? []}
            />
            <View className="flex-row items-center justify-between mt-2">
              <View className="flex-row gap-2 flex-1 flex-wrap">
                {occurrence.event.category?.map((cat) => (
                  <View
                    key={cat}
                    className="border-[1.5px] rounded-full px-4 py-[7px]"
                    style={{ borderColor: AppColors.borderLight }}
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
          </View>

          {/* Divider */}
          <View
            className="border-b"
            style={{ borderColor: AppColors.borderLight }}
          />

          {/* Bottom section */}
          <View
            className="flex-1 bg-white px-[22px] pt-[22px] pb-7 elevation-2"
            style={{
              shadowColor: "#000",
              shadowOpacity: 0.06,
              shadowRadius: 12,
            }}
          >
            <View className="flex-row items-center justify-between mb-[22px]">
              <Text
                className="text-[30px] font-bold tracking-tight"
                style={{ color: AppColors.primaryText }}
              >
                {duration}
              </Text>
              <View className="flex-row items-center gap-[5px] bg-[#F3F4F6] rounded-full px-3 py-[7px]">
                <MaterialIcons
                  name="directions-walk"
                  size={14}
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
