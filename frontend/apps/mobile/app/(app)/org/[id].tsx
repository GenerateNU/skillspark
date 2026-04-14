import { Image } from "expo-image";
import {
  ActivityIndicator,
  Pressable,
  ScrollView,
  Text,
  TouchableOpacity,
  View,
} from "react-native";
import { useState } from "react";
import { SafeAreaView } from "react-native-safe-area-context";
import { useLocalSearchParams, useRouter } from "expo-router";
import { useGetLocationById, useGetOrganization } from "@skillspark/api-client";
import type { Location, Organization } from "@skillspark/api-client";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { ThemedText } from "@/components/themed-text";
import { AppColors } from "@/constants/theme";
import { useThemeColor } from "@/hooks/use-theme-color";
import { useOrgLinks } from "@/hooks/useOrgLinks";
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
  const { openLink, hasLinks } = useOrgLinks(org.links ?? []);
  const [aboutExpanded, setAboutExpanded] = useState(false);
  const [aboutTruncated, setAboutTruncated] = useState(false);

  const cardStyle = {
    shadowColor: "#000",
    shadowOpacity: 0.08,
    shadowRadius: 12,
    shadowOffset: { width: 0, height: 2 },
    elevation: 3,
  };

  return (
    <SafeAreaView
      className="flex-1"
      style={{ backgroundColor }}
      edges={["top", "bottom"]}
    >
      {/* Header */}
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

      <ScrollView showsVerticalScrollIndicator={false}>
        {/* Hero image */}
        <View
          className="h-[200px]"
          style={{ backgroundColor: AppColors.imagePlaceholder }}
        >
          {org.presigned_url ? (
            <Image
              source={{ uri: org.presigned_url }}
              style={{ width: "100%", height: "100%" }}
              contentFit="cover"
            />
          ) : (
            <View className="flex-1 items-center justify-center">
              <IconSymbol name="photo" size={48} color={AppColors.mutedText} />
            </View>
          )}
        </View>

        {/* Org info */}
        <View className="px-4 pb-4 pt-4">
          <View className="flex-row items-start justify-between">
            <View className="mr-3 flex-1">
              <Text className="mb-1 text-[24px] font-nunito-bold">
                {org.name}
              </Text>
              {location && (
                <Text
                  className="mb-1 text-[14px] font-nunito"
                  style={{ color: AppColors.mutedText }}
                >
                  {location.district}, {location.province}
                </Text>
              )}
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
            <TouchableOpacity
              activeOpacity={0.7}
              className="mt-1 h-9 w-9 items-center justify-center rounded-full border-2"
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

        {/* Divider */}
        <View
          className="mx-4 mb-5 border-b border-dashed"
          style={{ borderColor: AppColors.divider }}
        />

        {/* About card */}
        <View className="mx-4 mb-4 rounded-2xl bg-white p-5" style={cardStyle}>
          <Text className="mb-2.5 text-[18px] font-nunito-bold">
            {translate("org.about")}
          </Text>
          <Text
            numberOfLines={aboutExpanded ? undefined : 4}
            onTextLayout={(e) => {
              if (!aboutExpanded)
                setAboutTruncated(e.nativeEvent.lines.length >= 4);
            }}
            className={`text-sm leading-[22px] ${aboutTruncated ? "mb-1" : "mb-4"}`}
            style={{ color: AppColors.secondaryText }}
          >
            {org.about ?? ""}
          </Text>
          {aboutTruncated && (
            <Pressable
              onPress={() => setAboutExpanded((prev) => !prev)}
              className="mb-4"
            >
              <Text
                className="text-[13px] font-semibold"
                style={{ color: AppColors.primaryText }}
              >
                {aboutExpanded
                  ? translate("event.seeLess")
                  : translate("event.seeMore")}
              </Text>
            </Pressable>
          )}
          {hasLinks && (
            <View className="flex-row flex-wrap gap-2.5">
              {(org.links ?? []).map((link, index) => (
                <Pressable
                  key={index}
                  onPress={() => openLink(link.href)}
                  className="rounded-full px-5 py-2.5 items-center"
                  style={{ backgroundColor: AppColors.borderLight }}
                >
                  <Text
                    className="text-[13px] font-semibold"
                    style={{ color: AppColors.primaryText }}
                  >
                    {link.label}
                  </Text>
                </Pressable>
              ))}
            </View>
          )}
        </View>

        {/* Rating card */}
        <TouchableOpacity
          activeOpacity={0.8}
          onPress={() => router.push(`/org/${org.id}/reviews`)}
        >
          <View
            className="mx-4 mb-4 rounded-2xl bg-white p-5"
            style={cardStyle}
          >
            <Text className="mb-3 text-[18px] font-nunito-bold">
              {translate("org.reviews")}
            </Text>
            <View className="flex-row items-center gap-4">
              <Text className="text-[42px] font-nunito-bold leading-[46px]">
                4.5
              </Text>
              <View className="items-start">
                <Image
                  source={require("@/assets/images/faces.png")}
                  className="h-10 w-[140px]"
                />
                <Text
                  className="mt-1 text-[13px] font-nunito"
                  style={{ color: AppColors.subtleText }}
                >
                  (140)
                </Text>
              </View>
            </View>
          </View>
        </TouchableOpacity>

        {/* See Schedule CTA */}
        <View className="px-4 pb-6 pt-1">
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
