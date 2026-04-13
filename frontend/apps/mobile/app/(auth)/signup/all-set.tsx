import React, { useEffect } from "react";
import { View, Image } from "react-native";
import { useSafeAreaInsets } from "react-native-safe-area-context";
import { useRouter } from "expo-router";
import { ThemedText } from "@/components/themed-text";
import { ThemedView } from "@/components/themed-view";
import { useTranslation } from "react-i18next";
import { AppColors, FontSizes } from "@/constants/theme";

// 7. done with onboarding
export default function AllSetScreen() {
	const router = useRouter();
	const { t: translate } = useTranslation();
	const insets = useSafeAreaInsets();

	useEffect(() => {
		const timer = setTimeout(() => {
			router.replace("/(app)/(tabs)");
		}, 2500);
		return () => clearTimeout(timer);
	}, [router]);

	return (
		<ThemedView className="flex-1" style={{ paddingTop: insets.top }}>
			<View className="px-6 pt-8 items-center">
				<ThemedText
					className="font-nunito-bold leading-[60px]"
					style={{
						letterSpacing: -0.5,
						fontSize: FontSizes.hero,
						color: AppColors.primaryText,
					}}
				>
					{translate("onboarding.allSet")}
				</ThemedText>
			</View>

			{/* need to add smiley here */}
			<View className="items-center justify-center flex-1">
				<Image
					source={require("@/assets/images/great.png")}
					className="w-36 h-36"
				/>
			</View>

			<View className="px-6 items-center">
				<ThemedText className="font-nunito">
					{translate("onboarding.settingUp")}
				</ThemedText>
			</View>
		</ThemedView>
	);
}
