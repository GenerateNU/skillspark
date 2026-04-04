import { ReviewAggregate } from "@skillspark/api-client";
import { Image, Text, View } from "react-native";

import { RATING_OPTIONS } from "@/constants/ratings";
import { AppColors } from "@/constants/theme";
import { useTranslation } from "react-i18next";

export const NO_REVIEWS_IMAGE = require("@/assets/images/noreview.png");

export function RatingAggregateCard({ aggregate }: { aggregate: ReviewAggregate }) {
  const total = aggregate.total_reviews ?? 0;
  const avg = aggregate.average_rating ?? 0;

  const { t: translate } = useTranslation();

  return (
    <View className="px-5 py-4">
      <View>
        {total === 0 ? (
            <View className="flex-row items-center gap-3 mb-1">
                <Image source={NO_REVIEWS_IMAGE} style={{ width: 40, height: 40 }} />
                <Text className="text-4xl font-bold mb-1" style={{ color: AppColors.primaryText }}>
                    {translate("review.noReviews")}
                </Text>
            </View>
            ) : (() => {
            const match = RATING_OPTIONS.find(r => r.rating === Math.round(avg));
            return (
                <View className="flex-row items-center gap-3 mb-1">
                <Image source={match?.image} style={{ width: 40, height: 40 }} />
                <Text className="text-4xl font-bold" style={{ color: AppColors.primaryText }}>
                    {translate(match?.labelKey ?? "")}
                </Text>
                </View>
            );
        })()}
        <Text
            className="text-sm mb-5"
            style={{ color: AppColors.subtleText }}
        >
            {total} reviews
        </Text>
      </View>

    <View className="mb-3" style={{ height: 0.5, backgroundColor: AppColors.borderLight }} />

      <View className="gap-3">
  {total > 0 && RATING_OPTIONS.slice().reverse().map(({ rating, image, labelKey }) => {
    const count = aggregate.breakdown.find(b => b.rating === rating)?.review_count ?? 0;
    const pct = total > 0 ? (count / total) * 100 : 0;
    return (
      <View key={rating} className="gap-1">
        <View className="flex-row items-center gap-2">
          <Image source={image} style={{ width: 18, height: 18 }} />
          <Text className="text-xs" style={{ color: AppColors.secondaryText }}>
            {translate(labelKey)}
          </Text>
        </View>
        <View
          className="w-full h-2 rounded-full overflow-hidden"
          style={{ backgroundColor: AppColors.borderLight }}
        >
          <View
            className="h-full rounded-full"
            style={{ width: `${pct}%`, backgroundColor: AppColors.primaryText }}
          />
        </View>
      </View>
    );
  })}
</View>
    </View>
  );
}