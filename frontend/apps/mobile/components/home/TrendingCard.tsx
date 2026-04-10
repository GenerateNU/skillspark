import { Image } from "expo-image";
import { View, Text, Pressable } from "react-native";
import { useRouter } from "expo-router";
import { type EventOccurrence } from "@skillspark/api-client";
import { AppColors, FontFamilies, FontSizes } from "@/constants/theme";
import { StarRating } from "@/components/StarRating";

const AVATAR_COLORS = [
  AppColors.purple,
  AppColors.emerald,
  AppColors.blue,
  AppColors.amber,
  AppColors.pink,
];

function hashColor(initials: string): string {
  let hash = 0;
  for (let i = 0; i < initials.length; i++) {
    hash = initials.charCodeAt(i) + ((hash << 5) - hash);
  }
  return AVATAR_COLORS[Math.abs(hash) % AVATAR_COLORS.length];
}

export function TrendingCard({ occurrence }: { occurrence: EventOccurrence }) {
  const router = useRouter();
  const ageLabel =
    occurrence.event.age_range_min != null
      ? `Ages ${occurrence.event.age_range_min}${
          occurrence.event.age_range_max != null
            ? ` - ${occurrence.event.age_range_max}`
            : "+"
        }`
      : null;
  // Placeholder initials — not real user data
  const letter1 = "E";
  const letter2 = "S";
  const color1 = hashColor(letter1);
  const color2 = hashColor(letter2);

  return (
    <View className="mr-[14px] pb-2">
      {/* Avatar circles floating above top-right of card */}
      <View className="flex-row justify-end pr-[14px] -mb-[18px] z-[1]">
        <View
          className="w-9 h-9 rounded-full border-[2.5px] items-center justify-center"
          style={{ backgroundColor: color1, borderColor: AppColors.white }}
        >
          <Text
            style={{
              color: AppColors.white,
              fontSize: FontSizes.md,
              fontFamily: FontFamilies.bold,
            }}
          >
            {letter1}
          </Text>
        </View>
        <View
          className="w-9 h-9 rounded-full border-[2.5px] items-center justify-center -ml-[10px]"
          style={{ backgroundColor: color2, borderColor: AppColors.white }}
        >
          <Text
            style={{
              color: AppColors.white,
              fontSize: FontSizes.md,
              fontFamily: FontFamilies.bold,
            }}
          >
            {letter2}
          </Text>
        </View>
      </View>

      {/* Card */}
      <Pressable
        onPress={() => router.push(`/event/${occurrence.id}`)}
        className="w-[218px] h-[121px] rounded-[12px] border flex-row items-center py-[6px] px-2 gap-2"
        style={{
          backgroundColor: AppColors.white,
          borderColor: AppColors.savedBackground,
          shadowColor: "#000",
          shadowOpacity: 0.1,
          shadowRadius: 3,
          shadowOffset: { width: 0, height: 4 },
          elevation: 2,
        }}
      >
        {/* Image */}
        <View className="w-[88px] h-[88px] rounded-[12px] overflow-hidden shrink-0">
          {occurrence.event.presigned_url ? (
            <Image
              source={{ uri: occurrence.event.presigned_url }}
              style={{ width: "100%", height: "100%" }}
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
        <View className="flex-1 gap-1">
          <Text
            style={{
              fontFamily: FontFamilies.bold,
              fontSize: FontSizes.base,
              color: AppColors.primaryText,
            }}
            numberOfLines={2}
          >
            {occurrence.event.title}
          </Text>
          <StarRating size={12} rating={0} />
          {ageLabel && (
            <Text
              style={{
                fontSize: FontSizes.sm,
                color: AppColors.mutedText,
                fontFamily: FontFamilies.regular,
              }}
            >
              {ageLabel}
            </Text>
          )}
        </View>
      </Pressable>
    </View>
  );
}
