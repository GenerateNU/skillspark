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


const GUARDIAN_ID = '88888888-8888-8888-8888-888888888888';

export default function SavedScreen() {
  const insets = useSafeAreaInsets();
  const colorScheme = useColorScheme();
  const router = useRouter();

  const { data: savedResponse, isLoading: savedLoading } =
} 