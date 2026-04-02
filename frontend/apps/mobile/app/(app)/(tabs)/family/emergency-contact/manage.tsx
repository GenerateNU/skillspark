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
import { useTranslation } from 'react-i18next';
import { useAuthContext } from '@/hooks/use-auth-context';
import { ErrorScreen } from '@/components/ErrorScreen';
import { EmergencyContactForm } from '@/components/EmergencyContactProfileForm';
import { queryClient } from '@/constants/query-client';
import { getGetEmergencyContactsByGuardianIdQueryKey, useCreateEmergencyContact, useDeleteEmergencyContact,  useUpdateEmergencyContact } from '@skillspark/api-client';
import { useGuardian } from '@/hooks/use-guardian';


// screen for adding an emergency contact
export default function ManageEmergencyContactScreen() {
  const router = useRouter();
  const params = useLocalSearchParams();
  const colorScheme = useColorScheme();
  const insets = useSafeAreaInsets();
  const theme = Colors[colorScheme ?? 'light'];

  //const { guardianId } = useAuthContext();
  const { guardianId } = useGuardian();
 

  const createEmergencyContactMutation = useCreateEmergencyContact();
  const updateEmergencyContactMutation = useUpdateEmergencyContact();
  const deleteEmergencyContactMutation = useDeleteEmergencyContact();

  const { t: translate } = useTranslation();
  const isEditing = !!params.id;

  const [firstName, setFirstName] = useState(
    params.name ? (params.name as string).split(' ')[0] : ''
  );
  const [lastName, setLastName] = useState(
    params.name ? (params.name as string).split(' ').slice(1).join(' ') : ''
  );
  const [phoneNumber, setPhoneNumber] = useState(params.phone_number as string || '');
  const [isSubmitting, setIsSubmitting] = useState(false);

  if (!guardianId) {
    return <ErrorScreen message="Illegal state: no guardian ID retrieved" />;
  }

  const isValidPhoneNumber = (phoneNumber: string) => {
    const phoneValidationRegex = /^[\+]?[(]?[0-9]{3}[)]?[-\s\.]?[0-9]{3}[-\s\.]?[0-9]{4,6}$/im;
    const isValid = (str: string) => phoneValidationRegex.test(str);

    return isValid(phoneNumber);
  }

  const name = [firstName, lastName].filter(Boolean).join(" ");
  const emergencyContactData = {
    guardian_id: guardianId,
    name: name,
    phone_number: phoneNumber,
  };

  const handleSave = async () => {
    if (!firstName || !lastName || !phoneNumber) {
        Alert.alert(
          translate("common.error"),
          translate("childProfile.requiredFieldsError"),
        );
        return;
    }

    if (!isValidPhoneNumber(phoneNumber)) {
      Alert.alert(
          translate("common.error"),
          "Please enter a valid phone number"
        );
      return;
    }

      setIsSubmitting(true);
      try {
        if (isEditing) {
          await updateEmergencyContactMutation.mutateAsync({
            id: params.id as string,
            data: emergencyContactData,
          });
        } else {
          await createEmergencyContactMutation.mutateAsync({ data: emergencyContactData });
        }
        await queryClient.invalidateQueries({
          queryKey: getGetEmergencyContactsByGuardianIdQueryKey(guardianId),
        });
        router.back();
      } catch (error) {
        Alert.alert(
          translate("common.errorOccurred"),
          translate("childProfile.saveError"),
        );
      } finally {
        setIsSubmitting(false);
      }
    };

  const handleDelete = () => {
      Alert.alert(
        translate("childProfile.deleteProfile"),
        translate("childProfile.deleteConfirm"),
        [
          { text: translate("common.cancel"), style: "cancel" },
          {
            text: translate("payment.delete"),
            style: "destructive",
            onPress: async () => {
              setIsSubmitting(true);
              try {
                await deleteEmergencyContactMutation.mutateAsync({
                  id: params.id as string,
                });
                await queryClient.invalidateQueries({
                  queryKey: getGetEmergencyContactsByGuardianIdQueryKey(guardianId),
                });
                router.back();
              } catch {
                Alert.alert(
                  translate("common.errorOccurred"),
                  translate("childProfile.deleteError"),
                );
                setIsSubmitting(false);
              }
            },
          },
        ],
      );
    };

  return (
    <ThemedView className="flex-1" style={{ paddingTop: insets.top }}>
      <Stack.Screen options={{ headerShown: false }} />
      <KeyboardAvoidingView
        behavior={Platform.OS === 'ios' ? 'padding' : 'height'}
        className="flex-1"
        keyboardVerticalOffset={0}
      >
        <ScrollView
          contentContainerStyle={{ paddingHorizontal: 20, paddingBottom: 40, paddingTop: 10 }}
          showsVerticalScrollIndicator={false}
        >
          <View className="flex-row items-center justify-between mb-6">
            <TouchableOpacity onPress={() => router.back()} className="w-8 h-8 justify-center items-start">
              <IconSymbol name="chevron.left" size={24} color={theme.text} />
            </TouchableOpacity>
            <ThemedText className="text-xl text-center font-nunito-bold mt-0.5">{translate('profile.familyInformation')}</ThemedText>
            {isEditing ? (
              <TouchableOpacity onPress={handleDelete}>
                <ThemedText className="font-nunito-semibold" style={{ color: AppColors.danger }}>{translate('emergencyContact.deleteContact')}</ThemedText>
              </TouchableOpacity>
            ) : (
              <View className="w-10" />
            )}
          </View>
          <ThemedText className="text-[22px] font-nunito-semibold mb-5">
            {isEditing ? translate('emergencyContact.editTitle') : translate('emergencyContact.addTitle')}
          </ThemedText>
          <EmergencyContactForm
            firstName={firstName}
            setFirstName={setFirstName}
            lastName={lastName}
            setLastName={setLastName}
            phoneNumber={phoneNumber}
            setPhoneNumber={setPhoneNumber}
          />
          <TouchableOpacity
            className={`py-4 rounded-xl items-center justify-center ${isSubmitting ? 'opacity-70' : 'opacity-100'}`}
            style={{ backgroundColor: theme.tint }}
            onPress={handleSave}
            disabled={isSubmitting}
          >
            <ThemedText className="text-white text-base font-nunito-semibold">
              {isSubmitting ? translate('emergencyContact.saving') : isEditing ? translate('emergencyContact.saveChanges') : translate('emergencyContact.addContact')}
            </ThemedText>
          </TouchableOpacity>
        </ScrollView>
      </KeyboardAvoidingView>
    </ThemedView>
  );
}
