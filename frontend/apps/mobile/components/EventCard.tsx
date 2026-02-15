import React from 'react';
import { StyleSheet, View, TouchableOpacity, Platform } from 'react-native';
import { Ionicons } from '@expo/vector-icons';
import { LocationPin } from '@/constants/mock-locations';
import { ThemedText } from '@/components/themed-text';
import { ThemedView } from '@/components/themed-view';
import { useThemeColor } from '@/hooks/use-theme-color';

interface EventCardProps {
  pin: LocationPin;
}

export function EventCard({ pin }: EventCardProps) {
  const placeholderColor = useThemeColor({ light: '#D0D0D0', dark: '#333333' }, 'background');

  return (
    <ThemedView style={styles.cardContainer}>
      <View style={styles.cardContentRow}>
        <View style={[styles.imagePlaceholder, { backgroundColor: placeholderColor }]}>
           <Ionicons name="image-outline" size={28} color="#888" />
        </View>
        <View style={styles.textContainer}>
          <ThemedText type="subtitle" style={styles.title}>{pin.title}</ThemedText>
          <ThemedText style={styles.members}>{pin.members} members</ThemedText>
          
          <View style={styles.ratingRow}>
            {[1, 2, 3, 4, 5].map((star) => (
              <Ionicons 
                key={star} 
                name="star" 
                size={16} 
                color={star <= Math.round(pin.rating) ? "#FFC107" : "#555"} 
              />
            ))}
          </View>

          <ThemedText numberOfLines={2} style={styles.description}>
            {pin.description}
          </ThemedText>
        </View>
        <View style={styles.radioButtonContainer}>
            <Ionicons name="radio-button-on" size={24} color="#888" />
        </View>
      </View>
      <TouchableOpacity 
        style={styles.reserveButton} 
        activeOpacity={1} 
        onPress={() => alert(`Reserved: ${pin.title}`)} 
      >
        <ThemedText style={styles.reserveButtonText}>Reserve</ThemedText>
      </TouchableOpacity>
    </ThemedView>
  );
}

const styles = StyleSheet.create({
  cardContainer: {
    position: 'absolute',
    bottom: 0,
    left: 0,
    right: 0,
    borderTopLeftRadius: 25,
    borderTopRightRadius: 25,
    padding: 20,
    paddingBottom: Platform.OS === 'ios' ? 40 : 20,
    
    shadowColor: "#000",
    shadowOffset: { width: 0, height: -2 },
    shadowOpacity: 0.1,
    shadowRadius: 5,
    elevation: 10,
  },
  cardContentRow: {
    flexDirection: 'row',
    marginBottom: 20,
  },
  imagePlaceholder: {
    width: 90,
    height: 90,
    borderRadius: 10,
    justifyContent: 'center',
    alignItems: 'center',
    marginRight: 15,
  },
  textContainer: {
    flex: 1,
    justifyContent: 'center',
  },
  radioButtonContainer: {
    justifyContent: 'flex-start',
    alignItems: 'flex-end',
    marginLeft: 10,
    marginTop: 5,
  },
  title: {
    marginBottom: 4,
    fontWeight: 'bold', 
  },
  members: {
    fontSize: 14,
    color: '#888',
    marginBottom: 6,
  },
  ratingRow: {
    flexDirection: 'row',
    alignItems: 'center',
    marginBottom: 8,
  },
  description: {
    fontSize: 14,
    color: '#888',
    lineHeight: 20,
  },
  reserveButton: {
    backgroundColor: '#333', 
    paddingVertical: 15,
    borderRadius: 12,
    alignItems: 'center',
    width: '100%',
  },
  reserveButtonText: {
    color: 'white',
    fontSize: 18,
    fontWeight: '600',
  },
});