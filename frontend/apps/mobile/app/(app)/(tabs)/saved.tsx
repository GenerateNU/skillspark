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
import { useAuthContext } from '@/hooks/use-auth-context';
import { ErrorScreen } from '@/components/ErrorScreen';

export default function SavedScreen() {
    const insets = useSafeAreaInsets();
    const colorScheme = useColorScheme();
    const router = useRouter();
    const theme = Colors[colorScheme ?? 'light'];

    const queryClient = useQueryClient();

    const { guardianId } = useAuthContext();

    const { data: response, isLoading, error } = useGetSavedByGuardianId(guardianId!, undefined, {
      query: {
        enabled: !!guardianId,
      }
    });
    const deleteSavedMutation  = useDeleteSaved();

    if (!guardianId) {
      return <ErrorScreen message="Illegal state: no guardian ID retrieved" />;
    }

    if (isLoading) {
        return (
        <View style={{ flex: 1, alignItems: "center", justifyContent: "center", gap: 8 }}>
            <ActivityIndicator size="large" />
            <ThemedText>Loading events...</ThemedText>
        </View>
        );
    }

    if (error) {
        return (
        <View style={{ flex: 1, alignItems: "center", justifyContent: "center", padding: 16 }}>
            <ThemedText style={{ color: "#EF4444", fontWeight: "600" }}>Error loading events</ThemedText>
            <ThemedText>{error.detail || "An error occurred"}</ThemedText>
        </View>
        );
    }

    if (!response || !Array.isArray(response.data)) {
        return (
        <View style={{ flex: 1, alignItems: "center", justifyContent: "center", padding: 16 }}>
            <ThemedText>No events available</ThemedText>
        </View>
        );
    }

    const savedEvents: Saved[] = response.status === 200 && Array.isArray(response.data)
    ? response.data
    : [];

    const handleDeleteSaved = (savedId: string) => {
      Alert.alert(
        'Delete saved event',
        'Are you sure you want remove this saved event?',
        [
          { text: 'Cancel', style: 'cancel' },
          {
            text: 'Delete', style: 'destructive',
            onPress: async () => {
              deleteSavedMutation.mutate(
              { id: savedId }, 
              {
                  onSuccess: () => {
                  queryClient.invalidateQueries({
                      queryKey: getGetSavedByGuardianIdQueryKey(guardianId as string)
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
        <ThemedText className="text-xl text-center font-nunito-bold">Saved</ThemedText>
        <View className="w-10" />
      </View>
        <ThemedView className="flex-1" style={{ paddingTop: insets.top }}>
    {savedEvents.length === 0 ? (
      <View style={{ flex: 1, justifyContent: 'center', alignItems: 'center', padding: 20 }}>
        <ThemedText className="text-center text-lg text-gray-500">
          You have no saved events.
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