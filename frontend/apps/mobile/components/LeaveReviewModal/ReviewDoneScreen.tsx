import { RATING_OPTIONS } from "@/constants/ratings";
import { AppColors } from "@/constants/theme";
import { Image } from "expo-image";
import { Text, TouchableOpacity, View } from "react-native";
import LogoBgWrapper from "../LogoBgWrapper";

interface Props {
  rating: number | null;
  onClose: () => void;
  translate: (key: string) => string;
}

export function ReviewDoneScreen({ rating, onClose, translate }: Props) {
  const submittedEmoji = RATING_OPTIONS.find((r) => r.rating === rating);

  return (
    <View className="flex-1 items-center justify-center px-8">
      <LogoBgWrapper>
        <Text
          className="text-3xl font-nunito-bold mb-6"
          style={{ color: AppColors.primaryText }}
        >
          {translate("review.thankYou")}
        </Text>
        {submittedEmoji && (
          <Image
            source={submittedEmoji.image}
            style={{ width: 80, height: 80, marginBottom: 24 }}
          />
        )}
        <Text
          className="text-base text-center mb-10"
          style={{ color: AppColors.secondaryText, lineHeight: 24 }}
        >
          {translate("review.thankYouMessage")}
        </Text>
        <TouchableOpacity
          onPress={onClose}
          className="w-full py-4 rounded-2xl items-center"
          style={{ backgroundColor: AppColors.primaryText }}
        >
          <Text className="text-base font-nunito-bold text-white">
            {translate("review.close")}
          </Text>
        </TouchableOpacity>
      </LogoBgWrapper>
    </View>
  );
}
