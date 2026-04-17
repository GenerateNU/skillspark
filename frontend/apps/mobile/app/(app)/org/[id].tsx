import { useState } from "react";
import { Image } from "expo-image";
import {
  ActivityIndicator,
  ScrollView,
  Text,
  TouchableOpacity,
  View,
} from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";
import { useLocalSearchParams, useRouter } from "expo-router";
import {
  useGetEventReviewsForOrganization,
  useGetLocationById,
  useGetOrganization,
} from "@skillspark/api-client";
import type {
  Location,
  Organization,
  SimpleReviewAggregate,
} from "@skillspark/api-client";
import { ErrorScreen } from "@/components/ErrorScreen";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { ThemedText } from "@/components/themed-text";
import { AppColors, Shadows } from "@/constants/theme";
import { useThemeColor } from "@/hooks/use-theme-color";
import { AboutPage } from "@/components/AboutPage";
import { EventRatingCard } from "@/components/EventRatingCard";
import { useTranslation } from "react-i18next";
import * as Linking from "expo-linking";
import { ShareModal } from "@/components/ShareModal";
import { RatingSmiley } from "@/components/RatingSmiley";
import LogoBgWrapper from "@/components/LogoBgWrapper";

function OrgDetail({
  org,
  location,
}: {
  org: Organization;
  location?: Location;
}) {
  const router = useRouter();
  const { t: translate } = useTranslation();
  const backgroundColor = useThemeColor({}, "background");
  const borderColor = useThemeColor({}, "borderColor");
  const [shareVisible, setShareVisible] = useState(false);

  const totalOrgReviews = org.review_summary?.total_reviews ?? 0;
  const { data: previewResp } = useGetEventReviewsForOrganization(
    org.id,
    { page: 1, page_size: 1, sort_by: "most_rated" },
    { query: { enabled: !!org.id && totalOrgReviews > 0 } },
  );
  const previewEvent =
    previewResp?.status === 200 &&
    Array.isArray(previewResp.data) &&
    previewResp.data.length > 0
      ? (previewResp.data[0] as SimpleReviewAggregate)
      : null;

  const cardStyle = {
    shadowColor: "#000",
    shadowOpacity: 0.08,
    shadowRadius: 12,
    shadowOffset: { width: 0, height: 2 },
    elevation: 3,
  };

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
        <ThemedText
          className="flex-1 text-center text-[16px] font-nunito-bold"
          numberOfLines={1}
        >
          {org.name}
        </ThemedText>
        <View className="w-8" />
      </View>

      <ScrollView showsVerticalScrollIndicator={false}>
        {/* Hero image */}
        <View
          className="h-[200px]"
          style={{ backgroundColor: AppColors.imagePlaceholder }}
        >
          {org.presigned_url ? (
            <Image
              source={{ uri: org.presigned_url }}
              style={{ width: "100%", height: "100%" }}
              contentFit="cover"
            />
          ) : (
            <View className="flex-1 items-center justify-center">
              <IconSymbol name="photo" size={48} color={AppColors.mutedText} />
            </View>
          )}
        </View>

        <LogoBgWrapper>
          {/* Org info */}
          <View className="px-4 pb-4 pt-4">
            <View className="flex-row items-start justify-between">
              <View className="mr-3 flex-1">
                <Text className="mb-1 text-[24px] font-nunito-bold">
                  {org.name}
                </Text>
                {location && (
                  <Text
                    className="mb-1 text-[14px] font-nunito"
                    style={{ color: AppColors.mutedText }}
                  >
                    {location.district}, {location.province}
                  </Text>
                )}
              </View>
              <TouchableOpacity
                onPress={() => setShareVisible(true)}
                activeOpacity={0.7}
                className="mt-1 h-9 w-9 items-center justify-center rounded-full border-2"
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

          {/* Divider */}
          <View
            className="mx-4 mb-5 border-b border-dashed"
            style={{ borderColor: AppColors.divider }}
          />

          {/* About card */}
          <View
            className="mx-4 mb-4 rounded-2xl bg-white p-5"
            style={cardStyle}
          >
            <AboutPage
              description={org.about ?? ""}
              links={org.links ?? []}
            />
          </View>

          {/* Reviews card */}
          <TouchableOpacity
            activeOpacity={0.8}
            onPress={() => router.push(`/org/${org.id}/reviews`)}
          >
            <View
              className="mx-4 mb-4 rounded-2xl bg-white p-5"
              style={cardStyle}
            >
              <Text
                className="mb-4 font-nunito-bold text-[18px]"
                style={{ color: AppColors.primaryText }}
              >
                {translate("org.reviews")}
              </Text>

              {totalOrgReviews > 0 ? (
                <View className="flex-row gap-3 items-start mb-2">
                  {/* Left: aggregate number + smiley + count */}
                  <View className="items-center">
                    <Text
                      className="font-nunito-bold"
                      style={{ color: AppColors.primaryText, fontSize: 40, lineHeight: 48 }}
                    >
                      {org.review_summary!.average_rating % 1 === 0
                        ? org.review_summary!.average_rating.toFixed(0)
                        : org.review_summary!.average_rating.toFixed(1)}
                    </Text>
                    <RatingSmiley
                      rating={org.review_summary!.average_rating}
                      width={44}
                      height={44}
                    />
                    <Text
                      className="text-[12px] font-nunito mt-1"
                      style={{ color: AppColors.subtleText }}
                    >
                      ({totalOrgReviews})
                    </Text>
                  </View>

                  {/* Right: preview event rating card */}
                  {previewEvent && (
                    <View className="flex-1">
                      <EventRatingCard
                        event={previewEvent.event}
                        aggregate={previewEvent}
                        onPress={() => router.push(`/org/${org.id}/reviews`)}
                      />
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
                style={{
                  height: 0.5,
                  backgroundColor: AppColors.borderLight,
                }}
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

          {/* See Schedule CTA */}
          <View className="px-4 pb-6 pt-1">
            <TouchableOpacity
              activeOpacity={0.85}
              className="w-full items-center rounded-full py-4"
              style={{ backgroundColor: "#000000" }}
              onPress={() => router.push(`/org/${org.id}/schedule`)}
            >
              <Text className="text-[17px] font-nunito-bold text-white">
                {translate("org.seeSchedule")}
              </Text>
            </TouchableOpacity>
          </View>

          <ShareModal
            visible={shareVisible}
            onClose={() => setShareVisible(false)}
            name={org.name}
            imageUrl={org.presigned_url ?? undefined}
            shareUrl={Linking.createURL(`org/${org.id}`)}
            message={translate("share.defaultMessage", { name: org.name })}
          />
        </LogoBgWrapper>
      </ScrollView>
    </SafeAreaView>
  );
}

export default function OrgScreen() {
  const { id } = useLocalSearchParams<{ id: string }>();
  const { data: response, isLoading, error } = useGetOrganization(id);
  const { t: translate } = useTranslation();
  const { data: locationResponse } = useGetLocationById(
    response?.status === 200 ? (response.data.location_id ?? "") : "",
    {
      query: {
        enabled: response?.status === 200 && !!response.data.location_id,
      },
    },
  );

  if (isLoading) {
    return (
      <View className="flex-1 items-center justify-center">
        <ActivityIndicator size="large" />
      </View>
    );
  }

  if (error || !response || response.status !== 200) {
    return <ErrorScreen message={translate("org.notFound")} />;
  }

  return (
    <OrgDetail
      org={response.data}
      location={
        locationResponse?.status === 200 ? locationResponse.data : undefined
      }
    />
  );
}
