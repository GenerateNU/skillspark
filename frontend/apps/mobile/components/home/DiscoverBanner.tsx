import { Image } from "expo-image";
import { View, Text, Pressable } from "react-native";
import { useRouter } from "expo-router";
import { type EventOccurrence } from "@skillspark/api-client";
import { AppColors, FontSizes } from "@/constants/theme";

export function DiscoverBanner({ event }: { event: EventOccurrence }) {
  const router = useRouter();
  const ageLabel =
    event.event.age_range_min != null
      ? `${event.event.age_range_min}${event.event.age_range_max != null ? `–${event.event.age_range_max}` : ""}+`
      : null;

  return (
    <Pressable
      onPress={() => router.push(`/event/${event.id}`)}
      className="mx-5 rounded-3xl overflow-hidden h-[188px]"
      style={{ backgroundColor: AppColors.primaryText }}
    >
      {event.event.presigned_url ? (
        <Image
          source={{ uri: event.event.presigned_url }}
          className="absolute inset-0 opacity-50"
          contentFit="cover"
        />
      ) : (
        <>
          <View className="absolute w-[140px] h-[140px] rounded-full opacity-95" style={{ backgroundColor: AppColors.purple, top: -20, left: 20 }} />
          <View className="absolute w-[120px] h-[120px] rounded-full opacity-95" style={{ backgroundColor: AppColors.primaryBlue, top: 10, left: 90 }} />
          <View className="absolute w-[100px] h-[100px] rounded-full opacity-95" style={{ backgroundColor: AppColors.green, top: -5, left: 170 }} />
          <View
            className="absolute w-[88px] h-[108px] bg-white rounded-2xl items-center justify-center gap-1.5 p-2.5"
            style={{
              top: "50%",
              left: "50%",
              transform: [{ translateX: -44 }, { translateY: -54 }],
              shadowColor: "#000",
              shadowOpacity: 0.25,
              shadowRadius: 12,
            }}
          >
            <View className="w-9 h-9 rounded-full" style={{ backgroundColor: AppColors.mintGreen }} />
            <View className="w-[52px] h-[7px] rounded-sm" style={{ backgroundColor: AppColors.divider }} />
            <View className="w-[38px] h-[7px] rounded-sm" style={{ backgroundColor: AppColors.surfaceGray }} />
          </View>
        </>
      )}
      <View
        className="absolute top-[18px] right-[18px] bg-white rounded-[10px] px-[11px] py-1.5"
        style={{ transform: [{ rotate: "12deg" }], shadowColor: "#000", shadowOpacity: 0.15, shadowRadius: 6 }}
      >
        <Text style={{ fontFamily: "NunitoSans_700Bold", fontSize: FontSizes.md, color: AppColors.primaryText }}>
          {ageLabel ?? event.event.title.slice(0, 6)}
        </Text>
      </View>
    </Pressable>
  );
}
