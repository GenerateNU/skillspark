import React, { useRef, useState } from 'react';
import MapView, { Marker, PROVIDER_GOOGLE } from 'react-native-maps';
import * as Location from 'expo-location';
import { IconSymbol } from '@/components/ui/icon-symbol';
import { ThemedView } from '@/components/themed-view';
import { EventCard } from '@/components/EventCard'; 

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
}

export function SkillSparkMap({ locations, userLocation }: SkillSparkMapProps) {
  const mapRef = useRef<MapView>(null);
  const [selectedPin, setSelectedPin] = useState<LocationPin | null>(null);

  const initialRegion = userLocation ? {
    latitude: userLocation.coords.latitude,
    longitude: userLocation.coords.longitude,
    latitudeDelta: 0.05,
    longitudeDelta: 0.05,
  } : undefined;

  return (
    <ThemedView className="flex-1">
      <MapView
        ref={mapRef}
        style={{ width: '100%', height: '100%' }}
        provider={PROVIDER_GOOGLE}
        initialRegion={initialRegion}
        showsUserLocation={true}
        showsMyLocationButton={true}
        onPress={() => setSelectedPin(null)}
        userInterfaceStyle="dark" 
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
      {selectedPin && (
        <EventCard pin={selectedPin} />
      )}
    </ThemedView>
  );
}