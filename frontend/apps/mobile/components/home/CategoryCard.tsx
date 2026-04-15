import { Image } from "expo-image";
import { View, Text, Pressable } from "react-native";
import { type EventOccurrence } from "@skillspark/api-client";
import { AppColors, FontFamilies, FontSizes } from "@/constants/theme";
import { useRouter } from "expo-router";

const CATEGORY_COLORS: Record<string, string> = {
  science: AppColors.emerald,
  technology: AppColors.blue,
  soccer: AppColors.amber,
  painting: AppColors.violet,
  dance: AppColors.pink,
  coding: AppColors.danger,
  basketball: AppColors.amber,
  swimming: AppColors.blue,
  "martial arts": AppColors.danger,
  "vocal music": AppColors.pink,
  "instrumental music": AppColors.pink,
  theater: AppColors.violet,
  drawing: AppColors.violet,
  photography: AppColors.emerald,
};



export function CategoryCard({
  category,
  occurrence,
}: {
  category: string;
  occurrence?: EventOccurrence;
}) {
  const router = useRouter();
  return (
    <Pressable
      onPress={() => router.push({ pathname: "/event-categories", params: { category } })}
      className="flex-1 m-[5px]"
      style={{
        shadowColor: "#000",
        shadowOpacity: 0.25,
        shadowRadius: 4,
        shadowOffset: { width: 0, height: 4 },
        elevation: 3,
      }}
    >
      <View className="h-[80px] rounded-[15px] overflow-hidden">
        {occurrence?.event.presigned_url ? (
          <Image
            source={{ uri: occurrence.event.presigned_url }}
            style={{ width: "100%", height: "100%" }}
            className="absolute inset-0 opacity-80"
            contentFit="cover"
          />
        ) : (
          <View
            className="absolute inset-0"
            style={{
              backgroundColor:
                CATEGORY_COLORS[category] ?? AppColors.categoryFallback,
            }}
          />
        )}
        <View
          className="absolute bottom-0 left-0 right-0 h-[55%] justify-end px-3 pb-[10px]"
          style={{ backgroundColor: AppColors.cardOverlay }}
        >
          <Text
            style={{
              color: AppColors.white,
              fontFamily: FontFamilies.bold,
              fontSize: FontSizes.base,
            }}
          >
            {category}
          </Text>
        </View>
      </View>
    </Pressable>
  );
}
