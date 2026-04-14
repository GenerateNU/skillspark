import { AppColors } from "@/constants/theme";
import { Text, TouchableOpacity } from "react-native";

interface Props {
  label: string;
  selected: boolean;
  onPress: () => void;
}

export function CategoryPill({ label, selected, onPress }: Props) {
  return (
    <TouchableOpacity
      onPress={onPress}
      className="px-4 py-2 rounded-full border"
      style={{
        borderColor: selected ? AppColors.primaryText : AppColors.borderLight,
        backgroundColor: selected ? AppColors.primaryText : "transparent",
      }}
    >
      <Text
        className="text-sm"
        style={{ color: selected ? "#fff" : AppColors.secondaryText }}
      >
        {label}
      </Text>
    </TouchableOpacity>
  );
}
