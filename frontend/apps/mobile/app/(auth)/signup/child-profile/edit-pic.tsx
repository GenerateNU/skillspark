import React, { useState } from "react";
import { View, TouchableOpacity } from "react-native";
import { Stack, useRouter, useLocalSearchParams } from "expo-router";
import { useSafeAreaInsets } from "react-native-safe-area-context";
import { ThemedText } from "@/components/themed-text";
import { AuthBackground } from "@/components/AuthBackground";
import { Colors, FontSizes } from "@/constants/theme";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { DEFAULT_AVATAR_COLOR } from "@/components/AvatarPicker";
import { setPendingAvatarCallback } from "@/constants/avatarPickerStore";
import { useTranslation } from "react-i18next";
import { useAuthContext } from "@/hooks/use-auth-context";
import { ErrorScreen } from "@/components/ErrorScreen";
import { ChildAvatar } from "@/components/ChildAvatar";
import { Button } from "@/components/Button";

export default function EditChildPictureScreen() {
	const router = useRouter();
	const params = useLocalSearchParams();
	const insets = useSafeAreaInsets();
	const theme = Colors.light;
	const { guardianId } = useAuthContext();
	const { t: translate } = useTranslation();

	const [firstName] = useState(
		params.name ? (params.name as string).split(" ")[0] : "",
	);
	const [lastName] = useState(
		params.name ? (params.name as string).split(" ").slice(1).join(" ") : "",
	);
	const [avatarFace, setAvatarFace] = useState<string | null>(
		(params.avatar_face as string) || null,
	);
	const [avatarBackground, setAvatarBackground] = useState(
		(params.avatar_background as string) || DEFAULT_AVATAR_COLOR,
	);

	if (!guardianId) {
		return <ErrorScreen message="Illegal state: no guardian ID retrieved" />;
	}

	const handleAvatarPress = () => {
		setPendingAvatarCallback(({ face, background }) => {
			setAvatarFace(face);
			setAvatarBackground(background);
		});
		const childName = [firstName, lastName].filter(Boolean).join(" ") || "?";
		router.push({
			pathname: "./avatar-picker",
			params: { avatarFace: avatarFace ?? "", avatarBackground, childName },
		});
	};

	const childName = [firstName, lastName].filter(Boolean).join(" ") || "?";

	return (
		<View className="flex-1" style={{ paddingTop: insets.top }}>
			<AuthBackground />
			<Stack.Screen options={{ headerShown: false }} />

			{/* Back button */}
			<TouchableOpacity
				onPress={() => router.back()}
				className="flex-row items-center px-5 py-3 gap-1"
				hitSlop={{ top: 10, bottom: 10, left: 10, right: 10 }}
			>
				<IconSymbol name="chevron.left" size={18} color={theme.text} />
				<ThemedText className="text-base font-nunito">
					{translate("onboarding.back")}
				</ThemedText>
			</TouchableOpacity>

			{/* Title */}
			<View className="px-6 pt-2">
				<ThemedText className="font-nunito-bold text-[#111]" style={{ fontSize: FontSizes.hero, lineHeight: FontSizes.hero + 8, letterSpacing: -0.5 }}>
					{translate("childProfile.setProfilePicture", {
						defaultValue: "Set your child's profile picture",
					})}
				</ThemedText>
			</View>

			{/* Avatar — centred in remaining space */}
			<View className="flex-1 items-center justify-center">
				<TouchableOpacity onPress={handleAvatarPress} activeOpacity={0.8}>
					<View className="relative">
						<ChildAvatar
							name={childName}
							avatarFace={avatarFace}
							avatarBackground={avatarBackground || DEFAULT_AVATAR_COLOR}
							size={160}
						/>
						<View className="absolute -top-2 -right-2 w-11 h-11 rounded-full bg-white shadow items-center justify-center">
							<IconSymbol name="pencil" size={18} color={theme.text} />
						</View>
					</View>
				</TouchableOpacity>
			</View>

			{/* Save button */}
			<View className="px-6 items-center" style={{ paddingBottom: insets.bottom + 16 }}>
				<Button
					label={translate("onboarding.save")}
					onPress={() => router.push("/(auth)/signup/child-profile")}
					disabled={false}
				/>
			</View>
		</View>
	);
}
