import { AppColors } from "@/constants/theme";
import MaterialIcons from "@expo/vector-icons/MaterialIcons";
import { Text, TouchableOpacity, View } from "react-native";

interface Props {
  onClose: () => void;
  translate: (key: string) => string;
}

export function ReviewHeader({ onClose, translate }: Props) {
  return (
    <View
      className="flex-row items-center justify-between px-5 py-4 border-b"
      style={{ borderColor: AppColors.borderLight }}
    >
      <TouchableOpacity
        onPress={onClose}
        hitSlop={{ top: 10, bottom: 10, left: 10, right: 10 }}
      >
        <MaterialIcons
          name="chevron-left"
          size={28}
          color={AppColors.primaryText}
        />
      </TouchableOpacity>
      <Text
        className="text-lg font-nunito-bold"
        style={{ color: AppColors.primaryText }}
      >
        {translate("review.leaveReview")}
      </Text>
      <View style={{ width: 28 }} />
    </View>
  );
}
