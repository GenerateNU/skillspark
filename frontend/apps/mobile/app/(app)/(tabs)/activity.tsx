import { useTranslation } from "react-i18next";
import { Text, View } from "react-native";

export default function ActivityScreen() {
    const { t: translate } = useTranslation();
  
  return (
    <View className="flex-1 items-center justify-center bg-white">
      <Text className="text-lg font-semibold text-[#11181C]">{translate("activity.activity")}</Text>
    </View>
  );
}
