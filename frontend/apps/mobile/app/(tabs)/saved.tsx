import { SavedEventCard } from '@/components/SavedEventCard';
import { ThemedText } from '@/components/themed-text';
import { ThemedView } from '@/components/themed-view';
import { IconSymbol } from '@/components/ui/icon-symbol';
import { Colors } from '@/constants/theme';
import { getGetSavedByGuardianIdQueryKey, Saved, useDeleteSaved, useGetSavedByGuardianId } from '@skillspark/api-client';
import { useQueryClient } from '@tanstack/react-query';
import { useRouter } from 'expo-router';
import React from 'react';
import { ActivityIndicator, Alert, FlatList, TouchableOpacity, useColorScheme, View } from 'react-native';
import { useSafeAreaInsets } from 'react-native-safe-area-context';
import { useGuardian } from '@/hooks/use-guardian';
import { useTranslation } from 'react-i18next';

export default function SavedScreen() {

    const insets = useSafeAreaInsets();
    const colorScheme = useColorScheme();
    const router = useRouter();
    const theme = Colors[colorScheme ?? 'light'];
    const { guardianId } = useGuardian();
    const { t: translate } = useTranslation();

    const queryClient = useQueryClient();

    const { data: response, isLoading, error } = useGetSavedByGuardianId(guardianId);
    const deleteSavedMutation  = useDeleteSaved();

    if (isLoading) {
        return (
        <View style={{ flex: 1, alignItems: "center", justifyContent: "center", gap: 8 }}>
            <ActivityIndicator size="large" />
            <ThemedText>{translate('common.loadingEvents')}</ThemedText>
        </View>
        );
    }

    if (error) {
      console.log(error)
        return (
        <View style={{ flex: 1, alignItems: "center", justifyContent: "center", padding: 16 }}>
            <ThemedText style={{ color: "#EF4444", fontWeight: "600" }}>{translate('common.errorLoadingEvents')}</ThemedText>
            <ThemedText>{error.detail || translate('common.errorOccurred')}</ThemedText>
        </View>
        );
    }

    if (!response || !Array.isArray(response.data)) {
        return (
        <View style={{ flex: 1, alignItems: "center", justifyContent: "center", padding: 16 }}>
            <ThemedText>{translate('common.noEventsAvailable')}</ThemedText>
        </View>
        );
    }

    const savedEvents: Saved[] = response.status === 200 && Array.isArray(response.data)
    ? response.data
    : [];

    const handleDeleteSaved = (savedId: string) => {
      Alert.alert(
        translate('saved.deleteTitle'),
        translate('saved.deleteConfirm'),
        [
          { text: translate('common.cancel'), style: 'cancel' },
          {
            text: translate('payment.delete'), style: 'destructive',
            onPress: async () => {
              deleteSavedMutation.mutate(
              { id: savedId }, 
              {
                  onSuccess: () => {
                  queryClient.invalidateQueries({
                      queryKey: getGetSavedByGuardianIdQueryKey(guardianId)
                  });
                  },
                  onError: (err) => console.error('Failed to delete saved event', err),
              }
              );
            }
          }
        ]
      );
    };

    return (
        <ThemedView className="flex-1" style={{ paddingTop: insets.top }}>
      <View className="flex-row items-center justify-between px-5 py-[14px]">
        <TouchableOpacity
          onPress={() => router.navigate('/profile')}
          className="w-10 justify-center items-start"
          hitSlop={{ top: 10, bottom: 10, left: 10, right: 10 }}
        >
          <IconSymbol name="chevron.left" size={24} color={theme.text} />
        </TouchableOpacity>
        <ThemedText className="text-xl text-center font-nunito-bold">{translate('saved.title')}</ThemedText>
        <View className="w-10" />
      </View>
        <ThemedView className="flex-1" style={{ paddingTop: insets.top }}>
    {savedEvents.length === 0 ? (
      <View style={{ flex: 1, justifyContent: 'center', alignItems: 'center', padding: 20 }}>
        <ThemedText className="text-center text-lg text-gray-500">
          {translate('saved.noEvents')}
        </ThemedText>
      </View>
    ) : (
      <FlatList
        data={savedEvents} // Saved[] array
        keyExtractor={(item) => item.id}
        renderItem={({ item }) => (
          <SavedEventCard
            event={item.event} // Event object inside Saved
            onBookmarkPress={() => handleDeleteSaved(item.id)} // Saved.id
          />
        )}
        contentContainerStyle={{ paddingTop: 10, paddingBottom: 20 }}
        showsVerticalScrollIndicator={false}
      />
    )}
  </ThemedView>
    </ThemedView>

    )

} 