import React, { useState } from "react";
import {
  View,
  TextInput,
  TouchableOpacity,
  Alert,
  ScrollView,
  KeyboardAvoidingView,
  Platform,
  Keyboard,
  Pressable,
} from "react-native";
import { AuthBackground } from "@/components/AuthBackground";
import { Stack, useRouter, useLocalSearchParams } from "expo-router";
import { useSafeAreaInsets } from "react-native-safe-area-context";
import { ThemedText } from "@/components/themed-text";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { useTranslation } from "react-i18next";
import { useAuthContext } from "@/hooks/use-auth-context";
import { ErrorScreen } from "@/components/ErrorScreen";
import { Button } from "@/components/Button";
import { AppColors, Colors, FontSizes } from "@/constants/theme";
import { queryClient } from "@/constants/query-client";
import {
  getGetEmergencyContactsByGuardianIdQueryKey,
  useCreateEmergencyContact,
  useDeleteEmergencyContact,
  useUpdateEmergencyContact,
} from "@skillspark/api-client";

export default function ManageEmergencyContactScreen() {
  const router = useRouter();
  const params = useLocalSearchParams();
  const insets = useSafeAreaInsets();
  const theme = Colors.light;
  const { guardianId } = useAuthContext();
  const { t: translate } = useTranslation();
  const isEditing = !!params.id;
  const [isSubmitting, setIsSubmitting] = useState(false);

  const existingName = (params.name as string) || "";
  const [firstName, setFirstName] = useState(existingName.split(" ")[0] ?? "");
  const [lastName, setLastName] = useState(
    existingName.split(" ").slice(1).join(" ") ?? "",
  );
  const [phoneNumber, setPhoneNumber] = useState(
    (params.phone_number as string) || "",
  );

  const createEmergencyContactMutation = useCreateEmergencyContact();
  const updateEmergencyContactMutation = useUpdateEmergencyContact();
  const deleteEmergencyContactMutation = useDeleteEmergencyContact();

  if (!guardianId) {
    return <ErrorScreen message={translate("common.noGuardianId")} />;
  }

  const isValidPhoneNumber = (phone: string) => {
    const phoneValidationRegex =
      /^[\+]?[(]?[0-9]{3}[)]?[-\s\.]?[0-9]{3}[-\s\.]?[0-9]{4,6}$/im;
    return phoneValidationRegex.test(phone);
  };

  const handleSave = async () => {
    if (!firstName || !phoneNumber) {
      Alert.alert(
        translate("common.error"),
        translate("emergencyContact.requiredFieldsError"),
      );
      return;
    }
    if (!isValidPhoneNumber(phoneNumber)) {
      Alert.alert(
        translate("common.error"),
        translate("emergencyContact.invalidPhoneNumber"),
      );
      return;
    }
    setIsSubmitting(true);
    const name = [firstName, lastName].filter(Boolean).join(" ");
    try {
      const emergencyContactData = {
        guardian_id: guardianId,
        name,
        phone_number: phoneNumber,
      };

      if (isEditing) {
        await updateEmergencyContactMutation.mutateAsync({
          id: params.id as string,
          data: emergencyContactData,
        });
      } else {
        await createEmergencyContactMutation.mutateAsync({
          data: emergencyContactData,
        });
      }

      await queryClient.invalidateQueries({
        queryKey: getGetEmergencyContactsByGuardianIdQueryKey(guardianId),
      });

      if (isEditing) {
        router.back();
      } else {
        router.push("/(auth)/signup/payment");
      }
    } catch {
      Alert.alert(
        translate("common.errorOccurred"),
        translate("emergencyContact.saveError"),
      );
    } finally {
      setIsSubmitting(false);
    }
  };

  const handleDelete = () => {
    Alert.alert(
      translate("emergencyContact.deleteProfile"),
      translate("emergencyContact.deleteConfirm"),
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
                queryKey:
                  getGetEmergencyContactsByGuardianIdQueryKey(guardianId),
              });
              router.back();
            } catch {
              Alert.alert(
                translate("common.errorOccurred"),
                translate("emergencyContact.deleteError"),
              );
              setIsSubmitting(false);
            }
          },
        },
      ],
    );
  };

  return (
    <View className="absolute inset-0">
      <AuthBackground />
      <View className="flex-1" style={{ paddingTop: insets.top }}>
        <Stack.Screen options={{ headerShown: false }} />
        <KeyboardAvoidingView
          behavior={Platform.OS === "ios" ? "padding" : "height"}
          className="flex-1"
        >
          <ScrollView
            contentContainerStyle={{ flexGrow: 1 }}
            keyboardShouldPersistTaps="handled"
            showsVerticalScrollIndicator={false}
          >
            <Pressable onPress={Keyboard.dismiss}>
              {/* Back button */}
              <TouchableOpacity
                onPress={() => router.back()}
                className="flex-row items-center px-5 py-3 gap-1"
                hitSlop={{ top: 10, bottom: 10, left: 10, right: 10 }}
              >
                <IconSymbol name="chevron.left" size={18} color={theme.text} />
                <ThemedText className="text-base font-nunito">
                  {translate("onboarding.back")}
                </ThemedText>
              </TouchableOpacity>

              {/* Title */}
              <View className="px-6 pt-2 pb-10 items-center">
                <ThemedText
                  className="font-nunito-bold text-[#111] text-center"
                  style={{
                    fontSize: FontSizes.hero,
                    lineHeight: FontSizes.hero + 8,
                    letterSpacing: -0.5,
                  }}
                >
                  {isEditing
                    ? translate("emergencyContact.editTitle")
                    : translate("emergencyContact.addTitle")}
                </ThemedText>
              </View>

              {/* Form fields */}
              <View className="px-6 gap-8">
                <View className="gap-1">
                  <ThemedText className="font-nunito-semibold text-base text-[#111]">
                    {translate("childProfile.firstName")}
                  </ThemedText>
                  <TextInput
                    className="border border-[#E5E7EB] rounded-[10px] px-4 py-[14px] bg-white text-base font-nunito text-[#11181C]"
                    value={firstName}
                    onChangeText={setFirstName}
                    autoCapitalize="words"
                    autoCorrect={false}
                    textContentType="givenName"
                    autoComplete="name-given"
                    placeholderTextColor={AppColors.placeholderText}
                  />
                </View>

                <View className="gap-1">
                  <ThemedText className="font-nunito-semibold text-base text-[#111]">
                    {translate("childProfile.lastName")}
                  </ThemedText>
                  <TextInput
                    className="border border-[#E5E7EB] rounded-[10px] px-4 py-[14px] bg-white text-base font-nunito text-[#11181C]"
                    value={lastName}
                    onChangeText={setLastName}
                    autoCapitalize="words"
                    autoCorrect={false}
                    textContentType="familyName"
                    autoComplete="name-family"
                    placeholderTextColor={AppColors.placeholderText}
                  />
                </View>

                <View className="gap-1">
                  <ThemedText className="font-nunito-semibold text-base text-[#111]">
                    {translate("onboarding.contactNumber")}
                  </ThemedText>
                  <TextInput
                    className="border border-[#E5E7EB] rounded-[10px] px-4 py-[14px] bg-white text-base font-nunito text-[#11181C]"
                    value={phoneNumber}
                    onChangeText={setPhoneNumber}
                    keyboardType="phone-pad"
                    autoCapitalize="none"
                    autoCorrect={false}
                    textContentType="telephoneNumber"
                    autoComplete="tel"
                    placeholderTextColor={AppColors.placeholderText}
                  />
                </View>
              </View>
            </Pressable>
          </ScrollView>
        </KeyboardAvoidingView>

        {/* Buttons pinned to bottom */}
        <View
          className="items-center px-6 pt-4"
          style={{ paddingBottom: insets.bottom + 16 }}
        >
          <Button
            label={
              isSubmitting
                ? translate("emergencyContact.saving")
                : isEditing
                  ? translate("emergencyContact.saveChanges")
                  : translate("onboarding.continue")
            }
            onPress={handleSave}
            disabled={isSubmitting}
          />
          {isEditing && (
            <TouchableOpacity className="mt-4" onPress={handleDelete}>
              <ThemedText className="font-nunito-semibold text-[#EF4444]">
                {translate("emergencyContact.deleteContact")}
              </ThemedText>
            </TouchableOpacity>
          )}
        </View>
      </View>
    </View>
  );
}
