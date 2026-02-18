import React from 'react';
import { StyleSheet, View, ScrollView, ActivityIndicator, useColorScheme } from 'react-native';
import { useSafeAreaInsets } from 'react-native-safe-area-context';
import { ThemedText } from '@/components/themed-text';
import { ThemedView } from '@/components/themed-view';
import { IconSymbol } from '@/components/ui/icon-symbol';
import { useGetGuardianById, useGetChildrenByGuardianId } from '@skillspark/api-client';
import { FamilyCard } from '@/components/FamilyCard';
import { ListItem } from '@/components/ListItem';

const GUARDIAN_ID = '88888888-8888-8888-8888-888888888888';

export default function ProfileScreen() {
  const insets = useSafeAreaInsets();
  const colorScheme = useColorScheme();

  const listBackgroundColor = colorScheme === 'dark' ? '#1c1c1e' : '#F9FAFB';
  const borderColor = colorScheme === 'dark' ? '#3f3f46' : '#E5E7EB';
  
  const { data: guardianResponse, isLoading: guardianLoading } = useGetGuardianById(GUARDIAN_ID);
  const { data: childrenResponse, isLoading: familyLoading } = useGetChildrenByGuardianId(GUARDIAN_ID);
  const guardian = guardianResponse?.status === 200 ? guardianResponse.data : null;
  const children = childrenResponse?.status === 200 ? childrenResponse.data : [];

  if (guardianLoading || familyLoading) {
    return (
      <ThemedView style={[styles.loadingContainer, { paddingTop: insets.top }]}>
        <ActivityIndicator size="large" />
      </ThemedView>
    );
  }

  return (
    <ThemedView style={[styles.container, { paddingTop: insets.top }]}>
      <ScrollView 
        showsVerticalScrollIndicator={false}
        contentContainerStyle={styles.scrollContent}
        bounces={false}
      >
        <View style={styles.profileSection}>
          <View style={[styles.avatarContainer, { backgroundColor: listBackgroundColor }]}>
            <IconSymbol name="photo" size={32} color="#9CA3AF" />
          </View>
          <ThemedText style={styles.userName}>
            {guardian?.name}
          </ThemedText>
          <ThemedText style={styles.userHandle}>
            @{guardian?.username}
          </ThemedText>
          <ThemedText style={styles.contactText}>Contact</ThemedText>
        </View>
        <View style={styles.section}>
          <ThemedText style={styles.sectionTitle}>Family</ThemedText>
          <View style={styles.familyRow}>
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
        <View style={styles.section}>
          <ThemedText style={styles.sectionTitle}>My Bookings</ThemedText>
          <View style={[styles.listGroup, { backgroundColor: listBackgroundColor, borderColor }]}>
            <ListItem label="Upcoming" />
            <ListItem label="Previous" />
            <ListItem label="Saved" isLast />
          </View>
        </View>
        <View style={styles.section}>
          <ThemedText style={styles.sectionTitle}>Preferences</ThemedText>
          <View style={[styles.listGroup, { backgroundColor: listBackgroundColor, borderColor }]}>
            <ListItem label="Payment" />
            <ListItem label="Family Information" />
            <ListItem label="Settings" isLast />
          </View>
        </View>
        <View style={{ height: 20 }} />
      </ScrollView>
    </ThemedView>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
  },
  loadingContainer: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
  },
  header: {
    paddingHorizontal: 20,
    paddingBottom: 5,
    paddingTop: 5,
  },
  headerTitle: {
    fontSize: 28,
    fontFamily: 'Archivo_700Bold',
  },
  scrollContent: {
    paddingTop: 10,
    paddingBottom: 20,
  },
  profileSection: {
    alignItems: 'center', 
    marginBottom: 20,
    marginTop: 5,
  },
  avatarContainer: {
    width: 72,
    height: 72,
    borderRadius: 36,
    justifyContent: 'center',
    alignItems: 'center',
    marginBottom: 10,
  },
  userName: {
    fontSize: 20,
    fontFamily: 'Archivo_600SemiBold',
    lineHeight: 24,
    marginBottom: 2,
    textAlign: 'center',
  },
  userHandle: {
    fontSize: 14,
    fontFamily: 'Archivo_400Regular',
    color: '#6B7280',
    lineHeight: 18,
    textAlign: 'center',
    marginBottom: 2,
  },
  contactText: {
    fontSize: 14,
    fontFamily: 'Archivo_400Regular',
    color: '#6B7280',
    lineHeight: 18,
    textAlign: 'center',
  },
  section: {
    paddingHorizontal: 20,
    marginBottom: 16,
  },
  sectionTitle: {
    fontSize: 16,
    marginBottom: 8,
    fontFamily: 'Archivo_600SemiBold',
  },
  familyRow: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    flexWrap: 'wrap',
    gap: 10,
  },
  listGroup: {
    borderRadius: 12,
    overflow: 'hidden',
    borderWidth: 1,
  },
});