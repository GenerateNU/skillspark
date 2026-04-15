import { IconSymbol } from "@/components/ui/icon-symbol";
import { RATING_OPTIONS } from "@/constants/ratings";
import { AppColors } from "@/constants/theme";
import { type Event, type SimpleReviewAggregate } from "@skillspark/api-client";
import { Image } from "expo-image";
import { useTranslation } from "react-i18next";
import { Text, TouchableOpacity, View } from "react-native";

interface EventRatingCardProps {
  event: Event;
  aggregate: SimpleReviewAggregate | null;
  onPress: () => void;
}

export function EventRatingCard({
  event,
  aggregate,
  onPress,
}: EventRatingCardProps) {
  const { t: translate } = useTranslation();
  const avg = aggregate?.average_rating ?? 0;
  const total = aggregate?.total_reviews ?? 0;
  const match = RATING_OPTIONS.find((r) => r.rating === Math.round(avg));

  return (
    <TouchableOpacity
      activeOpacity={0.8}
      onPress={onPress}
      className="flex-row items-center rounded-2xl bg-white p-3 gap-3"
      style={{
        shadowColor: "#000",
        shadowOpacity: 0.06,
        shadowRadius: 8,
        shadowOffset: { width: 0, height: 2 },
        elevation: 2,
      }}
    >
      <View
        className="w-[72px] h-[72px] rounded-xl overflow-hidden"
        style={{ backgroundColor: AppColors.imagePlaceholder }}
      >
        {event.presigned_url ? (
          <Image
            source={{ uri: event.presigned_url }}
            style={{ width: "100%", height: "100%" }}
            contentFit="cover"
          />
        ) : (
          <View className="flex-1 items-center justify-center">
            <IconSymbol name="photo" size={24} color={AppColors.mutedText} />
          </View>
        )}
      </View>

      <View className="flex-1">
        <Text
          className="text-[15px] font-nunito-bold mb-1"
          style={{ color: AppColors.primaryText }}
        >
          {event.title}
        </Text>
        {total > 0 && match ? (
          <View className="flex-row items-center gap-1.5">
            <Image source={match.image} style={{ width: 18, height: 18 }} />
            <Text
              className="text-[13px] font-nunito"
              style={{ color: AppColors.secondaryText }}
            >
              {translate(match.labelKey)}{" "}
              <Text style={{ color: AppColors.subtleText }}>({total})</Text>
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
      </View>
    </TouchableOpacity>
  );
}
