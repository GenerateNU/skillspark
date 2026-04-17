import { useState } from "react";
import { View, TouchableOpacity, Text } from "react-native";
import { useRouter } from "expo-router";
import { useTranslation } from "react-i18next";
import { SvgXml } from "react-native-svg";
import { ThemedText } from "./themed-text";
import { IconSymbol } from "./ui/icon-symbol";

interface ErrorScreenProps {
  message?: string;
}

const superSadSvg = `<svg width="192" height="192" viewBox="0 0 192 192" fill="none" xmlns="http://www.w3.org/2000/svg">
<rect width="191.7" height="191.7" rx="95.85" fill="#F09B9E"/>
<path d="M66.7505 78.7505C78.8453 78.7507 88.6498 88.5551 88.6499 100.65V111.879C88.6496 112.802 87.9018 113.55 86.979 113.55H46.521C45.5983 113.55 44.8504 112.802 44.8501 111.879V100.65C44.8502 88.555 54.6555 78.7505 66.7505 78.7505ZM33.6753 53.3999C36.0362 53.4 37.9497 55.3143 37.9497 57.6753C37.9496 60.0362 36.0361 61.9496 33.6753 61.9497C31.3143 61.9497 29.4 60.0362 29.3999 57.6753C29.3999 55.3143 31.3143 53.3999 33.6753 53.3999ZM99.8247 53.3999C102.186 53.3999 104.1 55.3143 104.1 57.6753C104.1 60.0362 102.186 61.9497 99.8247 61.9497C97.4639 61.9495 95.5504 60.0361 95.5503 57.6753C95.5503 55.3144 97.4638 53.4001 99.8247 53.3999Z" fill="black" stroke="black" stroke-width="3.6"/>
</svg>`;

export const ErrorScreen = ({ message }: ErrorScreenProps) => {
  const router = useRouter();
  const { t: translate } = useTranslation();
  const [expanded, setExpanded] = useState(false);

  return (
    <View className="flex-1 items-center justify-center bg-white">
      <ThemedText type="title" className="text-black text-center mb-2">
        {translate("errorScreen.title")}
      </ThemedText>
      <ThemedText className="text-black font-semibold text-center mb-12">
        {translate("errorScreen.subtitle")}
      </ThemedText>

      <SvgXml xml={superSadSvg} width={220} height={220} />

      <View className="mt-12 w-full px-6">
        <TouchableOpacity
          onPress={() =>
            router.canGoBack() ? router.back() : router.replace("/")
          }
          className="bg-black rounded-full py-4 items-center"
        >
          <Text className="text-white text-base font-medium">
            {translate("errorScreen.backToSafety")}
          </Text>
        </TouchableOpacity>
      </View>

      {message && (
        <View className="mt-6 w-full px-6">
          <TouchableOpacity
            onPress={() => setExpanded((prev) => !prev)}
            className="flex-row items-center justify-center gap-1"
          >
            <Text className="text-gray-500 text-sm">
              {expanded ? translate("errorScreen.hideDetails") : translate("errorScreen.showDetails")}
            </Text>
            <IconSymbol
              name={expanded ? "chevron.up" : "chevron.down"}
              size={16}
              color="#6b7280"
            />
          </TouchableOpacity>
          {expanded && (
            <View className="mt-3 bg-gray-100 rounded-xl p-4">
              <Text className="text-gray-700 text-sm ">{message}</Text>
            </View>
          )}
        </View>
      )}
    </View>
  );
};
