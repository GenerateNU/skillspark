import { AuthFormInput } from "@/components/AuthFormInput";
import { Button } from "@/components/Button";
import { ErrorMessage } from "@/components/ErrorMessage";
import { ThemedText } from "@/components/themed-text";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { AppColors, FontSizes } from "@/constants/theme";
import { useRouter } from "expo-router";
import { useState } from "react";
import { useFormContext } from "react-hook-form";
import { useTranslation } from "react-i18next";
import { SignupFormData } from "@/constants/signup-types";
import {
	Alert,
	KeyboardAvoidingView,
	Platform,
	ScrollView,
	StyleSheet,
	TouchableOpacity,
	View,
} from "react-native";
import { useSafeAreaInsets } from "react-native-safe-area-context";
import { useAuthContext } from "@/hooks/use-auth-context";
import { AuthBackground } from "@/components/AuthBackground";

// 2. email and password
export default function AccountScreen() {
	const router = useRouter();
	const { t: translate } = useTranslation();
	const insets = useSafeAreaInsets();
	const [errorText, setErrorText] = useState("");
	const { signup } = useAuthContext();
	const { handleSubmit, control, getValues } = useFormContext<SignupFormData>();

	const onSubmit = (formData: SignupFormData) => {
		signup(
			formData.name,
			formData.email,
			formData.username,
			formData.password,
			formData.language_preference,
			undefined,
			setErrorText,
			() => router.push("/(auth)/signup/photo"),
		);
	};

	const handleCreateAccount = () => {
		const email = getValues("email");
		const password = getValues("password");
		const confirmPassword = getValues("confirm_password");

		if (!email || !password || !confirmPassword) {
			Alert.alert(
				translate("common.error"),
				translate("childProfile.requiredFieldsError"),
			);
			return;
		}
		if (password !== confirmPassword) {
			Alert.alert(
				translate("common.error"),
				translate("onboarding.passwordMismatch"),
			);
			return;
		}
		handleSubmit(onSubmit)();
	};

	return (
		<View style={StyleSheet.absoluteFill}>
			<AuthBackground />
			<View style={{ flex: 1, paddingTop: insets.top }}>
			<KeyboardAvoidingView
				behavior={Platform.OS === "ios" ? "padding" : "height"}
				style={{ flex: 1 }}
			>
				<ScrollView
					contentContainerStyle={{ flexGrow: 1 }}
					keyboardShouldPersistTaps="handled"
					showsVerticalScrollIndicator={false}
				>
					{/* Back button */}
					<TouchableOpacity
						onPress={() => router.back()}
						style={{
							flexDirection: "row",
							alignItems: "center",
							paddingHorizontal: 20,
							paddingVertical: 12,
							gap: 4,
						}}
						hitSlop={{ top: 10, bottom: 10, left: 10, right: 10 }}
					>
						<IconSymbol name="chevron.left" size={18} color="#11181C" />
						<ThemedText style={{ fontSize: 16, fontFamily: "NunitoSans_400Regular" }}>
							{translate("onboarding.back")}
						</ThemedText>
					</TouchableOpacity>

					{/* Title */}
					<View style={{ paddingHorizontal: 24, paddingTop: 8, alignItems: "center" }}>
						<ThemedText
							style={{
								fontFamily: "NunitoSans_700Bold",
								fontSize: FontSizes.hero,
								lineHeight: 60,
								color: AppColors.primaryText,
								letterSpacing: -0.5,
								textAlign: "center",
							}}
						>
							{translate("onboarding.makeAccount")}
						</ThemedText>
					</View>

					{/* Form fields */}
					<View style={{ paddingHorizontal: 24, gap: 24, paddingTop: 80 }}>
						<View style={{ gap: 8 }}>
							<ThemedText style={{ fontSize: 16, fontFamily: "NunitoSans_600SemiBold" }}>
								{translate("onboarding.email")}
							</ThemedText>
							<AuthFormInput
								control={control}
								name="email"
								keyboardType="email-address"
								autoCapitalize="none"
							/>
						</View>

						<View style={{ gap: 8 }}>
							<ThemedText style={{ fontSize: 16, fontFamily: "NunitoSans_600SemiBold" }}>
								{translate("onboarding.password")}
							</ThemedText>
							<AuthFormInput
								control={control}
								name="password"
								secureTextEntry
							/>
						</View>

						<View style={{ gap: 8 }}>
							<ThemedText style={{ fontSize: 16, fontFamily: "NunitoSans_600SemiBold" }}>
								{translate("onboarding.confirmPassword")}
							</ThemedText>
							<AuthFormInput
								control={control}
								name="confirm_password"
								secureTextEntry
							/>
						</View>
					</View>
				</ScrollView>
			</KeyboardAvoidingView>

			{/* Button pinned to bottom */}
			<View
				style={{
					alignItems: "center",
					paddingHorizontal: 24,
					paddingTop: 16,
					paddingBottom: insets.bottom + 56,
				}}
			>
				<Button
					label={translate("onboarding.createAccount")}
					onPress={handleCreateAccount}
					disabled={false}
				/>
				<ErrorMessage message={errorText} />
			</View>
			</View>
		</View>
	);
}
