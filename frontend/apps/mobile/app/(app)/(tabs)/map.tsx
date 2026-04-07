import React, { useState, useEffect, useMemo } from "react";
import { Linking, Button, ActivityIndicator, View } from "react-native";
import * as Location from "expo-location";
import { useListOrganizations, useGetAllLocations } from "@skillspark/api-client";
import type { Organization, Location as OrgLocation } from "@skillspark/api-client";
import { ThemedView } from "@/components/themed-view";
import { ThemedText } from "@/components/themed-text";
import { SkillSparkMap } from "@/components/SkillSparkMap";
import { useTranslation } from "react-i18next";

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
  const { data: orgsData, isLoading: isOrgsLoading, error } = useListOrganizations();
  const { data: locsData, isLoading: isLocsLoading } = useGetAllLocations();

  const organizations: Organization[] = orgsData?.status === 200 ? orgsData.data : [];
  const locations: OrgLocation[] = locsData?.status === 200 ? locsData.data : [];

  const isApiLoading = isOrgsLoading || isLocsLoading;

  const mapLocations: LocationPin[] = useMemo(() => {
    if (!Array.isArray(organizations)) return [];

    return organizations
      .map((org) => {
        const location = locations.find((l) => l.id === org.location_id);
        if (!location) return null;
        return {
          id: org.id,
          title: org.name,
          description: "",
          latitude: location.latitude,
          longitude: location.longitude,
          rating: 5.0,
          members: 0,
          image: org.pfp_s3_key,
        };
      })
      .filter(Boolean) as LocationPin[];
  }, [organizations, locations]);

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
        <ThemedText className="mt-[10px]">
          {translate("common.loadingMapData")}
        </ThemedText>
      </ThemedView>
    );
  }

  if (locationPermissionDenied) {
    return (
      <ThemedView className="flex-1 items-center justify-center p-5">
        <ThemedText className="mb-5 text-center text-base">
          {translate("common.locationDenied")}
        </ThemedText>
        <Button
          title={translate("common.openSettings")}
          onPress={() => Linking.openSettings()}
        />
      </ThemedView>
    );
  }

  return (
    <ThemedView className="flex-1">
      <SkillSparkMap locations={mapLocations} userLocation={userLocation} />

      {!isApiLoading && error && (
        <View className="absolute top-[60px] self-center rounded-[20px] bg-[rgba(0,0,0,0.6)] px-5 py-[10px]">
          <ThemedText className="text-sm font-semibold text-white">
            {translate("common.errorFetchingEvents")}
          </ThemedText>
        </View>
      )}

      {!isApiLoading && !error && mapLocations.length === 0 && (
        <View className="absolute top-[60px] self-center rounded-[20px] bg-[rgba(0,0,0,0.6)] px-5 py-[10px]">
          <ThemedText className="text-sm font-semibold text-white">
            {translate("common.noEventsNearby")}
          </ThemedText>
        </View>
      )}
    </ThemedView>
  );
}
