import { AppColors, Colors } from "@/constants/theme";
import { View } from "react-native";
import { IconSymbol } from "./ui/icon-symbol";

interface NoProfilePicProps {
	width: number;
	height: number;
}

export const NoProfilePic = ({ width, height }: NoProfilePicProps) => {
	return (
		<View
			className="rounded-full border items-center justify-center overflow-hidden"
			style={{
				borderColor: Colors.light.borderColor,
				width: width,
				height: height,
			}}
		>
			<IconSymbol
				name="square.and.arrow.up"
				size={width / 1.5}
				color={AppColors.mutedText}
			/>
		</View>
	);
};
