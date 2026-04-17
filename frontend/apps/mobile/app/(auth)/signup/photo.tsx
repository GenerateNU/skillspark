import { Button } from "@/components/Button";
import { ErrorMessage } from "@/components/ErrorMessage";
import { ImageSelector } from "@/components/ImageSelector";
import { ThemedText } from "@/components/themed-text";
import { useAuthContext } from "@/hooks/use-auth-context";
import { useRouter } from "expo-router";
import { useState } from "react";
import { useTranslation } from "react-i18next";
import { View } from "react-native";
import { AuthBackground } from "@/components/AuthBackground";
import { useSafeAreaInsets } from "react-native-safe-area-context";
import { useGuardian } from "@/hooks/use-guardian";
import { AppColors, Colors, FontSizes } from "@/constants/theme";
import { IconSymbol } from "@/components/ui/icon-symbol";

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
		<View className="flex-1" style={{ paddingTop: insets.top }}>
			<AuthBackground />
			<View className="px-6 pt-8 items-center">
				<ThemedText
					className="font-nunito-bold text-[#111] text-center"
					style={{ fontSize: FontSizes.hero, lineHeight: FontSizes.hero + 8, letterSpacing: -0.5 }}
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
					placeholder={
						<View
							className="rounded-full border items-center justify-center overflow-hidden"
							style={{
								borderColor: Colors.light.borderColor,
								width: 150,
								height: 150,
							}}
						>
							<IconSymbol
								name="square.and.arrow.up"
								size={100}
								color={AppColors.mutedText}
							/>
						</View>
					}
				/>
			</View>

			<View className="px-6 items-center">
				<ImageSelector
					setImage={setImage}
					image={image}
					width={150}
					height={150}
					placeholder={
						<Button
							label={translate("onboarding.choosePhoto")}
							onPress={() => null}
							disabled={false}
							bgColor={"#FFFFFF"}
							width={"40%"}
							textColor={"#1B1B1B"}
						/>
					}
				/>
				<ThemedText className="font-nunito">
					{translate("onboarding.personalize")}
				</ThemedText>
			</View>

			<View className="px-6 items-center">
				{image ? (
					<View className="flex-row gap-3 w-full">
						<Button
							label={translate("onboarding.cancel")}
							onPress={() => setImage(undefined)}
							disabled={false}
							bgColor="#FFFFFF"
							textColor="#1B1B1B"
							width="48%"
						/>
						<Button
							label={translate("onboarding.save")}
							onPress={onSubmit}
							disabled={false}
							bgColor="#1B1B1B"
							textColor="#FFFFFF"
							width="48%"
						/>
					</View>
				) : (
					<Button
						label={translate("onboarding.skip")}
						onPress={onSubmit}
						disabled={false}
						bgColor="#FFFFFF"
						textColor="#1B1B1B"
					/>
				)}
				<ErrorMessage message={errorText} />
			</View>
		</View>
	);
}
