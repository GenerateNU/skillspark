import React, { useState, useEffect, useMemo } from 'react';
import { Linking, Button, ActivityIndicator, View } from 'react-native';
import * as Location from 'expo-location';
import { useGetAllEventOccurrences } from '@skillspark/api-client';
import type { EventOccurrence } from '@skillspark/api-client';
import { ThemedView } from '@/components/themed-view';
import { ThemedText } from '@/components/themed-text';
import { SkillSparkMap } from '@/components/SkillSparkMap';
import { useTranslation } from 'react-i18next';
import AsyncStorage from '@react-native-async-storage/async-storage';

export interface LocationPin {
  id: string;
  title: string;
  description: string;
  latitude: number;
  longitude: number;
  rating: number;
  members: number;
  image?: string;
}

export default function MapScreen() {
  const { t: translate } = useTranslation();
  const { data, isLoading: isApiLoading, error } = useGetAllEventOccurrences();
  
  const occurrences: EventOccurrence[] = data?.status === 200 ? data.data : [];

  const mapLocations: LocationPin[] = useMemo(() => {
    if (!Array.isArray(occurrences)) return [];

    return occurrences
      .filter((occ) => occ.location && occ.event)
      .map((occ) => ({
        id: occ.id,
        title: occ.event.title,
        description: occ.event.description,
        latitude: occ.location.latitude,
        longitude: occ.location.longitude,
        rating: 5.0,
        members: occ.curr_enrolled,
        image: occ.event.header_image_s3_key,
      }));
  }, [occurrences]);

  const [userLocation, setUserLocation] =
    useState<Location.LocationObject | null>(null);
  const [locationPermissionDenied, setLocationPermissionDenied] =
    useState(false);
  const [isLocationLoading, setIsLocationLoading] = useState(true);

  useEffect(() => {
    (async () => {
      const { status } = await Location.requestForegroundPermissionsAsync();

      if (status !== "granted") {
        setLocationPermissionDenied(true);
        setIsLocationLoading(false);
        return;
      }

      try {
        const location = await Location.getCurrentPositionAsync({});
        setUserLocation(location);
      } catch (err) {
        console.warn("Could not fetch location:", err);
      } finally {
        setIsLocationLoading(false);
      }
    })();
  }, []);

  if (isLocationLoading || isApiLoading) {
    return (
      <ThemedView className="flex-1 items-center justify-center p-5">
        <ActivityIndicator size="large" />
        <ThemedText className="mt-[10px]">{translate('common.loadingMapData')}</ThemedText>
      </ThemedView>
    );
  }

  if (locationPermissionDenied) {
    return (
      <ThemedView className="flex-1 items-center justify-center p-5">
        <ThemedText className="mb-5 text-center text-base">
          {translate('common.locationDenied')}
        </ThemedText>
        <Button title={translate('common.openSettings')} onPress={() => Linking.openSettings()} />
      </ThemedView>
    );
  }

  return (
    <ThemedView className="flex-1">
      <SkillSparkMap locations={mapLocations} userLocation={userLocation} />

      {!isApiLoading && error && (
         <View className="absolute top-[60px] self-center rounded-[20px] bg-[rgba(0,0,0,0.6)] px-5 py-[10px]">
            <ThemedText className="text-sm font-semibold text-white">{translate('common.errorFetchingEvents')}</ThemedText>
         </View>
      )}

      {!isApiLoading && !error && mapLocations.length === 0 && (
         <View className="absolute top-[60px] self-center rounded-[20px] bg-[rgba(0,0,0,0.6)] px-5 py-[10px]">
            <ThemedText className="text-sm font-semibold text-white">{translate('common.noEventsNearby')}</ThemedText>
         </View>
      )}
    </ThemedView>
  );
}
