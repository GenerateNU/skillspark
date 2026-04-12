import { AuthFormInput } from "@/components/AuthFormInput";
import { Button } from "@/components/Button";
import { ErrorMessage } from "@/components/ErrorMessage";
import { ImageSelector } from "@/components/ImageSelector";
import { PageRedirectButton } from "@/components/PageRedirectButton";
import { ThemedText } from "@/components/themed-text";
import { ThemedView } from "@/components/themed-view";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { AppColors, FontSizes } from "@/constants/theme";
import { useRouter } from "expo-router";
import { SetStateAction, useState } from "react";
import { useForm } from "react-hook-form";
import { useTranslation } from "react-i18next";
import { Image, TouchableOpacity, View } from "react-native";

// 3. add your profile photo or skip for now
export default function PhotoScreen() {
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
					{translate("onboarding.addPhoto")}
				</ThemedText>
			</View>

			<ImageSelector
				setImage={function (value: SetStateAction<string | undefined>): void {
					throw new Error("Function not implemented.");
				}}
				image={undefined}
				width={10}
				height={10}
			/>

			<View className="px-6 items-center">
				<Button
					label={translate("onboarding.choosePhoto")}
					onPress={() => console.log("choose photo")}
					disabled={false}
				/>
				<ThemedText>{translate("onboarding.personalize")}</ThemedText>
			</View>

			<View className="px-6 items-center">
				<Button
					label={translate("onboarding.skip")}
					onPress={() => router.push("/(auth)/signup/child-profile")}
					disabled={false}
				/>
				<ErrorMessage message={errorText} />
			</View>
		</ThemedView>
	);
}
