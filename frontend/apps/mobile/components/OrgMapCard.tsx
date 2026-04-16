import React from "react";
import { View, TouchableOpacity, Text } from "react-native";
import { Image } from "expo-image";
import type { LocationObject } from "expo-location";
import { useSafeAreaInsets } from "react-native-safe-area-context";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { ThemedText } from "@/components/themed-text";
import { ThemedView } from "@/components/themed-view";
import { useTranslation } from "react-i18next";
import { useRouter } from "expo-router";
import { haversineDistance } from "@/utils/distance";
import { RatingSmiley } from "@/components/RatingSmiley";
import { useGetReviewAggregateOrganization, type ReviewAggregate } from "@skillspark/api-client";

export interface LocationPin {
  id: string;
  title: string;
  description: string;
  latitude: number;
  longitude: number;
  rating: number;
  members: number;
  image?: string;
  district?: string;
}

interface OrgMapCardProps {
  pin: LocationPin;
  userLocation: LocationObject | null;
}

export function OrgMapCard({ pin, userLocation }: OrgMapCardProps) {
  const router = useRouter();
  const { t: translate } = useTranslation();
  const insets = useSafeAreaInsets();

  const { data: aggregateResp } = useGetReviewAggregateOrganization(pin.id, {
    query: { enabled: !!pin.id },
  });
  const aggregate =
    aggregateResp?.status === 200
      ? (aggregateResp.data as ReviewAggregate)
      : null;
  const avgRating = aggregate != null ? aggregate.average_rating : null;

  const distanceKm =
    userLocation != null
      ? haversineDistance(
          userLocation.coords.latitude,
          userLocation.coords.longitude,
          pin.latitude,
          pin.longitude,
        )
      : null;
  const distanceMi = distanceKm != null ? distanceKm * 0.621371 : null;

  const locationLine = [
    pin.district,
    distanceMi != null
      ? `${distanceMi.toFixed(1)} ${translate("map.mi")}`
      : null,
  ]
    .filter(Boolean)
    .join(" ");

  // Position card just above the floating tab bar pill
  const cardBottom = Math.max(insets.bottom, 8) + 92;

  return (
    <ThemedView
      className="absolute left-4 right-4 rounded-3xl p-4"
      style={{
        bottom: cardBottom,
        shadowColor: "#000",
        shadowOffset: { width: 0, height: 4 },
        shadowOpacity: 0.15,
        shadowRadius: 12,
        elevation: 8,
      }}
    >
      <View className="mb-3 flex-row">
        <View className="mr-3 h-[72px] w-[72px] overflow-hidden rounded-2xl bg-[#D9D9D9]">
          {pin.image ? (
            <Image
              source={{ uri: pin.image }}
              className="h-full w-full"
              contentFit="cover"
            />
          ) : (
            <View className="flex-1 items-center justify-center">
              <IconSymbol name="photo" size={26} color="#6B7280" />
            </View>
          )}
        </View>
        <View className="flex-1 justify-center gap-[2px]">
          <ThemedText className="font-nunito-bold text-[18px] leading-snug">
            {pin.title}
          </ThemedText>
          {locationLine.length > 0 && (
            <ThemedText className="font-nunito text-sm text-[#6B7280]">
              {locationLine}
            </ThemedText>
          )}
          {avgRating != null && (
            <View className="flex-row items-center gap-1">
              <RatingSmiley
                rating={aggregate!.total_reviews > 0 ? avgRating : null}
                width={16}
                height={16}
              />
              <ThemedText className="font-nunito text-sm text-[#6B7280]">
                {avgRating.toFixed(1)} {translate("map.smiles")}
              </ThemedText>
            </View>
          )}
        </View>
      </View>
      <TouchableOpacity
        className="w-full items-center rounded-xl bg-[#1C1C1E] py-[13px]"
        activeOpacity={0.85}
        onPress={() => router.push(`/org/${pin.id}`)}
      >
        <Text className="font-nunito-semibold text-base text-white">
          {translate("map.learnMore")}
        </Text>
      </TouchableOpacity>
    </ThemedView>
  );
}
