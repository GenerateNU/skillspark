import { SavedEventCard } from "@/components/SavedEventCard";
import { ThemedText } from "@/components/themed-text";
import { ThemedView } from "@/components/themed-view";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { Colors } from "@/constants/theme";
import {
  getGetSavedByGuardianIdQueryKey,
  Saved,
  useDeleteSaved,
} from "@skillspark/api-client";
import { useInfiniteSavedByGuardianId } from "@/hooks/use-infinite-saved";
import { useQueryClient } from "@tanstack/react-query";
import { useRouter } from "expo-router";
import React, { useMemo } from "react";
import {
  ActivityIndicator,
  Alert,
  FlatList,
  TouchableOpacity,
  View,
} from "react-native";
import { useSafeAreaInsets } from "react-native-safe-area-context";
import { useAuthContext } from "@/hooks/use-auth-context";
import { ErrorScreen } from "@/components/ErrorScreen";
import { useTranslation } from "react-i18next";

export default function SavedScreen() {
  const insets = useSafeAreaInsets();
  const router = useRouter();
  const theme = Colors.light;
  const queryClient = useQueryClient();
  const { guardianId } = useAuthContext();
  const { t: translate } = useTranslation();

  const {
    data,
    isLoading,
    error,
    fetchNextPage,
    hasNextPage,
    isFetchingNextPage,
  } = useInfiniteSavedByGuardianId(guardianId ?? undefined);

  const deleteSavedMutation = useDeleteSaved();

  const savedEvents: Saved[] = useMemo(
    () =>
      data?.pages.flatMap((page) =>
        Array.isArray(page.data) ? (page.data as Saved[]) : [],
      ) ?? [],
    [data],
  );

  if (!guardianId) {
    return <ErrorScreen message={translate("common.noGuardianId")} />;
  }

  if (isLoading) {
    return (
      <View className="flex-1 items-center justify-center gap-2">
        <ActivityIndicator size="large" />
        <ThemedText>{translate("common.loadingEvents")}</ThemedText>
      </View>
    );
  }

  if (error) {
    return (
      <ErrorScreen
        message={
          (error as { detail?: string }).detail ??
          translate("common.errorOccurred")
        }
      />
    );
  }

  const handleDeleteSaved = (savedId: string) => {
    Alert.alert(
      translate("saved.deleteTitle"),
      translate("saved.deleteConfirm"),
      [
        { text: translate("common.cancel"), style: "cancel" },
        {
          text: translate("payment.delete"),
          style: "destructive",
          onPress: async () => {
            deleteSavedMutation.mutate(
              { id: savedId },
              {
                onSuccess: () => {
                  queryClient.invalidateQueries({
                    queryKey: getGetSavedByGuardianIdQueryKey(guardianId),
                  });
                },
                onError: (err) =>
                  console.error("Failed to delete saved event", err),
              },
            );
          },
        },
      ],
    );
  };

  return (
    <ThemedView className="flex-1" style={{ paddingTop: insets.top }}>
      <View className="flex-row items-center justify-between px-5 py-[14px]">
        <TouchableOpacity
          onPress={() => router.navigate("/profile")}
          className="w-10 justify-center items-start"
          hitSlop={{ top: 10, bottom: 10, left: 10, right: 10 }}
        >
          <IconSymbol name="chevron.left" size={24} color={theme.text} />
        </TouchableOpacity>
        <ThemedText className="text-xl text-center font-nunito-bold">
          {translate("saved.title")}
        </ThemedText>
        <View className="w-10" />
      </View>
      <ThemedView className="flex-1">
        {savedEvents.length === 0 ? (
          <View className="flex-1 justify-center items-center p-5">
            <ThemedText className="text-center text-lg text-gray-500">
              {translate("saved.noEvents")}
            </ThemedText>
          </View>
        ) : (
          <FlatList
            data={savedEvents}
            keyExtractor={(item) => item.id}
            renderItem={({ item }) => (
              <SavedEventCard
                event={item.event}
                onBookmarkPress={() => handleDeleteSaved(item.id)}
              />
            )}
            contentContainerStyle={{ paddingTop: 10, paddingBottom: 20 }}
            showsVerticalScrollIndicator={false}
            onEndReached={() => {
              if (hasNextPage && !isFetchingNextPage) {
                fetchNextPage();
              }
            }}
            onEndReachedThreshold={0.3}
            ListFooterComponent={
              isFetchingNextPage ? (
                <View className="py-4 items-center">
                  <ActivityIndicator size="small" />
                </View>
              ) : null
            }
          />
        )}
      </ThemedView>
    </ThemedView>
  );
}
