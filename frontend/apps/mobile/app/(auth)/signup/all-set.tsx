import React, { useEffect } from "react";
import { StyleSheet, View } from "react-native";
import { AuthBackground } from "@/components/AuthBackground";
import { JumpingCharacter } from "@/components/JumpingCharacter";
import { useSafeAreaInsets } from "react-native-safe-area-context";
import { ThemedText } from "@/components/themed-text";
import { useTranslation } from "react-i18next";
import { AppColors } from "@/constants/theme";
import { useAuthContext } from "@/hooks/use-auth-context";

// 7. done with onboarding
export default function AllSetScreen() {
	const { t: translate } = useTranslation();
	const insets = useSafeAreaInsets();
	const { completeOnboarding } = useAuthContext();

	useEffect(() => {
		const timer = setTimeout(async () => {
			await completeOnboarding();
			// LoginRedirect handles navigation once hasAccount=true and inOnboarding=false
		}, 2500);
		return () => clearTimeout(timer);
	}, [completeOnboarding]);

	return (
		<View style={StyleSheet.absoluteFill}>
			<AuthBackground />
			<View style={{ flex: 1, paddingTop: insets.top }}>
			{/* Title */}
			<View style={{ alignItems: "center", paddingHorizontal: 24, paddingTop: 40, paddingBottom: 20 }}>
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
					{translate("onboarding.allSet")}
				</ThemedText>
			</View>

			{/* Character */}
			<View style={{ flex: 1, alignItems: "center", justifyContent: "center" }}>
				<JumpingCharacter />
			</View>

			{/* Status text */}
			<View style={{ paddingHorizontal: 24, paddingBottom: insets.bottom + 24, alignItems: "center" }}>
				<ThemedText
					style={{
						fontSize: 15,
						fontFamily: "NunitoSans_400Regular",
						color: AppColors.mutedText,
						textAlign: "center",
					}}
				>
					{translate("onboarding.settingUp")}
				</ThemedText>
			</View>
			</View>
		</View>
	);
}
