import { Image } from "expo-image";
import {
  ActivityIndicator,
  Pressable,
  ScrollView,
  Text,
  TouchableOpacity,
  View,
} from "react-native";
import * as Linking from "expo-linking";
import { SafeAreaView } from "react-native-safe-area-context";
import { useLocalSearchParams, useRouter } from "expo-router";
import {
  useGetEventOccurrencesByEventId,
  useGetOrganization,
  useGetReviewAggregate,
  useGetReviewByEventId,
} from "@skillspark/api-client";
import type { EventOccurrence, Organization, Review } from "@skillspark/api-client";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { RATING_OPTIONS } from "@/constants/ratings";
import { AppColors, Shadows } from "@/constants/theme";
import { useOrgLinks } from "@/hooks/useOrgLinks";
import { BookmarkButton } from "@/components/BookmarkButton";
import { ReviewCard } from "@/components/ReviewCard";
import { useTranslation } from "react-i18next";
import { AboutPage } from "@/components/AboutPage";
import { ShareModal } from "@/components/ShareModal";
import { formatLocation } from "@/utils/format";
import { getRatingOption } from "@/utils/ratings";
import { EventImage } from "@/components/EventImage";
import { ErrorScreen } from "@/components/ErrorScreen";
import LogoBgWrapper from "@/components/LogoBgWrapper";
import { useState } from "react";

