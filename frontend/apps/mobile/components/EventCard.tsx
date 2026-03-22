import React from 'react';
import { Alert, View, TouchableOpacity, Platform } from 'react-native';
import { IconSymbol } from '@/components/ui/icon-symbol';
import { ThemedText } from '@/components/themed-text';
import { ThemedView } from '@/components/themed-view';
import { useThemeColor } from '@/hooks/use-theme-color';
import { useTranslation } from 'react-i18next';

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

interface EventCardProps {
  pin: LocationPin;
}

export function EventCard({ pin }: EventCardProps) {
  const placeholderColor = useThemeColor({ light: '#D0D0D0', dark: '#333333' }, 'background');
  const { t } = useTranslation();

  return (
    <ThemedView
      className="absolute bottom-0 left-0 right-0 rounded-t-[25px] p-5"
      style={{
        paddingBottom: Platform.OS === 'ios' ? 40 : 20,
        shadowColor: "#000",
        shadowOffset: { width: 0, height: -2 },
        shadowOpacity: 0.1,
        shadowRadius: 5,
        elevation: 10,
      }}
    >
      <View className="mb-5 flex-row">
        <View
          className="mr-[15px] h-[90px] w-[90px] items-center justify-center rounded-[10px]"
          style={{ backgroundColor: placeholderColor }}
        >
           <IconSymbol name="photo" size={28} color="#888" />
        </View>
        <View className="flex-1 justify-center">
          <ThemedText type="subtitle" className="mb-1 font-bold">{pin.title}</ThemedText>
          <ThemedText className="mb-[6px] text-sm text-[#888]">{pin.members} {t('dashboard.members')}</ThemedText>
          
          <View className="mb-2 flex-row items-center">
            {[1, 2, 3, 4, 5].map((star) => (
              <IconSymbol
                key={star}
                name="star.fill"
                size={16}
                color={star <= Math.round(pin.rating) ? "#FFC107" : "#555"}
              />
            ))}
          </View>

          <ThemedText numberOfLines={2} className="text-sm leading-5 text-[#888]">
            {pin.description}
          </ThemedText>
        </View>
        <View className="ml-[10px] mt-[5px] items-end justify-start">
            <IconSymbol name="record.circle" size={24} color="#888" />
        </View>
      </View>
      <TouchableOpacity 
        className="w-full items-center rounded-xl bg-[#333] py-[15px]"
        activeOpacity={1} 
        onPress={() => Alert.alert(t('dashboard.reserve'), `${t('dashboard.reserved')}: ${pin.title}`)}
      >
        <ThemedText className="text-[18px] font-semibold text-white">{t('dashboard.reserve')}</ThemedText>
      </TouchableOpacity>
    </ThemedView>
  );
}