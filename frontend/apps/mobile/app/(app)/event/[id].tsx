import { Image } from "expo-image";
import { useState } from "react";
import {
  ActivityIndicator,
  Pressable,
  ScrollView,
  Text,
  TouchableOpacity,
  View,
} from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";
import { useLocalSearchParams, useRouter } from "expo-router";
import {
  useGetEventOccurrencesByEventId,
  useGetReviewAggregate,
} from "@skillspark/api-client";
import type { EventOccurrence } from "@skillspark/api-client";
import { ErrorScreen } from "@/components/ErrorScreen";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { AppColors } from "@/constants/theme";
import { RATING_OPTIONS } from "@/constants/ratings";
import { useOrgLinks } from "@/hooks/useOrgLinks";
import { BookmarkButton } from "@/components/BookmarkButton";
import { useTranslation } from "react-i18next";
import MaterialIcons from "@expo/vector-icons/MaterialIcons";
import { formatLocation } from "@/utils/format";

function EventOccurrenceDetail({
  occurrence,
}: {
  occurrence: EventOccurrence;
}) {
  const router = useRouter();
  const { t: translate } = useTranslation();
  const { openLink, hasLinks } = useOrgLinks(occurrence.org_links ?? []);
  const [aboutExpanded, setAboutExpanded] = useState(false);
  const [aboutTruncated, setAboutTruncated] = useState(false);

  const location = formatLocation(occurrence);
  const categories = occurrence.event.category?.join(" / ") ?? "";
  const orgId = occurrence.event.organization_id;

  const { data: aggregateResp } = useGetReviewAggregate(occurrence.event.id, {
    query: { enabled: !!occurrence.event.id },
  });
  const aggregate = aggregateResp?.status === 200 ? aggregateResp.data : null;
  const avgRating = aggregate?.average_rating ?? 0;
  const totalReviews = aggregate?.total_reviews ?? 0;
  const ratingMatch = RATING_OPTIONS.find(
    (r) => r.rating === Math.round(avgRating),
  );

  const cardShadow = {
    shadowColor: "#000",
    shadowOpacity: 0.08,
    shadowRadius: 12,
    shadowOffset: { width: 0, height: 2 },
    elevation: 3,
  };

  return (
    <SafeAreaView className="flex-1 bg-white" edges={["top", "bottom"]}>
      {/* Header */}
      <View
        className="flex-row items-center border-b px-4 pb-2.5 pt-3"
        style={{ borderBottomColor: AppColors.divider }}
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
          {translate("org.class")}
        </Text>
        <View className="w-8" />
      </View>

      <ScrollView
        showsVerticalScrollIndicator={false}
        contentContainerStyle={{ paddingBottom: 24 }}
      >
        {/* Hero image */}
        <View
          className="h-[250px]"
          style={{ backgroundColor: AppColors.imagePlaceholder }}
        >
          {occurrence.event.presigned_url ? (
            <Image
              source={{ uri: occurrence.event.presigned_url }}
              style={{ width: "100%", height: "100%" }}
              contentFit="cover"
            />
          ) : null}
        </View>

        {/* White content card overlapping image */}
        <View className="bg-white rounded-t-[28px] -mt-7 px-[22px] pb-6">
          {/* Drag handle */}
          <View
            className="w-[38px] h-1 rounded-sm self-center mt-3 mb-3.5"
            style={{ backgroundColor: AppColors.borderLight }}
          />

          {/* Title row with bookmark */}
          <View className="flex-row items-start justify-between mb-3">
            <Text
              className="flex-1 mr-3 text-[26px] font-nunito-bold leading-8"
              style={{ color: AppColors.primaryText }}
            >
              {occurrence.event.title}
            </Text>
            <BookmarkButton eventId={occurrence.event.id} />
          </View>

          {/* Location */}
          <View className="flex-row items-center gap-1.5 mb-2">
            <MaterialIcons
              name="location-on"
              size={16}
              color={AppColors.mutedText}
            />
            <Text
              className="text-[14px] font-nunito"
              style={{ color: AppColors.mutedText }}
            >
              {location}
            </Text>
          </View>

          {/* Category */}
          {!!categories && (
            <Text
              className="text-[14px] font-nunito mb-2"
              style={{ color: AppColors.mutedText }}
            >
              {categories}
            </Text>
          )}

          {/* Bookings this week */}
          <Text
            className="text-[14px] font-nunito"
            style={{ color: AppColors.primaryBlue }}
          >
            {occurrence.curr_enrolled}+ {translate("event.bookingsThisWeek")}
          </Text>
        </View>

        {/* About card */}
        <View className="mx-4 mb-4 rounded-2xl bg-white p-5" style={cardShadow}>
          <Text
            className="mb-2.5 text-[18px] font-nunito-bold"
            style={{ color: AppColors.primaryText }}
          >
            {translate("event.about")}
          </Text>
          <Text
            numberOfLines={aboutExpanded ? undefined : 4}
            onTextLayout={(e) => {
              if (!aboutExpanded)
                setAboutTruncated(e.nativeEvent.lines.length >= 4);
            }}
            className={`text-sm leading-[22px] font-nunito ${aboutTruncated ? "mb-1" : "mb-4"}`}
            style={{ color: AppColors.secondaryText }}
          >
            {occurrence.event.description}
          </Text>
          {aboutTruncated && (
            <Pressable
              onPress={() => setAboutExpanded((prev) => !prev)}
              className="mb-4"
            >
              <Text
                className="text-[13px] font-semibold"
                style={{ color: AppColors.primaryText }}
              >
                {aboutExpanded
                  ? translate("event.seeLess")
                  : translate("event.seeMore")}
              </Text>
            </Pressable>
          )}
          {hasLinks && (
            <View className="flex-row flex-wrap gap-2.5">
              {occurrence.org_links.map((link, index) => (
                <Pressable
                  key={index}
                  onPress={() => openLink(link.href)}
                  className="rounded-full px-5 py-2.5 items-center"
                  style={{ backgroundColor: AppColors.borderLight }}
                >
                  <Text
                    className="text-[13px] font-semibold"
                    style={{ color: AppColors.primaryText }}
                  >
                    {link.label}
                  </Text>
                </Pressable>
              ))}
            </View>
          )}
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
            style={cardShadow}
          >
            {/* Title */}
            <Text
              className="mb-3 font-nunito-bold text-[18px]"
              style={{ color: AppColors.primaryText }}
            >
              {translate("event.reviews")}
            </Text>

            {/* Aggregate rating */}
            {totalReviews > 0 && ratingMatch ? (
              <View className="flex-row items-center gap-2">
                <Image
                  source={ratingMatch.image}
                  style={{ width: 22, height: 22 }}
                />
                <Text
                  className="text-[15px] font-nunito-bold"
                  style={{ color: AppColors.primaryText }}
                >
                  {translate(ratingMatch.labelKey!)}
                </Text>
                <Text
                  className="text-[13px] font-nunito"
                  style={{ color: AppColors.subtleText }}
                >
                  ({totalReviews})
                </Text>
              </View>
            ) : (
              <Text
                className="text-[13px] font-nunito"
                style={{ color: AppColors.subtleText }}
              >
                {translate("review.noReviews")}
              </Text>
            )}

            {/* See more */}
            <View className="mt-4 items-center">
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
            className="w-full items-center rounded-full py-4"
            style={{ backgroundColor: AppColors.checkboxSelected }}
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

  if (isLoading) {
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
    response.data.length === 0
  ) {
    return <ErrorScreen message={translate("event.notFound")} />;
  }

  return <EventOccurrenceDetail occurrence={response.data[0]} />;
}
