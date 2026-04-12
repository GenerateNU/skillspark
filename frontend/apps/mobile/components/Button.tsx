import { TouchableOpacity, Text } from "react-native";

interface ButtonProps {
	label: string;
	onPress: () => void;
	disabled: boolean;
}

export const Button = ({ label, onPress, disabled }: ButtonProps) => {
	return (
		<TouchableOpacity
			className="rounded-full py-4 w-11/12 items-center bg-[#1b1b1b]"
			style={{
				opacity: disabled ? 0.5 : 1,
			}}
			onPress={onPress}
			activeOpacity={0.7}
			disabled={disabled}
		>
			<Text className="text-lg font-nunito-bold text-white">{label}</Text>
		</TouchableOpacity>
	);
};
