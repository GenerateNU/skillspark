import { Image } from "expo-image";
import { View, Text, Pressable } from "react-native";
import { useRouter } from "expo-router";
import { type EventOccurrence, useGetReviewAggregate } from "@skillspark/api-client";
import { AppColors, FontFamilies, FontSizes } from "@/constants/theme";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { RATING_OPTIONS } from "@/constants/ratings";

const AVATAR_COLORS = [AppColors.emerald, "#5B21B6", AppColors.amber];

export function DiscoverBanner({ event }: { event: EventOccurrence }) {
  const router = useRouter();
  const eventId = event.event.id;

  const { data: aggregateResp } = useGetReviewAggregate(eventId, {
    query: { enabled: !!eventId },
  });
  const aggregate = (aggregateResp as { status: number; data: { average_rating: number; total_reviews: number } } | undefined);
  const aggregateData = aggregate?.status === 200 ? aggregate.data : null;
  const avgRating = aggregateData?.average_rating ?? 0;
  const totalReviews = aggregateData?.total_reviews ?? 0;
  const ratingMatch = RATING_OPTIONS.find((r) => r.rating === Math.round(avgRating));

  return (
    <Pressable
      onPress={() => router.push(`/event/${eventId}`)}
      style={{
        flex: 1,
        borderRadius: 24,
        overflow: "hidden",
        backgroundColor: AppColors.primaryText,
      }}
    >
      {/* Background image */}
      {event.event.presigned_url ? (
        <Image
          source={{ uri: event.event.presigned_url }}
          style={{ position: "absolute", width: "100%", height: "100%" }}
          contentFit="cover"
        />
      ) : (
        <>
          <View
            style={{
              position: "absolute",
              width: 140,
              height: 140,
              borderRadius: 70,
              top: -20,
              left: 20,
              backgroundColor: AppColors.purple,
              opacity: 0.95,
            }}
          />
          <View
            style={{
              position: "absolute",
              width: 120,
              height: 120,
              borderRadius: 60,
              top: 10,
              left: 90,
              backgroundColor: AppColors.primaryBlue,
              opacity: 0.95,
            }}
          />
          <View
            style={{
              position: "absolute",
              width: 100,
              height: 100,
              borderRadius: 50,
              top: -5,
              left: 170,
              backgroundColor: AppColors.green,
              opacity: 0.95,
            }}
          />
        </>
      )}

      {/* Bottom overlay */}
      <View
        style={{
          position: "absolute",
          bottom: 0,
          left: 0,
          right: 0,
          backgroundColor: "rgba(0,0,0,0.60)",
          borderBottomLeftRadius: 24,
          borderBottomRightRadius: 24,
          paddingHorizontal: 16,
          paddingTop: 14,
          paddingBottom: 14,
        }}
      >
        <View
          style={{
            flexDirection: "row",
            alignItems: "flex-end",
            justifyContent: "space-between",
          }}
        >
          {/* Left: title + rating */}
          <View style={{ flex: 1, marginRight: 12 }}>
            <Text
              style={{
                color: "white",
                fontFamily: FontFamilies.bold,
                fontSize: 18,
                marginBottom: 5,
                lineHeight: 22,
              }}
              numberOfLines={1}
            >
              {event.event.title}
            </Text>
            {totalReviews > 0 && ratingMatch && (
              <View style={{ flexDirection: "row", alignItems: "center", gap: 5 }}>
                <Image
                  source={ratingMatch.image}
                  style={{ width: 20, height: 20 }}
                />
                <Text
                  style={{
                    color: "white",
                    fontFamily: FontFamilies.semiBold,
                    fontSize: FontSizes.sm,
                  }}
                >
                  {avgRating.toFixed(1)}
                </Text>
                <Text
                  style={{
                    color: "rgba(255,255,255,0.65)",
                    fontFamily: FontFamilies.regular,
                    fontSize: FontSizes.sm,
                  }}
                >
                  ({totalReviews} {totalReviews === 1 ? "Review" : "Reviews"})
                </Text>
              </View>
            )}
          </View>

          {/* See More button */}
          <Pressable
            onPress={() => router.push(`/event/${eventId}`)}
            style={{
              flexDirection: "row",
              alignItems: "center",
              backgroundColor: "white",
              borderRadius: 20,
              paddingHorizontal: 14,
              paddingVertical: 8,
              gap: 4,
            }}
          >
            <Text
              style={{
                fontFamily: FontFamilies.bold,
                fontSize: FontSizes.sm,
                color: AppColors.primaryText,
              }}
            >
              See More
            </Text>
            <IconSymbol
              name="chevron.right"
              size={12}
              color={AppColors.primaryText}
            />
          </Pressable>
        </View>
      </View>
    </Pressable>
  );
}
