import { View, Text } from "react-native";
import { AppColors } from "@/constants/theme";

export function StarRating({
  rating = 4,
  size = 13,
}: {
  rating?: number;
  size?: number;
}) {
  return (
    <View style={{ flexDirection: "row", gap: 1 }}>
      {Array.from({ length: 5 }).map((_, i) => (
        <Text
          key={i}
          style={{
            fontSize: size,
            color: i < rating ? AppColors.starFilled : AppColors.borderLight,
          }}
        >
          ★
        </Text>
      ))}
    </View>
  );
}
