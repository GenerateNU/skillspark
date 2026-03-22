import React from 'react';
import { View, ScrollView, ActivityIndicator, useColorScheme } from 'react-native';
import { useSafeAreaInsets } from 'react-native-safe-area-context';
import { useRouter } from 'expo-router';
import { ThemedText } from '@/components/themed-text';
import { ThemedView } from '@/components/themed-view';
import { IconSymbol } from '@/components/ui/icon-symbol';
import { FamilyCard } from '@/components/FamilyCard';
import { ListItem } from '@/components/ListItem';
import { useTranslation } from 'react-i18next';
import { useGuardian } from '@/hooks/use-guardian';

export default function ProfileScreen() {
  const insets = useSafeAreaInsets();
  const colorScheme = useColorScheme();
  const router = useRouter();
  const {t} = useTranslation();

  const listBackgroundColor = colorScheme === 'dark' ? '#1c1c1e' : '#F9FAFB';
  const borderColor = colorScheme === 'dark' ? '#3f3f46' : '#E5E7EB';

  const { guardian, children, isLoading } = useGuardian();

  if (isLoading) {
    return (
      <ThemedView className="flex-1 items-center justify-center" style={{ paddingTop: insets.top }}>
        <ActivityIndicator size="large" />
      </ThemedView>
    );
  }

  return (
    <ThemedView className="flex-1" style={{ paddingTop: insets.top }}>
      <ScrollView
        showsVerticalScrollIndicator={false}
        contentContainerStyle={{ paddingTop: 10, paddingBottom: 20 }}
        bounces={false}
      >
        <View className="items-center mb-5 mt-[5px]">
          <View
            className="w-[72px] h-[72px] rounded-full items-center justify-center mb-[10px]"
            style={{ backgroundColor: listBackgroundColor }}
          >
            <IconSymbol name="photo" size={32} color="#9CA3AF" />
          </View>
          <ThemedText className="text-xl leading-6 mb-[2px] text-center font-nunito-semibold">
            {guardian?.name}
          </ThemedText>
          <ThemedText className="text-sm text-[#6B7280] leading-[18px] text-center mb-[2px] font-nunito">
            @{guardian?.username}
          </ThemedText>
          <ThemedText className="text-sm text-[#6B7280] leading-[18px] text-center font-nunito">
            {t('profile.contact')}
          </ThemedText>
        </View>
        <View className="px-5 mb-4">
          <ThemedText className="text-base mb-2 font-nunito-semibold">{t('profile.family')}</ThemedText>
          <View className="flex-row flex-wrap justify-between gap-[10px]">
            {children.length > 0 ? (
              children.map((child: any) => (
                <FamilyCard
                  key={child.id}
                  initials={child.name?.charAt(0)}
                  name={child.name}
                  date={t('profile.born', { year: child.birth_year })}
                  
                />
              ))
            ) : (
              <ThemedText style={{ color: '#999', padding: 10 }}>{t('common.noChildrenFound')}</ThemedText>
            )}
          </View>
        </View>
        <View className="px-5 mb-4">
          <ThemedText className="text-base mb-2 font-nunito-semibold">{t('profile.myBookings')}</ThemedText>
          <View
            className="rounded-xl overflow-hidden border"
            style={{ backgroundColor: listBackgroundColor, borderColor }}
          >
            <ListItem label="Saved" isLast onPress={() => router.push('/saved')}/>
          </View>
        </View>
        <View className="px-5 mb-4">
          <ThemedText className="text-base mb-2 font-nunito-semibold">{t('profile.preferences')}</ThemedText>
          <View
            className="rounded-xl overflow-hidden border"
            style={{ backgroundColor: listBackgroundColor, borderColor }}
          >
            <ListItem label={t('profile.payment')} onPress={() => router.push('/payment')} />
            <ListItem
              label={t('profile.familyInformation')}
              onPress={() => router.push('/family')}
            />
            <ListItem label={t('profile.settings')} isLast onPress={() => router.push('/settings')} />
          </View>
        </View>
        <View className="h-5" />
      </ScrollView>
    </ThemedView>
  );
}
