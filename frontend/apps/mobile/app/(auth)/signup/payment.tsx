import React, { useState } from "react";
import {
	View,
	TextInput,
	TouchableOpacity,
	KeyboardAvoidingView,
	Platform,
	Keyboard,
	Pressable,
	ScrollView,
} from "react-native";
import { AuthBackground } from "@/components/AuthBackground";
import { useRouter } from "expo-router";
import { useSafeAreaInsets } from "react-native-safe-area-context";
import { ThemedText } from "@/components/themed-text";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { useTranslation } from "react-i18next";
import { AppColors, Colors, FontSizes } from "@/constants/theme";
import { Button } from "@/components/Button";

export default function PaymentScreen() {
	const router = useRouter();
	const insets = useSafeAreaInsets();
	const theme = Colors.light;
	const { t: translate } = useTranslation();

	const [cardNumber, setCardNumber] = useState("");
	const [cvv, setCvv] = useState("");
	const [zipCode, setZipCode] = useState("");

	const handleFinish = () => {
		router.push("/(auth)/signup/all-set");
	};

	return (
		<View className="flex-1" style={{ paddingTop: insets.top }}>
			<AuthBackground />
			<KeyboardAvoidingView
				behavior={Platform.OS === "ios" ? "padding" : "height"}
				className="flex-1"
			>
				<ScrollView
					contentContainerStyle={{ flexGrow: 1 }}
					keyboardShouldPersistTaps="handled"
					showsVerticalScrollIndicator={false}
				>
					<Pressable onPress={Keyboard.dismiss}>
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
						<View className="px-6 pt-2 pb-10">
							<ThemedText className="font-nunito-bold text-[#111]" style={{ fontSize: FontSizes.hero, lineHeight: FontSizes.hero + 8, letterSpacing: -0.5 }}>
								{translate("payment.title", { defaultValue: "Payment Information" })}
							</ThemedText>
						</View>

						{/* Form fields */}
						<View className="px-6 gap-8">
							<View className="gap-1">
								<ThemedText className="font-nunito-semibold text-base text-[#111]">
									{translate("payment.cardNumber", { defaultValue: "Card Number" })}
								</ThemedText>
								<TextInput
									className="border border-[#E5E7EB] rounded-[10px] px-4 py-[14px] bg-white text-base font-nunito text-[#11181C]"
									value={cardNumber}
									onChangeText={setCardNumber}
									keyboardType="number-pad"
									placeholderTextColor={AppColors.placeholderText}
								/>
							</View>

							<View className="flex-row gap-3">
								<View className="flex-1 gap-1">
									<ThemedText className="font-nunito-semibold text-base text-[#111]">
										{translate("payment.cvv", { defaultValue: "CVV" })}
									</ThemedText>
									<TextInput
										className="border border-[#E5E7EB] rounded-[10px] px-4 py-[14px] bg-white text-base font-nunito text-[#11181C]"
										value={cvv}
										onChangeText={setCvv}
										keyboardType="number-pad"
										placeholderTextColor={AppColors.placeholderText}
									/>
								</View>

								<View className="flex-1 gap-1">
									<ThemedText className="font-nunito-semibold text-base text-[#111]">
										{translate("payment.zipCode", { defaultValue: "Zip Code" })}
									</ThemedText>
									<TextInput
										className="border border-[#E5E7EB] rounded-[10px] px-4 py-[14px] bg-white text-base font-nunito text-[#11181C]"
										value={zipCode}
										onChangeText={setZipCode}
										keyboardType="number-pad"
										placeholderTextColor={AppColors.placeholderText}
									/>
								</View>
							</View>
						</View>
					</Pressable>
				</ScrollView>
			</KeyboardAvoidingView>

			{/* Finish button pinned to bottom */}
			<View
				className="items-center px-6 pt-4"
				style={{ paddingBottom: insets.bottom + 16 }}
			>
				<Button
					label={translate("payment.finish", { defaultValue: "Finish" })}
					onPress={handleFinish}
					disabled={false}
				/>
			</View>
		</View>
	);
}
