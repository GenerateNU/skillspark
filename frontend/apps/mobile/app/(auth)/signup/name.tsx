import { AuthFormInput } from "@/components/AuthFormInput";
import { Button } from "@/components/Button";
import { ThemedText } from "@/components/themed-text";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { useRouter } from "expo-router";
import { useFormContext } from "react-hook-form";
import { useTranslation } from "react-i18next";
import { SignupFormData } from "@/constants/signup-types";
import { useAuthContext } from "@/hooks/use-auth-context";
import {
	Alert,
	KeyboardAvoidingView,
	Platform,
	ScrollView,
	TouchableOpacity,
	View,
} from "react-native";
import { JumpingCharacter } from "@/components/JumpingCharacter";
import { useSafeAreaInsets } from "react-native-safe-area-context";
import { PageRedirectButton } from "@/components/PageRedirectButton";
import { AuthBackground } from "@/components/AuthBackground";

// 1. name and username
export default function NameScreen() {
	const router = useRouter();
	const { t: translate } = useTranslation();
	const insets = useSafeAreaInsets();
	const { control, getValues } = useFormContext<SignupFormData>();
	const { usernameExists } = useAuthContext();

	const handleContinue = async () => {
		if (!getValues("name") || !getValues("username")) {
			Alert.alert(
				translate("common.error"),
				translate("childProfile.requiredFieldsError"),
			);
			return;
		}

		const isAvailable = await usernameExists(getValues("username"), (msg) => {
			Alert.alert(translate("common.error"), msg);
		});
		if (!isAvailable) return;

		router.push("/(auth)/signup/account");
	};

	return (
		<View className="absolute inset-0">
			<AuthBackground />
			<View className="flex-1" style={{ paddingTop: insets.top + 4 }}>
				<KeyboardAvoidingView
					behavior={Platform.OS === "ios" ? "padding" : "height"}
					className="flex-1"
				>
					<ScrollView
						contentContainerStyle={{ flexGrow: 1 }}
						keyboardShouldPersistTaps="handled"
						showsVerticalScrollIndicator={false}
					>
						{/* Back button */}
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

						{/* Title */}
						<View className="px-6 pt-2 items-center">
							<ThemedText
								className="font-nunito-bold leading-[60px] text-[#111] text-[30px] text-center"
								style={{ letterSpacing: -0.5 }}
							>
								{translate("onboarding.whatName")}
							</ThemedText>
						</View>

						{/* Character image */}
						<View className="flex-1 items-center justify-center py-6">
							<JumpingCharacter />
						</View>

						{/* Form fields */}
						<View className="px-6 gap-6">
							<View className="gap-2">
								<ThemedText className="text-base font-nunito-semibold">
									{translate("onboarding.name")}
								</ThemedText>
								<AuthFormInput
									control={control}
									name="name"
									autoCapitalize="none"
								/>
							</View>

							<View className="gap-2">
								<ThemedText className="text-base font-nunito-semibold">
									{translate("onboarding.username")}
								</ThemedText>
								<AuthFormInput
									control={control}
									name="username"
									autoCapitalize="none"
								/>
							</View>
						</View>
					</ScrollView>
				</KeyboardAvoidingView>

				{/* Buttons pinned to bottom */}
				<View
					className="items-center px-6 pt-4"
					style={{ paddingBottom: insets.bottom + 4 }}
				>
					<Button
						label={translate("onboarding.continue")}
						onPress={handleContinue}
						disabled={false}
					/>
					<View className="mt-3">
						<PageRedirectButton
							label={translate("onboarding.alreadyHaveAccount")}
							onPress={() => router.navigate("/(auth)/login")}
						/>
					</View>
				</View>
			</View>
		</View>
	);
}
