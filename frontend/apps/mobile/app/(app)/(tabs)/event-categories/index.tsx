import React from "react";
import {
  View,
  TouchableOpacity,
  ScrollView,
  Text,
} from "react-native";
import { Stack, useRouter, useLocalSearchParams } from "expo-router";
import { useSafeAreaInsets } from "react-native-safe-area-context";
import { ThemedView } from "@/components/themed-view";
import { Colors, AppColors, FontSizes, FontFamilies } from "@/constants/theme";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { ErrorScreen } from "@/components/ErrorScreen";
import { useGetAllEventOccurrences } from "@skillspark/api-client";
import { useGeolocation } from "@/hooks/use-geolocation";
import { EventCategoriesListItem } from "@/components/EventCategoriesListItem";

export default function EventCategoryScreen() {
  const router = useRouter();
  const params = useLocalSearchParams();
  const insets = useSafeAreaInsets();
  const theme = Colors.light;

  const category = params.category as string;
  const { lat: geoLocationLat, lng: geoLocationLong } = useGeolocation();

  const { data: localizedOccurrencesResp } = useGetAllEventOccurrences({
    category,
    lat: geoLocationLat,
    lng: geoLocationLong,
    radius_km: 50,
    limit: 20,
    soldout: false,
  });

  if (!category) {
    return <ErrorScreen message="Illegal state: no category supplied" />;
  }

  if (!geoLocationLat || !geoLocationLong) {
    return <ErrorScreen message="Illegal state: no location found" />;
  }

  const occurrences = localizedOccurrencesResp?.status === 200 ? localizedOccurrencesResp.data : [];

  return (
    <ThemedView className="flex-1" style={{ paddingTop: insets.top }}>
      <Stack.Screen options={{ headerShown: false }} />

      {/* Header */}
      <View style={{ paddingHorizontal: 20, paddingTop: 10, paddingBottom: 6 }}>
        <TouchableOpacity
          onPress={() => router.back()}
          className="w-8 h-8 justify-center items-start"
        >
          <IconSymbol name="chevron.left" size={24} color={theme.text} />
        </TouchableOpacity>
        <Text
          style={{
            fontFamily: FontFamilies.bold,
            fontSize: FontSizes.lg,
            color: AppColors.primaryText,
            marginTop: 8,
          }}
        >
          {category}
        </Text>
      </View>

      {occurrences.length === 0 ? (
        <Text
          style={{
            fontFamily: FontFamilies.bold,
            fontSize: FontSizes.lg,
            color: AppColors.primaryText,
            paddingHorizontal: 20,
            marginTop: 20,
          }}
        >
          No events found for this category.
        </Text>
      ) : (
        <ScrollView
          contentContainerStyle={{ paddingHorizontal: 20, paddingBottom: 40 }}
          showsVerticalScrollIndicator={false}
        >
          {occurrences.map((o) => (
            <EventCategoriesListItem
              key={o.id}
              eventOccurrence={o}
              onPress={() => router.push(`/event/${o.event.id}`)}
            />
          ))}
        </ScrollView>
      )}
    </ThemedView>
  );
}
