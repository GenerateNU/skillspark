import React from 'react';
import { StyleSheet, View, ScrollView, TouchableOpacity, ActivityIndicator, useColorScheme } from 'react-native';
import { useRouter } from 'expo-router';
import { useSafeAreaInsets } from 'react-native-safe-area-context';
import { ThemedText } from '@/components/themed-text';
import { ThemedView } from '@/components/themed-view';
import { IconSymbol } from '@/components/ui/icon-symbol';
import { useGetChildrenByGuardianId, useGetGuardianById } from '@skillspark/api-client';
import { Colors } from '@/constants/theme';
import { ChildListItem } from '@/components/ChildListItem';
import { SectionHeader } from '@/components/SectionHeader';

const GUARDIAN_ID = '88888888-8888-8888-8888-888888888888';

export default function FamilyListScreen() {
  const router = useRouter();
  const insets = useSafeAreaInsets();
  const colorScheme = useColorScheme();
  const theme = Colors[colorScheme ?? 'light'];

  const { data: guardianResponse, isLoading: guardianLoading } = useGetGuardianById(GUARDIAN_ID);
  const { data: childrenResponse, isLoading: childrenLoading } = useGetChildrenByGuardianId(GUARDIAN_ID);

  const guardian = guardianResponse?.status === 200 ? guardianResponse.data : null;
  const children = childrenResponse?.status === 200 ? childrenResponse.data : [];

  if (guardianLoading || childrenLoading) {
    return (
      <ThemedView style={styles.loadingContainer}>
        <ActivityIndicator size="large" />
      </ThemedView>
    );
  }

  return (
    <ThemedView style={[styles.container, { paddingTop: insets.top }]}>
      <View style={styles.header}>
        <TouchableOpacity 
          onPress={() => router.back()} 
          style={styles.backButton} 
          hitSlop={{ top: 10, bottom: 10, left: 10, right: 10 }}
        >
          <IconSymbol name="chevron.left" size={24} color={theme.text} />
        </TouchableOpacity>
        <ThemedText style={styles.headerTitle}>Family Information</ThemedText>
        <View style={styles.headerRight} />
      </View>

      <ScrollView contentContainerStyle={styles.content} showsVerticalScrollIndicator={false}>
        {/* Guardian Profile Row */}
        <TouchableOpacity style={styles.row} activeOpacity={0.7}>
          <View style={styles.iconAvatar}>
            <IconSymbol name="person.circle" size={40} color={theme.text} />
          </View>
          <View style={styles.rowInfo}>
            <ThemedText style={styles.rowTitle}>{guardian?.name}</ThemedText>
            <ThemedText style={styles.rowSub}>@{guardian?.username}</ThemedText>
            <ThemedText style={styles.rowSub}>{guardian?.email}</ThemedText>
          </View>
          <IconSymbol name="chevron.right" size={18} color="#9CA3AF" />
        </TouchableOpacity>
        
        <View style={styles.divider} />
        
        <SectionHeader 
          title="Child Profile" 
          actionLabel="add profile +" 
          onAction={() => {}} 
        />
        
        {children.length === 0 && (
          <ThemedText style={styles.emptyText}>No child profiles added yet.</ThemedText>
        )}
        
        {children.map((child: any, idx: number) => (
          <React.Fragment key={child.id}>
            <ChildListItem child={child} onPress={() => {}} />
            {idx < children.length - 1 && <View style={styles.divider} />}
          </React.Fragment>
        ))}
        <View style={styles.divider} />
        <SectionHeader 
          title="Emergency Contact" 
          actionLabel="add contact +" 
          onAction={() => {}} 
        />
        <TouchableOpacity style={styles.row} activeOpacity={0.7}>
          <View style={styles.iconAvatar}>
            <IconSymbol name="person.circle" size={40} color={theme.text} />
          </View>
          <View style={styles.rowInfo}>
            <ThemedText style={styles.rowTitle}>Martha Smith</ThemedText>
            <ThemedText style={styles.rowSub}>(555) 123-4567</ThemedText>
          </View>
          <IconSymbol name="chevron.right" size={18} color="#9CA3AF" />
        </TouchableOpacity>
        <View style={{ height: 40 }} />
      </ScrollView>
    </ThemedView>
  );
}

const styles = StyleSheet.create({
  container: { flex: 1 },
  loadingContainer: { flex: 1, justifyContent: 'center', alignItems: 'center' },
  header: {
    flexDirection: 'row',
    alignItems: 'center',
    justifyContent: 'space-between',
    paddingHorizontal: 20,
    paddingVertical: 12,
  },
  backButton: { 
    width: 40, 
    justifyContent: 'center',
    alignItems: 'flex-start' 
  },
  headerTitle: { 
    fontSize: 18, 
    fontFamily: 'Archivo_600SemiBold', 
    textAlign: 'center' 
  },
  headerRight: { 
    width: 40 
  },
  content: { 
    paddingHorizontal: 20, 
    paddingTop: 8 
  },
  divider: { 
    height: 1, 
    backgroundColor: '#E5E7EB', 
    marginVertical: 4 
  },
  row: {
    flexDirection: 'row',
    alignItems: 'flex-start',
    paddingVertical: 12,
    gap: 12,
  },
  iconAvatar: { 
    width: 44, 
    height: 44, 
    justifyContent: 'center', 
    alignItems: 'center' 
  },
  rowInfo: { 
    flex: 1, 
    gap: 2 
  },
  rowTitle: { 
    fontSize: 16, 
    fontFamily: 'Archivo_600SemiBold' 
  },
  rowSub: { 
    fontSize: 13, 
    color: '#6B7280', 
    fontFamily: 'Archivo_400Regular' 
  },
  emptyText: { 
    color: '#9CA3AF', 
    fontSize: 14, 
    paddingBottom: 12 
  },
});