import React from 'react';
import {
  View,
  ScrollView,
  TouchableOpacity,
  ActivityIndicator,
  useColorScheme,
} from "react-native";
import { useRouter } from 'expo-router';
import { useSafeAreaInsets } from 'react-native-safe-area-context';
import { ThemedText } from '@/components/themed-text';
import { ThemedView } from '@/components/themed-view';
import { IconSymbol } from '@/components/ui/icon-symbol';
import { Colors, AppColors } from '@/constants/theme';
import { ChildListItem } from '@/components/ChildListItem';
import { SectionHeader } from '@/components/SectionHeader';
import { useTranslation } from 'react-i18next';
import { useGuardian } from '@/hooks/use-guardian';
import { useAuthContext } from '@/hooks/use-auth-context';
import { ErrorScreen } from '@/components/ErrorScreen';
import { EmergencyContactListItem } from '@/components/EmergencyContactListItem';

export default function FamilyListScreen() {
  const router = useRouter();
  const insets = useSafeAreaInsets();
  const colorScheme = useColorScheme();
  const theme = Colors[colorScheme ?? "light"];
  const { t: translate } = useTranslation();

  const { guardian, children, emergencyContacts, isLoading } = useGuardian();
  const { guardianId } = useAuthContext();


  const handleAddChild = () => {
    router.push("/family/manage");
  };

  const handleAddEmergencyContact = () => {
    router.push('/family/emergency-contact/manage');
  }

  const handleEditEmergencyContact = (emergencyContact: any) => {
    router.push({
      pathname: '/family/emergency-contact/manage',
      params: {
        id: emergencyContact.id,
        guardian_id: emergencyContact.guardian_id,
        name: emergencyContact.name,
        phone_number: emergencyContact.phone_number,
      },
    });
  };

  const handleEditChild = (child: any) => {
    router.push({
      pathname: "/family/manage",
      params: {
        id: child.id,
        name: child.name,
        birth_month: child.birth_month,
        birth_year: child.birth_year,
        school_id: child.school_id ?? "",
        interests: child.interests ?? [],
      },
    });
  };

  if (!guardianId) {
    return <ErrorScreen message="Illegal state: no guardian ID retrieved" />;
  }

  if (isLoading) {
    return (
      <ThemedView className="flex-1 items-center justify-center">
        <ActivityIndicator size="large" />
      </ThemedView>
    );
  }

  return (
    <ThemedView className="flex-1" style={{ paddingTop: insets.top }}>
      <View className="flex-row items-center justify-between px-5 py-3">
        <TouchableOpacity
          onPress={() => router.navigate("/profile")}
          className="w-10 justify-center items-start"
          hitSlop={{ top: 10, bottom: 10, left: 10, right: 10 }}
        >
          <IconSymbol name="chevron.left" size={24} color={theme.text} />
        </TouchableOpacity>
        <ThemedText className="text-xl text-center font-nunito-bold">
          {translate("familyInformation.title")}
        </ThemedText>
        <View className="w-10" />
      </View>

      <ScrollView
        contentContainerStyle={{ paddingHorizontal: 20, paddingTop: 12 }}
        showsVerticalScrollIndicator={false}
      >
        <TouchableOpacity
          className="flex-row items-start py-4 gap-3"
          activeOpacity={0.7}
        >
          <View className="w-11 h-11 items-center justify-center">
            <IconSymbol name="person.circle" size={40} color={theme.text} />
          </View>
          <View className="flex-1 gap-1">
            <ThemedText className="text-base font-nunito-semibold">
              {guardian?.name}
            </ThemedText>
            <ThemedText
              className="text-[13px] font-nunito"
              style={{ color: AppColors.mutedText }}
            >
              @{guardian?.username}
            </ThemedText>
            <ThemedText
              className="text-[13px] font-nunito"
              style={{ color: AppColors.mutedText }}
            >
              {guardian?.email}
            </ThemedText>
          </View>
        </TouchableOpacity>
        <View
          className="h-px my-3"
          style={{ backgroundColor: AppColors.divider }}
        />
        <SectionHeader
          title={translate("familyInformation.childProfile")}
          actionLabel={translate("familyInformation.addProfile")}
          onAction={handleAddChild}
        />
        {children.length === 0 && (
          <ThemedText
            className="text-sm pb-4 font-nunito"
            style={{ color: AppColors.subtleText }}
          >
            {translate("common.noChildProfilesAdded")}
          </ThemedText>
        )}
        {children.map((child: any, idx: number) => (
          <React.Fragment key={child.id}>
            <ChildListItem
              child={child}
              onPress={() => handleEditChild(child)}
            />
            {idx < children.length - 1 && (
              <View
                className="h-px my-3"
                style={{ backgroundColor: AppColors.divider }}
              />
            )}
          </React.Fragment>
        ))}
        <View
          className="h-px my-3"
          style={{ backgroundColor: AppColors.divider }}
        />
        <SectionHeader
          title={translate('familyInformation.emergencyContact')}
          actionLabel={translate('familyInformation.addContact')}
          onAction={() => handleAddEmergencyContact()}
        />

        {emergencyContacts.length === 0 && (
        <ThemedText className="text-sm pb-4 font-nunito" style={{ color: AppColors.subtleText }}>No emergency contacts added</ThemedText>
        )}
        {emergencyContacts.map((emergencyContact: any, idx: number) => (
          <React.Fragment key={emergencyContact.id}>
            <EmergencyContactListItem
              emergencyContact={emergencyContact}
              onPress={() => handleEditEmergencyContact(emergencyContact)}
            />
            {idx < emergencyContacts.length - 1 && <View className="h-px my-3" style={{ backgroundColor: AppColors.divider }} />}
          </React.Fragment>
        ))}
        <View className="h-10" />
      </ScrollView>
    </ThemedView>
  );
}
