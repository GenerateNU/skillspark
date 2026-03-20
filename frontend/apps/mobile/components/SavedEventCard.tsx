import { Event } from "@skillspark/api-client";
import React from 'react';
import { Image, Text, TouchableOpacity, View } from 'react-native';

import { Ionicons } from '@expo/vector-icons';


interface BookmarkIconProps {
  onPress?: () => void; 
}

export function BookmarkIcon({ onPress }: BookmarkIconProps) {
  return (
    <TouchableOpacity onPress={onPress}>
      <Ionicons 
        name="bookmark" 
        size={24} 
        color="#FFC107" 
      />
    </TouchableOpacity>
  );
}

interface SavedEventCardProps {
  event: Event;
  onBookmarkPress?: (event: Event) => void; 
}


export function SavedEventCard({ event, onBookmarkPress }: SavedEventCardProps) {  return (
    <View
      style={{
        marginHorizontal: 20,
        marginBottom: 12,
        flexDirection: 'row',
        backgroundColor: '#99C0EE4D',
        borderRadius: 12,
        padding: 16,
        height: 150,
        alignItems: 'center',
        shadowColor: '#000',
        shadowOpacity: 0.05,
        shadowOffset: { width: 0, height: 2 },
        shadowRadius: 4,
        elevation: 2,
      }}
    >
      <View
        style={{
          width: 80,
          height: 80,
          borderRadius: 40,
          overflow: 'hidden',
          marginRight: 16,
          alignItems: 'center',
          justifyContent: 'center',
          backgroundColor: '#E5E7EB',
        }}
      >
        {event.presigned_url && (
          <Image
            source={{ uri: event.presigned_url }}
            style={{ width: '100%', height: '100%' }}
          />
        )}
      </View>

      <View style={{ flex: 1, justifyContent: 'center' }}>
        <View style={{ flexDirection: 'row', alignItems: 'center' }}>
          <Text style={{ fontSize: 16, fontWeight: '600', color: '#111', flexShrink: 1 }}>
            {event.title}
          </Text>

          <View style={{ marginLeft: 12 }}>
            <BookmarkIcon onPress={() => onBookmarkPress?.(event)} />
          </View>
        </View>

        {event.category && event.category.length > 0 && (
          <View style={{ flexDirection: 'row', flexWrap: 'wrap', marginTop: 6 }}>
            {event.category.map((cat: string) => (
              <View
                key={cat}
                style={{
                  backgroundColor: '#0E9888',
                  paddingHorizontal: 10,
                  paddingVertical: 4,
                  borderRadius: 999,
                  marginRight: 6,
                  marginBottom: 4,
                }}
              >
                <Text style={{ fontSize: 12, color: '#FFFFFF', fontWeight: '500' }}>
                  {cat}
                </Text>
              </View>
            ))}
          </View>
        )}

        {event.description && (
          <Text style={{ fontSize: 14, color: '#555', marginTop: 6 }} numberOfLines={2}>
            {event.description}
          </Text>
        )}

        {event.age_range_min != null && (
          <Text style={{ fontSize: 12, color: '#777', marginTop: 4 }}>
            🧑 {event.age_range_min}{event.age_range_max != null ? `–${event.age_range_max}` : ''}+
          </Text>
        )}
      </View>
    </View>
  );
}