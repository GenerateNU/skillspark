import React, { useState } from "react";
import {
	View,
	TouchableOpacity,
	Image,
	KeyboardAvoidingView,
	ScrollView,
	Platform,
} from "react-native";
import { useRouter } from "expo-router";
import { useSafeAreaInsets } from "react-native-safe-area-context";
import { ThemedText } from "@/components/themed-text";
import { ThemedView } from "@/components/themed-view";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { useTranslation } from "react-i18next";
import { setCurrentLanguage } from "@skillspark/api-client";
import { useAuthContext } from "@/hooks/use-auth-context";
import { Button } from "@/components/Button";
import { AppColors, FontSizes } from "@/constants/theme";
import { useFormContext } from "react-hook-form";
import { SignupFormData } from "@/constants/signup-types";
import { PageRedirectButton } from "@/components/PageRedirectButton";

const LANGUAGES = [
	{ code: "en", label: "English", flag: "🇺🇸" },
	{ code: "th", label: "Thai", flag: "🇹🇭" },
];

export default function WelcomeScreen() {
	const router = useRouter();
	const insets = useSafeAreaInsets();
	const { t: translate, i18n } = useTranslation();

	const dividerColor = "#E5E7EB";

	const { setValue } = useFormContext<SignupFormData>();
	const [selected, setSelected] = useState(i18n.language ?? "en");
	setValue("language_preference", selected);
	const { setLanguage } = useAuthContext();

	const updateLanguageData = async (langCode: string) => {
		setSelected(langCode);
		await i18n.changeLanguage(langCode);
		setCurrentLanguage(langCode);
		setLanguage(langCode);
		setValue("language_preference", langCode);
		console.log(langCode);
	};

	return (
		<ThemedView className="flex-1" style={{ paddingTop: insets.top }}>
			<KeyboardAvoidingView
				behavior={Platform.OS === "ios" ? "padding" : "height"}
				className="flex-1"
			>
				<ScrollView
					contentContainerStyle={{ flexGrow: 1 }}
					keyboardShouldPersistTaps="handled"
					scrollEnabled={false}
				>
					<View className="px-6 pt-8 items-center">
						<ThemedText
							className="font-nunito-bold leading-[60px]"
							style={{
								letterSpacing: -0.5,
								fontSize: FontSizes.hero,
								color: AppColors.primaryText,
							}}
						>
							{translate("onboarding.welcome")}
						</ThemedText>
					</View>

					{/* need to add smiley here */}
					<View className="items-center justify-center flex-1">
						<Image
							source={require("@/assets/images/great.png")}
							className="w-36 h-36"
						/>
					</View>

					<ThemedText className="text-xl font-nunito-bold text-center mb-6">
						{translate("onboarding.chooseLanguage")}
					</ThemedText>
					<View className="px-6 mb-8">
						{LANGUAGES.map((lang, index) => (
							<React.Fragment key={lang.code}>
								<TouchableOpacity
									className="flex-row items-center py-5 gap-4"
									onPress={() => updateLanguageData(lang.code)}
									activeOpacity={0.6}
								>
									<ThemedText className="text-[38px] leading-[46px]">
										{lang.flag}
									</ThemedText>
									<ThemedText className="flex-1 text-lg font-nunito">
										{translate(`settings.languages.${lang.code}`)}
									</ThemedText>
									<IconSymbol
										name={
											selected === lang.code
												? "checkmark.circle.fill"
												: "circle"
										}
										size={26}
										color={selected === lang.code ? "#1C1C1E" : "#C7C7CC"}
									/>
								</TouchableOpacity>
								{index < LANGUAGES.length - 1 && (
									<View
										className="h-px ml-[58px]"
										style={{ backgroundColor: dividerColor }}
									/>
								)}
							</React.Fragment>
						))}
					</View>

					<View
						className="px-6 items-center"
						style={{ paddingBottom: insets.bottom + 16 }}
					>
						<Button
							label={translate("common.submit")}
							onPress={() => router.push("/(auth)/signup/name")}
							disabled={false}
						/>
					</View>
				</ScrollView>
			</KeyboardAvoidingView>
		</ThemedView>
	);
}
