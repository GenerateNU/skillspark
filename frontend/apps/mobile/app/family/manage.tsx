import React, { useState } from 'react';
import {
  StyleSheet,
  View,
  TextInput,
  TouchableOpacity,
  Alert,
  useColorScheme,
  ScrollView,
  KeyboardAvoidingView,
  Platform,
} from 'react-native';
import { Stack, useRouter, useLocalSearchParams } from 'expo-router';
import { useSafeAreaInsets } from 'react-native-safe-area-context';
import { ThemedText } from '@/components/themed-text';
import { ThemedView } from '@/components/themed-view';
import { Colors } from '@/constants/theme';
import { IconSymbol } from '@/components/ui/icon-symbol';
import { useCreateChild, useUpdateChild, useDeleteChild } from '@skillspark/api-client';

const GUARDIAN_ID = '88888888-8888-8888-8888-888888888888';

const INTEREST_OPTIONS = [
  'Soccer', 'Basketball', 'Baseball', 'Swimming', 'Tennis',
  'Music', 'Art', 'Dance', 'Drama', 'Coding',
  'Reading', 'Science', 'Math', 'Chess', 'Cooking',
];

const TAG_COLORS = [
  { bg: '#E6F4EA', border: '#4CAF50', text: '#2E7D32' },
  { bg: '#FFF8E1', border: '#FFC107', text: '#F57F17' },
  { bg: '#FCE4EC', border: '#E91E63', text: '#880E4F' },
  { bg: '#E3F2FD', border: '#2196F3', text: '#0D47A1' },
  { bg: '#F3E5F5', border: '#9C27B0', text: '#4A148C' },
];

const MONTHS = [
  'January','February','March','April','May','June',
  'July','August','September','October','November','December'
];

const YEARS = Array.from({ length: 20 }, (_, i) => String(new Date().getFullYear() - i));

