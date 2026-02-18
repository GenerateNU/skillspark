import React from 'react';
import { StyleSheet, View, ScrollView, TouchableOpacity, ActivityIndicator, useColorScheme } from 'react-native';
import { useRouter } from 'expo-router';
import { useSafeAreaInsets } from 'react-native-safe-area-context';
import { ThemedText } from '@/components/themed-text';
import { ThemedView } from '@/components/themed-view';
import { IconSymbol } from '@/components/ui/icon-symbol';
import { useGetChildrenByGuardianId, useGetGuardianById } from '@skillspark/api-client';
import { Colors } from '@/constants/theme';

const GUARDIAN_ID = '88888888-8888-8888-8888-888888888888';

const TAG_COLORS = [
  { bg: '#E6F4EA', border: '#4CAF50', text: '#2E7D32' },
  { bg: '#FFF8E1', border: '#FFC107', text: '#F57F17' },
  { bg: '#FCE4EC', border: '#E91E63', text: '#880E4F' },
  { bg: '#E3F2FD', border: '#2196F3', text: '#0D47A1' },
  { bg: '#F3E5F5', border: '#9C27B0', text: '#4A148C' },
];

const MAX_VISIBLE_TAGS = 3;

function InterestTags({ interests }: { interests?: string[] | string }) {
  const tags: string[] = Array.isArray(interests)
    ? interests
    : typeof interests === 'string' && interests
    ? interests.split(',').map(s => s.trim()).filter(Boolean)
    : [];

  if (!tags.length) return null;

  const visible = tags.slice(0, MAX_VISIBLE_TAGS);
  const overflow = tags.length - MAX_VISIBLE_TAGS;

  return (
    <View style={styles.tagsRow}>
      {visible.map((tag, i) => {
        const c = TAG_COLORS[i % TAG_COLORS.length];
        return (
          <View key={tag} style={[styles.tag, { backgroundColor: c.bg, borderColor: c.border }]}>
            <IconSymbol name="camera.filters" size={13} color={c.border} />
            <ThemedText style={[styles.tagText, { color: c.text }]}>{tag}</ThemedText>
          </View>
        );
      })}
      {overflow > 0 && (
        <ThemedText style={styles.overflowText}>+{overflow}</ThemedText>
      )}
    </View>
  );
}

