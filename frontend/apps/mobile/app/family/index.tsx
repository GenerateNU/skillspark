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

export default function FamilyListScreen() {
  const router = useRouter();
  const insets = useSafeAreaInsets();
  const colorScheme = useColorScheme();
  const theme = Colors[colorScheme ?? 'light'];
  
  const { data: guardianResponse, isLoading: guardianLoading } = useGetGuardianById(GUARDIAN_ID);
  const { data: childrenResponse, isLoading: childrenLoading } = useGetChildrenByGuardianId(GUARDIAN_ID);
  
  const guardian = guardianResponse?.status === 200 ? guardianResponse.data : null;
  const children = childrenResponse?.status === 200 ? childrenResponse.data : [];

  const handleAddChild = () => {
    router.push('/family/manage');
  };

  const handleEditChild = (child: any) => {
    router.push({
      pathname: '/family/manage',
      params: { 
        id: child.id,
        name: child.name,
        birth_year: child.birth_year,
        birth_month: child.birth_month,
        school_id: child.school_id,
        interests: child.interests
      }
    });
  };

  const cardBackgroundColor = colorScheme === 'dark' ? '#27272a' : '#F3F4F6';
  const borderColor = colorScheme === 'dark' ? '#3f3f46' : '#E5E7EB';

  if (guardianLoading || childrenLoading) {
    return (
      <ThemedView style={styles.loadingContainer}>
        <ActivityIndicator size="large" />
      </ThemedView>
    );
  }

  return (
    <ThemedView style={[styles.container, { paddingTop: insets.top }]}>
      
      {/* Custom Header */}
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
        
        {/* 1. User Info (Top) */}
        <View style={styles.section}>
          <ThemedText style={styles.sectionTitle}>My Profile</ThemedText>
          <View style={[styles.card, { backgroundColor: cardBackgroundColor, borderColor }]}>
            <View style={styles.cardRow}>
              <View style={[styles.avatarCircle, { backgroundColor: theme.tint }]}>
                 {/* Initials or Icon */}
                <ThemedText style={styles.avatarText}>
                  {guardian?.name ? guardian.name.charAt(0).toUpperCase() : 'U'}
                </ThemedText>
              </View>
              <View style={styles.cardInfo}>
                <ThemedText style={styles.cardTitle}>{guardian?.name || 'User'}</ThemedText>
                <ThemedText style={styles.cardSubtitle}>Guardian</ThemedText>
              </View>
              {/* Optional Edit Button for User */}
              {/* <TouchableOpacity>
                <IconSymbol name="pencil" size={20} color={theme.icon} />
              </TouchableOpacity> */}
            </View>
          </View>
        </View>

        {/* 2. Family Members (Middle) */}
        <View style={styles.section}>
          <View style={styles.sectionHeaderRow}>
            <ThemedText style={styles.sectionTitle}>Family Members</ThemedText>
          </View>
          
          <View style={styles.listContainer}>
            {children.map((child: any) => (
              <TouchableOpacity 
                key={child.id} 
                onPress={() => handleEditChild(child)}
                activeOpacity={0.7}
                style={[styles.card, { backgroundColor: cardBackgroundColor, borderColor }]}
              >
                <View style={styles.cardRow}>
                  <View style={[styles.initialsCircle, { borderColor: theme.tint }]}>
                    <ThemedText style={[styles.initialsText, { color: theme.tint }]}>
                      {child.name?.charAt(0)}
                    </ThemedText>
                  </View>
                  <View style={styles.cardInfo}>
                    <ThemedText style={styles.cardTitle}>{child.name}</ThemedText>
                    <ThemedText style={styles.cardSubtitle}>Born {child.birth_year}</ThemedText>
                  </View>
                  <IconSymbol name="chevron.right" size={20} color="#9CA3AF" />
                </View>
              </TouchableOpacity>
            ))}
            
            {children.length === 0 && (
              <View style={styles.emptyState}>
                <ThemedText style={{ color: '#9CA3AF' }}>No family members added yet.</ThemedText>
              </View>
            )}
          </View>

          <TouchableOpacity 
            style={[styles.addButton, { backgroundColor: theme.tint }]}
            onPress={handleAddChild}
          >
            <IconSymbol name="plus" size={20} color="#FFF" />
            <ThemedText style={styles.addButtonText}>Add Family Member</ThemedText>
          </TouchableOpacity>
        </View>

        {/* 3. Emergency Contact (Bottom) */}
        <View style={styles.section}>
          <ThemedText style={styles.sectionTitle}>Emergency Contact</ThemedText>
          <View style={[styles.card, { backgroundColor: cardBackgroundColor, borderColor, padding: 16 }]}>
            <View style={styles.cardRow}>
              <View style={[styles.avatarCircle, { backgroundColor: '#EF4444' }]}>
                <ThemedText style={styles.avatarText}>MS</ThemedText>
              </View>
              <View style={styles.cardInfo}>
                <ThemedText style={styles.cardTitle}>Martha Smith</ThemedText>
                <ThemedText style={styles.cardSubtitle}>Grandmother</ThemedText>
              </View>
              <TouchableOpacity style={styles.phoneButton}>
                <IconSymbol name="phone.fill" size={20} color={theme.tint} />
              </TouchableOpacity>
            </View>
            <View style={styles.divider} />
            <View style={styles.contactRow}>
              <IconSymbol name="phone" size={16} color="#9CA3AF" />
              <ThemedText style={styles.contactText}>(555) 123-4567</ThemedText>
            </View>
          </View>
        </View>

        <View style={{ height: 40 }} />
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
    flexDirection: 'row',
    alignItems: 'center',
    justifyContent: 'space-between',
    paddingHorizontal: 20,
    paddingVertical: 12,
  },
  backButton: {
    width: 40,
    justifyContent: 'center',
    alignItems: 'flex-start',
  },
  headerTitle: {
    fontSize: 18,
    fontFamily: 'Archivo_600SemiBold',
    textAlign: 'center',
  },
  headerRight: {
    width: 40,
  },
  content: {
    padding: 20,
    paddingTop: 10,
  },
  section: {
    marginBottom: 24,
  },
  sectionHeaderRow: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
    marginBottom: 12,
  },
  sectionTitle: {
    fontSize: 16,
    color: '#6B7280',
    fontFamily: 'Archivo_600SemiBold',
    marginBottom: 8,
    textTransform: 'uppercase',
    letterSpacing: 0.5,
  },
  card: {
    borderRadius: 16,
    padding: 16,
    borderWidth: 1,
  },
  cardRow: {
    flexDirection: 'row',
    alignItems: 'center',
  },
  cardInfo: {
    flex: 1,
    paddingHorizontal: 12,
  },
  cardTitle: {
    fontSize: 16,
    fontFamily: 'Archivo_600SemiBold',
  },
  cardSubtitle: {
    fontSize: 14,
    color: '#6B7280',
    fontFamily: 'Archivo_400Regular',
  },
  avatarCircle: {
    width: 48,
    height: 48,
    borderRadius: 24,
    justifyContent: 'center',
    alignItems: 'center',
  },
  avatarText: {
    color: '#FFF',
    fontSize: 18,
    fontFamily: 'Archivo_700Bold',
  },
  initialsCircle: {
    width: 48,
    height: 48,
    borderRadius: 24,
    borderWidth: 1.5,
    justifyContent: 'center',
    alignItems: 'center',
  },
  initialsText: {
    fontSize: 18,
    fontFamily: 'Archivo_600SemiBold',
  },
  phoneButton: {
    padding: 8,
  },
  divider: {
    height: 1,
    backgroundColor: '#E5E7EB',
    marginVertical: 12,
  },
  contactRow: {
    flexDirection: 'row',
    alignItems: 'center',
    gap: 8,
  },
  contactText: {
    fontSize: 14,
    fontFamily: 'Archivo_500Medium',
  },
  listContainer: {
    gap: 12,
    marginBottom: 16,
  },
  emptyState: {
    alignItems: 'center',
    padding: 20,
    borderWidth: 1,
    borderColor: '#E5E7EB',
    borderRadius: 12,
    borderStyle: 'dashed',
    marginBottom: 16,
  },
  addButton: {
    flexDirection: 'row',
    alignItems: 'center',
    justifyContent: 'center',
    padding: 16,
    borderRadius: 12,
    gap: 8,
  },
  addButtonText: {
    color: '#FFF',
    fontSize: 16,
    fontFamily: 'Archivo_600SemiBold',
  },
});