export default function ManageChildScreen() {
  const router = useRouter();
  const params = useLocalSearchParams();
  const colorScheme = useColorScheme();
  const insets = useSafeAreaInsets();
  const theme = Colors[colorScheme ?? 'light'];

  const isEditing = !!params.id;

  // Initial State Setup
  const [firstName, setFirstName] = useState(
    params.name ? (params.name as string).split(' ')[0] : ''
  );
  const [lastName, setLastName] = useState(
    params.name ? (params.name as string).split(' ').slice(1).join(' ') : ''
  );

  // Convert numeric month (1-12) to String Name if editing
  const initialMonthStr = params.birth_month
    ? MONTHS[parseInt(params.birth_month as string) - 1]
    : '';

  const [birthMonth, setBirthMonth] = useState(initialMonthStr);
  const [birthYear, setBirthYear] = useState(params.birth_year as string || '');
  const [schoolId] = useState(params.school_id as string || '');

  const initialInterests = Array.isArray(params.interests)
    ? params.interests
    : params.interests
    ? (params.interests as string).split(',').map(s => s.trim()).filter(Boolean)
    : [];
  const [interests, setInterests] = useState<string[]>(initialInterests);

  // Interest picker state
  const [showPicker, setShowPicker] = useState(false);
  const [searchQuery, setSearchQuery] = useState('');
  const [pendingInterests, setPendingInterests] = useState<string[]>([]);

  // Dropdown state
  const [showMonthDrop, setShowMonthDrop] = useState(false);
  const [showYearDrop, setShowYearDrop] = useState(false);

  const [isSubmitting, setIsSubmitting] = useState(false);

  const createChildMutation = useCreateChild();
  const updateChildMutation = useUpdateChild();
  const deleteChildMutation = useDeleteChild();

  // Dynamic Colors based on Theme
  const inputBg = colorScheme === 'dark' ? '#27272a' : '#F3F4F6';
  const dropdownPopupBg = colorScheme === 'dark' ? '#1c1c1e' : '#FFFFFF';
  const borderColor = colorScheme === 'dark' ? '#3f3f46' : '#E5E7EB';
  const placeholderColor = '#9CA3AF';

  const handleSave = async () => {
    if (!firstName || !birthYear || !birthMonth || !schoolId) {
      Alert.alert('Error', 'Please fill in all required fields (Name, Birth Date, School ID)');
      return;
    }
    const name = [firstName, lastName].filter(Boolean).join(' ');
    setIsSubmitting(true);
    try {
      const childData = {
        name,
        birth_year: parseInt(birthYear, 10),
        birth_month: MONTHS.indexOf(birthMonth) + 1,
        guardian_id: GUARDIAN_ID,
        school_id: schoolId,
        interests,
      };
      if (isEditing) {
        await updateChildMutation.mutateAsync({ id: params.id as string, data: childData });
      } else {
        await createChildMutation.mutateAsync({ data: childData });
      }
      router.back();
    } catch (error) {
      console.error(error);
      Alert.alert('Error', 'Failed to save. Please try again.');
    } finally {
      setIsSubmitting(false);
    }
  };

  const handleDelete = () => {
    Alert.alert(
      'Delete Profile',
      'Are you sure you want to remove this child profile?',
      [
        { text: 'Cancel', style: 'cancel' },
        {
          text: 'Delete', style: 'destructive',
          onPress: async () => {
            setIsSubmitting(true);
            try {
              await deleteChildMutation.mutateAsync({ id: params.id as string });
              router.back();
            } catch {
              Alert.alert('Error', 'Failed to delete.');
              setIsSubmitting(false);
            }
          }
        }
      ]
    );
  };

  const removeInterest = (tag: string) => setInterests(prev => prev.filter(i => i !== tag));

  const openPicker = () => {
    setPendingInterests([...interests]);
    setSearchQuery('');
    setShowPicker(true);
  };

  const togglePending = (item: string) => {
    setPendingInterests(prev =>
      prev.includes(item) ? prev.filter(i => i !== item) : [...prev, item]
    );
  };

  const confirmPicker = () => {
    setInterests(pendingInterests);
    setShowPicker(false);
  };

  const filteredOptions = INTEREST_OPTIONS.filter(o =>
    o.toLowerCase().includes(searchQuery.toLowerCase())
  );

  return (
    <ThemedView style={[styles.container, { paddingTop: insets.top }]}>
      <Stack.Screen options={{ headerShown: false }} />
      <KeyboardAvoidingView
        behavior={Platform.OS === 'ios' ? 'padding' : 'height'}
        style={{ flex: 1 }}
        keyboardVerticalOffset={0}
      >
        <ScrollView contentContainerStyle={styles.scrollContent} showsVerticalScrollIndicator={false}>
          <View style={styles.familyBanner}>
            <TouchableOpacity onPress={() => router.back()} style={styles.bannerBack}>
              <IconSymbol name="chevron.left" size={24} color={theme.text} />
            </TouchableOpacity>
            <ThemedText style={styles.bannerTitle}>Family Information</ThemedText>
            {isEditing ? (
              <TouchableOpacity onPress={handleDelete}>
                <ThemedText style={{ color: '#EF4444', fontFamily: 'Archivo_600SemiBold' }}>Delete</ThemedText>
              </TouchableOpacity>
            ) : (
              <View style={{ width: 40 }} />
            )}
          </View>
          <ThemedText style={styles.pageTitle}>
            {isEditing ? 'Edit Child Profile' : 'Create Child Profile'}
          </ThemedText>
          <TextInput
            style={[styles.textInput, { backgroundColor: inputBg, color: theme.text }]}
            value={firstName}
            onChangeText={setFirstName}
            placeholder="First Name"
            placeholderTextColor={placeholderColor}
          />
          <TextInput
            style={[styles.textInput, { backgroundColor: inputBg, color: theme.text }]}
            value={lastName}
            onChangeText={setLastName}
            placeholder="Last Name"
            placeholderTextColor={placeholderColor}
          />
          <View style={styles.row}>
            <View style={styles.dropdownWrapper}>
              <TouchableOpacity
                style={[styles.dropdown, { backgroundColor: inputBg }]}
                onPress={() => { setShowMonthDrop(!showMonthDrop); setShowYearDrop(false); }}
              >
                <ThemedText style={birthMonth ? { color: theme.text } : styles.dropdownPlaceholder}>
                  {birthMonth || 'Month'}
                </ThemedText>
                <IconSymbol name="chevron.down" size={16} color="#6B7280" />
              </TouchableOpacity>

              {showMonthDrop && (
                <View style={[styles.dropdownList, { backgroundColor: dropdownPopupBg, borderColor }]}>
                  <ScrollView nestedScrollEnabled style={{ maxHeight: 180 }}>
                    {MONTHS.map(m => (
                      <TouchableOpacity
                        key={m}
                        style={[styles.dropdownItem, { borderBottomColor: borderColor }]}
                        onPress={() => { setBirthMonth(m); setShowMonthDrop(false); }}
                      >
                        <ThemedText>{m}</ThemedText>
                      </TouchableOpacity>
                    ))}
                  </ScrollView>
                </View>
              )}
            </View>
            <View style={styles.dropdownWrapper}>
              <TouchableOpacity
                style={[styles.dropdown, { backgroundColor: inputBg }]}
                onPress={() => { setShowYearDrop(!showYearDrop); setShowMonthDrop(false); }}
              >
                <ThemedText style={birthYear ? { color: theme.text } : styles.dropdownPlaceholder}>
                  {birthYear || 'Year'}
                </ThemedText>
                <IconSymbol name="chevron.down" size={16} color="#6B7280" />
              </TouchableOpacity>

              {showYearDrop && (
                <View style={[styles.dropdownList, { backgroundColor: dropdownPopupBg, borderColor }]}>
                  <ScrollView nestedScrollEnabled style={{ maxHeight: 180 }}>
                    {YEARS.map(y => (
                      <TouchableOpacity
                        key={y}
                        style={[styles.dropdownItem, { borderBottomColor: borderColor }]}
                        onPress={() => { setBirthYear(y); setShowYearDrop(false); }}
                      >
                        <ThemedText>{y}</ThemedText>
                      </TouchableOpacity>
                    ))}
                  </ScrollView>
                </View>
              )}
            </View>
          </View>
          <ThemedText style={styles.sectionLabel}>Interests</ThemedText>
          {interests.length > 0 && (
            <View style={styles.tagsRow}>
              {interests.map((tag, idx) => {
                const color = TAG_COLORS[idx % TAG_COLORS.length];
                return (
                  <TouchableOpacity
                    key={tag}
                    style={[styles.tag, { backgroundColor: color.bg, borderColor: color.border }]}
                    onPress={() => removeInterest(tag)}
                  >
                    <IconSymbol name="camera.filters" size={13} color={color.border} />
                    <ThemedText style={[styles.tagText, { color: color.text }]}>{tag}</ThemedText>
                  </TouchableOpacity>
                );
              })}
            </View>
          )}
          <TouchableOpacity style={[styles.addInterestBtn, { borderColor }]} onPress={openPicker}>
            <ThemedText style={[styles.addInterestText, { color: theme.text }]}>Add Interest</ThemedText>
            <IconSymbol name="chevron.down" size={16} color="#6B7280" />
          </TouchableOpacity>
          {showPicker && (
            <View style={[styles.pickerPanel, { borderColor }]}>
              <View style={styles.searchRow}>
                <TextInput
                  style={[styles.searchInput, { color: theme.text }]}
                  value={searchQuery}
                  onChangeText={setSearchQuery}
                  placeholder="Input"
                  placeholderTextColor={placeholderColor}
                />
                <IconSymbol name="magnifyingglass" size={20} color="#6B7280" />
              </View>
              <View style={[styles.pickerDivider, { backgroundColor: borderColor }]} />
              {filteredOptions.map(item => (
                <TouchableOpacity
                  key={item}
                  style={[styles.pickerItem, { borderBottomColor: inputBg }]}
                  onPress={() => togglePending(item)}
                >
                  <ThemedText style={styles.pickerItemText}>{item}</ThemedText>
                  <View style={[styles.checkbox, pendingInterests.includes(item) && styles.checkboxChecked]}>
                    {pendingInterests.includes(item) && (
                      <IconSymbol name="checkmark" size={12} color="#1F2937" />
                    )}
                  </View>
                </TouchableOpacity>
              ))}
              <View style={styles.pickerActions}>
                <TouchableOpacity style={styles.addBtn} onPress={confirmPicker}>
                  <ThemedText style={styles.addBtnText}>Add</ThemedText>
                </TouchableOpacity>
                <TouchableOpacity onPress={() => setShowPicker(false)}>
                  <ThemedText style={styles.cancelText}>Cancel</ThemedText>
                </TouchableOpacity>
              </View>
            </View>
          )}
          <View style={styles.availabilityRow}>
            <ThemedText style={styles.availabilityTitle}>Availability</ThemedText>
            <TouchableOpacity style={[styles.editBtn, { borderColor }]}>
              <ThemedText style={styles.editBtnText}>Edit</ThemedText>
            </TouchableOpacity>
          </View>
          <TouchableOpacity
            style={[styles.saveButton, { backgroundColor: theme.tint, opacity: isSubmitting ? 0.7 : 1 }]}
            onPress={handleSave}
            disabled={isSubmitting}
          >
            <ThemedText style={styles.saveButtonText}>
              {isSubmitting ? 'Saving...' : 'Save Changes'}
            </ThemedText>
          </TouchableOpacity>

        </ScrollView>
      </KeyboardAvoidingView>
    </ThemedView>
  );
}

