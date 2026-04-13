import { AppColors } from "@/constants/theme";
import MaterialIcons from "@expo/vector-icons/MaterialIcons";
import { Image } from "expo-image";
import { Text, View } from "react-native";

interface Props {
  eventName: string;
  eventLocation: string;
  eventImageUrl?: string;
}

export function EventPreview({
  eventName,
  eventLocation,
  eventImageUrl,
}: Props) {
  return (
    <View className="flex-row items-center gap-3 mb-6">
      <View
        className="w-12 h-12 rounded-xl overflow-hidden"
        style={{ backgroundColor: AppColors.imagePlaceholder }}
      >
        {!!eventImageUrl && (
          <Image
            source={{ uri: eventImageUrl }}
            style={{ width: 48, height: 48 }}
            contentFit="cover"
          />
        )}
      </View>
      <View className="flex-1">
        <Text
          className="text-base font-nunito-bold"
          style={{ color: AppColors.primaryText }}
        >
          {eventName}
        </Text>
        <View className="flex-row items-center gap-1 mt-0.5">
          <MaterialIcons
            name="location-on"
            size={12}
            color={AppColors.subtleText}
          />
          <Text
            className="text-xs"
            style={{ color: AppColors.subtleText }}
            numberOfLines={1}
          >
            {eventLocation}
          </Text>
        </View>
      </View>
    </View>
  );
}
