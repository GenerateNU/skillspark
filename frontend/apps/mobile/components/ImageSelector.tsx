import {
	Image,
	TouchableOpacity,
	TouchableOpacityProps,
	View,
} from "react-native";
import * as ImagePicker from "expo-image-picker";
import { useColorScheme } from "@/hooks/use-color-scheme.web";
import { AppColors, Colors } from "@/constants/theme";
import { ThemedText } from "./themed-text";
import { NoProfilePic } from "./NoProfilePic";
import { useTranslation } from "react-i18next";

interface ImageSelectorProps extends TouchableOpacityProps {
	setImage: React.Dispatch<React.SetStateAction<string | null>>;
	image: string | null;
	width: number;
	height: number;
}

export const ImageSelector = ({
	setImage,
	image,
	width,
	height,
	...props
}: ImageSelectorProps) => {
	const colorScheme = useColorScheme();
	const theme = Colors[colorScheme ?? "light"];
	const { t: translate } = useTranslation();

	const pickImage = async () => {
		const result = await ImagePicker.launchImageLibraryAsync({
			mediaTypes: ["images"],
			allowsEditing: true,
			aspect: [3, 3],
		});

		if (!result.canceled) {
			setImage(result.assets[0].uri);
		}
	};

	return (
		<TouchableOpacity onPress={pickImage} {...props}>
			<View
				className="rounded-full border items-center justify-center overflow-hidden"
				style={{ borderColor: theme.borderColor, width: width, height: height }}
			>
				{image && (
					<Image
						source={{ uri: image }}
						className="w-full h-full"
						resizeMode="cover"
					/>
				)}
				{!image && <NoProfilePic width={width} height={height} />}
			</View>
			<ThemedText className="text-sm" style={{ color: AppColors.mutedText }}>
				{translate("editProfile.changeImage")}
			</ThemedText>
		</TouchableOpacity>
	);
};
