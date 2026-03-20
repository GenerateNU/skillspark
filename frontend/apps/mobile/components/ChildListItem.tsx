import React from 'react';
import { View, TouchableOpacity } from 'react-native';
import { ThemedText } from '@/components/themed-text';
import { IconSymbol } from '@/components/ui/icon-symbol';
import { InterestTags } from './InterestTags';
import { Child } from '@skillspark/api-client';
import { AppColors } from '@/constants/theme';
import { useTranslation } from 'react-i18next';

type ChildListItemProps = {
  child: Child;
  onPress?: () => void;
};

export function ChildListItem({ child, onPress }: ChildListItemProps) {
  const { t } = useTranslation();
  return (
    <TouchableOpacity className="flex-row items-start py-3 gap-3" onPress={onPress} activeOpacity={0.7}>
      <View className="w-11 h-11 rounded-[22px] border-[1.5px] items-center justify-center" style={{ borderColor: AppColors.subtleText }}>
        <ThemedText className="text-[15px] font-nunito-semibold">
          {child.name?.slice(0, 2).toUpperCase() || '??'}
        </ThemedText>
      </View>
      <View className="flex-1 gap-[2px]">
        <ThemedText className="text-base font-nunito-semibold">{child.name}</ThemedText>
        <ThemedText className="text-[13px] font-nunito" style={{ color: AppColors.mutedText }}>
          {child.birth_month ? `${child.birth_month}, ` : ''}
          {child.birth_year}
        </ThemedText>
        {child.interests && (
          <ThemedText className="text-[13px] font-nunito" style={{ color: AppColors.mutedText }}>{t('familyInformation.interests')}</ThemedText>
        )}
        <InterestTags interests={child.interests} />
      </View>
      <IconSymbol name="chevron.right" size={18} color={AppColors.subtleText} />
    </TouchableOpacity>
  );
}
