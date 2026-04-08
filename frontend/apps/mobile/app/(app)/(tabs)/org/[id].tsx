import { Image } from "expo-image";
import {
  ActivityIndicator,
  ScrollView,
  Text,
  TouchableOpacity,
  View,
} from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";
import { useLocalSearchParams, useRouter } from "expo-router";
import { useGetLocationById, useGetOrganization } from "@skillspark/api-client";
import type { Location, Organization } from "@skillspark/api-client";
import { AboutPage } from "@/components/AboutPage";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { ThemedText } from "@/components/themed-text";
import { AppColors } from "@/constants/theme";
import { useThemeColor } from "@/hooks/use-theme-color";
import { useTranslation } from "react-i18next";

function OrgDetail({
  org,
  location,
}: {
  org: Organization;
  location?: Location;
}) {
  const router = useRouter();
  const { t: translate } = useTranslation();
  const backgroundColor = useThemeColor({}, "background");
  const borderColor = useThemeColor({}, "borderColor");

  return (
    <SafeAreaView
      className="flex-1"
      style={{ backgroundColor }}
      edges={["top", "bottom"]}
    >
      <View
        className="flex-row items-center border-b px-4 pb-2.5 pt-3"
        style={{ backgroundColor, borderBottomColor: borderColor }}
      >
        <TouchableOpacity
          onPress={() => router.back()}
          activeOpacity={0.7}
          className="h-8 w-8 items-center justify-center"
        >
          <IconSymbol
            name="chevron.left"
            size={28}
            color={AppColors.primaryText}
          />
        </TouchableOpacity>
        <ThemedText
          className="flex-1 text-center text-[16px] font-nunito-bold"
          numberOfLines={1}
        >
          {org.name}
        </ThemedText>
        <View className="w-8" />
      </View>
      <ScrollView showsVerticalScrollIndicator={false} className="pb-6">
        <View
          className="h-[160px]"
          style={{ backgroundColor: AppColors.imagePlaceholder }}
        >
          {org.presigned_url ? (
            <Image
              source={{ uri: org.presigned_url }}
              className="h-full w-full"
              contentFit="cover"
            />
          ) : (
            <View className="flex-1 items-center justify-center">
              <IconSymbol name="photo" size={48} color={AppColors.mutedText} />
            </View>
          )}
        </View>
        <View className="px-4 pb-2 pt-4">
          <View className="flex-row items-start justify-between">
            <View className="mr-3 flex-1">
              <Text className="mb-0.5 text-[24px] font-nunito-bold">
                {org.name}
              </Text>
              <Text
                className="mb-[5px] text-[14px] font-nunito"
                style={{ color: AppColors.mutedText }}
              >
                {location && (
                  <Text
                    className="mb-[5px] text-[14px] font-nunito"
                    style={{ color: AppColors.mutedText }}
                  >
                    {location.district}, {location.province}
                  </Text>
                )}
              </Text>
              <View className="flex-row items-center gap-1.5">
                <Text className="text-[14px]">🔥</Text>
                <Text
                  className="text-[14px] font-nunito"
                  style={{ color: AppColors.mutedText }}
                >
                  100+ {translate("org.bookingsThisWeek")}
                </Text>
              </View>
            </View>
            <View className="mt-1 flex-col items-center gap-2">
              <TouchableOpacity
                activeOpacity={0.7}
                className="h-9 w-9 items-center justify-center rounded-full border-2"
                style={{ borderColor: AppColors.borderLight }}
              >
                <IconSymbol
                  name="square.and.arrow.up"
                  size={18}
                  color={AppColors.secondaryText}
                />
              </TouchableOpacity>
            </View>
          </View>
        </View>
        
        <TouchableOpacity
          activeOpacity={0.7}
          className="mb-3 mt-2 mx-[15px] rounded-[32px] bg-white p-5 shadow"
          onPress={() => router.push({
            pathname: `/(app)/(tabs)/org/[id]/reviews`,
            params: { id: org.id },
          })}
        >
          <View className="items-center">
            <Text className="mb-1 text-[22px] font-nunito-bold">
              {translate("org.reviews")}
            </Text>
            <Text className="text-[42px] font-nunito-bold leading-[46px]">
              4.5
            </Text>
            <Image
              source={require("@/assets/images/faces.png")}
              className="my-3 h-10 w-[140px]"
            />
            <Text
              className="mt-1.5 text-[14px] font-nunito"
              style={{ color: AppColors.subtleText }}
            >
              (140)
            </Text>
          </View>
        </View>
        {org.links.length > 0 && (
          <View className="mx-4 mb-3">
            <AboutPage description="" links={org.links} />
          </View>
        )}
        <View className="px-4 pb-2.5 pt-1">
          <TouchableOpacity
            activeOpacity={0.85}
            className="w-full items-center rounded-full py-4"
            style={{ backgroundColor: AppColors.checkboxSelected }}
            onPress={() => {}}
          >
            <Text className="text-[17px] font-nunito-bold text-white">
              {translate("org.seeSchedule")}
            </Text>
          </TouchableOpacity>
        </View>
      </ScrollView>
    </SafeAreaView>
  );
}

export default function OrgScreen() {
  const { id } = useLocalSearchParams<{ id: string }>();
  const { data: response, isLoading, error } = useGetOrganization(id);
  const { t: translate } = useTranslation();
  const { data: locationResponse } = useGetLocationById(
    response?.status === 200 ? (response.data.location_id ?? "") : "",
    {
      query: {
        enabled: response?.status === 200 && !!response.data.location_id,
      },
    },
  );

  if (isLoading) {
    return (
      <View className="flex-1 items-center justify-center">
        <ActivityIndicator size="large" />
      </View>
    );
  }

  if (error || !response || response.status !== 200) {
    return (
      <View className="flex-1 items-center justify-center p-6">
        <Text className="text-base font-semibold text-red-500">
          {translate("org.notFound")}
        </Text>
      </View>
    );
  }

  return (
    <OrgDetail
      org={response.data}
      location={
        locationResponse?.status === 200 ? locationResponse.data : undefined
      }
    />
  );
}
