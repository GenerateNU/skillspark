import { Image } from "expo-image";
import { View, Text, Pressable } from "react-native";
import { useRouter } from "expo-router";
import { type EventOccurrence } from "@skillspark/api-client";
import { AppColors, FontFamilies, FontSizes } from "@/constants/theme";
import { formatEventDate, formatEventTime } from "@/utils/format";

export function UpcomingClassCard({
  occurrence,
}: {
  occurrence: EventOccurrence;
}) {
  const router = useRouter();
  const location = [
    occurrence.location?.address_line1,
    occurrence.location?.district,
  ]
    .filter(Boolean)
    .join(", ");
  const badge = occurrence.event.category?.[0];

  return (
    <Pressable
      onPress={() => router.push(`/event/${occurrence.id}`)}
      className="mr-4 w-[310px] rounded-2xl"
      style={{
        backgroundColor: AppColors.white,
        shadowColor: "#000",
        shadowOpacity: 0.07,
        shadowRadius: 10,
        shadowOffset: { width: 0, height: 2 },
        elevation: 3,
      }}
    >
      <View className="flex-row p-3 gap-3 items-center">
        {/* Image */}
        <View className="w-[88px] h-[88px] rounded-[12px] overflow-hidden">
          {occurrence.event.presigned_url ? (
            <Image
              source={{ uri: occurrence.event.presigned_url }}
              className="w-[88px] h-[88px]"
              contentFit="cover"
            />
          ) : (
            <View
              className="w-[88px] h-[88px]"
              style={{ backgroundColor: AppColors.divider }}
            />
          )}
        </View>

        {/* Text */}
        <View className="flex-1 gap-0.5">
          <Text
            style={{
              fontFamily: FontFamilies.bold,
              fontSize: FontSizes.base,
              color: AppColors.primaryText,
            }}
            numberOfLines={1}
          >
            {occurrence.event.title}
          </Text>
          <Text
            style={{
              fontSize: FontSizes.sm,
              color: AppColors.mutedText,
              fontFamily: FontFamilies.regular,
            }}
          >
            {formatEventDate(occurrence.start_time)}
          </Text>
          <Text
            style={{
              fontSize: FontSizes.sm,
              color: AppColors.mutedText,
              fontFamily: FontFamilies.regular,
            }}
          >
            {formatEventTime(occurrence.start_time, occurrence.end_time)}
          </Text>
          {!!location && (
            <Text
              style={{
                fontSize: FontSizes.xs,
                color: AppColors.subtleText,
                fontFamily: FontFamilies.regular,
              }}
              numberOfLines={1}
            >
              {location}
            </Text>
          )}
        </View>

        {/* Badge */}
        {!!badge && (
          <View
            className="rounded-full px-[10px] py-1 self-start mt-0.5"
            style={{ backgroundColor: AppColors.badgeGreenBg }}
          >
            <Text
              style={{
                fontSize: FontSizes.xs,
                color: AppColors.badgeGreenText,
                fontFamily: FontFamilies.semiBold,
              }}
            >
              {badge}
            </Text>
          </View>
        )}
      </View>
    </Pressable>
  );
}
