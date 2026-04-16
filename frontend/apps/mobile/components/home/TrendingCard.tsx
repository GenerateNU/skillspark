import { View, Text, Pressable, Image as RNImage } from "react-native";
import { useRouter } from "expo-router";
import { type EventOccurrence, type ReviewAggregate, useGetReviewAggregate } from "@skillspark/api-client";
import { AppColors, FontFamilies, FontSizes } from "@/constants/theme";
import { haversineDistance } from "@/utils/distance";
import { formatAgeRange } from "@/utils/format";
import { getRatingOption } from "@/utils/ratings";
import { EventImage } from "@/components/EventImage";

export function TrendingCard({
  occurrence,
  userLat,
  userLng,
  width = 300,
}: {
  occurrence: EventOccurrence;
  userLat?: number;
  userLng?: number;
  width?: number | "100%";
}) {
  const router = useRouter();

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
      ? `${reviewAggregate.average_rating.toFixed(1)} / 5 Smiles`
      : "No reviews yet";

  return (
    <View style={{ marginRight: width === "100%" ? 0 : 14, paddingBottom: 8 }}>
      <Pressable
        onPress={() => router.push(`/event/${occurrence.event.id}`)}
        style={{
          width,
          height: 140,
          flexDirection: "row",
          backgroundColor: AppColors.white,
          borderRadius: 16,
          borderWidth: 1,
          borderColor: AppColors.savedBackground,
          overflow: "hidden",
          shadowColor: "#000",
          shadowOpacity: 0.08,
          shadowRadius: 4,
          shadowOffset: { width: 0, height: 2 },
          elevation: 2,
        }}
      >
        {/* Image */}
        <EventImage
          uri={occurrence.event.presigned_url}
          style={{ width: 110, height: 140 }}
        />

        {/* Content */}
        <View
          style={{ flex: 1, padding: 12, justifyContent: "center", gap: 4 }}
        >
          <Text
            style={{
              fontFamily: FontFamilies.bold,
              fontSize: 16,
              color: AppColors.primaryText,
            }}
            numberOfLines={2}
          >
            {occurrence.event.title}
          </Text>

          {/* Smiley rating */}
          <View
            style={{ flexDirection: "row", alignItems: "center", gap: 6 }}
          >
            <RNImage source={ratingOption.image} style={{ width: 18, height: 18 }} />
            <Text
              style={{
                fontFamily: FontFamilies.regular,
                fontSize: FontSizes.sm,
                color: AppColors.mutedText,
              }}
            >
              {reviewLabel}
            </Text>
          </View>

          {ageLabel && (
            <Text
              style={{
                fontFamily: FontFamilies.regular,
                fontSize: FontSizes.sm,
                color: AppColors.mutedText,
              }}
            >
              {ageLabel}
            </Text>
          )}

          {distance != null && (
            <Text
              style={{
                fontFamily: FontFamilies.regular,
                fontSize: FontSizes.sm,
                color: AppColors.mutedText,
              }}
            >
              {distance.toFixed(1)} km away
            </Text>
          )}
        </View>
      </Pressable>
    </View>
  );
}
