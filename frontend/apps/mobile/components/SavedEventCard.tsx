import { IconSymbol } from "@/components/ui/icon-symbol";
import { AppColors, TAG_COLORS } from "@/constants/theme";
import { Event } from "@skillspark/api-client";
import { router } from "expo-router";
import React from "react";
import { useTranslation } from "react-i18next";

import { Image } from "expo-image";
import { Pressable, Text, TouchableOpacity, View } from "react-native";

interface BookmarkIconProps {
  onPress?: () => void;
}

export function BookmarkIcon({ onPress }: BookmarkIconProps) {
  return (
    <TouchableOpacity onPress={onPress}>
      <IconSymbol name="bookmark.fill" size={24} color="#FFC107" />
    </TouchableOpacity>
  );
}

interface SavedEventCardProps {
  event: Event;
  onBookmarkPress?: (event: Event) => void;
}

export function SavedEventCard({
  event,
  onBookmarkPress,
}: SavedEventCardProps) {
  const { t: translate } = useTranslation();

  console.log(event)

  return (
    <Pressable
      onPress={() => router.push(`/event/${event.id}`)} //TODO: fix event details to be either event or occurrence based on design
      className="mx-5 mb-3 flex-row rounded-xl p-4 h-[150px] items-center shadow-sm elevation-2"
      style={{
        backgroundColor: AppColors.savedBackground,
        shadowColor: "#000",
        shadowOpacity: 0.05,
        shadowOffset: { width: 0, height: 2 },
        shadowRadius: 4,
      }}
    >
      <View
        className="w-20 h-20 rounded-full overflow-hidden mr-4 items-center justify-center"
        style={{ backgroundColor: AppColors.divider }}
      >
        {event.presigned_url && (
          <Image
            source={{ uri: event.presigned_url }}
            style={{ width: 80, height: 80 }}
            contentFit="cover"
          />
        )}
      </View>
      <View className="flex-1 justify-center">
        <View className="flex-row items-center">
          <Text className="text-base font-semibold text-[#111] shrink">
            {event.title}
          </Text>
          <View className="ml-3">
            <BookmarkIcon onPress={() => onBookmarkPress?.(event)} />
          </View>
        </View>
        {event.category && event.category.length > 0 && (
          <View className="flex-row flex-wrap mt-1.5">
            {event.category.map((cat: string) => (
              <View
                key={cat}
                className="px-2.5 py-1 rounded-full mr-1.5 mb-1"
                style={{ backgroundColor: TAG_COLORS[0].bg }}
              >
                <Text
                  className="text-xs font-medium"
                  style={{ color: TAG_COLORS[0].text }}
                >
                  {translate(`interests.${cat}`, { defaultValue: cat })}
                </Text>
              </View>
            ))}
          </View>
        )}
      </View>
    </Pressable>
  );
}
