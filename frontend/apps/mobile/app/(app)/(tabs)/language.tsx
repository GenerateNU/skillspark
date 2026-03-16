import React, { useState } from 'react';
import { View, TouchableOpacity, useColorScheme } from 'react-native';
import { useRouter } from 'expo-router';
import { useSafeAreaInsets } from 'react-native-safe-area-context';
import { ThemedText } from '@/components/themed-text';
import { ThemedView } from '@/components/themed-view';
import { IconSymbol } from '@/components/ui/icon-symbol';
import { Colors } from '@/constants/theme';

const LANGUAGES = [
  { code: 'en', label: 'English', flag: '🇺🇸' },
  { code: 'th', label: 'Thai',    flag: '🇹🇭' },
];

export default function LanguageScreen() {
  const router = useRouter();
  const insets = useSafeAreaInsets();
  const colorScheme = useColorScheme();
  const theme = Colors[colorScheme ?? 'light'];

  const dividerColor = colorScheme === 'dark' ? '#3a3a3c' : '#E5E7EB';

  const [selected, setSelected] = useState('en');

  return (
    <ThemedView className="flex-1" style={{ paddingTop: insets.top }}>
      <View className="flex-row items-center justify-between px-5 py-[14px]">
        <TouchableOpacity
          onPress={() => router.navigate('/profile')}
          className="w-10 items-start justify-center"
          hitSlop={{ top: 10, bottom: 10, left: 10, right: 10 }}
        >
          <IconSymbol name="chevron.left" size={24} color={theme.text} />
        </TouchableOpacity>
        <ThemedText className="text-xl text-center font-nunito-bold">Settings</ThemedText>
        <View className="w-10" />
      </View>
      <ThemedText className="text-2xl px-5 pt-4 pb-5 font-nunito-bold">Language</ThemedText>
      <View className="px-5">
        {LANGUAGES.map((lang, index) => (
          <React.Fragment key={lang.code}>
            <TouchableOpacity
              className="flex-row items-center py-[18px] gap-[14px]"
              onPress={() => setSelected(lang.code)}
              activeOpacity={0.6}
            >
              <ThemedText className="text-[38px] leading-[46px]">{lang.flag}</ThemedText>
              <ThemedText className="flex-1 text-lg font-nunito">{lang.label}</ThemedText>
              <IconSymbol
                name={selected === lang.code ? 'checkmark.circle.fill' : 'circle'}
                size={26}
                color={selected === lang.code ? theme.text : '#C7C7CC'}
              />
            </TouchableOpacity>
            {index < LANGUAGES.length - 1 && (
              <View className="h-px ml-[66px]" style={{ backgroundColor: dividerColor }} />
            )}
          </React.Fragment>
        ))}
      </View>
    </ThemedView>
  );
}
