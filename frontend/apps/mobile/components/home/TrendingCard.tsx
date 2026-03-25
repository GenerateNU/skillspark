import { Image } from "expo-image";
import { View, Text, Pressable } from "react-native";
import { useRouter } from "expo-router";
import { type EventOccurrence } from "@skillspark/api-client";
import { AppColors, FontSizes } from "@/constants/theme";
import { StarRating } from "@/components/StarRating";

const AVATAR_COLORS = [AppColors.purple, AppColors.emerald, AppColors.blue, AppColors.amber, AppColors.pink];

export function TrendingCard({ occurrence, index }: { occurrence: EventOccurrence; index: number }) {
  const router = useRouter();
  const ageLabel =
    occurrence.event.age_range_min != null
      ? `Ages ${occurrence.event.age_range_min}${
          occurrence.event.age_range_max != null ? ` - ${occurrence.event.age_range_max}` : "+"
        }`
      : null;
  const color1 = AVATAR_COLORS[index % AVATAR_COLORS.length];
  const color2 = AVATAR_COLORS[(index + 2) % AVATAR_COLORS.length];
  // Placeholder initials — not real user data
  const letter1 = "E";
  const letter2 = "S";

  return (
    <View style={{ marginRight: 14 }}>
      {/* Avatar circles floating above top-right of card */}
      <View style={{ flexDirection: "row", justifyContent: "flex-end", paddingRight: 14, marginBottom: -18, zIndex: 1 }}>
        <View style={{ width: 36, height: 36, borderRadius: 18, backgroundColor: color1, borderWidth: 2.5, borderColor: AppColors.white, alignItems: "center", justifyContent: "center" }}>
          <Text style={{ color: AppColors.white, fontSize: FontSizes.md, fontFamily: "NunitoSans_700Bold" }}>{letter1}</Text>
        </View>
        <View style={{ width: 36, height: 36, borderRadius: 18, backgroundColor: color2, borderWidth: 2.5, borderColor: AppColors.white, marginLeft: -10, alignItems: "center", justifyContent: "center" }}>
          <Text style={{ color: AppColors.white, fontSize: FontSizes.md, fontFamily: "NunitoSans_700Bold" }}>{letter2}</Text>
        </View>
      </View>

      {/* Card */}
      <Pressable
        onPress={() => router.push(`/event/${occurrence.id}`)}
        style={{
          width: 200,
          backgroundColor: AppColors.white,
          borderRadius: 20,
          borderWidth: 1,
          borderColor: AppColors.savedBackground,
          flexDirection: "row",
          alignItems: "center",
          paddingVertical: 6,
          paddingHorizontal: 8,
          gap: 8,
        }}
      >
        {/* Image */}
        <View style={{ width: 88, height: 88, borderRadius: 12, overflow: "hidden", flexShrink: 0 }}>
          {occurrence.event.presigned_url ? (
            <Image
              source={{ uri: occurrence.event.presigned_url }}
              style={{ width: 88, height: 88 }}
              contentFit="cover"
            />
          ) : (
            <View style={{ width: 88, height: 88, backgroundColor: AppColors.divider }} />
          )}
        </View>

        {/* Text */}
        <View style={{ flex: 1, gap: 4 }}>
          <Text style={{ fontFamily: "NunitoSans_700Bold", fontSize: FontSizes.base, color: AppColors.primaryText }} numberOfLines={2}>
            {occurrence.event.title}
          </Text>
          <StarRating size={12} rating={0} />
          {ageLabel && (
            <Text style={{ fontSize: FontSizes.sm, color: AppColors.mutedText, fontFamily: "NunitoSans_400Regular" }}>{ageLabel}</Text>
          )}
        </View>
      </Pressable>
    </View>
  );
}
