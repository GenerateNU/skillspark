import { AuthFormInput } from "@/components/AuthFormInput";
import { Button } from "@/components/Button";
import { ErrorMessage } from "@/components/ErrorMessage";
import { ThemedText } from "@/components/themed-text";
import { ThemedView } from "@/components/themed-view";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { AppColors, FontSizes } from "@/constants/theme";
import { useRouter } from "expo-router";
import { useEffect, useState } from "react";
import { useFormContext } from "react-hook-form";
import { useTranslation } from "react-i18next";
import { SignupFormData } from "@/constants/signup-types";
import { TouchableOpacity, View } from "react-native";
import { useSafeAreaInsets } from "react-native-safe-area-context";

// 2. name and username
export default function NameScreen() {
	const router = useRouter();
	const { t: translate } = useTranslation();
	const insets = useSafeAreaInsets();
	const [errorText, setErrorText] = useState("");
	const [isDisabled, setIsDisabled] = useState(true);
	const { control, watch } = useFormContext<SignupFormData>();

	const watchName = watch("name");
	const watchUsername = watch("username");

	useEffect(() => {
		setIsDisabled(!watchName || !watchUsername);
	}, [watchName, watchUsername]);

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
					disabled={isDisabled}
				/>

				<ErrorMessage message={errorText} />
			</View>
		</ThemedView>
	);
}
