import { AppColors, FontFamilies, FontSizes } from "@/constants/theme";
import { Text, TouchableOpacity, View } from "react-native";

type Props = {
  label: string;
  value: boolean;
  onToggle: () => void;
};

export function SoldOutToggle({ label, value, onToggle }: Props) {
  return (
    <TouchableOpacity
      onPress={onToggle}
      activeOpacity={0.7}
      className="flex-row items-center justify-between"
    >
      <Text
        style={{
          fontFamily: FontFamilies.bold,
          fontSize: FontSizes.lg,
          color: AppColors.primaryText,
        }}
      >
        {label}
      </Text>
      <View
        style={{
          width: 46,
          height: 26,
          borderRadius: 13,
          backgroundColor: value ? AppColors.primaryText : AppColors.borderLight,
          justifyContent: "center",
          paddingHorizontal: 3,
        }}
      >
        <View
          style={{
            width: 20,
            height: 20,
            borderRadius: 10,
            backgroundColor: AppColors.white,
            alignSelf: value ? "flex-end" : "flex-start",
          }}
        />
      </View>
    </TouchableOpacity>
  );
}
