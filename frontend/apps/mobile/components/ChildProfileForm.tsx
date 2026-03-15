import React from 'react';
import { View, TextInput, TouchableOpacity, ScrollView, useColorScheme } from 'react-native';
import { ThemedText } from '@/components/themed-text';
import { IconSymbol } from '@/components/ui/icon-symbol';
import { Colors, AppColors, TAG_COLORS } from '@/constants/theme';
import { SchoolPicker } from '@/components/SchoolPicker';

const INTEREST_OPTIONS = [
  'science', 'math', 'music', 'art', 'sports', 'technology', 'language', 'other',
];

const capitalize = (s: string) => s.charAt(0).toUpperCase() + s.slice(1);

export const MONTHS = [
  'January','February','March','April','May','June',
  'July','August','September','October','November','December',
];

export const YEARS = Array.from({ length: 20 }, (_, i) => String(new Date().getFullYear() - i));

export type ChildProfileFormProps = {
  firstName: string;
  setFirstName: (v: string) => void;
  lastName: string;
  setLastName: (v: string) => void;
  birthMonth: string;
  setBirthMonth: (v: string) => void;
  birthYear: string;
  setBirthYear: (v: string) => void;
  schoolId: string;
  setSchoolId: (v: string) => void;
  interests: string[];
  setInterests: React.Dispatch<React.SetStateAction<string[]>>;
  searchQuery: string;
  setSearchQuery: (v: string) => void;
  showMonthDrop: boolean;
  setShowMonthDrop: (v: boolean) => void;
  showYearDrop: boolean;
  setShowYearDrop: (v: boolean) => void;
};

