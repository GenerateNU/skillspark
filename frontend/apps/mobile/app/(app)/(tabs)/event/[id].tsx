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
import { useSafeAreaInsets } from "react-native-safe-area-context";
import { useLocalSearchParams, useRouter } from "expo-router";
import { useGetEventOccurrencesById } from "@skillspark/api-client";
import type { EventOccurrence } from "@skillspark/api-client";
import MaterialIcons from "@expo/vector-icons/MaterialIcons";
import { AppColors } from "@/constants/theme";
import { StarRating } from "@/components/StarRating";
import { BookmarkButton } from "@/components/BookmarkButton";
import { formatDuration } from "@/utils/format";
import { useTranslation } from "react-i18next";
import { ListItem } from "@/components/ListItem";
import { useOrgLinks } from "@/hooks/useOrgLinks";

function formatAddress(occurrence: EventOccurrence) {
  const loc = occurrence.location;
  const parts = [loc.address_line1, loc.district].filter(Boolean);
  return parts.join(", ") || "Location";
}

function EventOccurrenceDetail({
  occurrence,
}: {
  occurrence: EventOccurrence;
}) {
  const router = useRouter();
  const insets = useSafeAreaInsets();
  const [descriptionExpanded, setDescriptionExpanded] = useState(false);
  const [descriptionTruncated, setDescriptionTruncated] = useState(false);
  const { t: translate } = useTranslation();
  const { openLink, hasLinks } = useOrgLinks(occurrence.org_links ?? []);
  const duration = formatDuration(occurrence.start_time, occurrence.end_time, {
    hr: translate("event.hr"),
    min: translate("event.min"),
  });
  const address = formatAddress(occurrence);

  return (
    <View className="flex-1 bg-[#F4F6F8]">
      <View
        className="flex-1 mx-3.5 rounded-[32px] overflow-hidden bg-white elevation-8"
        style={{
          marginTop: insets.top + 8,
          marginBottom: insets.bottom - 20,
          shadowColor: "#000",
          shadowOpacity: 0.12,
          shadowRadius: 20,
        }}
      >
        <ScrollView
          showsVerticalScrollIndicator={false}
          bounces={false}
          contentContainerStyle={{ flexGrow: 1 }}
        >
          {/* Hero image */}
          <View className="h-[250px] bg-[#1a1a1a]">
            {occurrence.event.presigned_url ? (
              <Image
                source={{ uri: occurrence.event.presigned_url }}
                className="w-full h-full"
                contentFit="cover"
              />
            ) : (
              <View className="flex-1 bg-[#C5C5C5]" />
            )}
            <TouchableOpacity
              onPress={() => router.navigate("/")}
              activeOpacity={0.7}
              className="absolute top-4 left-4 z-10 flex-row items-center bg-white rounded-full px-4 py-2.5 elevation-10"
              style={{
                shadowColor: "#000",
                shadowOpacity: 0.15,
                shadowRadius: 8,
              }}
            >
              <MaterialIcons
                name="chevron-left"
                size={20}
                color={AppColors.primaryText}
              />
              <Text
                className="text-[15px] font-medium"
                style={{ color: AppColors.primaryText }}
              >
                {translate("event.back")}
              </Text>
            </TouchableOpacity>
          </View>

          {/* Content card */}
          <View
            className="bg-white rounded-t-[28px] -mt-7 px-[22px] pb-6 elevation-2"
            style={{
              shadowColor: "#000",
              shadowOpacity: 0.06,
              shadowRadius: 12,
            }}
          >
            {/* Drag handle */}
            <View
              className="w-[38px] h-1 rounded-sm self-center mt-3 mb-3.5"
              style={{ backgroundColor: AppColors.borderLight }}
            />
            <View className="absolute -top-[5px] left-2.5">
              <BookmarkButton eventId={occurrence.event.id} />
            </View>
            <Text
              className="text-[28px] font-bold tracking-tight mb-2"
              style={{ color: AppColors.primaryText }}
            >
              {occurrence.event.title}
            </Text>
            <View className="flex-row items-center gap-1.5 mb-3.5">
              <MaterialIcons
                name="location-on"
                size={16}
                color={AppColors.primaryText}
              />
              <Text
                className="text-[13px] flex-1"
                style={{ color: AppColors.secondaryText }}
                numberOfLines={1}
              >
                {address}
              </Text>
              <StarRating size={17} />
              <ListItem
                label={translate("review.title")}
                isLast
                onPress={() =>
                  router.push(`/event/review?id=${occurrence.event.id}`)
                }
              />
            </View>
            <Text
              numberOfLines={descriptionExpanded ? undefined : 5}
              onTextLayout={(e) => {
                if (!descriptionExpanded) {
                  setDescriptionTruncated(e.nativeEvent.lines.length >= 5);
                }
              }}
              className={`text-sm leading-[22px] ${descriptionTruncated ? "mb-1" : "mb-[18px]"}`}
              style={{ color: AppColors.secondaryText }}
            >
              {occurrence.event.description}
            </Text>
            {descriptionTruncated && (
              <Pressable
                onPress={() => setDescriptionExpanded((prev) => !prev)}
                className="mb-3.5"
              >
                <Text
                  className="text-[13px] font-semibold"
                  style={{ color: AppColors.primaryText }}
                >
                  {descriptionExpanded
                    ? translate("event.seeLess")
                    : translate("event.seeMore")}
                </Text>
              </Pressable>
            )}
            <View className="flex-row items-center justify-between">
              <View className="flex-row gap-2 flex-1 flex-wrap">
                {occurrence.event.category?.map((cat) => (
                  <View
                    key={cat}
                    className="border-[1.5px] rounded-full px-4 py-[7px]"
                    style={{ borderColor: AppColors.borderLight }}
                  >
                    <Text
                      className="text-[13px]"
                      style={{ color: AppColors.secondaryText }}
                    >
                      {translate(`interests.${cat}`, { defaultValue: cat })}
                    </Text>
                  </View>
                ))}
              </View>
              <View className="items-end ml-3.5">
                <Text
                  className="text-xl font-bold"
                  style={{ color: AppColors.primaryText }}
                >
                  {occurrence.price} THB
                </Text>
                <Text
                  className="text-xs"
                  style={{ color: AppColors.subtleText }}
                >
                  {translate("event.perSession")}
                </Text>
              </View>
            </View>
            {hasLinks && (
              <View className="flex-row flex-wrap gap-2.5 mt-4">
                {(occurrence.org_links ?? []).map((link, index) => (
                  <Pressable
                    key={index}
                    onPress={() => openLink(link.href)}
                    className="rounded-full px-5 py-2.5 items-center"
                    style={{
                      backgroundColor: AppColors.borderLight,
                    }}
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

          {/* Divider */}
          <View
            className="border-b border-dashed"
            style={{ borderColor: AppColors.borderLight }}
          />

          {/* Bottom section */}
          <View
            className="flex-1 bg-white px-[22px] pt-[22px] pb-7 elevation-2"
            style={{
              shadowColor: "#000",
              shadowOpacity: 0.06,
              shadowRadius: 12,
            }}
          >
            <View className="flex-row items-center justify-between mb-[22px]">
              <Text
                className="text-[30px] font-bold tracking-tight"
                style={{ color: AppColors.primaryText }}
              >
                {duration}
              </Text>
              <View className="flex-row items-center gap-[5px] bg-[#F3F4F6] rounded-full px-3 py-[7px]">
                <MaterialIcons
                  name="directions-walk"
                  size={14}
                  color={AppColors.secondaryText}
                />
                <Text
                  className="text-[13px]"
                  style={{ color: AppColors.secondaryText }}
                >
                  8 {translate("event.minWalk")}
                </Text>
                <MaterialIcons
                  name="arrow-forward"
                  size={12}
                  color={AppColors.subtleText}
                />
                <MaterialIcons
                  name="directions-bus"
                  size={16}
                  color={AppColors.secondaryText}
                />
              </View>
            </View>
            <View className="flex-row items-center justify-between">
              <View>
                <View className="flex-row items-center gap-3">
                  <View
                    className="w-4 h-4 rounded-full items-center justify-center"
                    style={{ backgroundColor: AppColors.primaryText }}
                  >
                    <View className="w-1.5 h-1.5 rounded-full bg-white" />
                  </View>
                  <Text
                    className="text-sm font-medium"
                    style={{ color: AppColors.secondaryText }}
                  >
                    {translate("event.home")}
                  </Text>
                </View>
                <View className="pl-1.5 py-0.5">
                  <Text
                    className="text-sm leading-[10px]"
                    style={{ color: AppColors.subtleText }}
                  >
                    •
                  </Text>
                  <Text
                    className="text-sm leading-[10px]"
                    style={{ color: AppColors.subtleText }}
                  >
                    •
                  </Text>
                  <Text
                    className="text-sm leading-[10px]"
                    style={{ color: AppColors.subtleText }}
                  >
                    •
                  </Text>
                </View>
                <View className="flex-row items-center gap-2.5">
                  <MaterialIcons
                    name="location-on"
                    size={16}
                    color={AppColors.secondaryText}
                  />
                  <Text
                    className="text-sm font-medium"
                    style={{ color: AppColors.secondaryText }}
                  >
                    {translate("event.location")}
                  </Text>
                </View>
              </View>
              <TouchableOpacity
                onPress={() => {}}
                activeOpacity={0.7}
                className="rounded-2xl px-[26px] py-3.5"
                style={{ backgroundColor: AppColors.primaryText }}
              >
                <Text className="text-white text-[17px] font-bold">
                  {translate("event.register")}
                </Text>
              </TouchableOpacity>
            </View>
          </View>
        </ScrollView>
      </View>
    </View>
  );
}

export default function EventOccurrenceScreen() {
  const { id } = useLocalSearchParams<{ id: string }>();
  const { data: response, isLoading, error } = useGetEventOccurrencesById(id);
  const { t: translate } = useTranslation();

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
        <Text
          className="text-base font-semibold"
          style={{ color: AppColors.danger }}
        >
          {translate("event.notFound")}
        </Text>
      </View>
    );
  }

  return <EventOccurrenceDetail occurrence={response.data} />;
}
