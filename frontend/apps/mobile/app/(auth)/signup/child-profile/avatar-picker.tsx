import React, { useState } from "react";
import { View, TouchableOpacity } from "react-native";
import { Stack, useRouter, useLocalSearchParams } from "expo-router";
import { useSafeAreaInsets } from "react-native-safe-area-context";
import { ThemedText } from "@/components/themed-text";
import { AuthBackground } from "@/components/AuthBackground";
import { Colors, FontSizes } from "@/constants/theme";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { AvatarPicker, DEFAULT_AVATAR_COLOR } from "@/components/AvatarPicker";
import { ChildAvatar, getInitials } from "@/components/ChildAvatar";
import { resolvePendingAvatarCallback } from "@/constants/avatarPickerStore";
import { useTranslation } from "react-i18next";
import { Button } from "@/components/Button";

export default function AvatarPickerScreen() {
	const router = useRouter();
	const params = useLocalSearchParams();
	const insets = useSafeAreaInsets();
	const theme = Colors.light;
	const { t: translate } = useTranslation();

	const childName = (params.childName as string) || "?";
	const initials = getInitials(childName);

	const [face, setFace] = useState<string | null>(
		params.avatarFace ? (params.avatarFace as string) : null,
	);
	const [background, setBackground] = useState<string>(
		(params.avatarBackground as string) || DEFAULT_AVATAR_COLOR,
	);

	const handleSave = () => {
		resolvePendingAvatarCallback({ face, background });
		router.back();
	};

	const handleCancel = () => {
		router.back();
	};

	return (
		<View className="flex-1" style={{ paddingTop: insets.top }}>
			<AuthBackground />
			<Stack.Screen options={{ headerShown: false }} />

			{/* Back button */}
			<TouchableOpacity
				onPress={handleCancel}
				className="flex-row items-center px-5 py-3 gap-1"
				hitSlop={{ top: 10, bottom: 10, left: 10, right: 10 }}
			>
				<IconSymbol name="chevron.left" size={18} color={theme.text} />
				<ThemedText className="text-base font-nunito">
					{translate("onboarding.back")}
				</ThemedText>
			</TouchableOpacity>

			{/* Title */}
			<View className="px-6 pt-2 pb-5">
				<ThemedText className="font-nunito-bold text-[#111]" style={{ fontSize: FontSizes.hero, lineHeight: FontSizes.hero + 8, letterSpacing: -0.5 }}>
					{translate("childProfile.setProfilePicture", {
						defaultValue: "Set your child's profile picture",
					})}
				</ThemedText>
			</View>

			{/* Avatar preview */}
			<View className="items-center mb-5">
				<ChildAvatar
					name={childName}
					avatarFace={face}
					avatarBackground={background || DEFAULT_AVATAR_COLOR}
					size={90}
				/>
			</View>

			{/* Picker card */}
			<View className="flex-1 px-6 justify-center">
				<AvatarPicker
					selectedFace={face}
					selectedBackground={background}
					onFaceChange={setFace}
					onBackgroundChange={setBackground}
					childInitials={initials}
				/>
			</View>

			{/* Cancel / Save */}
			<View className="flex-row gap-3 px-6" style={{ paddingBottom: insets.bottom + 16 }}>
				<Button
					label={translate("common.cancel")}
					onPress={handleCancel}
					disabled={false}
					bgColor="#FFFFFF"
					textColor="#1B1B1B"
					width="48%"
				/>
				<Button
					label={translate("common.save", { defaultValue: "Save" })}
					onPress={handleSave}
					disabled={false}
					bgColor="#1B1B1B"
					textColor="#FFFFFF"
					width="48%"
				/>
			</View>
		</View>
	);
}
