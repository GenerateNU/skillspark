import { AuthFormInput } from "@/components/AuthFormInput";
import { Button } from "@/components/Button";
import { ThemedText } from "@/components/themed-text";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { AppColors, FontSizes } from "@/constants/theme";
import { useRouter } from "expo-router";
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
import { JumpingCharacter } from "@/components/JumpingCharacter";
import { useSafeAreaInsets } from "react-native-safe-area-context";
import { PageRedirectButton } from "@/components/PageRedirectButton";
import { AuthBackground } from "@/components/AuthBackground";

// 1. name and username
export default function NameScreen() {
	const router = useRouter();
	const { t: translate } = useTranslation();
	const insets = useSafeAreaInsets();
	const { control, getValues } = useFormContext<SignupFormData>();

	const handleContinue = () => {
		if (!getValues("name") || !getValues("username")) {
			Alert.alert(
				translate("common.error"),
				//translate("childProfile.requiredFieldsError"),
			);
			return;
		}
		router.push("/(auth)/signup/account");
	};

	return (
		<View style={StyleSheet.absoluteFill}>
			<AuthBackground />
			<View
				style={{ flex: 1, paddingTop: insets.top + 4 }}
			>
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
							<ThemedText
								style={{ fontSize: 16, fontFamily: "NunitoSans_400Regular" }}
							>
								{translate("onboarding.back")}
							</ThemedText>
						</TouchableOpacity>

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
								{translate("onboarding.whatName")}
							</ThemedText>
						</View>

						{/* Character image */}
						<View
							style={{
								flex: 1,
								alignItems: "center",
								justifyContent: "center",
								paddingVertical: 24,
							}}
						>
							<JumpingCharacter />
						</View>

						{/* Form fields */}
						<View style={{ paddingHorizontal: 24, gap: 24 }}>
							<View style={{ gap: 8 }}>
								<ThemedText
									style={{ fontSize: 16, fontFamily: "NunitoSans_600SemiBold" }}
								>
									{translate("onboarding.name")}
								</ThemedText>
								<AuthFormInput
									control={control}
									name="name"
									autoCapitalize="none"
								/>
							</View>

							<View style={{ gap: 8 }}>
								<ThemedText
									style={{ fontSize: 16, fontFamily: "NunitoSans_600SemiBold" }}
								>
									{translate("onboarding.username")}
								</ThemedText>
								<AuthFormInput
									control={control}
									name="username"
									autoCapitalize="none"
								/>
							</View>
						</View>
					</ScrollView>
				</KeyboardAvoidingView>

				{/* Buttons pinned to bottom */}
				<View
					style={{
						alignItems: "center",
						paddingHorizontal: 24,
						paddingTop: 16,
						paddingBottom: insets.bottom + 4,
					}}
				>
					<Button
						label={translate("onboarding.continue")}
						onPress={handleContinue}
						disabled={false}
					/>
					<View style={{ marginTop: 12 }}>
						<PageRedirectButton
							label={translate("onboarding.alreadyHaveAccount")}
							onPress={() => router.navigate("/(auth)/login")}
						/>
					</View>
				</View>
			</View>
		</View>
	);
}
