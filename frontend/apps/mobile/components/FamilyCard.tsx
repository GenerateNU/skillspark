import React from 'react';
import { View, useColorScheme } from 'react-native';
import { ThemedText } from '@/components/themed-text';

type FamilyCardProps = {
  initials: string;
  name: string;
  date: string;
};

export function FamilyCard({ initials, name, date }: FamilyCardProps) {
  const colorScheme = useColorScheme();

  const backgroundColor = colorScheme === 'dark' ? '#27272a' : '#F3F4F6';
  const borderColor = colorScheme === 'dark' ? '#3f3f46' : '#E5E7EB';

  return (
    <View className="w-[48%] rounded-xl p-[10px] flex-row items-center" style={{ backgroundColor }}>
      <View className="w-8 h-8 rounded-2xl border items-center justify-center mr-2" style={{ borderColor }}>
        <ThemedText className="text-xs font-nunito-semibold">{initials}</ThemedText>
      </View>
      <View>
        <ThemedText className="text-sm font-nunito-medium">{name}</ThemedText>
        <ThemedText className="text-[10px] text-[#6B7280] font-nunito">{date}</ThemedText>
      </View>
    </View>
  );
}
