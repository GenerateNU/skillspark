import { AuthFormInput } from "@/components/AuthFormInput";
import { Button } from "@/components/Button";
import { ErrorMessage } from "@/components/ErrorMessage";
import { ThemedText } from "@/components/themed-text";
import { ThemedView } from "@/components/themed-view";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { AppColors, FontSizes } from "@/constants/theme";
import { useRouter } from "expo-router";
import { useState } from "react";
import { useFormContext } from "react-hook-form";
import { useTranslation } from "react-i18next";
import { SignupFormData } from "@/constants/signup-types";
import {
	Image,
	KeyboardAvoidingView,
	Platform,
	ScrollView,
	TouchableOpacity,
	View,
} from "react-native";
import { useSafeAreaInsets } from "react-native-safe-area-context";
import { useAuthContext } from "@/hooks/use-auth-context";

// 2. email and password
export default function AccountScreen() {
	const router = useRouter();
	const { t: translate } = useTranslation();
	const insets = useSafeAreaInsets();
	const [errorText, setErrorText] = useState("");
	const [isDisabled, setIsDisabled] = useState(true);
	const { signup } = useAuthContext();
	const { handleSubmit, control, getValues } = useFormContext<SignupFormData>();

	const checkFields = () => {
		const passwordField = getValues("password");
		const confirmPasswordField = getValues("confirm_password");
		const allFilled =
			!!getValues("email") && !!passwordField && !!confirmPasswordField;
		const passwordsMatch = passwordField === confirmPasswordField;

		if (!passwordsMatch) {
			setErrorText(translate("onboarding.passwordMismatch"));
		} else {
			setErrorText("");
		}

		setIsDisabled(!allFilled || !passwordsMatch);
	};

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

	return (
		<ThemedView className="flex-1" style={{ paddingTop: insets.top }}>
			<KeyboardAvoidingView
				behavior={Platform.OS === "ios" ? "padding" : "height"}
				className="flex-1"
				keyboardVerticalOffset={insets.top}
			>
				<ScrollView
					contentContainerStyle={{ flexGrow: 1 }}
					keyboardShouldPersistTaps="handled"
				>
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

					<View className="px-6 pt-8 items-center">
						<ThemedText
							className="font-nunito-bold leading-[60px]"
							style={{
								letterSpacing: -0.5,
								fontSize: FontSizes.hero,
								color: AppColors.primaryText,
							}}
						>
							{translate("onboarding.makeAccount")}
						</ThemedText>
					</View>

					{/* need to add smiley here */}
					<View className="items-center justify-center flex-1">
						<Image
							source={require("@/assets/images/great.png")}
							className="w-36 h-36"
						/>
					</View>

					<View className="px-6 gap-5">
						<View className="gap-2">
							<ThemedText className="text-lg font-nunito-semibold">
								{translate("onboarding.email")}
							</ThemedText>
							<AuthFormInput
								control={control}
								name="email"
								keyboardType="email-address"
								autoCapitalize="none"
								onBlur={(e) => checkFields()}
							/>
						</View>

						<View className="gap-2">
							<ThemedText className="text-lg font-nunito-semibold">
								{translate("onboarding.password")}
							</ThemedText>
							<AuthFormInput
								control={control}
								name="password"
								secureTextEntry
								onBlur={(e) => checkFields()}
							/>
						</View>

						<View className="gap-2">
							<ThemedText className="text-lg font-nunito-semibold">
								{translate("onboarding.confirmPassword")}
							</ThemedText>
							<AuthFormInput
								control={control}
								name="confirm_password"
								secureTextEntry
								onBlur={(e) => checkFields()}
							/>
						</View>
					</View>

					<View className="px-6 items-center">
						<Button
							label={translate("onboarding.createAccount")}
							onPress={handleSubmit(onSubmit)}
							disabled={isDisabled}
						/>
						<ErrorMessage message={errorText} />
					</View>
				</ScrollView>
			</KeyboardAvoidingView>
		</ThemedView>
	);
}