export function ChildProfileForm({
  firstName, setFirstName,
  lastName, setLastName,
  birthMonth, setBirthMonth,
  birthYear, setBirthYear,
  schoolId, setSchoolId,
  interests, setInterests,
  searchQuery, setSearchQuery,
  showMonthDrop, setShowMonthDrop,
  showYearDrop, setShowYearDrop,
}: ChildProfileFormProps) {
  const colorScheme = useColorScheme();
  const theme = Colors[colorScheme ?? 'light'];

  const removeInterest = (tag: string) => setInterests(prev => prev.filter(i => i !== tag));

  const toggleInterest = (item: string) => {
    setInterests(prev =>
      prev.includes(item) ? prev.filter(i => i !== item) : [...prev, item]
    );
  };

  const filteredOptions = INTEREST_OPTIONS.filter(o =>
    o.toLowerCase().includes(searchQuery.toLowerCase())
  );

  return (
    <>
      <TextInput
        className="rounded-[10px] px-4 py-[14px] text-base font-nunito mb-3"
        style={{ backgroundColor: theme.inputBg, color: theme.text }}
        value={firstName}
        onChangeText={setFirstName}
        placeholder="First Name"
        placeholderTextColor={AppColors.placeholderText}
      />
      <TextInput
        className="rounded-[10px] px-4 py-[14px] text-base font-nunito mb-3"
        style={{ backgroundColor: theme.inputBg, color: theme.text }}
        value={lastName}
        onChangeText={setLastName}
        placeholder="Last Name"
        placeholderTextColor={AppColors.placeholderText}
      />
      <View className="flex-row gap-3 mb-6" style={{ zIndex: 10 }}>
        <View className="flex-1" style={{ zIndex: 10 }}>
          <TouchableOpacity
            className="rounded-[10px] px-4 py-[14px] flex-row items-center justify-between"
            style={{ backgroundColor: theme.inputBg }}
            onPress={() => { setShowMonthDrop(!showMonthDrop); setShowYearDrop(false); }}
          >
            <ThemedText className={birthMonth ? '' : 'font-nunito'} style={birthMonth ? {} : { color: AppColors.placeholderText }}>
              {birthMonth || 'Month'}
            </ThemedText>
            <IconSymbol name="chevron.down" size={16} color={AppColors.mutedText} />
          </TouchableOpacity>

          {showMonthDrop && (
            <View
              className="absolute left-0 right-0 rounded-[10px] border"
              style={{ top: 52, backgroundColor: theme.dropdownBg, borderColor: theme.borderColor, zIndex: 100, elevation: 5, shadowColor: '#000', shadowOpacity: 0.1, shadowRadius: 8, shadowOffset: { width: 0, height: 2 } }}
            >
              <ScrollView nestedScrollEnabled style={{ maxHeight: 180 }}>
                {MONTHS.map(m => (
                  <TouchableOpacity
                    key={m}
                    className="px-4 py-3 border-b"
                    style={{ borderBottomColor: theme.borderColor }}
                    onPress={() => { setBirthMonth(m); setShowMonthDrop(false); }}
                  >
                    <ThemedText>{m}</ThemedText>
                  </TouchableOpacity>
                ))}
              </ScrollView>
            </View>
          )}
        </View>
        <View className="flex-1" style={{ zIndex: 10 }}>
          <TouchableOpacity
            className="rounded-[10px] px-4 py-[14px] flex-row items-center justify-between"
            style={{ backgroundColor: theme.inputBg }}
            onPress={() => { setShowYearDrop(!showYearDrop); setShowMonthDrop(false); }}
          >
            <ThemedText className={birthYear ? '' : 'font-nunito'} style={birthYear ? {} : { color: AppColors.placeholderText }}>
              {birthYear || 'Year'}
            </ThemedText>
            <IconSymbol name="chevron.down" size={16} color={AppColors.mutedText} />
          </TouchableOpacity>

          {showYearDrop && (
            <View
              className="absolute left-0 right-0 rounded-[10px] border"
              style={{ top: 52, backgroundColor: theme.dropdownBg, borderColor: theme.borderColor, zIndex: 100, elevation: 5, shadowColor: '#000', shadowOpacity: 0.1, shadowRadius: 8, shadowOffset: { width: 0, height: 2 } }}
            >
              <ScrollView nestedScrollEnabled style={{ maxHeight: 180 }}>
                {YEARS.map(y => (
                  <TouchableOpacity
                    key={y}
                    className="px-4 py-3 border-b"
                    style={{ borderBottomColor: theme.borderColor }}
                    onPress={() => { setBirthYear(y); setShowYearDrop(false); }}
                  >
                    <ThemedText>{y}</ThemedText>
                  </TouchableOpacity>
                ))}
              </ScrollView>
            </View>
          )}
        </View>
      </View>
      <SchoolPicker value={schoolId} onChange={setSchoolId} />
      <ThemedText className="text-base font-nunito-semibold mb-3">Interests</ThemedText>
      {interests.length > 0 && (
        <ScrollView horizontal showsHorizontalScrollIndicator={false} className="mb-3" contentContainerStyle={{ gap: 8, paddingRight: 4 }}>
          {interests.map((tag, idx) => {
            const color = TAG_COLORS[idx % TAG_COLORS.length];
            return (
              <TouchableOpacity
                key={tag}
                className="flex-row items-center px-2 py-1 rounded-full border gap-1"
                style={{ backgroundColor: color.bg, borderColor: color.border }}
                onPress={() => removeInterest(tag)}
              >
                <IconSymbol name="camera.filters" size={13} color={color.border} />
                <ThemedText className="text-xs font-nunito-medium" style={{ color: color.text }}>{capitalize(tag)}</ThemedText>
              </TouchableOpacity>
            );
          })}
        </ScrollView>
      )}
      <View className="border rounded-[10px] overflow-hidden mb-6" style={{ borderColor: theme.borderColor }}>
        <View className="flex-row items-center px-4 py-3 gap-2">
          <TextInput
            className="flex-1 text-base font-nunito"
            style={{ color: theme.text }}
            value={searchQuery}
            onChangeText={setSearchQuery}
            placeholder="Search interests..."
            placeholderTextColor={AppColors.placeholderText}
          />
          <IconSymbol name="magnifyingglass" size={20} color={AppColors.mutedText} />
        </View>
        <View className="h-px" style={{ backgroundColor: theme.borderColor }} />
        <View onStartShouldSetResponder={() => true} onMoveShouldSetResponder={() => true}>
          <ScrollView nestedScrollEnabled showsVerticalScrollIndicator style={{ maxHeight: 150 }}>
            {filteredOptions.map(item => (
              <TouchableOpacity
                key={item}
                className="flex-row items-center justify-between px-4 py-4 border-b"
                style={{ borderBottomColor: theme.inputBg }}
                onPress={() => toggleInterest(item)}
              >
                <ThemedText className="text-base font-nunito">{capitalize(item)}</ThemedText>
                <View
                  className="w-[22px] h-[22px] rounded-[4px] border-[1.5px] items-center justify-center"
                  style={{ borderColor: interests.includes(item) ? AppColors.checkboxSelected : AppColors.subtleText }}
                >
                  {interests.includes(item) && (
                    <IconSymbol name="checkmark" size={12} color={AppColors.checkboxSelected} />
                  )}
                </View>
              </TouchableOpacity>
            ))}
          </ScrollView>
        </View>
      </View>
    </>
  );
}
