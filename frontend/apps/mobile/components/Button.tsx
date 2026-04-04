import { AppColors } from "@/constants/theme";
import { TouchableOpacity, Text } from "react-native";

interface ButtonProps {
	label: string;
	onPress: () => void;
	disabled: boolean;
}

export const Button = ({ label, onPress, disabled }: ButtonProps) => {
	return (
		<TouchableOpacity
			className="rounded-lg p-[10px] w-full items-center"
			style={{
				backgroundColor: AppColors.primaryBlue,
				opacity: disabled ? 0.5 : 1,
			}}
			onPress={onPress}
			activeOpacity={0.5}
			disabled={disabled}
		>
			<Text className="text-base font-medium text-white">{label}</Text>
		</TouchableOpacity>
	);
};
