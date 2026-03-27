import { Image } from "expo-image";
import { View, Text, Pressable } from "react-native";
import { useRouter } from "expo-router";
import { type EventOccurrence } from "@skillspark/api-client";
import { AppColors, FontSizes } from "@/constants/theme";

export function RecommendedCard({ occurrence, childName }: { occurrence: EventOccurrence; childName: string }) {
  const router = useRouter();
  return (
    <Pressable
      onPress={() => router.push(`/event/${occurrence.id}`)}
      className="mr-5 items-center w-[79px]"
    >
      <View
        className="w-[79px] h-[72px] rounded-[12px] overflow-hidden"
        style={{
          shadowColor: "#000",
          shadowOpacity: 0.25,
          shadowRadius: 4,
          shadowOffset: { width: 0, height: 4 },
          elevation: 3,
        }}
      >
        {occurrence.event.presigned_url ? (
          <Image
            source={{ uri: occurrence.event.presigned_url }}
            className="w-[79px] h-[72px]"
            contentFit="cover"
          />
        ) : (
          <View className="w-[79px] h-[72px]" style={{ backgroundColor: AppColors.imagePlaceholder }} />
        )}
      </View>
      <Text
        className="mt-[5px] text-center font-nunito"
        style={{ fontSize: FontSizes.xs, color: AppColors.mutedText }}
        numberOfLines={1}
      >
        {childName}
      </Text>
    </Pressable>
  );
}
