import { Button } from "@/components/Button";
import { AuthBackground } from "@/components/AuthBackground";
import { ThemedText } from "@/components/themed-text";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { AppColors, FontSizes } from "@/constants/theme";
import { useRouter } from "expo-router";
import { useTranslation } from "react-i18next";
import { ScrollView, Text, TouchableOpacity, View } from "react-native";
import { useSafeAreaInsets } from "react-native-safe-area-context";
import { useAuthContext } from "@/hooks/use-auth-context";
import { useGuardian } from "@/hooks/use-guardian";
import { ChildAvatar } from "@/components/ChildAvatar";
import { MONTHS } from "@/components/ChildProfileForm";

// 4. set up your child's profile
export default function ChildProfileScreen() {
	const router = useRouter();
	const { t: translate } = useTranslation();
	const insets = useSafeAreaInsets();
	const { guardianId } = useAuthContext();
	const { children } = useGuardian(guardianId);

	return (
		<View style={{ flex: 1, paddingTop: insets.top }}>
			<AuthBackground />
			{/* Back button */}
			<TouchableOpacity
				onPress={() => router.back()}
				style={{
					flexDirection: "row",
					alignItems: "center",
					paddingHorizontal: 20,
					paddingVertical: 12,
					gap: 4,
				}}
				hitSlop={{ top: 10, bottom: 10, left: 10, right: 10 }}
			>
				<IconSymbol name="chevron.left" size={18} color="#11181C" />
				<ThemedText
					style={{ fontSize: 16, fontFamily: "NunitoSans_400Regular" }}
				>
					{translate("onboarding.back")}
				</ThemedText>
			</TouchableOpacity>

			{/* Title */}
			<View
				style={{
					paddingHorizontal: 24,
					paddingTop: 8,
					paddingBottom: 24,
					alignItems: "center",
				}}
			>
				<ThemedText
					style={{
						fontFamily: "NunitoSans_700Bold",
						fontSize: FontSizes.hero,
						lineHeight: 40,
						color: AppColors.primaryText,
						letterSpacing: -0.5,
						textAlign: "center",
					}}
				>
					{translate("onboarding.setUpChild")}
				</ThemedText>
			</View>

			{/* Scrollable content */}
			<ScrollView
				style={{ flex: 1 }}
				contentContainerStyle={{
					paddingHorizontal: 24,
					paddingBottom: 16,
					gap: 12,
				}}
				showsVerticalScrollIndicator={false}
			>
				{/* Added children */}
				{children.map((child: any) => (
					<TouchableOpacity
						className="shadow-sm"
						key={child.id}
						onPress={() =>
							router.push({
								pathname: "/(auth)/signup/child-profile/add-child",
								params: {
									id: child.id,
									name: child.name,
									birth_month: child.birth_month,
									birth_year: child.birth_year,
									school_id: child.school_id,
									interests: child.interests?.join(","),
									avatar_face: child.avatar_face ?? "",
									avatar_background: child.avatar_background ?? "",
								},
							})
						}
						activeOpacity={0.75}
						style={{
							flexDirection: "row",
							alignItems: "center",
							backgroundColor: "#FFFFFF",
							borderRadius: 16,
							paddingVertical: 14,
							paddingHorizontal: 16,
							gap: 12,
						}}
					>
						<View className="mr-2">
							<ChildAvatar
								name={child.name ?? ""}
								avatarFace={child.avatar_face}
								avatarBackground={child.avatar_background}
								size={32}
							/>
						</View>
						<View className="flex-1">
							<ThemedText
								className="text-sm font-nunito-medium"
								numberOfLines={1}
							>
								{child.name}
							</ThemedText>
							<ThemedText
								className="text-[10px] font-nunito"
								style={{ color: AppColors.mutedText }}
							>
								{[
									child.birth_month ? MONTHS[child.birth_month - 1] : null,
									child.birth_year,
								]
									.filter(Boolean)
									.join(", ")}
							</ThemedText>
						</View>
						<IconSymbol
							name="chevron.right"
							size={18}
							color={AppColors.mutedText}
						/>
					</TouchableOpacity>
				))}

				{/* Add a new child profile button */}
				<TouchableOpacity
					className="shadow-sm"
					onPress={() => router.push("/(auth)/signup/child-profile/add-child")}
					activeOpacity={0.7}
					style={{
						backgroundColor: "#FFFFFF",
						borderRadius: 16,
						paddingVertical: 20,
						alignItems: "center",
					}}
				>
					<Text
						style={{
							fontFamily: "NunitoSans_600SemiBold",
							fontSize: 15,
							color: AppColors.primaryText,
						}}
					>
						{translate("onboarding.addNewChild")}
					</Text>
					<Text
						style={{
							fontFamily: "NunitoSans_400Regular",
							fontSize: 22,
							color: AppColors.primaryText,
							lineHeight: 28,
						}}
					>
						+
					</Text>
				</TouchableOpacity>
			</ScrollView>

			<View
				style={{
					alignItems: "center",
					paddingHorizontal: 24,
					paddingTop: 16,
					paddingBottom: insets.bottom + 16,
				}}
			>
				<Button
					label={translate("onboarding.continue")}
					onPress={() => router.push("/(auth)/signup/emergency-contact")}
					disabled={false}
				/>
			</View>
		</View>
	);
}
