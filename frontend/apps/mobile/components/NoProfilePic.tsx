import { AppColors, Colors } from "@/constants/theme";
import { useColorScheme } from "@/hooks/use-color-scheme";
import { View } from "react-native";
import { IconSymbol } from "./ui/icon-symbol";

interface NoProfilePicProps {
	width: number;
	height: number;
}

export const NoProfilePic = ({ width, height }: NoProfilePicProps) => {
	const colorScheme = useColorScheme();
	const theme = Colors[colorScheme ?? "light"];

	return (
		<View
			className="rounded-full border items-center justify-center overflow-hidden"
			style={{ borderColor: theme.borderColor, width: width, height: height }}
		>
			<IconSymbol
				name="person"
				size={width / 1.5}
				color={AppColors.mutedText}
			/>
		</View>
	);
};
