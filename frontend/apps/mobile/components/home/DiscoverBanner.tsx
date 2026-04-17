import { Image } from "expo-image";
import { View, Text, Pressable } from "react-native";
import { useRouter } from "expo-router";
import { type Event, useGetReviewAggregate } from "@skillspark/api-client";
import { AppColors, FontFamilies } from "@/constants/theme";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { RATING_OPTIONS } from "@/constants/ratings";

export function DiscoverBanner({ event }: { event: Event }) {
  const router = useRouter();
  const eventId = event.id;

  const { data: aggregateResp } = useGetReviewAggregate(eventId, {
    query: {
      enabled: !!eventId,
      staleTime: 5 * 60 * 1000,
      gcTime: 30 * 60 * 1000,
      refetchOnMount: false,
    },
  });
  const aggregateData =
    aggregateResp?.status === 200 ? aggregateResp.data : null;
  const avgRating = aggregateData?.average_rating ?? 0;
  const totalReviews = aggregateData?.total_reviews ?? 0;
  const ratingMatch = RATING_OPTIONS.find(
    (r) => r.rating === Math.round(avgRating),
  );

  return (
    <View className="flex-1 rounded-3xl overflow-hidden bg-[#1a1a1a]">
      {/* Background image */}
      {event.presigned_url ? (
        <Image
          source={{ uri: event.presigned_url }}
          style={{ position: "absolute", width: "100%", height: "100%" }}
          contentFit="cover"
        />
      ) : (
        <>
          <View className="absolute w-[140px] h-[140px] rounded-full -top-5 left-5 bg-purple-700 opacity-95" />
          <View className="absolute w-[120px] h-[120px] rounded-full top-[10px] left-[90px] bg-blue-600 opacity-95" />
          <View className="absolute w-[100px] h-[100px] rounded-full -top-[5px] left-[170px] bg-green-500 opacity-95" />
        </>
      )}

      {/* Bottom overlay */}
      <View className="absolute bottom-0 left-0 right-0 bg-black/60 rounded-b-3xl px-4 pt-[14px] pb-[14px]">
        <View className="flex-row items-end justify-between">
          {/* Left: title + rating */}
          <View className="flex-1 mr-3">
            <Text
              className="text-white text-lg mb-[5px] leading-[22px]"
              style={{ fontFamily: FontFamilies.bold }}
              numberOfLines={1}
            >
              {event.title}
            </Text>
            {totalReviews > 0 && ratingMatch && (
              <View className="flex-row items-center gap-[5px]">
                <Image
                  source={ratingMatch.image}
                  style={{ width: 20, height: 20 }}
                />
                <Text
                  className="text-white text-sm"
                  style={{ fontFamily: FontFamilies.semiBold }}
                >
                  {avgRating.toFixed(1)}
                </Text>
                <Text
                  className="text-white/65 text-sm"
                  style={{ fontFamily: FontFamilies.regular }}
                >
                  ({totalReviews} {totalReviews === 1 ? "Review" : "Reviews"})
                </Text>
              </View>
            )}
          </View>

          {/* See More button */}
          <Pressable
            onPress={() => router.push(`/event/${eventId}`)}
            className="flex-row items-center bg-white rounded-full px-[14px] py-2 gap-1"
          >
            <Text
              className="text-sm"
              style={{
                fontFamily: FontFamilies.bold,
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
    </View>
  );
}
