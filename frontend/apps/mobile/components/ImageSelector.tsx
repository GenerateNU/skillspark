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
import { IconSymbol } from "./ui/icon-symbol";

interface ImageSelectorProps extends TouchableOpacityProps {
	setImage: React.Dispatch<React.SetStateAction<string | undefined>>;
	image: string | undefined;
}

export const ImageSelector = ({
	setImage,
	image,
	...props
}: ImageSelectorProps) => {
	const colorScheme = useColorScheme();
	const theme = Colors[colorScheme ?? "light"];

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
				className="w-40 h-40 rounded-full border items-center justify-center overflow-hidden"
				style={{ borderColor: theme.borderColor }}
			>
				{image && (
					<Image
						source={{ uri: image }}
						className="w-full h-full"
						resizeMode="cover"
					/>
				)}
				{!image && (
					<IconSymbol name="person" size={60} color={AppColors.mutedText} />
				)}
			</View>
			<ThemedText className="text-sm" style={{ color: AppColors.mutedText }}>
				Change Image
			</ThemedText>
		</TouchableOpacity>
	);
};
