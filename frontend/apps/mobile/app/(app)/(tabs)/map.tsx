import React, { useState, useEffect, useMemo } from "react";
import { Linking, Button, ActivityIndicator, View } from "react-native";
import * as Location from "expo-location";
import { useRouter, useLocalSearchParams } from "expo-router";
import {
  useListOrganizations,
  useGetAllLocations,
  useGetAllEventOccurrences,
} from "@skillspark/api-client";
import type {
  Organization,
  Location as OrgLocation,
} from "@skillspark/api-client";
import { ThemedView } from "@/components/themed-view";
import { ThemedText } from "@/components/themed-text";
import { SkillSparkMap } from "@/components/SkillSparkMap";
import { useTranslation } from "react-i18next";
import { haversineDistance } from "@/utils/distance";

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

const MAX_DISTANCE = 40;
const MAX_AGE = 12;

export default function MapScreen() {
  const { t: translate } = useTranslation();
  const router = useRouter();
  const params = useLocalSearchParams<{
    distanceKm?: string;
    minStartMinutes?: string;
    maxStartMinutes?: string;
    age?: string;
    categories?: string;
  }>();

  const {
    data: orgsData,
    isLoading: isOrgsLoading,
    error,
  } = useListOrganizations();
  const { data: locsData, isLoading: isLocsLoading } = useGetAllLocations();

  const bangkok: Location.LocationObject = {
    coords: {
      latitude: 13.7563,
      longitude: 100.5018,
      altitude: null,
      altitudeAccuracy: null,
      accuracy: null,
      heading: null,
      speed: null,
    },
    timestamp: 34,
  };

  const organizations: Organization[] =
    orgsData?.status === 200 ? orgsData.data : [];
  const locations: OrgLocation[] =
    locsData?.status === 200 ? locsData.data : [];

  const isApiLoading = isOrgsLoading || isLocsLoading;

  // ── Parse active filter params ──────────────────────────────────────────
  const activeDistanceKm = params.distanceKm
    ? parseInt(params.distanceKm, 10)
    : MAX_DISTANCE;
  const activeMinStartMinutes =
    params.minStartMinutes && params.minStartMinutes !== ""
      ? parseInt(params.minStartMinutes, 10)
      : null;
  const activeMaxStartMinutes =
    params.maxStartMinutes && params.maxStartMinutes !== ""
      ? parseInt(params.maxStartMinutes, 10)
      : null;
  const activeAge = params.age ? parseInt(params.age, 10) : 0;
  const activeCategories =
    params.categories && params.categories.length > 0
      ? params.categories.split(",").filter(Boolean)
      : [];

  const filtersActive =
    activeDistanceKm < MAX_DISTANCE ||
    activeMinStartMinutes != null ||
    activeMaxStartMinutes != null ||
    activeAge > 0 ||
    activeCategories.length > 0;

  // ── Fetch event occurrences when filters are active ──────────────────────
  const { data: occurrencesData, isLoading: isOccurrencesLoading } =
    useGetAllEventOccurrences(
      {
        limit: 1000,
        categories: activeCategories.length > 0 ? activeCategories : undefined,
      },
      { query: { enabled: filtersActive } }
    );

  // ── Build full map pins ──────────────────────────────────────────────────
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
          rating: org.review_summary?.average_rating ?? 0,
          members: 0,
          image: org.pfp_s3_key,
        };
      })
      .filter(Boolean) as LocationPin[];
  }, [organizations, locations]);

  // ── Apply filters client-side ────────────────────────────────────────────
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
        await Location.getCurrentPositionAsync({});
        setUserLocation(bangkok);
      } catch (err) {
        console.warn("Could not fetch location:", err);
      } finally {
        setIsLocationLoading(false);
      }
    })();
  }, []);

  const filteredLocations = useMemo(() => {
    if (!filtersActive) return mapLocations;

    // Build set of matching org IDs from event occurrence filters
    let matchingOrgIds: Set<string> | null = null;

    if (filtersActive) {
      const occurrences =
        occurrencesData?.status === 200 ? occurrencesData.data : [];

      const filtered = occurrences.filter((occ) => {
        // Time-of-day filter
        if (
          activeMinStartMinutes != null ||
          activeMaxStartMinutes != null
        ) {
          const start = new Date(occ.start_time);
          const startMins = start.getHours() * 60 + start.getMinutes();
          if (
            activeMinStartMinutes != null &&
            startMins < activeMinStartMinutes
          )
            return false;
          if (
            activeMaxStartMinutes != null &&
            startMins > activeMaxStartMinutes
          )
            return false;
        }

        // Age filter: selected age must fall within event's age range
        if (activeAge > 0) {
          const minAge = occ.event.age_range_min;
          const maxAge = occ.event.age_range_max;
          if (minAge != null && minAge > activeAge) return false;
          if (maxAge != null && maxAge < activeAge) return false;
        }

        return true;
      });

      matchingOrgIds = new Set(
        filtered.map((o) => o.event.organization_id)
      );
    }

    return mapLocations.filter((loc) => {
      // Org-ID filter from occurrence matching
      if (matchingOrgIds != null && !matchingOrgIds.has(loc.id)) return false;

      // Distance filter
      if (activeDistanceKm < MAX_DISTANCE && userLocation) {
        const dist = haversineDistance(
          userLocation.coords.latitude,
          userLocation.coords.longitude,
          loc.latitude,
          loc.longitude
        );
        if (dist > activeDistanceKm) return false;
      }

      return true;
    });
  }, [
    filtersActive,
    mapLocations,
    occurrencesData,
    activeDistanceKm,
    activeMinStartMinutes,
    activeMaxStartMinutes,
    activeAge,
    userLocation,
  ]);

  // ── Navigate to filter screen ────────────────────────────────────────────
  function handleFilterPress() {
    router.push({
      pathname: "/(app)/(tabs)/map-filter",
      params: {
        distanceKm: String(activeDistanceKm),
        minStartMinutes:
          activeMinStartMinutes != null ? String(activeMinStartMinutes) : "",
        maxStartMinutes:
          activeMaxStartMinutes != null ? String(activeMaxStartMinutes) : "",
        age: String(activeAge),
        categories: activeCategories.join(","),
      },
    });
  }

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

  const isFilterLoading = filtersActive && isOccurrencesLoading;

  return (
    <ThemedView className="flex-1">
      <SkillSparkMap
        locations={filteredLocations}
        userLocation={userLocation}
        onFilterPress={handleFilterPress}
      />

      {isFilterLoading && (
        <View className="absolute top-[60px] self-center rounded-[20px] bg-[rgba(0,0,0,0.6)] px-5 py-[10px]">
          <ThemedText className="text-sm font-semibold text-white">
            Applying filters…
          </ThemedText>
        </View>
      )}

      {!isApiLoading && error && (
        <View className="absolute top-[60px] self-center rounded-[20px] bg-[rgba(0,0,0,0.6)] px-5 py-[10px]">
          <ThemedText className="text-sm font-semibold text-white">
            {translate("common.errorFetchingEvents")}
          </ThemedText>
        </View>
      )}

      {!isApiLoading && !error && !isFilterLoading && filteredLocations.length === 0 && (
        <View className="absolute top-[60px] self-center rounded-[20px] bg-[rgba(0,0,0,0.6)] px-5 py-[10px]">
          <ThemedText className="text-sm font-semibold text-white">
            {filtersActive
              ? "No organizations match your filters"
              : translate("common.noEventsNearby")}
          </ThemedText>
        </View>
      )}
    </ThemedView>
  );
}
