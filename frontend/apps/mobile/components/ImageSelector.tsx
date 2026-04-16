import {
  Image,
  TouchableOpacity,
  TouchableOpacityProps,
  View,
} from "react-native";
import * as ImagePicker from "expo-image-picker";
import { AppColors, Colors } from "@/constants/theme";
import { ThemedText } from "./themed-text";
import { NoProfilePic } from "./NoProfilePic";
import { useTranslation } from "react-i18next";
import { IconSymbol } from "@/components/ui/icon-symbol";

interface ImageSelectorProps extends TouchableOpacityProps {
  setImage: React.Dispatch<React.SetStateAction<string | undefined>>;
  image: string | undefined;
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
  const theme = Colors.light;
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
      <View style={{ position: "relative" }}>
        <View
          className="rounded-full border items-center justify-center overflow-hidden"
          style={{
            borderColor: theme.borderColor,
            width: width,
            height: height,
          }}
        >
          {image && (
            <Image
              source={{ uri: image }}
              style={{ width: "100%", height: "100%" }}
              resizeMode="cover"
            />
          )}
          {!image && <NoProfilePic width={width} height={height} />}
        </View>
        <View
          style={{
            position: "absolute",
            top: 2,
            right: 2,
            width: 32,
            height: 32,
            borderRadius: 16,
            backgroundColor: theme.background,
            alignItems: "center",
            justifyContent: "center",
            borderWidth: 1,
            borderColor: theme.borderColor,
          }}
        >
          <IconSymbol name="pencil" size={20} color={theme.text} />
        </View>
      </View>
      <ThemedText className="text-sm" style={{ color: AppColors.mutedText }}>
        {translate("editProfile.changeImage")}
      </ThemedText>
    </TouchableOpacity>
  );
};
