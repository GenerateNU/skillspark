import React, { useState, useEffect, useMemo } from 'react';
import { StyleSheet, Linking, Button, ActivityIndicator, View } from 'react-native';
import * as Location from 'expo-location';
import { useGetAllEventOccurrences } from '@skillspark/api-client';
import { ThemedView } from '@/components/themed-view';
import { ThemedText } from '@/components/themed-text';
import { SkillSparkMap } from '@/components/SkillSparkMap';

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
  const { data: occurrences, isLoading: isApiLoading, error } = useGetAllEventOccurrences();
  
  // mapLocations now strictly uses API data
  const mapLocations: LocationPin[] = useMemo(() => {
    return (occurrences || []) 
      .filter(occ => occ.location && occ.event)
      .map(occ => ({
        id: occ.id,
        title: occ.event.title,
        description: occ.event.description,
        latitude: occ.location.latitude,
        longitude: occ.location.longitude,
        // Defaulting rating as it is not currently provided by the API
        rating: 5.0, 
        members: occ.curr_enrolled,
        image: occ.event.header_image_s3_key
      }));
  }, [occurrences]);

  const [userLocation, setUserLocation] = useState<Location.LocationObject | null>(null);
  const [locationPermissionDenied, setLocationPermissionDenied] = useState(false);
  const [isLocationLoading, setIsLocationLoading] = useState(true);

  useEffect(() => {
    (async () => {
      const { status } = await Location.requestForegroundPermissionsAsync();
      
      if (status !== 'granted') {
        setLocationPermissionDenied(true);
        setIsLocationLoading(false);
        return;
      }

      try {
        const location = await Location.getCurrentPositionAsync({});
        setUserLocation(location);
      } catch (err) {
        console.warn('Could not fetch location:', err);
      } finally {
        setIsLocationLoading(false);
      }
    })();
  }, []);

  if (isLocationLoading || isApiLoading) {
    return (
      <ThemedView style={styles.centerContainer}>
        <ActivityIndicator size="large" />
        <ThemedText style={{ marginTop: 10 }}>Loading Map Data...</ThemedText>
      </ThemedView>
    );
  }

  if (locationPermissionDenied) {
    return (
      <ThemedView style={styles.centerContainer}>
        <ThemedText style={styles.errorText}>
          Permission to access location was denied. Please enable it in settings.
        </ThemedText>
        <Button title="Open Settings" onPress={() => Linking.openSettings()} />
      </ThemedView>
    );
  }

  return (
    <ThemedView style={styles.container}>
      <SkillSparkMap 
        locations={mapLocations} 
        userLocation={userLocation} 
      />

      {!isApiLoading && mapLocations.length === 0 && (
         <View style={styles.emptyStateContainer}>
            <ThemedText style={styles.emptyStateText}>No events found nearby.</ThemedText>
         </View>
      )}
    </ThemedView>
  );
}

const styles = StyleSheet.create({
  container: { flex: 1 },
  centerContainer: { 
    flex: 1, 
    justifyContent: 'center', 
    alignItems: 'center', 
    padding: 20 
  },
  errorText: {
    textAlign: 'center',
    marginBottom: 20,
    fontSize: 16,
  },
  emptyStateContainer: {
    position: 'absolute',
    top: 60,
    alignSelf: 'center',
    backgroundColor: 'rgba(0,0,0,0.6)',
    paddingHorizontal: 20,
    paddingVertical: 10,
    borderRadius: 20,
  },
  emptyStateText: {
    color: 'white',
    fontSize: 14,
    fontWeight: '600',
  }
});