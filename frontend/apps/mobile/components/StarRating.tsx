import { View, Text } from "react-native";
import { AppColors, FontFamilies } from "@/constants/theme";

export function StarRating({
  rating = 4,
  size = 13,
  filledColor = AppColors.starFilled,
}: {
  rating?: number;
  size?: number;
  filledColor?: string;
}) {
  return (
    <View style={{ flexDirection: "row", gap: 1 }}>
      {Array.from({ length: 5 }).map((_, i) => (
        <Text
          key={i}
          style={{
            fontSize: size,
            fontFamily: FontFamilies.regular,
            color: i < rating ? filledColor : AppColors.borderLight,
          }}
        >
          ★
        </Text>
      ))}
    </View>
  );
}
