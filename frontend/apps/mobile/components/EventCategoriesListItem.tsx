import React from "react";
import { View, Text } from "react-native";
import { Image } from "expo-image";
import { TouchableOpacity } from "react-native";

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
  const { location, event } = eventOccurrence;
  const organization = useGetOrganization(event.organization_id);
  const orgName = organization.data?.status === 200 ? organization.data.data.name : "";
  const address = `${location.address_line1}, ${location.district} ${location.postal_code}`;
  const ageRange = `Ages ${event.age_range_min} - ${event.age_range_max}`;

  return (
    <TouchableOpacity
      onPress={onPress}
      activeOpacity={0.7}
      style={{
        flexDirection: "row",
        alignItems: "center",
        backgroundColor: "#fff",
        borderRadius: 12,
        padding: 12,
        marginBottom: 12,
        gap: 12,
        shadowColor: "#000",
        shadowOpacity: 0.08,
        shadowRadius: 6,
        shadowOffset: { width: 0, height: 2 },
        elevation: 2,
      }}
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

      <View style={{ flex: 1, gap: 3 }}>
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
