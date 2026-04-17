import { RATING_OPTIONS } from "@/constants/ratings";
import { AppColors } from "@/constants/theme";
import { Review, useGetGuardianById } from "@skillspark/api-client";
import { useTranslation } from "react-i18next";
import { Image, Text, View } from "react-native";

function timeAgo(dateStr: string, translate: (key: string) => string) {
  const diff = Date.now() - new Date(dateStr).getTime();
  const days = Math.floor(diff / (1000 * 60 * 60 * 24));
  if (days === 0) return translate("time.today");
  if (days === 1) return translate("time.oneDayAgo");
  if (days < 30) return `${days} ${translate("time.daysAgo")}`;
  const months = Math.floor(days / 30);
  if (months === 1) return translate("time.oneMonthAgo");
  if (months < 12) return `${months} ${translate("time.monthsAgo")}`;
  const years = Math.floor(months / 12);
  return years === 1
    ? translate("time.oneYearAgo")
    : `${years} ${translate("time.yearsAgo")}`;
}

export function ReviewCard({ review, hideAvatar }: { review: Review; hideAvatar?: boolean }) {
  const { t: translate } = useTranslation();
  const match = RATING_OPTIONS.find((r) => r.rating === review.rating);

  const { data: guardianResp } = useGetGuardianById(review.guardian_id ?? "", {
    query: { enabled: !!review.guardian_id },
  });
  const guardianName =
    (guardianResp as unknown as { data: { name: string } } | undefined)?.data
      ?.name ?? null;

  const displayName = review.guardian_id
    ? (guardianName ?? "...")
    : translate("review.anonymous");

  return (
    <View className="flex-row gap-3 p-4 rounded-2xl">
      {!hideAvatar && (
        <View className="items-center pt-1">
          {match && <Image source={match.image} className="w-9 h-9" />}
        </View>
      )}

      <View className="flex-1 gap-2">
        <Text
          className="text-sm leading-5"
          style={{ color: AppColors.primaryText }}
        >
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
                <Text
                  className="text-xs"
                  style={{ color: AppColors.secondaryText }}
                >
                  {cat}
                </Text>
              </View>
            ))}
          </View>
        )}

        <View className="flex-row justify-between items-center">
          <Text className="text-xs" style={{ color: AppColors.subtleText }}>
            {displayName}
          </Text>
          <Text className="text-xs" style={{ color: AppColors.subtleText }}>
            {timeAgo(review.created_at, translate)}
          </Text>
        </View>
      </View>
    </View>
  );
}
