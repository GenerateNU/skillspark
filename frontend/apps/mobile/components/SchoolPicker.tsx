import React, { useState } from 'react';
import { View, TouchableOpacity, ScrollView } from 'react-native';
import { ThemedText } from '@/components/themed-text';
import { IconSymbol } from '@/components/ui/icon-symbol';
import { AppColors, Colors } from '@/constants/theme';
import { useColorScheme } from '@/hooks/use-color-scheme';
import { useGetAllSchools, School } from '@skillspark/api-client';
import { useTranslation } from 'react-i18next';

type SchoolPickerProps = {
  value: string;
  onChange: (schoolId: string) => void;
};

export function SchoolPicker({ value, onChange }: SchoolPickerProps) {
  const [showDrop, setShowDrop] = useState(false);
  const colorScheme = useColorScheme();
  const theme = Colors[colorScheme ?? 'light'];
  const { t: translate } = useTranslation();

  const { data, isLoading, isError } = useGetAllSchools();
  const schools = Array.isArray(data?.data) ? data.data : [];
  const selectedSchool = schools.find((s: School) => s.id === value);

  const placeholderLabel = isLoading
    ? translate('childProfile.loadingSchools')
    : isError
    ? translate('childProfile.failedToLoadSchools')
    : translate('childProfile.selectSchool');

  return (
    <View className="z-[20]">
      <TouchableOpacity
        className="rounded-[10px] px-4 py-[14px] flex-row items-center justify-between mb-6 bg-[#F3F4F6] dark:bg-[#27272a]"
        onPress={() => setShowDrop(prev => !prev)}
        disabled={isLoading || isError}
      >
        <ThemedText
          className={`font-nunito ${selectedSchool ? '' : 'text-[#9CA3AF]'}`}
        >
          {selectedSchool ? selectedSchool.name : placeholderLabel}
        </ThemedText>
        <IconSymbol name="chevron.down" size={16} color={AppColors.mutedText} />
      </TouchableOpacity>
      {showDrop && (
        <View
          className="absolute left-0 right-0 top-[52px] rounded-[10px] border z-[100] elevation-5"
          style={{
            backgroundColor: theme.dropdownBg,
            borderColor: theme.borderColor,
            shadowColor: '#000',
            shadowOpacity: 0.1,
            shadowRadius: 8,
            shadowOffset: { width: 0, height: 2 },
          }}
        >
          <ScrollView nestedScrollEnabled className="max-h-[200px]">
            {schools.map(school => (
              <TouchableOpacity
                key={school.id}
                className="px-4 py-3 border-b border-b-[#E5E7EB] dark:border-b-[#3f3f46]"
                onPress={() => { onChange(school.id); setShowDrop(false); }}
              >
                <ThemedText>{school.name}</ThemedText>
              </TouchableOpacity>
            ))}
            {schools.length === 0 && !isLoading && (
              <View className="px-4 py-3">
                <ThemedText className="text-[#6B7280]">{translate('childProfile.noSchoolsFound')}</ThemedText>
              </View>
            )}
          </ScrollView>
        </View>
      )}
    </View>
  );
}