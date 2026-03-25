import { View } from "react-native";
import { ThemedText } from "./themed-text";
import { AppColors } from "@/constants/theme";

interface ErrorScreenProps {
  message: string;
}

export const ErrorScreen = ({ message }: ErrorScreenProps) => {
    return (
        <View className="flex-1 items-center justify-center p-4">
            <ThemedText className="font-semibold" style={{ color: AppColors.danger }}>
            ERROR: {message}
            </ThemedText>
        </View>
    )
};