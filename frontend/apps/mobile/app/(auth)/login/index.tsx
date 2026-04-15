import { ThemedText } from "@/components/themed-text";
import { router } from "expo-router";
import React, { useState } from "react";
import {
	Alert,
	KeyboardAvoidingView,
	Platform,
	ScrollView,
	TouchableOpacity,
	View,
} from "react-native";
import { useAuthContext } from "@/hooks/use-auth-context";
import { useForm } from "react-hook-form";
import { ErrorMessage } from "@/components/ErrorMessage";
import { PageRedirectButton } from "@/components/PageRedirectButton";
import { Button } from "@/components/Button";
import { AuthFormInput } from "@/components/AuthFormInput";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { useTranslation } from "react-i18next";
import { useSafeAreaInsets } from "react-native-safe-area-context";
import { AppColors, FontSizes } from "@/constants/theme";
import { JumpingCharacter } from "@/components/JumpingCharacter";

type LoginFormData = {
	email: string;
	password: string;
};

const BG = "#EDE8FF";

export default function LoginScreen() {
	const insets = useSafeAreaInsets();
	const [errorText, setErrorText] = useState("");
	const { t: translate } = useTranslation();
	const { login } = useAuthContext();

	const { control, handleSubmit } = useForm<LoginFormData>({
		defaultValues: {
			email: "",
			password: "",
		},
	});

	const onSubmit = (formData: LoginFormData) => {
		if (!formData.email || !formData.password) {
			Alert.alert(
				translate("common.error"),
				translate("childProfile.requiredFieldsError"),
			);
			return;
		}
		login(formData.email, formData.password, setErrorText, () =>
			router.push("/(app)/(tabs)"),
		);
	};

	return (
		<View style={{ flex: 1, backgroundColor: BG, paddingTop: insets.top }}>
			<KeyboardAvoidingView
				behavior={Platform.OS === "ios" ? "padding" : "height"}
				style={{ flex: 1 }}
			>
				<ScrollView
					contentContainerStyle={{
						flexGrow: 1,
						paddingBottom: insets.bottom + 24,
					}}
					keyboardShouldPersistTaps="handled"
					showsVerticalScrollIndicator={false}
				>
					{/* Back button — only shown when there is a previous screen to return to; space always reserved */}
					<View style={{ height: 44, justifyContent: "center" }}>
						{router.canGoBack() && (
							<TouchableOpacity
								onPress={() => router.back()}
								style={{
									flexDirection: "row",
									alignItems: "center",
									paddingHorizontal: 20,
									gap: 4,
								}}
								hitSlop={{ top: 10, bottom: 10, left: 10, right: 10 }}
							>
								<IconSymbol name="chevron.left" size={18} color="#11181C" />
								<ThemedText
									style={{ fontSize: 16, fontFamily: "NunitoSans_400Regular" }}
								>
									{translate("onboarding.back")}
								</ThemedText>
							</TouchableOpacity>
						)}
					</View>

					{/* Title */}
					<View
						style={{
							paddingHorizontal: 24,
							paddingTop: 8,
							alignItems: "center",
						}}
					>
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
							{translate("onboarding.signIn")}
						</ThemedText>
					</View>

					{/* Character image */}
					<View style={{ alignItems: "center", paddingVertical: 24 }}>
						<JumpingCharacter />
					</View>

					{/* Form fields */}
					<View style={{ paddingHorizontal: 24, gap: 24 }}>
						<View style={{ gap: 8 }}>
							<ThemedText
								style={{ fontSize: 16, fontFamily: "NunitoSans_600SemiBold" }}
							>
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
							<ThemedText
								style={{ fontSize: 16, fontFamily: "NunitoSans_600SemiBold" }}
							>
								{translate("onboarding.password")}
							</ThemedText>
							<AuthFormInput
								control={control}
								name="password"
								secureTextEntry
							/>
						</View>
					</View>

					{/* Buttons */}
					<View
						style={{
							paddingHorizontal: 24,
							paddingTop: 32,
							alignItems: "center",
						}}
					>
						<Button
							label={translate("onboarding.signUp")}
							onPress={handleSubmit(onSubmit)}
							disabled={false}
						/>
						<PageRedirectButton
							label={translate("onboarding.dontHaveAccount")}
							onPress={() => router.navigate("/(auth)/signup/name")}
						/>
						<ErrorMessage message={errorText} />
					</View>
				</ScrollView>
			</KeyboardAvoidingView>
		</View>
	);
}
