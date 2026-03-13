import React from 'react';
import { View, TouchableOpacity } from 'react-native';
import { ThemedText } from '@/components/themed-text';
import { IconSymbol } from '@/components/ui/icon-symbol';
import { InterestTags } from './InterestTags';
import { Child } from '@skillspark/api-client';

type ChildListItemProps = {
  child: Child;
  onPress?: () => void;
};

export function ChildListItem({ child, onPress }: ChildListItemProps) {
  return (
    <TouchableOpacity className="flex-row items-start py-3 gap-3" onPress={onPress} activeOpacity={0.7}>
      <View className="w-11 h-11 rounded-[22px] border-[1.5px] border-[#9CA3AF] items-center justify-center">
        <ThemedText className="text-[15px] font-nunito-semibold">
          {child.name?.slice(0, 2).toUpperCase() || '??'}
        </ThemedText>
      </View>
      <View className="flex-1 gap-[2px]">
        <ThemedText className="text-base font-nunito-semibold">{child.name}</ThemedText>
        <ThemedText className="text-[13px] text-[#6B7280] font-nunito">
          {child.birth_month ? `${child.birth_month}, ` : ''}
          {child.birth_year}
        </ThemedText>
        {child.interests && (
          <ThemedText className="text-[13px] text-[#6B7280] font-nunito">Interests</ThemedText>
        )}
        <InterestTags interests={child.interests} />
      </View>
      <IconSymbol name="chevron.right" size={18} color="#9CA3AF" />
    </TouchableOpacity>
  );
}
