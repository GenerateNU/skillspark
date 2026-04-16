import { Image } from "expo-image";
import { View } from "react-native";
import type { StyleProp, ViewStyle } from "react-native";
import { AppColors } from "@/constants/theme";

interface EventImageProps {
  uri?: string | null;
  style?: StyleProp<ViewStyle>;
  contentFit?: "cover" | "contain";
}

/** Renders an event image from a presigned URL, falling back to a placeholder background. */
export function EventImage({ uri, style, contentFit = "cover" }: EventImageProps) {
  if (uri) {
    return <Image source={{ uri }} style={style} contentFit={contentFit} />;
  }
  return <View style={[{ backgroundColor: AppColors.imagePlaceholder }, style]} />;
}
