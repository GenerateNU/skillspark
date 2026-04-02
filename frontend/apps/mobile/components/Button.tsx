import { AppColors } from "@/constants/theme";
import { TouchableOpacity, Text } from "react-native";

interface ButtonProps {
  label: string;
  onPress: () => void;
}

export const Button = ({ label, onPress }: ButtonProps) => {
  return (
    <TouchableOpacity
      className="rounded-lg p-[10px] w-full items-center"
      style={{ backgroundColor: AppColors.primaryBlue }}
      onPress={onPress}
      activeOpacity={0.5}
    >
      <Text className="text-base font-medium text-white">{label}</Text>
    </TouchableOpacity>
  );
};
