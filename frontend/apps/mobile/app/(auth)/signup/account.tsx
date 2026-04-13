import { AuthFormInput } from "@/components/AuthFormInput";
import { Button } from "@/components/Button";
import { ErrorMessage } from "@/components/ErrorMessage";
import { PageRedirectButton } from "@/components/PageRedirectButton";
import { ThemedText } from "@/components/themed-text";
import { ThemedView } from "@/components/themed-view";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { AppColors, FontSizes } from "@/constants/theme";
import { useRouter } from "expo-router";
import { useEffect, useState } from "react";
import { useFormContext } from "react-hook-form";
import { useTranslation } from "react-i18next";
import { SignupFormData } from "@/constants/signup-types";
import { Image, TouchableOpacity, View } from "react-native";
import { useSafeAreaInsets } from "react-native-safe-area-context";

// 1. email and password
export default function AccountScreen() {
	const router = useRouter();
	const { t: translate } = useTranslation();
	const insets = useSafeAreaInsets();
	const [errorText, setErrorText] = useState("");
	const [isDisabled, setIsDisabled] = useState(true);
	const { control, watch } = useFormContext<SignupFormData>();

	const watchEmail = watch("email");
	const watchPassword = watch("password");
	const watchConfirmPassword = watch("confirm_password");

	useEffect(() => {
		const allFilled = !!watchEmail && !!watchPassword && !!watchConfirmPassword;
		const passwordsMatch = watchPassword === watchConfirmPassword;

		if (allFilled && passwordsMatch) {
			setErrorText("");
		}

		setIsDisabled(!allFilled || !passwordsMatch);
	}, [watchEmail, watchPassword, watchConfirmPassword]);

	const checkPasswordMatch = (password: string, confirmPassword: string) => {
		if (watchEmail && watchPassword && watchConfirmPassword) {
			if (password !== confirmPassword) {
				setErrorText(translate("onboarding.passwordMismatch"));
			}
		}
	};

	return (
		<ThemedView className="flex-1" style={{ paddingTop: insets.top }}>
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
						onBlur={() =>
							checkPasswordMatch(watchPassword, watchConfirmPassword)
						}
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
						onBlur={() =>
							checkPasswordMatch(watchPassword, watchConfirmPassword)
						}
					/>
				</View>
			</View>

			<View className="px-6 items-center">
				<Button
					label={translate("onboarding.createAccount")}
					onPress={() => router.push("/(auth)/signup/name")}
					disabled={isDisabled}
				/>

				<PageRedirectButton
					label={translate("onboarding.alreadyHaveAccount")}
					onPress={() => router.navigate("/(auth)/login")}
				/>

				<ErrorMessage message={errorText} />
			</View>
		</ThemedView>
	);
}
