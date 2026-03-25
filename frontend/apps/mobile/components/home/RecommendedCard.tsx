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
      style={{ marginRight: 20, alignItems: "center", width: 76 }}
    >
      <View style={{ width: 88, height: 88, borderRadius: 18, overflow: "hidden" }}>
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
      <Text
        style={{ fontSize: FontSizes.xs, color: AppColors.mutedText, marginTop: 5, textAlign: "center", fontFamily: "NunitoSans_400Regular" }}
        numberOfLines={1}
      >
        {childName}
      </Text>
    </Pressable>
  );
}
