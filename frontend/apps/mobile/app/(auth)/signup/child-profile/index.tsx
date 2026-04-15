import { Button } from "@/components/Button";
import { ErrorMessage } from "@/components/ErrorMessage";
import { ThemedText } from "@/components/themed-text";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { AppColors, FontSizes } from "@/constants/theme";
import { useRouter } from "expo-router";
import { useState } from "react";
import { useTranslation } from "react-i18next";
import { Text, TouchableOpacity, View } from "react-native";
import { JumpingCharacter } from "@/components/JumpingCharacter";
import { useSafeAreaInsets } from "react-native-safe-area-context";

const BG = "#EDE8FF";

// 4. set up your child's profile
export default function ChildProfileScreen() {
	const router = useRouter();
	const { t: translate } = useTranslation();
	const insets = useSafeAreaInsets();
	const [errorText, setErrorText] = useState("");

	return (
		<View style={{ flex: 1, backgroundColor: BG, paddingTop: insets.top }}>
			{/* Back button in lavender area */}
			<TouchableOpacity
				onPress={() => router.back()}
				style={{ flexDirection: "row", alignItems: "center", paddingHorizontal: 20, paddingVertical: 12, gap: 4 }}
				hitSlop={{ top: 10, bottom: 10, left: 10, right: 10 }}
			>
				<IconSymbol name="chevron.left" size={18} color="#11181C" />
				<ThemedText style={{ fontSize: 16, fontFamily: "NunitoSans_400Regular" }}>
					{translate("onboarding.back")}
				</ThemedText>
			</TouchableOpacity>

			{/* Top lavender section: title + character */}
			<View style={{ flex: 1, alignItems: "center", justifyContent: "center", paddingHorizontal: 24 }}>
				<ThemedText
					style={{
						fontFamily: "NunitoSans_700Bold",
						fontSize: FontSizes.hero,
						color: AppColors.primaryText,
						letterSpacing: -0.5,
						textAlign: "center",
						marginBottom: 24,
					}}
				>
					{translate("onboarding.setUpChild")}
				</ThemedText>
				<JumpingCharacter />
			</View>

			{/* Bottom white card: buttons */}
			<View
				style={{
					backgroundColor: "#FFFFFF",
					borderTopLeftRadius: 28,
					borderTopRightRadius: 28,
					paddingHorizontal: 24,
					paddingTop: 28,
					paddingBottom: insets.bottom + 20,
					gap: 12,
				}}
			>
				{/* Outlined "Add New Child" button */}
				<TouchableOpacity
					onPress={() => router.push("/(auth)/signup/child-profile/add-child")}
					activeOpacity={0.7}
					style={{
						borderWidth: 1.5,
						borderColor: "#1B1B1B",
						borderRadius: 24,
						paddingVertical: 16,
						alignItems: "center",
						width: "95%",
						alignSelf: "center",
					}}
				>
					<Text style={{ fontFamily: "NunitoSans_700Bold", fontSize: 18, color: "#1B1B1B" }}>
						{translate("onboarding.addNewChild")}
					</Text>
				</TouchableOpacity>
				<Button
					label={translate("onboarding.continue")}
					onPress={() => router.push("/(auth)/signup/emergency-contact")}
					disabled={false}
				/>
				<ErrorMessage message={errorText} />
			</View>
		</View>
	);
}
