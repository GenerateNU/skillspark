import React from "react";
import { FontSizes } from "@/constants/theme";
import { View } from "react-native";
import { AuthBackground } from "@/components/AuthBackground";
import { SvgXml } from "react-native-svg";
import { Button } from "@/components/Button";
import { useSafeAreaInsets } from "react-native-safe-area-context";
import { ThemedText } from "@/components/themed-text";
import { useTranslation } from "react-i18next";
import { useAuthContext } from "@/hooks/use-auth-context";
import { useRouter } from "expo-router";

const ALL_SET_SVG = `<svg width="192" height="192" viewBox="0 0 192 192" fill="none" xmlns="http://www.w3.org/2000/svg">
<rect width="191.7" height="191.7" rx="95.85" fill="#B0F19B"/>
<path d="M136.18 77.9219C137.102 77.9114 137.859 78.6505 137.87 79.5732L137.997 90.8018C138.134 102.896 128.442 112.812 116.348 112.949C104.253 113.086 94.3375 103.393 94.2002 91.2988L94.0732 80.0703C94.063 79.1477 94.802 78.3915 95.7246 78.3809L136.18 77.9219ZM83.0752 52.7998C85.4361 52.7999 87.3496 54.7142 87.3496 57.0752C87.3495 59.4361 85.4361 61.3495 83.0752 61.3496C80.7142 61.3496 78.7999 59.4361 78.7998 57.0752C78.7998 54.7142 80.7142 52.7998 83.0752 52.7998ZM149.225 52.7998C151.586 52.7998 153.5 54.7142 153.5 57.0752C153.5 59.4361 151.586 61.3496 149.225 61.3496C146.864 61.3494 144.95 59.436 144.95 57.0752C144.95 54.7143 146.864 52.8 149.225 52.7998Z" fill="black" stroke="black" stroke-width="3.6"/>
</svg>`;

// 7. done with onboarding
export default function AllSetScreen() {
	const { t: translate } = useTranslation();
	const insets = useSafeAreaInsets();
	const { completeOnboarding } = useAuthContext();
	const router = useRouter();

	const handleContinue = async () => {
		await completeOnboarding();
		router.replace("/(app)/(tabs)");
	};

	return (
		<View className="absolute inset-0">
			<AuthBackground />
			<View className="flex-1" style={{ paddingTop: insets.top }}>
				{/* Title */}
				<View className="items-center px-6 pt-10 pb-5">
					<ThemedText
						className="font-nunito-bold text-[#111] text-center"
						style={{ fontSize: FontSizes.hero, lineHeight: FontSizes.hero + 8, letterSpacing: -0.5 }}
						numberOfLines={1}
						adjustsFontSizeToFit
					>
						{translate("onboarding.allSet")}
					</ThemedText>
				</View>

				{/* Character */}
				<View className="flex-1 items-center justify-center">
					<SvgXml xml={ALL_SET_SVG} width={192} height={192} />
				</View>

				{/* Continue button */}
				<View
					className="px-6 items-center"
					style={{ paddingBottom: insets.bottom + 24 }}
				>
					<Button
						label={translate("onboarding.continue")}
						onPress={handleContinue}
						disabled={false}
					/>
				</View>
			</View>
		</View>
	);
}