function EventOccurrenceDetail({
  occurrence,
  org,
}: {
  occurrence: EventOccurrence;
  org: Organization | null;
}) {
  const router = useRouter();
  const { t: translate } = useTranslation();
  const [shareVisible, setShareVisible] = useState(false);
  const handleBack = () => router.back();
  const { openLink, hasLinks } = useOrgLinks(occurrence.org_links ?? []);

  const location = formatLocation(occurrence);
  const categories = (occurrence.event.category || [])
    .map((elem) =>
      //Capitalize the first char of every word
      elem
        .toLowerCase()
        .split(" ")
        .map((word) => word.charAt(0).toUpperCase() + word.slice(1))
        .join(" "),
    )
    .join(" / ");
  const orgId = occurrence.event.organization_id;
  const orgName = org?.name ?? "";

  const { data: aggregateResp } = useGetReviewAggregate(occurrence.event.id, {
    query: { enabled: !!occurrence.event.id },
  });
  const aggregate = aggregateResp?.status === 200 ? aggregateResp.data : null;
  const avgRating = aggregate?.average_rating ?? 0;
  const totalReviews = aggregate?.total_reviews ?? 0;
  const ratingOption = getRatingOption(avgRating);
  const ratingMatch = RATING_OPTIONS.find(
    (r) => r.rating === Math.round(avgRating)
  );

  const cardShadow = {
    shadowColor: "#000",
    shadowOpacity: 0.08,
    shadowRadius: 12,
    shadowOffset: { width: 0, height: 2 },
    elevation: 3,
  };

  const { data: reviewsResp } = useGetReviewByEventId(
    occurrence.event.id,
    { page: 1, page_size: 5, sort_by: "highest" },
    { query: { enabled: !!occurrence.event.id && totalReviews > 0 } },
  );
  const rawReviews =
    reviewsResp?.status === 200 ? (reviewsResp.data as Review[]) : [];
  const previewReview =
    rawReviews.length > 0
      ? rawReviews.reduce<Review>(
          (best, r) =>
            Math.abs(r.rating - avgRating) < Math.abs(best.rating - avgRating)
              ? r
              : best,
          rawReviews[0],
        )
      : null;

  return (
    <SafeAreaView className="flex-1 bg-white" edges={["top", "bottom"]}>
      {/* Header */}
      <View
        className="flex-row items-center border-b px-4 pb-2.5 pt-3"
        style={{ borderBottomColor: AppColors.divider }}
      >
        <TouchableOpacity
          onPress={handleBack}
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
          {translate("org.class")}
        </Text>
        <View className="w-8" />
      </View>

      <ScrollView
        showsVerticalScrollIndicator={false}
        contentContainerStyle={{ paddingBottom: 24 }}
      >
        {/* Hero image */}
        <View className="h-[250px]">
          <EventImage
            uri={occurrence.event.presigned_url}
            style={{ width: "100%", height: "100%" }}
          />
        </View>

        {/* White content card overlapping image */}
        <View className="bg-white -mt-7 px-[22px] pb-6">
          {/* Drag handle */}
          <View
            className="w-[38px] h-1 rounded-sm self-center mt-3 mb-3.5"
          />

          {/* Title row with bookmark + share */}
          <View className="mb-1">
            <View className="absolute top-0 right-0 flex-row items-center gap-2 z-10">
              <TouchableOpacity
                onPress={() => setShareVisible(true)}
                activeOpacity={0.7}
                className="h-9 w-9 items-center justify-center rounded-full border-2"
                style={{ borderColor: AppColors.borderLight }}
              >
                <IconSymbol
                  name="square.and.arrow.up"
                  size={18}
                  color={AppColors.secondaryText}
                />
              </TouchableOpacity>
              <BookmarkButton
                eventId={occurrence.event.id}
                iconSize={18}
                className="h-9 w-9 items-center justify-center rounded-full border-2"
                style={{ borderColor: AppColors.borderLight }}
              />
            </View>
            <Text
              className="mr-24 text-[26px] font-nunito-bold leading-8"
              style={{ color: AppColors.primaryText }}
            >
              {occurrence.event.title}
            </Text>
          </View>
<LogoBgWrapper>
          {/* Location */}
          <View className="flex-row items-center gap-1 mb-2">
            <IconSymbol name="location" size={22} color={AppColors.mutedText} />
            <Text
              className="text-[14px] font-nunito"
              style={{ color: AppColors.mutedText }}
            >
              {location}
            </Text>
          </View>

          {/* Category */}
          {!!categories && (
            <View className="flex-row gap-1">
              <IconSymbol
                name="star.fill"
                size={22}
                color={AppColors.mutedText}
              />
              <Text
                className="text-[14px] font-nunito mb-2"
                style={{ color: AppColors.mutedText }}
              >
                {categories}
              </Text>
            </View>
          )}

          {/* Bookings this week */}
          <View className="flex-row items-center gap-1 mb-2">
            <IconSymbol
              name="flame.fill"
              size={18}
              color={AppColors.primaryBlue}
            />
            <Text
              className="text-[14px] font-nunito"
              style={{ color: AppColors.primaryBlue }}
            >
              {occurrence.curr_enrolled}+ {translate("event.bookingsThisWeek")}
            </Text>
          </View>

          {/* Org badge */}
          {!!orgName && !!orgId && (
            <TouchableOpacity
              onPress={() => router.push(`../org/${orgId}`)}
              activeOpacity={0.7}
              className="self-start px-4 py-1.5 rounded-full mt-1"
              style={{ backgroundColor: AppColors.savedBackground }}
            >
              <Text
                className="text-[13px] font-nunito-bold"
                style={{ color: AppColors.primaryBlue }}
              >
                {orgName}
              </Text>
            </TouchableOpacity>
          )}
          </LogoBgWrapper>
        </View>

        
        {/* About card */}
        <View
          className="mx-4 mb-4 rounded-2xl bg-white p-5"
          style={Shadows.card}
        >
          <AboutPage
            description={occurrence.event.description}
            links={occurrence.org_links ?? []}
          />
        </View>

        {/* Reviews card */}
        <TouchableOpacity
          activeOpacity={0.8}
          onPress={() =>
            router.push({
              pathname: "/event/[id]/reviews",
              params: {
                id: occurrence.event.id,
                occurrenceId: occurrence.id,
                canReview: "true",
                eventName: occurrence.event.title,
                eventLocation: location,
                eventImageUrl: occurrence.event.presigned_url ?? "",
              },
            })
          }
        >
          <View
            className="mx-4 mb-4 rounded-2xl bg-white p-5"
            style={Shadows.card}
          >
            <Text
              className="mb-4 font-nunito-bold text-[18px]"
              style={{ color: AppColors.primaryText }}
            >
              {translate("event.reviews")}
            </Text>

            {totalReviews > 0 ? (
              <View className="flex-row gap-3 items-start mb-2">
                {/* Left: aggregate number + smiley + count */}
                <View className="items-center">
                  <Text
                    className="font-nunito-bold"
                    style={{ color: AppColors.primaryText, fontSize: 40, lineHeight: 48 }}
                  >
                    {avgRating % 1 === 0
                      ? avgRating.toFixed(0)
                      : avgRating.toFixed(1)}
                  </Text>
                  <Image
                    source={ratingOption.image}
                    style={{ width: 44, height: 44 }}
                  />
                  <Text
                    className="text-[12px] font-nunito mt-1"
                    style={{ color: AppColors.subtleText }}
                  >
                    ({totalReviews})
                  </Text>
                </View>

                {/* Right: preview review (no avatar — aggregate smiley already on left) */}
                {previewReview && (
                  <View className="flex-1 -mr-5">
                    <ReviewCard review={previewReview} hideAvatar />
                  </View>
                )}
              </View>
            ) : (
              <Text
                className="text-[13px] font-nunito mb-2"
                style={{ color: AppColors.subtleText }}
              >
                {translate("review.noReviews")}
              </Text>
            )}

            <View
              style={{ height: 0.5, backgroundColor: AppColors.borderLight }}
              className="mt-2 mb-3"
            />
            <View className="items-center">
              <Text
                className="text-[13px] font-nunito underline"
                style={{ color: AppColors.primaryText }}
              >
                {translate("event.seeMoreReviews")}
              </Text>
            </View>
          </View>
        </TouchableOpacity>

        {/* Reserve button */}
        <View className="px-4 pb-2 pt-1">
          <TouchableOpacity
            activeOpacity={0.85}
            className="w-full items-center rounded-full py-4 bg-black"
            onPress={() =>
              router.push({
                pathname: "/org/[id]/schedule",
                params: {
                  id: orgId,
                  filterClass: occurrence.event.title,
                },
              })
            }
          >
            <Text className="text-[17px] font-nunito-bold text-white">
              {translate("event.reserve")}
            </Text>
          </TouchableOpacity>
        </View>
        
        
      </ScrollView>

      <ShareModal
        visible={shareVisible}
        onClose={() => setShareVisible(false)}
        name={occurrence.event.title}
        imageUrl={occurrence.event.presigned_url ?? undefined}
        shareUrl={Linking.createURL(`event/${occurrence.event.id}`)}
        message={translate("share.defaultMessage", {
          name: occurrence.event.title,
        })}
      />
    </SafeAreaView>
  );
}

export default function EventOccurrenceScreen() {
  const { id } = useLocalSearchParams<{ id: string }>();
  const {
    data: response,
    isLoading,
    error,
  } = useGetEventOccurrencesByEventId(id);
  const { t: translate } = useTranslation();

  const occurrence = response?.status === 200 ? response.data[0] : null;
  const orgId = occurrence?.event.organization_id;

  const { data: orgResp, isLoading: orgLoading } = useGetOrganization(
    orgId ?? "",
    { query: { enabled: !!orgId } },
  );
  const org = orgResp?.status === 200 ? orgResp.data : null;

  if (isLoading || orgLoading) {
    return (
      <View className="flex-1 items-center justify-center">
        <ActivityIndicator size="large" />
      </View>
    );
  }

  if (
    error ||
    !response ||
    response.status !== 200 ||
    response.data.length === 0 ||
    !occurrence
  ) {
    return <ErrorScreen message={translate("event.notFound")} />;
  }

  return <EventOccurrenceDetail occurrence={occurrence} org={org} />;
}