const styles = StyleSheet.create({
  container: { flex: 1 },
  scrollContent: { paddingHorizontal: 20, paddingBottom: 40, paddingTop: 10 },

  familyBanner: {
    flexDirection: 'row',
    alignItems: 'center',
    justifyContent: 'space-between',
    marginBottom: 24,
  },
  bannerBack: {
    width: 32,
    height: 32,
    justifyContent: 'center',
    alignItems: 'flex-start',
  },
  bannerTitle: {
    fontSize: 20,
    fontFamily: 'Archivo_700Bold',
    textAlign: 'center',
  },

  pageTitle: {
    fontSize: 22,
    fontFamily: 'Archivo_600SemiBold',
    marginBottom: 20,
  },

  textInput: {
    borderRadius: 10,
    paddingHorizontal: 16,
    paddingVertical: 14,
    fontSize: 16,
    fontFamily: 'Archivo_400Regular',
    marginBottom: 12,
  },

  row: {
    flexDirection: 'row',
    gap: 12,
    marginBottom: 24,
    zIndex: 10,
  },
  dropdownWrapper: {
    flex: 1,
    zIndex: 10,
  },
  dropdown: {
    borderRadius: 10,
    paddingHorizontal: 16,
    paddingVertical: 14,
    flexDirection: 'row',
    alignItems: 'center',
    justifyContent: 'space-between',
  },
  dropdownPlaceholder: {
    fontSize: 16,
    color: '#9CA3AF',
    fontFamily: 'Archivo_400Regular',
  },
  dropdownList: {
    position: 'absolute',
    top: 52,
    left: 0,
    right: 0,
    borderRadius: 10,
    borderWidth: 1,
    zIndex: 100,
    elevation: 5,
    shadowColor: '#000',
    shadowOpacity: 0.1,
    shadowRadius: 8,
    shadowOffset: { width: 0, height: 2 },
  },
  dropdownItem: {
    paddingHorizontal: 16,
    paddingVertical: 12,
    borderBottomWidth: 1,
  },

  sectionLabel: {
    fontSize: 16,
    fontFamily: 'Archivo_600SemiBold',
    marginBottom: 12,
  },

  tagsRow: {
    flexDirection: 'row',
    flexWrap: 'wrap',
    gap: 8,
    marginBottom: 12,
  },
  tag: {
    flexDirection: 'row',
    alignItems: 'center',
    paddingHorizontal: 8,
    paddingVertical: 4,
    borderRadius: 20,
    borderWidth: 1,
    gap: 4,
  },
  tagText: {
    fontSize: 12,
    fontFamily: 'Archivo_500Medium',
  },

  addInterestBtn: {
    borderWidth: 1,
    borderRadius: 10,
    paddingHorizontal: 16,
    paddingVertical: 14,
    flexDirection: 'row',
    alignItems: 'center',
    justifyContent: 'space-between',
    marginBottom: 12,
  },
  addInterestText: {
    fontSize: 16,
    fontFamily: 'Archivo_400Regular',
  },

  pickerPanel: {
    borderWidth: 1,
    borderRadius: 10,
    overflow: 'hidden',
    marginBottom: 24,
  },
  searchRow: {
    flexDirection: 'row',
    alignItems: 'center',
    paddingHorizontal: 16,
    paddingVertical: 12,
    gap: 8,
  },
  searchInput: {
    flex: 1,
    fontSize: 16,
    fontFamily: 'Archivo_400Regular',
  },
  pickerDivider: {
    height: 1,
  },
  pickerItem: {
    flexDirection: 'row',
    alignItems: 'center',
    justifyContent: 'space-between',
    paddingHorizontal: 16,
    paddingVertical: 16,
    borderBottomWidth: 1,
  },
  pickerItemText: {
    fontSize: 16,
    fontFamily: 'Archivo_400Regular',
  },
  checkbox: {
    width: 22,
    height: 22,
    borderRadius: 4,
    borderWidth: 1.5,
    borderColor: '#9CA3AF',
    justifyContent: 'center',
    alignItems: 'center',
  },
  checkboxChecked: {
    borderColor: '#1F2937',
  },
  pickerActions: {
    flexDirection: 'row',
    alignItems: 'center',
    padding: 16,
    gap: 16,
  },
  addBtn: {
    backgroundColor: '#1D4ED8',
    paddingHorizontal: 28,
    paddingVertical: 12,
    borderRadius: 8,
  },
  addBtnText: {
    color: '#FFF',
    fontSize: 15,
    fontFamily: 'Archivo_600SemiBold',
  },
  cancelText: {
    fontSize: 15,
    fontFamily: 'Archivo_400Regular',
    color: '#6B7280',
  },

  availabilityRow: {
    flexDirection: 'row',
    alignItems: 'center',
    justifyContent: 'space-between',
    marginBottom: 24,
    marginTop: 8,
  },
  availabilityTitle: {
    fontSize: 18,
    fontFamily: 'Archivo_600SemiBold',
  },
  editBtn: {
    borderWidth: 1,
    borderRadius: 8,
    paddingHorizontal: 20,
    paddingVertical: 8,
  },
  editBtnText: {
    fontSize: 15,
    fontFamily: 'Archivo_400Regular',
    color: '#6B7280',
  },

  saveButton: {
    paddingVertical: 16,
    borderRadius: 12,
    alignItems: 'center',
    justifyContent: 'center',
  },
  saveButtonText: {
    color: '#FFF',
    fontSize: 16,
    fontFamily: 'Archivo_600SemiBold',
  },
});
