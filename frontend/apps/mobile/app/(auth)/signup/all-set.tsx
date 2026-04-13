import React, { useState } from "react";
import { View, TouchableOpacity, Image } from "react-native";
import { useRouter } from "expo-router";
import { useSafeAreaInsets } from "react-native-safe-area-context";
import { ThemedText } from "@/components/themed-text";
import { ThemedView } from "@/components/themed-view";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { useTranslation } from "react-i18next";
import { setCurrentLanguage } from "@skillspark/api-client";
import { useAuthContext } from "@/hooks/use-auth-context";
import { ErrorMessage } from "@/components/ErrorMessage";
import { Button } from "@/components/Button";
import { AppColors, FontSizes } from "@/constants/theme";

// 7. done with onboarding
export default function AllSetScreen() {
	const router = useRouter();
	const insets = useSafeAreaInsets();
	const { t: translate, i18n } = useTranslation();

	const [selected, setSelected] = useState(i18n.language ?? "en");
	const [errorText, setErrorText] = useState("");

	return (
		<ThemedView className="flex-1">
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
				<ErrorMessage message={errorText} />
			</View>
		</ThemedView>
	);
}
