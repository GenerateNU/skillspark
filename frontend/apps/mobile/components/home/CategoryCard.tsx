import { Image } from "expo-image";
import { View, Text, Pressable } from "react-native";
import { type EventOccurrence } from "@skillspark/api-client";
import { AppColors, FontSizes } from "@/constants/theme";

const CATEGORY_COLORS: Record<string, string> = {
  Sport: AppColors.blue,
  Arts: AppColors.violet,
  Music: AppColors.pink,
  Tech: AppColors.emerald,
  Activity: AppColors.amber,
  Tutoring: AppColors.danger,
};

export function CategoryCard({ category, occurrence }: { category: string; occurrence?: EventOccurrence }) {
  return (
    <Pressable
      style={{
        flex: 1,
        height: 110,
        margin: 5,
        borderRadius: 18,
        overflow: "hidden",
      }}
    >
      {occurrence?.event.presigned_url ? (
        <Image
          source={{ uri: occurrence.event.presigned_url }}
          style={{ position: "absolute", width: "100%", height: "100%" }}
          contentFit="cover"
        />
      ) : (
        <View
          style={{
            position: "absolute",
            width: "100%",
            height: "100%",
            backgroundColor: CATEGORY_COLORS[category] ?? AppColors.categoryFallback,
          }}
        />
      )}
      {/* Gradient-style overlay */}
      <View
        style={{
          position: "absolute",
          bottom: 0,
          left: 0,
          right: 0,
          height: "55%",
          backgroundColor: AppColors.cardOverlay,
          justifyContent: "flex-end",
          paddingHorizontal: 12,
          paddingBottom: 10,
        }}
      >
        <Text style={{ color: AppColors.white, fontFamily: "NunitoSans_700Bold", fontSize: FontSizes.base }}>{category}</Text>
      </View>
    </Pressable>
  );
}
