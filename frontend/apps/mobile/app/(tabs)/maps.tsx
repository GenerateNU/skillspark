import React, { useState } from 'react';
import { StyleSheet, View, Text } from 'react-native';
import MapView, { Marker, PROVIDER_GOOGLE, Region } from 'react-native-maps';
import { ThemedView } from '@/components/themed-view';

const MOCK_LOCATIONS = [
    {
      id: '10000000-0000-0000-0000-000000000001',
      title: 'Junior Robotics Workshop', 
      description: '123 Sukhumvit Road, Khlong Toei, Bangkok',
      latitude: 13.7563,
      longitude: 100.5018,
    },
    {
      id: '10000000-0000-0000-0000-000000000002',
      title: 'Soccer Skills Training', 
      description: '456 Rama IV Road, Pathum Wan, Bangkok',
      latitude: 13.7467,
      longitude: 100.5350,
    },
    {
      id: '10000000-0000-0000-0000-000000000003',
      title: 'Astronomy Club',
      description: '789 Vibhavadi Rangsit Road, Chatuchak, Bangkok',
      latitude: 13.8200,
      longitude: 100.5600,
    },
    {
      id: '10000000-0000-0000-0000-000000000004',
      title: 'Painting & Drawing Workshop',
      description: '321 Phetchaburi Road, Ratchathewi, Bangkok',
      latitude: 13.7650,
      longitude: 100.5380,
    },
    {
      id: '10000000-0000-0000-0000-000000000005',
      title: 'Piano for Beginners',
      description: '654 Sathorn Road, Yan Nawa, Bangkok',
      latitude: 13.7300,
      longitude: 100.5240,
    },
  ];
const INITIAL_REGION: Region = {
  latitude: 13.7563, 
  longitude: 100.5018,
  latitudeDelta: 0.0922,
  longitudeDelta: 0.0421,
};

export default function MapScreen() {
  return (
    <ThemedView style={styles.container}>
      <MapView
        style={styles.map}
        provider={PROVIDER_GOOGLE}
        initialRegion={INITIAL_REGION}
        showsUserLocation={true}
        showsMyLocationButton={true}
      >
        {MOCK_LOCATIONS.map((loc) => (
          <Marker
            key={loc.id}
            coordinate={{ latitude: loc.latitude, longitude: loc.longitude }}
            title={loc.title}
            description={loc.description}
          />
        ))}
      </MapView>
    </ThemedView>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
  },
  map: {
    width: '100%',
    height: '100%',
  },
});