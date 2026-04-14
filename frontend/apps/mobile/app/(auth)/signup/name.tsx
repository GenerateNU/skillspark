import { AuthFormInput } from "@/components/AuthFormInput";
import { Button } from "@/components/Button";
import { ErrorMessage } from "@/components/ErrorMessage";
import { ThemedText } from "@/components/themed-text";
import { ThemedView } from "@/components/themed-view";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { AppColors, FontSizes } from "@/constants/theme";
import { useRouter } from "expo-router";
import { useEffect, useState } from "react";
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
import { PageRedirectButton } from "@/components/PageRedirectButton";

// 1. name and username
export default function NameScreen() {
	const router = useRouter();
	const { t: translate } = useTranslation();
	const insets = useSafeAreaInsets();
	const [errorText, setErrorText] = useState("");
	const [isDisabled, setIsDisabled] = useState(true);
	const { control, getValues } = useFormContext<SignupFormData>();

	const checkEmptyField = () => {
		if (!getValues("name") || !getValues("username")) {
			setIsDisabled(true);
		} else {
			setIsDisabled(false);
		}
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
							{translate("onboarding.whatName")}
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
								{translate("onboarding.name")}
							</ThemedText>
							<AuthFormInput
								control={control}
								name="name"
								autoCapitalize="none"
								onBlur={checkEmptyField}
							/>
						</View>

						<View className="gap-2">
							<ThemedText className="text-lg font-nunito-semibold">
								{translate("onboarding.username")}
							</ThemedText>
							<AuthFormInput
								control={control}
								name="username"
								autoCapitalize="none"
								onBlur={checkEmptyField}
							/>
						</View>
					</View>

					<View className="px-6 items-center">
						<Button
							label={translate("onboarding.continue")}
							onPress={() => router.push("/(auth)/signup/account")}
							disabled={isDisabled}
						/>

						<PageRedirectButton
							label={translate("onboarding.alreadyHaveAccount")}
							onPress={() => router.navigate("/(auth)/login")}
						/>

						<ErrorMessage message={errorText} />
					</View>
				</ScrollView>
			</KeyboardAvoidingView>
		</ThemedView>
	);
}
