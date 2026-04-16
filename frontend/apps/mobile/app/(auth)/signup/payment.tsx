import React from "react";
import { View, TouchableOpacity } from "react-native";
import { AuthBackground } from "@/components/AuthBackground";
import { useRouter } from "expo-router";
import { useSafeAreaInsets } from "react-native-safe-area-context";
import { ThemedText } from "@/components/themed-text";
import { IconSymbol } from "@/components/ui/icon-symbol";
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
		<View className="flex-1" style={{ paddingTop: insets.top }}>
			<AuthBackground />
			{/* Header row in lavender area */}
			<View className="flex-row items-center justify-between px-5 py-3.5">
				<TouchableOpacity
					onPress={() => router.back()}
					className="flex-row items-center gap-1"
					hitSlop={{ top: 10, bottom: 10, left: 10, right: 10 }}
				>
					<IconSymbol name="chevron.left" size={18} color="#11181C" />
					<ThemedText className="text-base font-nunito">
						{translate("onboarding.back")}
					</ThemedText>
				</TouchableOpacity>

				<ThemedText className="text-lg font-nunito-bold">
					{translate("payment.title")}
				</ThemedText>

				<View className="w-10" />
			</View>

			{/* White card with billing content */}
			<View
				className="flex-1 bg-white rounded-tl-[28px] rounded-tr-[28px] px-6 pt-7"
				style={{ paddingBottom: insets.bottom + 24 }}
			>
				<ThemedText className="text-xl font-nunito-bold text-[#111] mb-6">
					{translate("payment.manageBilling")}
				</ThemedText>

				{/* Card info block */}
				<View className="bg-[#F9F9F9] rounded-2xl p-5 mb-7 border border-[#EBEBEB]">
					{/* TODO: Replace with real payment method data from billing API */}
					<ThemedText className="text-[13px] font-nunito text-[#6B7280] mb-1">
						{translate("payment.creditCard")}
					</ThemedText>
					<ThemedText className="text-[15px] font-nunito-semibold text-[#111] mb-0.5">
						{translate("payment.name")}
					</ThemedText>
					<ThemedText className="text-[15px] font-nunito text-[#111] tracking-[2px]">
						**** **** **** XXXX
					</ThemedText>
				</View>

				{/* Action buttons */}
				<View className="flex-row gap-3">
					<TouchableOpacity
						className="flex-1 bg-[#1C1C1E] rounded-3xl py-4 items-center"
						onPress={handleUpdateBilling}
						activeOpacity={0.8}
					>
						<ThemedText className="text-white text-[15px] font-nunito-semibold">
							{translate("payment.updateBilling")}
						</ThemedText>
					</TouchableOpacity>

					<TouchableOpacity
						className="flex-1 rounded-3xl py-4 items-center border-[1.5px] border-[#1C1C1E]"
						onPress={handleDelete}
						activeOpacity={0.8}
					>
						<ThemedText className="text-[#1C1C1E] text-[15px] font-nunito">
							{translate("payment.delete")}
						</ThemedText>
					</TouchableOpacity>
				</View>
			</View>
		</View>
	);
}
