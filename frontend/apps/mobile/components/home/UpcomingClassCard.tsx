import { Image } from "expo-image";
import { View, Text, Pressable } from "react-native";
import { useRouter } from "expo-router";
import { type EventOccurrence } from "@skillspark/api-client";
import { AppColors, FontSizes } from "@/constants/theme";
import { formatEventDate, formatEventTime } from "@/utils/format";

export function UpcomingClassCard({ occurrence }: { occurrence: EventOccurrence }) {
  const router = useRouter();
  const location = [occurrence.location?.address_line1, occurrence.location?.district]
    .filter(Boolean)
    .join(", ");
  const badge = occurrence.event.category?.[0];

  return (
    <Pressable
      onPress={() => router.push(`/event/${occurrence.id}`)}
      className="mr-4"
      style={{
        width: 310,
        backgroundColor: AppColors.white,
        borderRadius: 16,
        shadowColor: "#000",
        shadowOpacity: 0.07,
        shadowRadius: 10,
        shadowOffset: { width: 0, height: 2 },
        elevation: 3,
      }}
    >
      <View style={{ flexDirection: "row", padding: 12, gap: 12, alignItems: "center" }}>
        {/* Image */}
        <View style={{ width: 88, height: 88, borderRadius: 12, overflow: "hidden" }}>
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
        <View style={{ flex: 1, gap: 2 }}>
          <Text style={{ fontFamily: "NunitoSans_700Bold", fontSize: FontSizes.base, color: AppColors.primaryText }} numberOfLines={1}>
            {occurrence.event.title}
          </Text>
          <Text style={{ fontSize: FontSizes.sm, color: AppColors.mutedText, fontFamily: "NunitoSans_400Regular" }}>
            {formatEventDate(occurrence.start_time)}
          </Text>
          <Text style={{ fontSize: FontSizes.sm, color: AppColors.mutedText, fontFamily: "NunitoSans_400Regular" }}>
            {formatEventTime(occurrence.start_time, occurrence.end_time)}
          </Text>
          {!!location && (
            <Text style={{ fontSize: FontSizes.xs, color: AppColors.subtleText, fontFamily: "NunitoSans_400Regular" }} numberOfLines={1}>
              {location}
            </Text>
          )}
        </View>

        {/* Badge */}
        {!!badge && (
          <View style={{ backgroundColor: AppColors.badgeGreenBg, borderRadius: 20, paddingHorizontal: 10, paddingVertical: 4, alignSelf: "flex-start", marginTop: 2 }}>
            <Text style={{ fontSize: FontSizes.xs, color: AppColors.badgeGreenText, fontFamily: "NunitoSans_600SemiBold" }}>{badge}</Text>
          </View>
        )}
      </View>
    </Pressable>
  );
}
