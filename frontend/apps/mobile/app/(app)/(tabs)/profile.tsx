import React from 'react';
import { View, ScrollView, ActivityIndicator, useColorScheme } from 'react-native';
import { useSafeAreaInsets } from 'react-native-safe-area-context';
import { useRouter } from 'expo-router';
import { ThemedText } from '@/components/themed-text';
import { ThemedView } from '@/components/themed-view';
import { IconSymbol } from '@/components/ui/icon-symbol';
import { useGetGuardianById, useGetChildrenByGuardianId } from '@skillspark/api-client';
import { FamilyCard } from '@/components/FamilyCard';
import { ListItem } from '@/components/ListItem';
import { useAuthContext } from '@/hooks/use-auth-context';

export default function ProfileScreen() {
  const insets = useSafeAreaInsets();
  const colorScheme = useColorScheme();
  const router = useRouter();

  const listBackgroundColor = colorScheme === 'dark' ? '#1c1c1e' : '#F9FAFB';
  const borderColor = colorScheme === 'dark' ? '#3f3f46' : '#E5E7EB';

  const { guardianId } = useAuthContext();

  if (!guardianId) {
    // error state
  }

  const { data: guardianResponse, isLoading: guardianLoading } = useGetGuardianById(guardianId);
  const { data: childrenResponse, isLoading: familyLoading } = useGetChildrenByGuardianId(guardianId);
  const guardian = guardianResponse?.status === 200 ? guardianResponse.data : null;
  const children = childrenResponse?.status === 200 ? childrenResponse.data : [];

  if (guardianLoading || familyLoading) {
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
            Contact
          </ThemedText>
        </View>
        <View className="px-5 mb-4">
          <ThemedText className="text-base mb-2 font-nunito-semibold">Family</ThemedText>
          <View className="flex-row flex-wrap justify-between gap-[10px]">
            {children.length > 0 ? (
              children.map((child: any) => (
                <FamilyCard
                  key={child.id}
                  initials={child.name?.charAt(0)}
                  name={child.name}
                  date={`Born ${child.birth_year}`}
                />
              ))
            ) : (
              <ThemedText style={{ color: '#999', padding: 10 }}>No children found</ThemedText>
            )}
          </View>
        </View>
        <View className="px-5 mb-4">
          <ThemedText className="text-base mb-2 font-nunito-semibold">My Bookings</ThemedText>
          <View
            className="rounded-xl overflow-hidden border"
            style={{ backgroundColor: listBackgroundColor, borderColor }}
          >
            <ListItem label="Upcoming" />
            <ListItem label="Previous" />
            <ListItem label="Saved" isLast />
          </View>
        </View>
        <View className="px-5 mb-4">
          <ThemedText className="text-base mb-2 font-nunito-semibold">Preferences</ThemedText>
          <View
            className="rounded-xl overflow-hidden border"
            style={{ backgroundColor: listBackgroundColor, borderColor }}
          >
            <ListItem label="Payment" onPress={() => router.push('/payment')} />
            <ListItem
              label="Family Information"
              onPress={() => router.push('/family')}
            />
            <ListItem label="Settings" isLast onPress={() => router.push('/settings')} />
          </View>
        </View>
        <View className="h-5" />
      </ScrollView>
    </ThemedView>
  );
}
