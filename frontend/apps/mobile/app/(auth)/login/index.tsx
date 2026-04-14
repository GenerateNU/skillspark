import { ThemedText } from "@/components/themed-text";
import { ThemedView } from "@/components/themed-view";
import { router } from "expo-router";
import React, { useState } from "react";
import { Image, TouchableOpacity, View } from "react-native";
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

type LoginFormData = {
	email: string;
	password: string;
};

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
		if (formData.email === "" || formData.password === "") {
			setErrorText("Missing email or password");
		} else {
			login(formData.email, formData.password, setErrorText, () =>
				router.push("/(app)/(tabs)"),
			);
		}
	};

	const handleGoToSignUp = () => {
		router.navigate("/(auth)/signup/name");
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
					{translate("onboarding.signIn")}
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
					<AuthFormInput control={control} name="password" secureTextEntry />
				</View>
			</View>

			<View className="px-6 items-center">
				<Button
					label={translate("onboarding.signUp")}
					onPress={handleSubmit(onSubmit)}
					disabled={false}
				/>
				<PageRedirectButton
					label={translate("onboarding.dontHaveAccount")}
					onPress={handleGoToSignUp}
				/>
				<ErrorMessage message={errorText} />
			</View>
		</ThemedView>
	);
}
