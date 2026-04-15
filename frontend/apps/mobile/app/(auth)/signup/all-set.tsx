import React, { useEffect } from "react";
import { View } from "react-native";
import { JumpingCharacter } from "@/components/JumpingCharacter";
import { useSafeAreaInsets } from "react-native-safe-area-context";
import { ThemedText } from "@/components/themed-text";
import { useTranslation } from "react-i18next";
import { AppColors, FontSizes } from "@/constants/theme";
import { useAuthContext } from "@/hooks/use-auth-context";

const BG = "#EDE8FF";

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
		<View style={{ flex: 1, backgroundColor: BG, paddingTop: insets.top }}>
			{/* Top section: title + character */}
			<View style={{ flex: 1, alignItems: "center", justifyContent: "center", paddingHorizontal: 24 }}>
				<ThemedText
					style={{
						fontFamily: "NunitoSans_700Bold",
						fontSize: FontSizes.hero,
						color: AppColors.primaryText,
						letterSpacing: -0.5,
						textAlign: "center",
						marginBottom: 28,
					}}
				>
					{translate("onboarding.allSet")}
				</ThemedText>
				<JumpingCharacter />
			</View>

			{/* Bottom white card: status text */}
			<View
				style={{
					backgroundColor: "#FFFFFF",
					borderTopLeftRadius: 28,
					borderTopRightRadius: 28,
					paddingHorizontal: 24,
					paddingTop: 28,
					paddingBottom: insets.bottom + 24,
					alignItems: "center",
				}}
			>
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
	);
}
