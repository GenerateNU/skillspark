import { AppColors, FontFamilies, FontSizes } from "@/constants/theme";
import { Child } from "@skillspark/api-client";
import { Text, TouchableOpacity, View } from "react-native";

export function ChildAvatar({
  child,
  selected,
  onPress,
}: {
  child: Child;
  selected: boolean;
  onPress: () => void;
}) {
  const initials = child.name
    .split(" ")
    .map((p) => p[0])
    .join("")
    .toUpperCase()
    .slice(0, 2);

  return (
    <TouchableOpacity
      onPress={onPress}
      activeOpacity={0.7}
      className="items-center gap-1.5"
    >
      <View
        className="w-[52px] h-[52px] rounded-full items-center justify-center"
        style={{
          backgroundColor: selected
            ? AppColors.checkboxSelected
            : AppColors.surfaceGray,
          borderWidth: selected ? 2.5 : 2,
          borderColor: selected
            ? AppColors.checkboxSelected
            : AppColors.borderLight,
          borderStyle: selected ? "solid" : "dashed",
        }}
      >
        <Text
          style={{
            fontFamily: FontFamilies.bold,
            fontSize: FontSizes.base,
            color: selected ? "#fff" : AppColors.mutedText,
          }}
        >
          {initials}
        </Text>
      </View>
      <Text
        style={{
          fontFamily: FontFamilies.regular,
          fontSize: FontSizes.xs,
          color: AppColors.mutedText,
        }}
        numberOfLines={1}
      >
        {child.name.split(" ")[0]}
      </Text>
    </TouchableOpacity>
  );
}
