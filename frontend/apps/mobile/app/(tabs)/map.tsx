import React, { useState, useEffect } from 'react';
import { StyleSheet, Linking, Button, ActivityIndicator } from 'react-native';
import * as Location from 'expo-location';
// import { useGetAllEventOccurrences } from '@skillspark/api-client'; <-- uncomment this when we will implement actual API calls
import { ThemedView } from '@/components/themed-view';
import { ThemedText } from '@/components/themed-text';
import { SkillSparkMap } from '@/components/SkillSparkMap';
import { MOCK_LOCATIONS, LocationPin } from '@/constants/mock-locations'; 

export default function MapScreen() {
  // const { data: occurrences, isLoading: isApiLoading, error } = useGetAllEventOccurrences(); <-- uncomment when API calls are implemented
  
  const isApiLoading = false; // Mock loading state
  const mapLocations: LocationPin[] = MOCK_LOCATIONS; 

  /* <--- uncomment this block when API calss are implemented
  const mapLocations: LocationPin[] = (occurrences || []) 
    .filter(occ => occ.location && occ.event)
    .map(occ => ({
      id: occ.id,
      title: occ.event.title,
      description: occ.location.address_line1,
      latitude: occ.location.latitude,
      longitude: occ.location.longitude,
    }));
  */

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
  }
});