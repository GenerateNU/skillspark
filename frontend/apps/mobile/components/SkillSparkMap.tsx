import React, { useRef, useState } from 'react';
import { StyleSheet } from 'react-native';
import MapView, { Marker, PROVIDER_GOOGLE } from 'react-native-maps';
import * as Location from 'expo-location';
import { Ionicons } from '@expo/vector-icons';
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
    <ThemedView style={styles.container}>
      <MapView
        ref={mapRef}
        style={styles.map}
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
            <Ionicons 
              name="location" 
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

const styles = StyleSheet.create({
  container: { flex: 1 },
  map: { width: '100%', height: '100%' },
});