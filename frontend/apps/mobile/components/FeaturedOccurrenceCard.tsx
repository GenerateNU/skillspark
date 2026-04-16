import { Image } from "expo-image";
import { Pressable, Text, View } from "react-native";
import { useRouter } from "expo-router";
import { useTranslation } from "react-i18next";
import type { EventOccurrence, Child } from "@skillspark/api-client";
import { AppColors } from "@/constants/theme";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { ChildAvatar } from "@/components/ChildAvatar";

type Props = {
  occurrence: EventOccurrence;
  orgName: string | null;
  ageLabel: string | null;
  address: string | null;
  childName: string;
  child: Child | undefined;
};

export function FeaturedOccurrenceCard({ occurrence, orgName, ageLabel, address, childName, child }: Props) {
  const router = useRouter();
  const { t: translate } = useTranslation();

  return (
    <View className="mx-5 mb-5 rounded-3xl p-4" style={{ backgroundColor: AppColors.bluePastel }}>
      <View className="flex-row items-center gap-2 mb-3">
        <Text className="font-nunito-bold text-base text-[#111]">
          {translate("dashboard.trendingRightNow")}
        </Text>
      </View>
      <Pressable
        onPress={() => router.push(`/event/${occurrence.event.id}`)}
        className="bg-white rounded-2xl p-3 flex-row items-center gap-3"
      >
        {occurrence.event.presigned_url ? (
          <Image
            source={{ uri: occurrence.event.presigned_url }}
            className="w-[110px] h-[80px] rounded-xl"
            contentFit="cover"
          />
        ) : (
          <View className="w-[110px] h-[80px] rounded-xl bg-[#D9D9D9]" />
        )}
        <View className="flex-1 gap-0.5">
          <Text className="font-nunito-bold text-base text-[#111]" numberOfLines={1}>
            {occurrence.event.title}
          </Text>
          {!!orgName && (
            <Text className="font-nunito text-xs text-gray-500">{orgName}</Text>
          )}
          {!!ageLabel && (
            <Text className="font-nunito text-xs text-gray-500">{ageLabel}</Text>
          )}
          {!!address && (
            <Text className="font-nunito text-xs text-gray-500" numberOfLines={2}>
              {address}
            </Text>
          )}
        </View>
        <View className="items-end justify-between self-stretch gap-2">
          <View className="flex-row items-center bg-gray-100 rounded-full px-2 py-1 gap-1">
            <ChildAvatar
              name={childName}
              avatarFace={child?.avatar_face}
              avatarBackground={child?.avatar_background}
              size={20}
            />
            <Text className="font-nunito-semibold text-xs text-[#111]">
              {childName}
            </Text>
          </View>
          <View
            className="w-9 h-9 rounded-full items-center justify-center"
            style={{ backgroundColor: AppColors.slateBlue }}
          >
            <IconSymbol name="chevron.right" size={14} color="white" />
          </View>
        </View>
      </Pressable>
    </View>
  );
}
