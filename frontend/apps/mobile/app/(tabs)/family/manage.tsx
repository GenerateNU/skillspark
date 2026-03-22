import React, { useState } from 'react';
import {
  View,
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
import { Colors, AppColors } from '@/constants/theme';
import { IconSymbol } from '@/components/ui/icon-symbol';
import { useQueryClient } from '@tanstack/react-query';
import { useCreateChild, useUpdateChild, useDeleteChild, getGetChildrenByGuardianIdQueryKey } from '@skillspark/api-client';
import { ChildProfileForm, MONTHS } from '@/components/ChildProfileForm';
import { useTranslation } from 'react-i18next';
import { useGuardian } from '@/hooks/use-guardian';


export default function ManageChildScreen() {
  const router = useRouter();
  const params = useLocalSearchParams();
  const colorScheme = useColorScheme();
  const insets = useSafeAreaInsets();
  const theme = Colors[colorScheme ?? 'light'];

  const { t: translate } = useTranslation();
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

  const { guardianId } = useGuardian();
  const queryClient = useQueryClient();
  const createChildMutation = useCreateChild();
  const updateChildMutation = useUpdateChild();
  const deleteChildMutation = useDeleteChild();

  const handleSave = async () => {
    if (!firstName || !birthYear || !birthMonth || !schoolId) {
      Alert.alert(translate('common.error'), translate('childProfile.requiredFieldsError'));
      return;
    }
    const name = [firstName, lastName].filter(Boolean).join(' ');
    setIsSubmitting(true);
    try {
      const childData = {
        name,
        birth_year: parseInt(birthYear, 10),
        birth_month: MONTHS.indexOf(birthMonth) + 1,
        guardian_id: guardianId,
        school_id: schoolId,
        interests,
      };
      if (isEditing) {
        await updateChildMutation.mutateAsync({ id: params.id as string, data: childData });
      } else {
        await createChildMutation.mutateAsync({ data: childData });
      }
      await queryClient.invalidateQueries({ queryKey: getGetChildrenByGuardianIdQueryKey(guardianId) });
      router.back();
    } catch (error) {
      console.error(error);
      Alert.alert(translate('common.errorOccurred'), translate('childProfile.saveError'));
    } finally {
      setIsSubmitting(false);
    }
  };

  const handleDelete = () => {
    Alert.alert(
      translate('childProfile.deleteProfile'),
      translate('childProfile.deleteConfirm'),
      [
        { text: translate('common.cancel'), style: 'cancel' },
        {
          text: translate('payment.delete'), style: 'destructive',
          onPress: async () => {
            setIsSubmitting(true);
            try {
              await deleteChildMutation.mutateAsync({ id: params.id as string });
              await queryClient.invalidateQueries({ queryKey: getGetChildrenByGuardianIdQueryKey(guardianId) });
              router.back();
            } catch {
              Alert.alert(translate('common.errorOccurred'), translate('childProfile.deleteError'));
              setIsSubmitting(false);
            }
          }
        }
      ]
    );
  };

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
            <ThemedText className="text-xl text-center font-nunito-bold">{translate('familyInformation.title')}</ThemedText>
            {isEditing ? (
              <TouchableOpacity onPress={handleDelete}>
                <ThemedText className="font-nunito-semibold" style={{ color: AppColors.danger }}>{translate('payment.delete')}</ThemedText>
              </TouchableOpacity>
            ) : (
              <View className="w-10" />
            )}
          </View>
          <ThemedText className="text-[22px] font-nunito-semibold mb-5">
            {isEditing ? translate('childProfile.editTitle') : translate('childProfile.createTitle')}
          </ThemedText>
          <ChildProfileForm
            firstName={firstName}
            setFirstName={setFirstName}
            lastName={lastName}
            setLastName={setLastName}
            birthMonth={birthMonth}
            setBirthMonth={setBirthMonth}
            birthYear={birthYear}
            setBirthYear={setBirthYear}
            schoolId={schoolId}
            setSchoolId={setSchoolId}
            interests={interests}
            setInterests={setInterests}
            searchQuery={searchQuery}
            setSearchQuery={setSearchQuery}
            showMonthDrop={showMonthDrop}
            setShowMonthDrop={setShowMonthDrop}
            showYearDrop={showYearDrop}
            setShowYearDrop={setShowYearDrop}
          />
          <TouchableOpacity
            className="py-4 rounded-xl items-center justify-center"
            style={{ backgroundColor: theme.tint, opacity: isSubmitting ? 0.7 : 1 }}
            onPress={handleSave}
            disabled={isSubmitting}
          >
            <ThemedText className="text-white text-base font-nunito-semibold">
              {isSubmitting ? translate('childProfile.saving') : translate('childProfile.saveChanges')}
            </ThemedText>
          </TouchableOpacity>

        </ScrollView>
      </KeyboardAvoidingView>
    </ThemedView>
  );
}
