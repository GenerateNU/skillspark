import { AppColors } from "@/constants/theme";
import { TouchableOpacity, Text } from "react-native";

interface SubmitButtonProps {
  label: string;
  onPress: () => void;
}

export const SubmitButton = ({ label, onPress }: SubmitButtonProps) => {
  return (
    <TouchableOpacity
      className="rounded-lg p-[10px] w-full items-center" style={{ backgroundColor: AppColors.primaryBlue }}
      onPress={onPress}
      activeOpacity={0.5}
    >
      <Text className="text-base font-medium text-white">
        {label}
      </Text>
    </TouchableOpacity>
  );
};
