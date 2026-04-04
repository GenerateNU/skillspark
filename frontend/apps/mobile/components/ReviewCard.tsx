import { RATING_OPTIONS } from "@/constants/ratings";
import { AppColors } from "@/constants/theme";
import { Review } from "@skillspark/api-client";
import { useTranslation } from "react-i18next";
import { Image, Text, View } from "react-native";

function timeAgo(dateStr: string) {
  const diff = Date.now() - new Date(dateStr).getTime();
  const days = Math.floor(diff / (1000 * 60 * 60 * 24));
  if (days === 0) return "Today";
  if (days === 1) return "1 day ago";
  if (days < 30) return `${days} days ago`;
  const months = Math.floor(days / 30);
  if (months === 1) return "1 month ago";
  if (months < 12) return `${months} months ago`;
  const years = Math.floor(months / 12);
  return years === 1 ? "1 year ago" : `${years} years ago`;
}

export function ReviewCard({ review }: { review: Review }) {

  const { t: translate } = useTranslation();
  const match = RATING_OPTIONS.find(r => r.rating === review.rating);

  return (
    <View
      className="flex-row gap-3 p-4 rounded-2xl"
    >
      <View className="items-center pt-1">
        {match && <Image source={match.image} style={{ width: 36, height: 36 }} />}
      </View>

      <View className="flex-1 gap-2">
        <Text className="text-sm leading-5" style={{ color: AppColors.primaryText }}>
          {review.description}
        </Text>

        {review.categories?.length > 0 && (
          <View className="flex-row flex-wrap gap-1">
            {review.categories.map((cat) => (
              <View
                key={cat}
                className="px-3 py-1 rounded-full"
                style={{ backgroundColor: AppColors.ratingPill }}
                >
                <Text className="text-xs" style={{ color: AppColors.secondaryText }}>
                  {cat}
                </Text>
              </View>
            ))}
          </View>
        )}

        <View className="flex-row justify-between items-center">
          <Text className="text-xs" style={{ color: AppColors.subtleText }}>
            Anonymous
          </Text> 
          <Text className="text-xs" style={{ color: AppColors.subtleText }}>
            {timeAgo(review.created_at)}
          </Text>
        </View>
      </View>
    </View>
  );
}