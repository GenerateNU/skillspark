import { View, Text, Pressable } from "react-native";
import { EventImage } from "@/components/EventImage";
import { useRouter } from "expo-router";
import { type Event } from "@skillspark/api-client";
import { AppColors, FontFamilies, FontSizes } from "@/constants/theme";

export function SearchResultCard({ event }: { event: Event }) {
  const router = useRouter();
  const badge = event.category?.[0];
  const ageLabel =
    event.age_range_min != null
      ? `Ages ${event.age_range_min}${
          event.age_range_max != null ? ` – ${event.age_range_max}` : "+"
        }`
      : null;

  return (
    <Pressable
      onPress={() => router.push(`/event/${event.id}`)}
      style={{
        height: 118,
        borderRadius: 12,
        backgroundColor: AppColors.white,
        flexDirection: "row",
        alignItems: "center",
        paddingVertical: 12,
        paddingHorizontal: 12,
        gap: 12,
        shadowColor: "#000",
        shadowOpacity: 0.07,
        shadowRadius: 8,
        shadowOffset: { width: 0, height: 2 },
        elevation: 2,
      }}
    >
      {/* Image */}
      <View
        style={{
          width: 94,
          height: 94,
          borderRadius: 10,
          overflow: "hidden",
          flexShrink: 0,
        }}
      >
        <EventImage
          uri={event.presigned_url}
          style={{ width: "100%", height: "100%" }}
        />
      </View>

      {/* Text */}
      <View style={{ flex: 1, gap: 4 }}>
        <Text
          style={{
            fontFamily: FontFamilies.bold,
            fontSize: FontSizes.base,
            color: AppColors.primaryText,
          }}
          numberOfLines={2}
        >
          {event.title}
        </Text>
        {ageLabel && (
          <Text
            style={{
              fontFamily: FontFamilies.regular,
              fontSize: FontSizes.xs,
              color: AppColors.subtleText,
            }}
          >
            {ageLabel}
          </Text>
        )}
        {!!badge && (
          <View
            style={{
              alignSelf: "flex-start",
              backgroundColor: AppColors.badgeGreenBg,
              borderRadius: 999,
              paddingHorizontal: 8,
              paddingVertical: 2,
            }}
          >
            <Text
              style={{
                fontFamily: FontFamilies.semiBold,
                fontSize: FontSizes.xs,
                color: AppColors.badgeGreenText,
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