export default function FamilyListScreen() {
  const router = useRouter();
  const insets = useSafeAreaInsets();
  const colorScheme = useColorScheme();
  const theme = Colors[colorScheme ?? 'light'];

  const { data: guardianResponse, isLoading: guardianLoading } = useGetGuardianById(GUARDIAN_ID);
  const { data: childrenResponse, isLoading: childrenLoading } = useGetChildrenByGuardianId(GUARDIAN_ID);

  const guardian = guardianResponse?.status === 200 ? guardianResponse.data : null;
  const children = childrenResponse?.status === 200 ? childrenResponse.data : [];

  // const handleAddChild = () => router.push('/family/manage');

  // const handleEditChild = (child: any) => {
  //   router.push({
  //     pathname: '/family/manage',
  //     params: {
  //       id: child.id,
  //       name: child.name,
  //       birth_year: child.birth_year,
  //       birth_month: child.birth_month,
  //       school_id: child.school_id,
  //       interests: child.interests,
  //     },
  //   });
  // };

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
        <TouchableOpacity onPress={() => router.back()} style={styles.backButton} hitSlop={{ top: 10, bottom: 10, left: 10, right: 10 }}>
          <IconSymbol name="chevron.left" size={24} color={theme.text} />
        </TouchableOpacity>
        <ThemedText style={styles.headerTitle}>Family Information</ThemedText>
        <View style={styles.headerRight} />
      </View>

      <ScrollView contentContainerStyle={styles.content} showsVerticalScrollIndicator={false}>
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
        <View style={styles.sectionHeaderRow}>
          <ThemedText style={styles.sectionTitle}>Child Profile</ThemedText>
          <TouchableOpacity onPress={() => {}}>
            <ThemedText style={[styles.addLink, { color: theme.tint }]}>add profile +</ThemedText>
          </TouchableOpacity>
        </View>
        {children.length === 0 && (
          <ThemedText style={styles.emptyText}>No child profiles added yet.</ThemedText>
        )}
        {children.map((child: any, idx: number) => (
          <React.Fragment key={child.id}>
            <TouchableOpacity style={styles.row} onPress={() => {}} activeOpacity={0.7}>
              <View style={styles.initialsCircle}>
                <ThemedText style={styles.initialsText}>
                  {child.name?.slice(0, 2).toUpperCase() || '??'}
                </ThemedText>
              </View>
              <View style={styles.rowInfo}>
                <ThemedText style={styles.rowTitle}>{child.name}</ThemedText>
                <ThemedText style={styles.rowSub}>
                  {child.birth_month ? `${child.birth_month}, ` : ''}{child.birth_year}
                </ThemedText>
                {child.interests && (
                  <ThemedText style={styles.rowSub}>Interests</ThemedText>
                )}
                <InterestTags interests={child.interests} />
              </View>
              <IconSymbol name="chevron.right" size={18} color="#9CA3AF" />
            </TouchableOpacity>
            {idx < children.length - 1 && <View style={styles.divider} />}
          </React.Fragment>
        ))}
        <View style={styles.divider} />
        <View style={styles.sectionHeaderRow}>
          <ThemedText style={styles.sectionTitle}>Emergency Contact</ThemedText>
          <TouchableOpacity>
            <ThemedText style={[styles.addLink, { color: theme.tint }]}>add contact +</ThemedText>
          </TouchableOpacity>
        </View>
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
  backButton: { width: 40, justifyContent: 'center', alignItems: 'flex-start' },
  headerTitle: { fontSize: 18, fontFamily: 'Archivo_600SemiBold', textAlign: 'center' },
  headerRight: { width: 40 },

  content: { paddingHorizontal: 20, paddingTop: 8 },

  divider: { height: 1, backgroundColor: '#E5E7EB', marginVertical: 4 },

  sectionHeaderRow: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
    paddingVertical: 14,
  },
  sectionTitle: { fontSize: 17, fontFamily: 'Archivo_700Bold' },
  addLink: { fontSize: 14, fontFamily: 'Archivo_500Medium' },

  row: {
    flexDirection: 'row',
    alignItems: 'flex-start',
    paddingVertical: 12,
    gap: 12,
  },
  iconAvatar: { width: 44, height: 44, justifyContent: 'center', alignItems: 'center' },
  initialsCircle: {
    width: 44,
    height: 44,
    borderRadius: 22,
    borderWidth: 1.5,
    borderColor: '#9CA3AF',
    justifyContent: 'center',
    alignItems: 'center',
  },
  initialsText: { fontSize: 15, fontFamily: 'Archivo_600SemiBold' },
  rowInfo: { flex: 1, gap: 2 },
  rowTitle: { fontSize: 16, fontFamily: 'Archivo_600SemiBold' },
  rowSub: { fontSize: 13, color: '#6B7280', fontFamily: 'Archivo_400Regular' },

  tagsRow: { flexDirection: 'row', flexWrap: 'wrap', gap: 6, marginTop: 4 },
  tag: {
    flexDirection: 'row',
    alignItems: 'center',
    gap: 4,
    paddingHorizontal: 8,
    paddingVertical: 4,
    borderRadius: 20,
    borderWidth: 1,
  },
  tagText: { fontSize: 12, fontFamily: 'Archivo_500Medium' },
  overflowText: { fontSize: 13, color: '#6B7280', fontFamily: 'Archivo_500Medium', alignSelf: 'center' },

  emptyText: { color: '#9CA3AF', fontSize: 14, paddingBottom: 12 },
});