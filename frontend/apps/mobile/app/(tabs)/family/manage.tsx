import React, { useState } from 'react';
import {
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
import { Colors, AppColors, TAG_COLORS } from '@/constants/theme';
import { IconSymbol } from '@/components/ui/icon-symbol';
import { useQueryClient } from '@tanstack/react-query';
import { useCreateChild, useUpdateChild, useDeleteChild, getGetChildrenByGuardianIdQueryKey } from '@skillspark/api-client';

// TODO: Replace with authenticated user's guardian ID
const GUARDIAN_ID = '88888888-8888-8888-8888-888888888888';

const INTEREST_OPTIONS = [
  'science', 'math', 'music', 'art', 'sports', 'technology', 'language', 'other',
];

const capitalize = (s: string) => s.charAt(0).toUpperCase() + s.slice(1);

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
  // TODO: Replace raw school_id text input with a school picker that shows the school name
  const [schoolId, setSchoolId] = useState(params.school_id as string || '');

  const initialInterests = Array.isArray(params.interests)
    ? params.interests
    : params.interests
    ? (params.interests as string).split(',').map(s => s.trim()).filter(Boolean)
    : [];
  const [interests, setInterests] = useState<string[]>(initialInterests);

  const [searchQuery, setSearchQuery] = useState('');

  const [showMonthDrop, setShowMonthDrop] = useState(false);
  const [showYearDrop, setShowYearDrop] = useState(false);

  const [isSubmitting, setIsSubmitting] = useState(false);

  const queryClient = useQueryClient();
  const createChildMutation = useCreateChild();
  const updateChildMutation = useUpdateChild();
  const deleteChildMutation = useDeleteChild();

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
      await queryClient.invalidateQueries({ queryKey: getGetChildrenByGuardianIdQueryKey(GUARDIAN_ID) });
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
              await queryClient.invalidateQueries({ queryKey: getGetChildrenByGuardianIdQueryKey(GUARDIAN_ID) });
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

  const toggleInterest = (item: string) => {
    setInterests(prev =>
      prev.includes(item) ? prev.filter(i => i !== item) : [...prev, item]
    );
  };

  const filteredOptions = INTEREST_OPTIONS.filter(o =>
    o.toLowerCase().includes(searchQuery.toLowerCase())
  );

  return (
    <ThemedView className="flex-1" style={{ paddingTop: insets.top }}>
      <Stack.Screen options={{ headerShown: false }} />
      <KeyboardAvoidingView
        behavior={Platform.OS === 'ios' ? 'padding' : 'height'}
        style={{ flex: 1 }}
        keyboardVerticalOffset={0}
      >
        <ScrollView contentContainerStyle={{ paddingHorizontal: 20, paddingBottom: 40, paddingTop: 10 }} showsVerticalScrollIndicator={false}>
          <View className="flex-row items-center justify-between mb-6">
            <TouchableOpacity onPress={() => router.back()} className="w-8 h-8 justify-center items-start">
              <IconSymbol name="chevron.left" size={24} color={theme.text} />
            </TouchableOpacity>
            <ThemedText className="text-xl text-center font-nunito-bold">Family Information</ThemedText>
            {isEditing ? (
              <TouchableOpacity onPress={handleDelete}>
                <ThemedText className="font-nunito-semibold" style={{ color: AppColors.danger }}>Delete</ThemedText>
              </TouchableOpacity>
            ) : (
              <View className="w-10" />
            )}
          </View>
          <ThemedText className="text-[22px] font-nunito-semibold mb-5">
            {isEditing ? 'Edit Child Profile' : 'Create Child Profile'}
          </ThemedText>
          <TextInput
            className="rounded-[10px] px-4 py-[14px] text-base font-nunito mb-3"
            style={{ backgroundColor: theme.inputBg, color: theme.text }}
            value={firstName}
            onChangeText={setFirstName}
            placeholder="First Name"
            placeholderTextColor={AppColors.placeholderText}
          />
          <TextInput
            className="rounded-[10px] px-4 py-[14px] text-base font-nunito mb-3"
            style={{ backgroundColor: theme.inputBg, color: theme.text }}
            value={lastName}
            onChangeText={setLastName}
            placeholder="Last Name"
            placeholderTextColor={AppColors.placeholderText}
          />
          <View className="flex-row gap-3 mb-6" style={{ zIndex: 10 }}>
            <View className="flex-1" style={{ zIndex: 10 }}>
              <TouchableOpacity
                className="rounded-[10px] px-4 py-[14px] flex-row items-center justify-between"
                style={{ backgroundColor: theme.inputBg }}
                onPress={() => { setShowMonthDrop(!showMonthDrop); setShowYearDrop(false); }}
              >
                <ThemedText className={birthMonth ? '' : 'font-nunito'} style={birthMonth ? {} : { color: AppColors.placeholderText }}>
                  {birthMonth || 'Month'}
                </ThemedText>
                <IconSymbol name="chevron.down" size={16} color={AppColors.mutedText} />
              </TouchableOpacity>

              {showMonthDrop && (
                <View
                  className="absolute left-0 right-0 rounded-[10px] border"
                  style={{ top: 52, backgroundColor: theme.dropdownBg, borderColor: theme.borderColor, zIndex: 100, elevation: 5, shadowColor: '#000', shadowOpacity: 0.1, shadowRadius: 8, shadowOffset: { width: 0, height: 2 } }}
                >
                  <ScrollView nestedScrollEnabled style={{ maxHeight: 180 }}>
                    {MONTHS.map(m => (
                      <TouchableOpacity
                        key={m}
                        className="px-4 py-3 border-b"
                        style={{ borderBottomColor: theme.borderColor }}
                        onPress={() => { setBirthMonth(m); setShowMonthDrop(false); }}
                      >
                        <ThemedText>{m}</ThemedText>
                      </TouchableOpacity>
                    ))}
                  </ScrollView>
                </View>
              )}
            </View>
            <View className="flex-1" style={{ zIndex: 10 }}>
              <TouchableOpacity
                className="rounded-[10px] px-4 py-[14px] flex-row items-center justify-between"
                style={{ backgroundColor: theme.inputBg }}
                onPress={() => { setShowYearDrop(!showYearDrop); setShowMonthDrop(false); }}
              >
                <ThemedText className={birthYear ? '' : 'font-nunito'} style={birthYear ? {} : { color: AppColors.placeholderText }}>
                  {birthYear || 'Year'}
                </ThemedText>
                <IconSymbol name="chevron.down" size={16} color={AppColors.mutedText} />
              </TouchableOpacity>

              {showYearDrop && (
                <View
                  className="absolute left-0 right-0 rounded-[10px] border"
                  style={{ top: 52, backgroundColor: theme.dropdownBg, borderColor: theme.borderColor, zIndex: 100, elevation: 5, shadowColor: '#000', shadowOpacity: 0.1, shadowRadius: 8, shadowOffset: { width: 0, height: 2 } }}
                >
                  <ScrollView nestedScrollEnabled style={{ maxHeight: 180 }}>
                    {YEARS.map(y => (
                      <TouchableOpacity
                        key={y}
                        className="px-4 py-3 border-b"
                        style={{ borderBottomColor: theme.borderColor }}
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
          {/* TODO: Replace with school picker that resolves school_id to a human-readable school name */}
          <TextInput
            className="rounded-[10px] px-4 py-[14px] text-base font-nunito mb-6"
            style={{ backgroundColor: theme.inputBg, color: theme.text }}
            value={schoolId}
            onChangeText={setSchoolId}
            placeholder="School ID"
            placeholderTextColor={AppColors.placeholderText}
          />
          <ThemedText className="text-base font-nunito-semibold mb-3">Interests</ThemedText>
          {interests.length > 0 && (
            <ScrollView horizontal showsHorizontalScrollIndicator={false} className="mb-3" contentContainerStyle={{ gap: 8, paddingRight: 4 }}>
              {interests.map((tag, idx) => {
                const color = TAG_COLORS[idx % TAG_COLORS.length];
                return (
                  <TouchableOpacity
                    key={tag}
                    className="flex-row items-center px-2 py-1 rounded-full border gap-1"
                    style={{ backgroundColor: color.bg, borderColor: color.border }}
                    onPress={() => removeInterest(tag)}
                  >
                    <IconSymbol name="camera.filters" size={13} color={color.border} />
                    <ThemedText className="text-xs font-nunito-medium" style={{ color: color.text }}>{capitalize(tag)}</ThemedText>
                  </TouchableOpacity>
                );
              })}
            </ScrollView>
          )}
          <View className="border rounded-[10px] overflow-hidden mb-6" style={{ borderColor: theme.borderColor }}>
            <View className="flex-row items-center px-4 py-3 gap-2">
              <TextInput
                className="flex-1 text-base font-nunito"
                style={{ color: theme.text }}
                value={searchQuery}
                onChangeText={setSearchQuery}
                placeholder="Search interests..."
                placeholderTextColor={AppColors.placeholderText}
              />
              <IconSymbol name="magnifyingglass" size={20} color={AppColors.mutedText} />
            </View>
            <View className="h-px" style={{ backgroundColor: theme.borderColor }} />
            <View onStartShouldSetResponder={() => true} onMoveShouldSetResponder={() => true}>
            <ScrollView nestedScrollEnabled showsVerticalScrollIndicator style={{ maxHeight: 150 }}>
              {filteredOptions.map(item => (
                <TouchableOpacity
                  key={item}
                  className="flex-row items-center justify-between px-4 py-4 border-b"
                  style={{ borderBottomColor: theme.inputBg }}
                  onPress={() => toggleInterest(item)}
                >
                  <ThemedText className="text-base font-nunito">{capitalize(item)}</ThemedText>
                  <View
                    className="w-[22px] h-[22px] rounded-[4px] border-[1.5px] items-center justify-center"
                    style={{ borderColor: interests.includes(item) ? AppColors.checkboxSelected : AppColors.subtleText }}
                  >
                    {interests.includes(item) && (
                      <IconSymbol name="checkmark" size={12} color={AppColors.checkboxSelected} />
                    )}
                  </View>
                </TouchableOpacity>
              ))}
            </ScrollView>
            </View>
          </View>
          <TouchableOpacity
            className="py-4 rounded-xl items-center justify-center"
            style={{ backgroundColor: theme.tint, opacity: isSubmitting ? 0.7 : 1 }}
            onPress={handleSave}
            disabled={isSubmitting}
          >
            <ThemedText className="text-white text-base font-nunito-semibold">
              {isSubmitting ? 'Saving...' : 'Save Changes'}
            </ThemedText>
          </TouchableOpacity>

        </ScrollView>
      </KeyboardAvoidingView>
    </ThemedView>
  );
}
