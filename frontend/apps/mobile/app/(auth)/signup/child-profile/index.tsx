import { Button } from "@/components/Button";
import { ErrorMessage } from "@/components/ErrorMessage";
import { ThemedText } from "@/components/themed-text";
import { ThemedView } from "@/components/themed-view";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { AppColors, FontSizes } from "@/constants/theme";
import { useRouter } from "expo-router";
import { SetStateAction, useState } from "react";
import { useForm } from "react-hook-form";
import { useTranslation } from "react-i18next";
import { Image, TouchableOpacity, View } from "react-native";
import { useSafeAreaInsets } from "react-native-safe-area-context";

// 4. set up your child's profile
// add new child profile -> go to "set up your child's profile page"
export default function ChildProfileScreen() {
	const router = useRouter();
	const { t: translate } = useTranslation();
	const insets = useSafeAreaInsets();
	const [errorText, setErrorText] = useState("");
	const { control } = useForm();

	return (
		<ThemedView className="flex-1" style={{ paddingTop: insets.top }}>
			<TouchableOpacity
				onPress={() => router.back()}
				className="flex-row items-center px-5 py-3 gap-1"
				hitSlop={{ top: 10, bottom: 10, left: 10, right: 10 }}
			>
				<IconSymbol name="chevron.left" size={18} color="#11181C" />
				<ThemedText className="text-base font-nunito">
					{translate("onboarding.back")}
				</ThemedText>
			</TouchableOpacity>

			<View className="px-6 pt-8 items-center">
				<ThemedText
					className="font-nunito-bold leading-[60px]"
					style={{
						letterSpacing: -0.5,
						fontSize: FontSizes.hero,
						color: AppColors.primaryText,
					}}
				>
					{translate("onboarding.setUpChild")}
				</ThemedText>
			</View>

			<View className="px-6 items-center">
				<Button
					label={translate("onboarding.addNewChild")}
					onPress={() => router.push("/(auth)/signup/child-profile/add-child")}
					disabled={false}
					bgColor={"#FFFFFF"}
					width={"60%"}
					textColor={"#1B1B1B"}
				/>
			</View>

			<View className="px-6 items-center">
				<Button
					label={translate("onboarding.continue")}
					onPress={() => router.push("/(auth)/signup/emergency-contact")}
					disabled={false}
				/>
				<ErrorMessage message={errorText} />
			</View>
		</ThemedView>
	);
}
