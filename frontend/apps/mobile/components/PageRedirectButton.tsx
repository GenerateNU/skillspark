import { AppColors } from "@/constants/theme";
import { TouchableOpacity, Text } from "react-native";

interface PageRedirectButtonProps {
  label: string;
  onPress: () => void;
}

export const PageRedirectButton = ({
  label,
  onPress,
}: PageRedirectButtonProps) => {
  return (
    <TouchableOpacity
      className="rounded-lg p-[10px] w-full items-center"
      onPress={onPress}
      activeOpacity={0.5}
    >
      <Text className="text-base font-medium" style={{ color: AppColors.primaryBlue }}>
        {label}
      </Text>
    </TouchableOpacity>
  );
};
