import { Event } from "@skillspark/api-client";
import React from 'react';
import { Image, Text, TouchableOpacity, View } from 'react-native';

import { AppColors, TAG_COLORS } from "@/constants/theme";
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
        backgroundColor: AppColors.savedBackground,
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
          backgroundColor: AppColors.divider,
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
                  backgroundColor: TAG_COLORS[0].bg,
                  paddingHorizontal: 10,
                  paddingVertical: 4,
                  borderRadius: 999,
                  marginRight: 6,
                  marginBottom: 4,
                }}
              >
                <Text style={{ fontSize: 12, color: TAG_COLORS[0].text, fontWeight: '500' }}>
                  {cat}
                </Text>
              </View>
            ))}
          </View>
        )}

      </View>
    </View>
  );
}