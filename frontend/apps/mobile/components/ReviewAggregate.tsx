import { RATING_OPTIONS } from "@/constants/ratings";
import { AppColors } from "@/constants/theme";
import { ReviewAggregate } from "@skillspark/api-client";
import { useTranslation } from "react-i18next";
import { Image, Text, View } from "react-native";

export const NO_REVIEWS_IMAGE = require("@/assets/images/ratings/noreview.png");

export function RatingAggregateCard({
  aggregate,
}: {
  aggregate: ReviewAggregate;
}) {
  const total = aggregate.total_reviews ?? 0;
  const avg = aggregate.average_rating ?? 0;
  const { t: translate } = useTranslation();
  return (
    <View className="px-5 py-4">
      <View>
        {total === 0 ? (
          <View className="flex-row items-center gap-3 mb-1">
            <Image source={NO_REVIEWS_IMAGE} className="w-10 h-10" />
            <Text
              className="text-4xl font-bold mb-1"
              style={{ color: AppColors.primaryText }}
            >
              {translate("review.noReviews")}
            </Text>
          </View>
        ) : (
          (() => {
            const match = RATING_OPTIONS.find(
              (r) => r.rating === Math.round(avg)
            );
            return (
              <View className="flex-row items-center gap-3 mb-1">
                <Image source={match?.image} className="w-10 h-10" />
                <Text
                  className="text-4xl font-bold"
                  style={{ color: AppColors.primaryText }}
                >
                  {translate(match?.labelKey ?? "")}
                </Text>
              </View>
            );
          })()
        )}
        <Text className="text-sm mb-5" style={{ color: AppColors.subtleText }}>
          {total} {translate("review.reviews")}
        </Text>
      </View>
      <View
        className="mb-3"
        style={{ height: 0.5, backgroundColor: AppColors.borderLight }}
      />
      <View className="gap-3">
        {total > 0 &&
          RATING_OPTIONS.filter((r) => r.rating !== null).map(({ rating, image, labelKey }) => {
            const count =
              aggregate.breakdown?.find((b) => b.rating === rating)
                ?.review_count ?? 0;
            const pct = total > 0 ? (count / total) * 100 : 0;
            return (
              <View key={rating} className="gap-1">
                <View className="flex-row items-center gap-2">
                  <Image source={image} className="w-4 h-4" />
                  <Text
                    className="text-xs"
                    style={{ color: AppColors.secondaryText }}
                  >
                    {translate(labelKey ?? "")}
                  </Text>
                </View>
                <View
                  className="w-full h-2 rounded-full overflow-hidden"
                  style={{ backgroundColor: AppColors.borderLight }}
                >
                  <View
                    className="h-full rounded-full"
                    style={{
                      width: `${pct}%`,
                      backgroundColor: AppColors.primaryText,
                    }}
                  />
                </View>
              </View>
            );
          })}
      </View>
    </View>
  );
}