import { AppColors } from "@/constants/theme";
import { TouchableOpacity, Text } from "react-native";
import { ThemedText } from "./themed-text";

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
      <ThemedText className="text-base font-nunito-bold underline">
        {label}
      </ThemedText>
    </TouchableOpacity>
  );
};
