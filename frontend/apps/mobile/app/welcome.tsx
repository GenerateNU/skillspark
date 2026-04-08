import React, { useState } from "react";
import { View, TouchableOpacity } from "react-native";
import { useRouter } from "expo-router";
import { useSafeAreaInsets } from "react-native-safe-area-context";
import { ThemedText } from "@/components/themed-text";
import { ThemedView } from "@/components/themed-view";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { useTranslation } from "react-i18next";
import { setCurrentLanguage } from "@skillspark/api-client";
import { useAuthContext } from "@/hooks/use-auth-context";
import { ErrorMessage } from "@/components/ErrorMessage";

const LANGUAGES = [
	{ code: "en", label: "English", flag: "🇺🇸" },
	{ code: "th", label: "Thai", flag: "🇹🇭" },
];

export default function WelcomeScreen() {
	const router = useRouter();
	const insets = useSafeAreaInsets();
	const { t: translate, i18n } = useTranslation();

	const dividerColor = "#E5E7EB";

	const [selected, setSelected] = useState(i18n.language ?? "en");
	const [errorText, setErrorText] = useState("");
	const { setLanguage } = useAuthContext();

	const updateLanguageData = async (langCode: string) => {
		setSelected(langCode);
		await i18n.changeLanguage(langCode);
		setCurrentLanguage(langCode);

		setLanguage(langCode);
	};

	return (
		<ThemedView className="flex-1" style={{ paddingTop: insets.top }}>
			<View className="flex-row items-center justify-between px-5 py-[14px]">
				<ThemedText className="text-xl text-center font-nunito-bold">
					Welcome to SkillSpark!{/*translate("settings.title")*/}
				</ThemedText>
				<View className="w-10" />
			</View>
			<ThemedText className="text-2xl px-5 pt-4 pb-5 font-nunito-bold">
				{translate("settings.language")}
			</ThemedText>
			<View className="px-5">
				{LANGUAGES.map((lang, index) => (
					<React.Fragment key={lang.code}>
						<TouchableOpacity
							className="flex-row items-center py-[18px] gap-[14px]"
							onPress={() => {
								updateLanguageData(lang.code);
							}}
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
									selected === lang.code ? "checkmark.circle.fill" : "circle"
								}
								size={26}
								color={selected === lang.code ? "#11181C" : "#C7C7CC"}
							/>
						</TouchableOpacity>
						{index < LANGUAGES.length - 1 && (
							<View
								className="h-px ml-[66px]"
								style={{ backgroundColor: dividerColor }}
							/>
						)}
					</React.Fragment>
				))}
			</View>

			<ErrorMessage message={errorText} />
		</ThemedView>
	);
}
