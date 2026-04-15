import React from "react";
import { View, Text, TouchableOpacity } from "react-native";
import { Image } from "expo-image";
import { useTranslation } from "react-i18next";

import { EventOccurrence, useGetOrganization } from "@skillspark/api-client";
import { AppColors, FontFamilies, FontSizes } from "@/constants/theme";

type EventCategoriesListItemProps = {
  eventOccurrence: EventOccurrence;
  onPress?: () => void;
};

export function EventCategoriesListItem({
  eventOccurrence,
  onPress,
}: EventCategoriesListItemProps) {
  const { t: translate } = useTranslation();
  const { location, event } = eventOccurrence;
  const organization = useGetOrganization(event.organization_id);
  const orgName = organization.data?.status === 200 ? organization.data.data.name : "";
  const address = `${location.address_line1}, ${location.district} ${location.postal_code}`;
  const ageRange = `${translate("eventCategories.ages")} ${event.age_range_min} - ${event.age_range_max}`;

  return (
    <TouchableOpacity
      onPress={onPress}
      activeOpacity={0.7}
      className="flex-row items-center bg-white rounded-xl p-3 mb-3 gap-3 shadow-sm"
    >
      {event.presigned_url ? (
        <Image
          source={{ uri: event.presigned_url }}
          style={{ width: 90, height: 90, borderRadius: 8 }}
          contentFit="cover"
        />
      ) : (
        <View style={{ width: 90, height: 90, borderRadius: 8, backgroundColor: AppColors.categoryFallback }} />
      )}

      <View className="flex-1 gap-1">
        <Text style={{ fontFamily: FontFamilies.bold, fontSize: FontSizes.base, color: AppColors.primaryText }}>
          {event.title}
        </Text>
        {orgName ? (
          <Text style={{ fontFamily: FontFamilies.regular, fontSize: FontSizes.sm, color: AppColors.mutedText }}>
            {orgName}
          </Text>
        ) : null}
        <Text style={{ fontFamily: FontFamilies.regular, fontSize: FontSizes.sm, color: AppColors.mutedText }}>
          {ageRange}
        </Text>
        <Text style={{ fontFamily: FontFamilies.regular, fontSize: FontSizes.sm, color: AppColors.mutedText }}>
          {address}
        </Text>
      </View>
    </TouchableOpacity>
  );
}
