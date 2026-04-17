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
} from "@skillspark/api-client";
import type { EventOccurrence, Organization } from "@skillspark/api-client";
import MaterialIcons from "@expo/vector-icons/MaterialIcons";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { RATING_OPTIONS } from "@/constants/ratings";
import { AppColors, Shadows } from "@/constants/theme";
import { useOrgLinks } from "@/hooks/useOrgLinks";
import { BookmarkButton } from "@/components/BookmarkButton";
import { useTranslation } from "react-i18next";
import { AboutPage } from "@/components/AboutPage";
import { ShareModal } from "@/components/ShareModal";
import { formatLocation } from "@/utils/format";
import { getRatingOption } from "@/utils/ratings";
import { EventImage } from "@/components/EventImage";
import { ExpandableText } from "@/components/ExpandableText";
import { ErrorScreen } from "@/components/ErrorScreen";
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
        <View className="bg-white rounded-t-[28px] -mt-7 px-[22px] pb-6">
          {/* Drag handle */}
          <View
            className="w-[38px] h-1 rounded-sm self-center mt-3 mb-3.5"
            style={{ backgroundColor: AppColors.borderLight }}
          />

          {/* Title row with bookmark + share */}
          <View className="mb-1">
            <View className="absolute top-0 right-0 items-center z-10">
              <TouchableOpacity
                onPress={() => setShareVisible(true)}
                activeOpacity={0.7}
              >
                <IconSymbol
                  name="square.and.arrow.up"
                  size={28}
                  color={AppColors.primaryText}
                />
              </TouchableOpacity>
              <BookmarkButton eventId={occurrence.event.id} />
            </View>
            <Text
              className="mr-10 text-[26px] font-nunito-bold leading-8"
              style={{ color: AppColors.primaryText }}
            >
              {occurrence.event.title}
            </Text>
          </View>

          {/* Location */}
          <View className="flex-row items-center gap-1 mb-2">
            <MaterialIcons
              name="location-on"
              size={22}
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

          {!!orgId && (
            <TouchableOpacity
              onPress={() => router.push(`../org/${orgId}`)}
              className="flex-row items-center gap-1 mb-2"
            >
              <IconSymbol
                name="person.fill"
                size={22}
                color={AppColors.mutedText}
              />
              <Text
                className="text-[14px] font-nunito underline"
                style={{ color: AppColors.mutedText }}
              >
                {orgName}
              </Text>
            </TouchableOpacity>
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
        <View
          className="mx-4 mb-4 rounded-2xl bg-white p-5"
          style={Shadows.card}
        >
          <AboutPage
            description={occurrence.event.description}
            links={occurrence.org_links ?? []}
          />
          <ExpandableText text={occurrence.event.description ?? ""} />
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
            style={Shadows.card}
          >
            <Text
              className="mb-3 font-nunito-bold text-[18px]"
              style={{ color: AppColors.primaryText }}
            >
              {translate("event.reviews")}
            </Text>

            {/* Aggregate rating */}
            {totalReviews > 0 ? (
              <View className="flex-row items-center gap-2">
                <Image
                  source={ratingOption.image}
                  style={{ width: 22, height: 22 }}
                />
                <Text
                  className="text-[15px] font-nunito-bold"
                  style={{ color: AppColors.primaryText }}
                >
                  {translate(ratingOption.labelKey!)}
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
