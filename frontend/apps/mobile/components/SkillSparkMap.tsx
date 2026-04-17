import React, { useRef, useState } from "react";
import MapView, { Marker, PROVIDER_GOOGLE } from "react-native-maps";
import * as Location from "expo-location";
import { View } from "react-native";
import { SvgXml } from "react-native-svg";
import { ThemedView } from "@/components/themed-view";
import { OrgMapCard } from "@/components/OrgMapCard";
import { OrgListSheet } from "@/components/OrgListSheet";
import { FLOATING_TAB_BAR_SCROLL_PADDING } from "@/components/floating-tab-bar";
import { pinSvg } from "@/constants/mapPins";

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

interface SkillSparkMapProps {
  locations: LocationPin[];
  userLocation: Location.LocationObject | null;
}

export function SkillSparkMap({ locations, userLocation }: SkillSparkMapProps) {
  const mapRef = useRef<MapView>(null);
  const [selectedPin, setSelectedPin] = useState<LocationPin | null>(null);

  const initialRegion = userLocation
    ? {
        latitude: userLocation.coords.latitude,
        longitude: userLocation.coords.longitude,
        latitudeDelta: 0.05,
        longitudeDelta: 0.05,
      }
    : undefined;

  return (
    <ThemedView className="flex-1">
      <MapView
        ref={mapRef}
        style={{ width: "100%", height: "100%" }}
        provider={PROVIDER_GOOGLE}
        initialRegion={initialRegion}
        showsUserLocation={true}
        showsMyLocationButton={true}
        onPress={() => setSelectedPin(null)}
        userInterfaceStyle="light"
      >
        {locations.map((loc) => (
          <Marker
            key={loc.id}
            coordinate={{ latitude: loc.latitude, longitude: loc.longitude }}
            onPress={(e) => {
              e.stopPropagation();
              setSelectedPin(loc);
            }}
          >
            <SvgXml xml={pinSvg} width={36} height={45} />
          </Marker>
        ))}
      </MapView>

      {/* Permanent white background behind the floating nav bar so the map
          never bleeds through around the tab bar pill or safe-area zone. */}
      <View
        pointerEvents="none"
        style={{
          position: "absolute",
          bottom: 0,
          left: 0,
          right: 0,
          height: FLOATING_TAB_BAR_SCROLL_PADDING-10,
          backgroundColor: "#fffFFF",
          zIndex: 100,
        }}
      />

      {selectedPin ? (
        <OrgMapCard pin={selectedPin} userLocation={userLocation} />
      ) : (
        <OrgListSheet locations={locations} userLocation={userLocation} />
      )}
    </ThemedView>
  );
}
