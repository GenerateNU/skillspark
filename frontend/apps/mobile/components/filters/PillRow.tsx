import { AppColors, FontFamilies, FontSizes } from "@/constants/theme";
import { Text, TouchableOpacity, View } from "react-native";

export function PillRow({
  options,
  activeIndex,
  onSelect,
}: {
  options: string[];
  activeIndex: number;
  onSelect: (idx: number) => void;
}) {
  return (
    <View className="flex-row flex-wrap gap-2">
      {options.map((label, idx) => {
        const active = idx === activeIndex;
        return (
          <TouchableOpacity
            key={label}
            onPress={() => onSelect(idx)}
            className="px-4 py-[7px] rounded-full border"
            style={{
              backgroundColor: active ? AppColors.primaryText : "transparent",
              borderColor: active
                ? AppColors.primaryText
                : AppColors.borderLight,
            }}
            activeOpacity={0.7}
          >
            <Text
              style={{
                fontFamily: FontFamilies.semiBold,
                fontSize: FontSizes.base,
                color: active ? AppColors.white : AppColors.secondaryText,
              }}
            >
              {label}
            </Text>
          </TouchableOpacity>
        );
      })}
    </View>
  );
}
