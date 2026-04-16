import React, { useState } from "react";
import { View, TouchableOpacity } from "react-native";
import { JumpingCharacter } from "@/components/JumpingCharacter";
import { useRouter } from "expo-router";
import { useSafeAreaInsets } from "react-native-safe-area-context";
import { ThemedText } from "@/components/themed-text";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { useTranslation } from "react-i18next";
import { setCurrentLanguage } from "@skillspark/api-client";
import { useAuthContext } from "@/hooks/use-auth-context";
import { Button } from "@/components/Button";
import { useFormContext } from "react-hook-form";
import { SignupFormData } from "@/constants/signup-types";
import { AuthBackground } from "@/components/AuthBackground";

const LANGUAGES = [
	{ code: "en", label: "English", flag: "🇺🇸" },
	{ code: "th", label: "Thai", flag: "🇹🇭" },
];

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
		<View className="absolute inset-0">
			<AuthBackground />
			<View className="flex-1" style={{ paddingTop: insets.top }}>
				<View className="h-11" />
				{/* Title */}
				<View className="items-center px-6 pt-10 pb-5">
					<ThemedText
						className="font-nunito-bold text-[30px] leading-[38px] text-[#111] text-center"
						style={{ letterSpacing: -0.5 }}
						numberOfLines={1}
						adjustsFontSizeToFit
					>
						{translate("onboarding.welcome")}
					</ThemedText>
				</View>

				{/* Character image */}
				<View className="items-center justify-center py-[30px]">
					<JumpingCharacter />
					<ThemedText className="font-nunito-semibold text-base text-[#374151] mt-6 mb-2 text-center">
						{translate("onboarding.chooseLanguage")}
					</ThemedText>
				</View>

				{/* Language section */}
				<View className="px-6 flex-1">
					<View>
						{LANGUAGES.map((lang, index) => {
							const isSelected = selected === lang.code;
							return (
								<React.Fragment key={lang.code}>
									<TouchableOpacity
										onPress={() => updateLanguageData(lang.code)}
										activeOpacity={0.6}
										className="flex-row items-center gap-[14px] py-4"
									>
										<ThemedText className="text-[30px] leading-9">
											{lang.flag}
										</ThemedText>
										<ThemedText className="flex-1 text-base font-nunito text-[#111]">
											{translate(`settings.languages.${lang.code}`)}
										</ThemedText>
										<IconSymbol
											name={isSelected ? "checkmark.circle.fill" : "circle"}
											size={24}
											color={isSelected ? "#1C1C1E" : "#C7C7CC"}
										/>
									</TouchableOpacity>
									{index < LANGUAGES.length - 1 && (
										<View className="h-px bg-[#D1C8F0] ml-[50px]" />
									)}
								</React.Fragment>
							);
						})}
					</View>
				</View>

				{/* Submit button */}
				<View
					className="items-center px-6 pt-4"
					style={{ paddingBottom: insets.bottom + 56 }}
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
