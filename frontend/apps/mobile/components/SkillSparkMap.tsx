import React, { useRef, useState } from "react";
import MapView, { Marker, PROVIDER_GOOGLE } from "react-native-maps";
import * as Location from "expo-location";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { ThemedView } from "@/components/themed-view";
import { EventCard } from "@/components/EventCard";
import { OrgListSheet } from "@/components/OrgListSheet";

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

interface SkillSparkMapProps {
  locations: LocationPin[];
  userLocation: Location.LocationObject | null;
  onFilterPress: () => void;
}

export function SkillSparkMap({ locations, userLocation, onFilterPress }: SkillSparkMapProps) {
  const mapRef = useRef<MapView>(null);
  const [selectedPin, setSelectedPin] = useState<LocationPin | null>(null);

  const initialRegion = userLocation
    ? {
        latitude: userLocation.coords.latitude,
        longitude: userLocation.coords.longitude,
        latitudeDelta: 0.05,
        longitudeDelta: 0.05,
      }
    : {
        latitude: 13.7563,
        longitude: 100.5018,
        latitudeDelta: 0.1,
        longitudeDelta: 0.1,
      };

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
            <IconSymbol
              name="location.fill"
              size={40}
              color={selectedPin?.id === loc.id ? "#FF4B4B" : "#FF6B6B"}
            />
          </Marker>
        ))}
      </MapView>
      {selectedPin ? (
        <EventCard pin={selectedPin} />
      ) : (
        <OrgListSheet locations={locations} userLocation={userLocation} onFilterPress={onFilterPress} />
      )}
    </ThemedView>
  );
}
