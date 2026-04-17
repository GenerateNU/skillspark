import { View, Text, Pressable, type StyleProp, type ViewStyle } from "react-native";
import { Image } from "expo-image";
import { useRouter } from "expo-router";
import { useTranslation } from "react-i18next";
import { type EventOccurrence, type ReviewAggregate, useGetReviewAggregate } from "@skillspark/api-client";
import { AppColors } from "@/constants/theme";
import { haversineDistance } from "@/utils/distance";
import { formatAgeRange } from "@/utils/format";
import { getRatingOption } from "@/utils/ratings";
import { EventImage } from "@/components/EventImage";

export function TrendingCard({
  occurrence,
  userLat,
  userLng,
  width = 300,
  style,
}: {
  occurrence: EventOccurrence;
  userLat?: number;
  userLng?: number;
  width?: number | "100%";
  style?: StyleProp<ViewStyle>;
}) {
  const router = useRouter();
  const { t } = useTranslation();

  const ageLabel =
    formatAgeRange(occurrence.event.age_range_min, occurrence.event.age_range_max) || null;

  const distance =
    userLat != null &&
    userLng != null &&
    occurrence.location?.latitude != null &&
    occurrence.location?.longitude != null
      ? haversineDistance(
          userLat,
          userLng,
          occurrence.location.latitude,
          occurrence.location.longitude,
        )
      : null;

  const { data: reviewResp } = useGetReviewAggregate(occurrence.event.id);
  const reviewAggregate =
    reviewResp?.status === 200 ? (reviewResp.data as ReviewAggregate) : null;

  const ratingOption = getRatingOption(reviewAggregate?.average_rating);

  const reviewLabel =
    reviewAggregate?.average_rating != null
      ? `${reviewAggregate.average_rating.toFixed(1)} ${t("occurrence.smiles")}`
      : t("event.noReviews");

  return (
    <Pressable
      onPress={() => router.push(`/event/${occurrence.event.id}`)}
      className="mb-2 h-[140px] flex-row rounded-2xl border overflow-hidden"
      style={[{
        width,
        backgroundColor: AppColors.white,
        borderColor: AppColors.savedBackground,
        shadowColor: "#000",
        shadowOpacity: 0.08,
        shadowRadius: 4,
        shadowOffset: { width: 0, height: 2 },
        elevation: 2,
      }, style]}
    >
      {/* Image */}
      <EventImage
        uri={occurrence.event.presigned_url}
        style={{ width: 110, height: 140 }}
      />

      {/* Content */}
      <View className="flex-1 p-3 justify-center gap-1">
        <Text
          className="font-nunito-bold text-base"
          style={{ color: AppColors.primaryText }}
          numberOfLines={2}
        >
          {occurrence.event.title}
        </Text>

        {/* Smiley rating */}
        <View className="flex-row items-center gap-1.5">
          <Image source={ratingOption.image} style={{ width: 18, height: 18 }} />
          <Text
            className="font-nunito text-[12px]"
            style={{ color: AppColors.mutedText }}
          >
            {reviewLabel}
          </Text>
        </View>

        {ageLabel && (
          <Text
            className="font-nunito text-[12px]"
            style={{ color: AppColors.mutedText }}
          >
            {ageLabel}
          </Text>
        )}

        {distance != null && (
          <Text
            className="font-nunito text-[12px]"
            style={{ color: AppColors.mutedText }}
          >
            {distance.toFixed(1)} {t("map.km")} {t("map.away")}
          </Text>
        )}
      </View>
    </Pressable>
  );
}
