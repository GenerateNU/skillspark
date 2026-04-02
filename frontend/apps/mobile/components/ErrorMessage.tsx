import { AppColors } from "@/constants/theme";
import { Text } from "react-native";

interface ErrorMessageProps {
  message: string;
}

export const ErrorMessage = ({ message }: ErrorMessageProps) => {
  return (
    <Text className="text-base text-center" style={{ color: AppColors.danger }}>
      {message}
    </Text>
  );
};
