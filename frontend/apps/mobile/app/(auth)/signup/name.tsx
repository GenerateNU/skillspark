import { AuthFormInput } from "@/components/AuthFormInput";
import { Button } from "@/components/Button";
import { ErrorMessage } from "@/components/ErrorMessage";
import { ThemedText } from "@/components/themed-text";
import { ThemedView } from "@/components/themed-view";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { AppColors, FontSizes } from "@/constants/theme";
import { useRouter } from "expo-router";
import { useState } from "react";
import { useForm } from "react-hook-form";
import { useTranslation } from "react-i18next";
import { TouchableOpacity, View } from "react-native";

// 2. name and username
export default function NameScreen() {
	const router = useRouter();
	const { t: translate } = useTranslation();
	const [errorText, setErrorText] = useState("");
	const { control } = useForm();

	return (
		<ThemedView className="flex-1">
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
					{translate("onboarding.whatName")}
				</ThemedText>
			</View>

			<View className="px-6 gap-5">
				<View className="gap-2">
					<ThemedText className="text-lg font-nunito-semibold">
						{translate("onboarding.name")}
					</ThemedText>
					<AuthFormInput control={control} name="name" autoCapitalize="none" />
				</View>

				<View className="gap-2">
					<ThemedText className="text-lg font-nunito-semibold">
						{translate("onboarding.username")}
					</ThemedText>
					<AuthFormInput
						control={control}
						name="username"
						autoCapitalize="none"
					/>
				</View>
			</View>

			<View className="px-6 items-center">
				<Button
					label={translate("onboarding.continue")}
					onPress={() => router.push("/(auth)/signup/photo")}
					disabled={false}
					bgColor={"#1B1B1B"}
					width={"91.666667%"}
					textColor={"#FFFFFF"}
				/>

				<ErrorMessage message={errorText} />
			</View>
		</ThemedView>
	);
}
