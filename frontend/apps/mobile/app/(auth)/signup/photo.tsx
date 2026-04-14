import { Button } from "@/components/Button";
import { ErrorMessage } from "@/components/ErrorMessage";
import { ImageSelector } from "@/components/ImageSelector";
import { ThemedText } from "@/components/themed-text";
import { ThemedView } from "@/components/themed-view";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { AppColors, FontSizes } from "@/constants/theme";
import { useAuthContext } from "@/hooks/use-auth-context";
import { useRouter } from "expo-router";
import { useState } from "react";
import { useTranslation } from "react-i18next";
import { TouchableOpacity, View } from "react-native";
import { useSafeAreaInsets } from "react-native-safe-area-context";
import { useGuardian } from "@/hooks/use-guardian";

// 3. add your profile photo or skip for now
export default function PhotoScreen() {
	const router = useRouter();
	const { t: translate } = useTranslation();
	const insets = useSafeAreaInsets();
	const [errorText, setErrorText] = useState("");
	const [image, setImage] = useState<string | undefined>(undefined);
	const { update, guardianId } = useAuthContext();
	const { guardian } = useGuardian(guardianId);

	const onSubmit = () => {
		if (!guardian) {
			setErrorText("ERROR: Could not fetch guardian ID");
		} else {
			update(
				() => router.push("/(auth)/signup/child-profile"),
				setErrorText,
				guardian.id,
				guardian.email,
				guardian.language_preference,
				guardian.name,
				guardian.username,
				image,
			);
		}
	};

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
					{translate("onboarding.addPhoto")}
				</ThemedText>
			</View>

			<View className="items-center shadow-sm">
				<ImageSelector
					setImage={setImage}
					image={image}
					width={150}
					height={150}
				/>
			</View>

			<View className="px-6 items-center">
				<Button
					label={translate("onboarding.choosePhoto")}
					onPress={() => console.log("choose photo")}
					disabled={false}
					bgColor={"#FFFFFF"}
					width={"40%"}
					textColor={"#1B1B1B"}
				/>
				<ThemedText className="font-nunito">
					{translate("onboarding.personalize")}
				</ThemedText>
			</View>

			<View className="px-6 items-center">
				<Button
					label={translate("onboarding.skip")}
					onPress={onSubmit}
					disabled={false}
				/>
				<ErrorMessage message={errorText} />
			</View>
		</ThemedView>
	);
}
