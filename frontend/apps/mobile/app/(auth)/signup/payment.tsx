import React from "react";
import { View, TouchableOpacity } from "react-native";
import { AuthBackground } from "@/components/AuthBackground";
import { useRouter } from "expo-router";
import { useSafeAreaInsets } from "react-native-safe-area-context";
import { ThemedText } from "@/components/themed-text";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { AppColors } from "@/constants/theme";
import { useTranslation } from "react-i18next";

export default function PaymentScreen() {
	const router = useRouter();
	const insets = useSafeAreaInsets();
	const { t: translate } = useTranslation();

	const handleUpdateBilling = () => {
		router.push("/(auth)/signup/all-set");
	};
	const handleDelete = () => {};

	return (
		<View style={{ flex: 1, paddingTop: insets.top }}>
			<AuthBackground />
			{/* Header row in lavender area */}
			<View
				style={{
					flexDirection: "row",
					alignItems: "center",
					justifyContent: "space-between",
					paddingHorizontal: 20,
					paddingVertical: 14,
				}}
			>
				<TouchableOpacity
					onPress={() => router.back()}
					style={{ flexDirection: "row", alignItems: "center", gap: 4 }}
					hitSlop={{ top: 10, bottom: 10, left: 10, right: 10 }}
				>
					<IconSymbol name="chevron.left" size={18} color="#11181C" />
					<ThemedText style={{ fontSize: 16, fontFamily: "NunitoSans_400Regular" }}>
						{translate("onboarding.back")}
					</ThemedText>
				</TouchableOpacity>

				<ThemedText style={{ fontSize: 18, fontFamily: "NunitoSans_700Bold" }}>
					{translate("payment.title")}
				</ThemedText>

				<View style={{ width: 40 }} />
			</View>

			{/* White card with billing content */}
			<View
				style={{
					flex: 1,
					backgroundColor: "#FFFFFF",
					borderTopLeftRadius: 28,
					borderTopRightRadius: 28,
					paddingHorizontal: 24,
					paddingTop: 28,
					paddingBottom: insets.bottom + 24,
				}}
			>
				<ThemedText
					style={{
						fontSize: 20,
						fontFamily: "NunitoSans_700Bold",
						color: AppColors.primaryText,
						marginBottom: 24,
					}}
				>
					{translate("payment.manageBilling")}
				</ThemedText>

				{/* Card info block */}
				<View
					style={{
						backgroundColor: "#F9F9F9",
						borderRadius: 16,
						padding: 20,
						marginBottom: 28,
						borderWidth: 1,
						borderColor: "#EBEBEB",
					}}
				>
					{/* TODO: Replace with real payment method data from billing API */}
					<ThemedText
						style={{
							fontSize: 13,
							fontFamily: "NunitoSans_400Regular",
							color: AppColors.mutedText,
							marginBottom: 4,
						}}
					>
						{translate("payment.creditCard")}
					</ThemedText>
					<ThemedText
						style={{
							fontSize: 15,
							fontFamily: "NunitoSans_600SemiBold",
							color: AppColors.primaryText,
							marginBottom: 2,
						}}
					>
						{translate("payment.name")}
					</ThemedText>
					<ThemedText
						style={{
							fontSize: 15,
							fontFamily: "NunitoSans_400Regular",
							color: AppColors.primaryText,
							letterSpacing: 2,
						}}
					>
						**** **** **** XXXX
					</ThemedText>
				</View>

				{/* Action buttons */}
				<View style={{ flexDirection: "row", gap: 12 }}>
					<TouchableOpacity
						style={{
							flex: 1,
							backgroundColor: "#1C1C1E",
							borderRadius: 24,
							paddingVertical: 16,
							alignItems: "center",
						}}
						onPress={handleUpdateBilling}
						activeOpacity={0.8}
					>
						<ThemedText style={{ color: "#FFFFFF", fontSize: 15, fontFamily: "NunitoSans_600SemiBold" }}>
							{translate("payment.updateBilling")}
						</ThemedText>
					</TouchableOpacity>

					<TouchableOpacity
						style={{
							flex: 1,
							borderRadius: 24,
							paddingVertical: 16,
							alignItems: "center",
							borderWidth: 1.5,
							borderColor: "#1C1C1E",
						}}
						onPress={handleDelete}
						activeOpacity={0.8}
					>
						<ThemedText style={{ color: "#1C1C1E", fontSize: 15, fontFamily: "NunitoSans_400Regular" }}>
							{translate("payment.delete")}
						</ThemedText>
					</TouchableOpacity>
				</View>
			</View>
		</View>
	);
}
