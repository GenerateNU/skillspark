import React, { useState } from "react";
import { StyleSheet, View, TouchableOpacity } from "react-native";
import { JumpingCharacter } from "@/components/JumpingCharacter";
import { useRouter } from "expo-router";
import { useSafeAreaInsets } from "react-native-safe-area-context";
import { ThemedText } from "@/components/themed-text";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { useTranslation } from "react-i18next";
import { setCurrentLanguage } from "@skillspark/api-client";
import { useAuthContext } from "@/hooks/use-auth-context";
import { Button } from "@/components/Button";
import { AppColors } from "@/constants/theme";
import { useFormContext } from "react-hook-form";
import { SignupFormData } from "@/constants/signup-types";
import { AuthBackground } from "@/components/AuthBackground";

const LANGUAGES = [
	{ code: "en", label: "English", flag: "🇺🇸" },
	{ code: "th", label: "Thai", flag: "🇹🇭" },
];

const DIVIDER = "#D1C8F0";

export default function WelcomeScreen() {
	const router = useRouter();
	const insets = useSafeAreaInsets();
	const { t: translate, i18n } = useTranslation();

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
	};

	return (
		<View style={StyleSheet.absoluteFill}>
			<AuthBackground />
			<View style={{ flex: 1, paddingTop: insets.top }}>
				<View style={{ height: 44, justifyContent: "center" }} />
				{/* Title */}
				<View
					style={{
						alignItems: "center",
						paddingHorizontal: 24,
						paddingTop: 40,
						paddingBottom: 20,
					}}
				>
					<ThemedText
						style={{
							fontFamily: "NunitoSans_700Bold",
							fontSize: 30,
							lineHeight: 38,
							color: AppColors.primaryText,
							letterSpacing: -0.5,
							textAlign: "center",
						}}
						numberOfLines={1}
						adjustsFontSizeToFit
					>
						{translate("onboarding.welcome")}
					</ThemedText>
				</View>

				{/* Character image */}
				<View
					style={{
						alignItems: "center",
						justifyContent: "center",
						paddingVertical: 30,
					}}
				>
					<JumpingCharacter />
					<ThemedText
						style={{
							fontFamily: "NunitoSans_600SemiBold",
							fontSize: 16,
							color: AppColors.secondaryText,
							marginTop: 24,
							marginBottom: 8,
							textAlign: "center",
						}}
					>
						{translate("onboarding.chooseLanguage")}
					</ThemedText>
				</View>

				{/* Language section */}
				<View style={{ paddingHorizontal: 24, flex: 1 }}>
					<View>
						{LANGUAGES.map((lang, index) => {
							const isSelected = selected === lang.code;
							return (
								<React.Fragment key={lang.code}>
									<TouchableOpacity
										onPress={() => updateLanguageData(lang.code)}
										activeOpacity={0.6}
										style={{
											flexDirection: "row",
											alignItems: "center",
											gap: 14,
											paddingVertical: 16,
										}}
									>
										<ThemedText style={{ fontSize: 30, lineHeight: 36 }}>
											{lang.flag}
										</ThemedText>
										<ThemedText
											style={{
												flex: 1,
												fontSize: 16,
												fontFamily: "NunitoSans_400Regular",
												color: AppColors.primaryText,
											}}
										>
											{translate(`settings.languages.${lang.code}`)}
										</ThemedText>
										<IconSymbol
											name={isSelected ? "checkmark.circle.fill" : "circle"}
											size={24}
											color={isSelected ? "#1C1C1E" : "#C7C7CC"}
										/>
									</TouchableOpacity>
									{index < LANGUAGES.length - 1 && (
										<View
											style={{
												height: 1,
												backgroundColor: DIVIDER,
												marginLeft: 50,
											}}
										/>
									)}
								</React.Fragment>
							);
						})}
					</View>
				</View>

				{/* Submit button */}
				<View
					style={{
						alignItems: "center",
						paddingHorizontal: 24,
						paddingBottom: insets.bottom + 56,
						paddingTop: 16,
					}}
				>
					<Button
						label={translate("common.submit")}
						onPress={() => router.push("/(auth)/signup/name")}
						disabled={false}
					/>
				</View>
			</View>
		</View>
	);
}
