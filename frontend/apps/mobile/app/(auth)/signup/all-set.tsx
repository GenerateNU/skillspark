import React, { useEffect } from "react";
import { View } from "react-native";
import { AuthBackground } from "@/components/AuthBackground";
import { JumpingCharacter } from "@/components/JumpingCharacter";
import { useSafeAreaInsets } from "react-native-safe-area-context";
import { ThemedText } from "@/components/themed-text";
import { useTranslation } from "react-i18next";
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
		<View className="absolute inset-0">
			<AuthBackground />
			<View className="flex-1" style={{ paddingTop: insets.top }}>
				{/* Title */}
				<View className="items-center px-6 pt-10 pb-5">
					<ThemedText
						className="font-nunito-bold text-[30px] leading-[38px] text-[#111] text-center"
						style={{ letterSpacing: -0.5 }}
						numberOfLines={1}
						adjustsFontSizeToFit
					>
						{translate("onboarding.allSet")}
					</ThemedText>
				</View>

				{/* Character */}
				<View className="flex-1 items-center justify-center">
					<JumpingCharacter />
				</View>

				{/* Status text */}
				<View
					className="px-6 items-center"
					style={{ paddingBottom: insets.bottom + 24 }}
				>
					<ThemedText className="text-[15px] font-nunito text-[#6B7280] text-center">
						{translate("onboarding.settingUp")}
					</ThemedText>
				</View>
			</View>
		</View>
	);
}
