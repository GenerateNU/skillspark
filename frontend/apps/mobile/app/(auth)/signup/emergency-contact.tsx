import React, { useState } from "react";
import {
	View,
	TouchableOpacity,
	Alert,
	ScrollView,
	KeyboardAvoidingView,
	Platform,
} from "react-native";
import { Stack, useRouter, useLocalSearchParams } from "expo-router";
import { useSafeAreaInsets } from "react-native-safe-area-context";
import { ThemedText } from "@/components/themed-text";
import { AppColors } from "@/constants/theme";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { useTranslation } from "react-i18next";
import { useAuthContext } from "@/hooks/use-auth-context";
import { ErrorScreen } from "@/components/ErrorScreen";
import { EmergencyContactForm } from "@/components/EmergencyContactProfileForm";
import { queryClient } from "@/constants/query-client";
import {
	getGetEmergencyContactsByGuardianIdQueryKey,
	useCreateEmergencyContact,
	useDeleteEmergencyContact,
	useUpdateEmergencyContact,
} from "@skillspark/api-client";

const BG = "#EDE8FF";

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
	const [phoneNumber, setPhoneNumber] = useState(
		(params.phone_number as string) || "",
	);
	const [name, setName] = useState((params.name as string) || "");
	const [isSubmitting, setIsSubmitting] = useState(false);

	if (!guardianId) {
		return <ErrorScreen message="Illegal state: no guardian ID retrieved" />;
	}

	const isValidPhoneNumber = (phoneNumber: string) => {
		const phoneValidationRegex =
			/^[\+]?[(]?[0-9]{3}[)]?[-\s\.]?[0-9]{3}[-\s\.]?[0-9]{4,6}$/im;
		return phoneValidationRegex.test(phoneNumber);
	};

	const emergencyContactData = {
		guardian_id: guardianId,
		name: name,
		phone_number: phoneNumber,
	};

	const handleSave = async () => {
		if (!name || !phoneNumber) {
			Alert.alert(
				translate("common.error"),
				translate("childProfile.requiredFieldsError"),
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
		try {
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
		<View style={{ flex: 1, backgroundColor: BG, paddingTop: insets.top }}>
			<Stack.Screen options={{ headerShown: false }} />
			<KeyboardAvoidingView
				behavior={Platform.OS === "ios" ? "padding" : "height"}
				style={{ flex: 1 }}
				keyboardVerticalOffset={0}
			>
				{/* Header row in lavender area */}
				<View
					style={{
						flexDirection: "row",
						alignItems: "center",
						justifyContent: "space-between",
						paddingHorizontal: 20,
						paddingVertical: 14,
					}}
				>
					<TouchableOpacity
						onPress={() => router.back()}
						style={{ flexDirection: "row", alignItems: "center", gap: 4 }}
						hitSlop={{ top: 10, bottom: 10, left: 10, right: 10 }}
					>
						<IconSymbol name="chevron.left" size={18} color="#11181C" />
						<ThemedText style={{ fontSize: 16, fontFamily: "NunitoSans_400Regular" }}>
							{translate("onboarding.back")}
						</ThemedText>
					</TouchableOpacity>

					<ThemedText style={{ fontSize: 18, fontFamily: "NunitoSans_700Bold" }}>
						{translate("profile.familyInformation")}
					</ThemedText>

					{isEditing ? (
						<TouchableOpacity onPress={handleDelete}>
							<ThemedText style={{ fontFamily: "NunitoSans_600SemiBold", color: AppColors.danger }}>
								{translate("emergencyContact.deleteContact")}
							</ThemedText>
						</TouchableOpacity>
					) : (
						<View style={{ width: 40 }} />
					)}
				</View>

				{/* White card with form */}
				<ScrollView
					contentContainerStyle={{ flexGrow: 1 }}
					keyboardShouldPersistTaps="handled"
					showsVerticalScrollIndicator={false}
				>
					<View
						style={{
							flex: 1,
							backgroundColor: "#FFFFFF",
							borderTopLeftRadius: 28,
							borderTopRightRadius: 28,
							paddingHorizontal: 24,
							paddingTop: 28,
							paddingBottom: insets.bottom + 24,
						}}
					>
						<ThemedText
							style={{
								fontSize: 20,
								fontFamily: "NunitoSans_600SemiBold",
								color: AppColors.primaryText,
								marginBottom: 20,
							}}
						>
							{isEditing
								? translate("emergencyContact.editTitle")
								: translate("emergencyContact.addTitle")}
						</ThemedText>

						<EmergencyContactForm
							name={name}
							setName={setName}
							phoneNumber={phoneNumber}
							setPhoneNumber={setPhoneNumber}
						/>

						<TouchableOpacity
							style={{
								backgroundColor: "#1C1C1E",
								borderRadius: 24,
								paddingVertical: 16,
								alignItems: "center",
								marginTop: 8,
								opacity: isSubmitting ? 0.7 : 1,
							}}
							onPress={handleSave}
							disabled={isSubmitting}
							activeOpacity={0.8}
						>
							<ThemedText style={{ color: "#FFFFFF", fontSize: 16, fontFamily: "NunitoSans_600SemiBold" }}>
								{isSubmitting
									? translate("emergencyContact.saving")
									: isEditing
										? translate("emergencyContact.saveChanges")
										: translate("emergencyContact.addContact")}
							</ThemedText>
						</TouchableOpacity>
					</View>
				</ScrollView>
			</KeyboardAvoidingView>
		</View>
	);
}
