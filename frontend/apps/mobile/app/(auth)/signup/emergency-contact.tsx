import React, { useState } from "react";
import {
	View,
	TouchableOpacity,
	Alert,
	ScrollView,
	KeyboardAvoidingView,
	Platform,
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
import { AuthFormInput } from "@/components/AuthFormInput";
import { PageRedirectButton } from "@/components/PageRedirectButton";
import { useForm } from "react-hook-form";
import { queryClient } from "@/constants/query-client";
import {
	getGetEmergencyContactsByGuardianIdQueryKey,
	useCreateEmergencyContact,
	useDeleteEmergencyContact,
	useUpdateEmergencyContact,
} from "@skillspark/api-client";


type EmergencyContactFormData = {
	name: string;
	phone_number: string;
};

// screen for adding an emergency contact
export default function ManageEmergencyContactScreen() {
	const router = useRouter();
	const params = useLocalSearchParams();
	const insets = useSafeAreaInsets();

	const { guardianId } = useAuthContext();

	const createEmergencyContactMutation = useCreateEmergencyContact();
	const updateEmergencyContactMutation = useUpdateEmergencyContact();
	const deleteEmergencyContactMutation = useDeleteEmergencyContact();

	const { t: translate } = useTranslation();
	const isEditing = !!params.id;
	const [isSubmitting, setIsSubmitting] = useState(false);

	const { control, getValues } = useForm<EmergencyContactFormData>({
		defaultValues: {
			name: (params.name as string) || "",
			phone_number: (params.phone_number as string) || "",
		},
	});

	if (!guardianId) {
		return <ErrorScreen message="Illegal state: no guardian ID retrieved" />;
	}

	const isValidPhoneNumber = (phone: string) => {
		const phoneValidationRegex =
			/^[\+]?[(]?[0-9]{3}[)]?[-\s\.]?[0-9]{3}[-\s\.]?[0-9]{4,6}$/im;
		return phoneValidationRegex.test(phone);
	};

	const handleSave = async () => {
		const name = getValues("name");
		const phone_number = getValues("phone_number");

		if (!name || !phone_number) {
			Alert.alert(
				translate("common.error"),
				translate("childProfile.requiredFieldsError"),
			);
			return;
		}

		if (!isValidPhoneNumber(phone_number)) {
			Alert.alert(
				translate("common.error"),
				translate("emergencyContact.invalidPhoneNumber"),
			);
			return;
		}

		setIsSubmitting(true);
		try {
			const emergencyContactData = {
				guardian_id: guardianId,
				name,
				phone_number,
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
								queryKey:
									getGetEmergencyContactsByGuardianIdQueryKey(guardianId),
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
						{/* Back button */}
						<TouchableOpacity
							onPress={() => router.back()}
							className="flex-row items-center px-5 py-3 gap-1"
							hitSlop={{ top: 10, bottom: 10, left: 10, right: 10 }}
						>
							<IconSymbol name="chevron.left" size={18} color="#11181C" />
							<ThemedText className="text-base font-nunito">
								{translate("onboarding.back")}
							</ThemedText>
						</TouchableOpacity>

						{/* Title */}
						<View className="px-6 pt-2 items-center">
							<ThemedText
								className="font-nunito-bold leading-[60px] text-[#111] text-[30px] text-center"
								style={{ letterSpacing: -0.5 }}
							>
								{isEditing
									? translate("emergencyContact.editTitle")
									: translate("emergencyContact.addTitle")}
							</ThemedText>
						</View>

						{/* Form fields */}
						<View className="px-6 gap-6 pt-20">
							<View className="gap-2">
								<ThemedText className="text-base font-nunito-semibold">
									{translate("emergencyContact.name")}
								</ThemedText>
								<AuthFormInput
									control={control}
									name="name"
									autoCapitalize="words"
								/>
							</View>

							<View className="gap-2">
								<ThemedText className="text-base font-nunito-semibold">
									{translate("emergencyContact.phoneNumber")}
								</ThemedText>
								<AuthFormInput
									control={control}
									name="phone_number"
									keyboardType="phone-pad"
									autoCapitalize="none"
								/>
							</View>
						</View>
					</ScrollView>
				</KeyboardAvoidingView>

				{/* Buttons pinned to bottom */}
				<View
					className="items-center px-6 pt-4"
					style={{ paddingBottom: insets.bottom + 56 }}
				>
					<Button
						label={
							isSubmitting
								? translate("emergencyContact.saving")
								: isEditing
									? translate("emergencyContact.saveChanges")
									: translate("emergencyContact.addContact")
						}
						onPress={handleSave}
						disabled={isSubmitting}
					/>
					{isEditing && (
						<View className="mt-3">
							<PageRedirectButton
								label={translate("emergencyContact.deleteContact")}
								onPress={handleDelete}
							/>
						</View>
					)}
				</View>
			</View>
		</View>
	);
}
