import { AuthFormInput } from "@/components/AuthFormInput";
import { FontSizes } from "@/constants/theme";
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
import { useSafeAreaInsets } from "react-native-safe-area-context";
import { AuthBackground } from "@/components/AuthBackground";

export default function NameScreen() {
	const router = useRouter();
	const { t: translate } = useTranslation();
	const insets = useSafeAreaInsets();
	const { control, getValues } = useFormContext<SignupFormData>();
	const { usernameExists, signup } = useAuthContext();

	const handleContinue = async () => {
		const firstName = getValues("first_name");
		const username = getValues("username");

		if (!firstName || !username) {
			Alert.alert(
				translate("common.error"),
				translate("childProfile.requiredFieldsError"),
			);
			return;
		}

		const isAvailable = await usernameExists(username, (msg) => {
			Alert.alert(translate("common.error"), msg);
		});
		if (!isAvailable) return;

		const lastName = getValues("last_name");
		const name = [firstName, lastName].filter(Boolean).join(" ");

		signup(
			name,
			getValues("email"),
			username,
			getValues("password"),
			getValues("language_preference"),
			undefined,
			(msg) => Alert.alert(translate("common.error"), msg),
			() => router.push("/(auth)/signup/photo"),
		);
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
						<View className="px-6 pt-10 pb-5 items-center">
							<ThemedText
								className="font-nunito-bold text-[#111] text-center"
								style={{ fontSize: FontSizes.hero, lineHeight: FontSizes.hero + 8, letterSpacing: -0.5 }}
							>
								{translate("onboarding.whatName")}
							</ThemedText>
						</View>

						{/* Form fields */}
						<View className="flex-1 justify-center px-6">
							<View className="gap-6">
								<View className="gap-2">
									<ThemedText className="text-base font-nunito-semibold">
										{translate("childProfile.firstName")}
									</ThemedText>
									<AuthFormInput
										control={control}
										name="first_name"
										autoCapitalize="words"
									/>
								</View>

								<View className="gap-2">
									<ThemedText className="text-base font-nunito-semibold">
										{translate("childProfile.lastName")}
									</ThemedText>
									<AuthFormInput
										control={control}
										name="last_name"
										autoCapitalize="words"
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
						</View>
					</ScrollView>
				</KeyboardAvoidingView>

				{/* Button pinned to bottom */}
				<View
					className="items-center px-6 pt-4"
					style={{ paddingBottom: insets.bottom + 16 }}
				>
					<Button
						label={translate("onboarding.continue")}
						onPress={handleContinue}
						disabled={false}
					/>
				</View>
			</View>
		</View>
	);
}
