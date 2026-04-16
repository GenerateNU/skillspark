import { AuthFormInput } from "@/components/AuthFormInput";
import { Button } from "@/components/Button";
import { ThemedText } from "@/components/themed-text";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { useRouter } from "expo-router";
import { useFormContext } from "react-hook-form";
import { useTranslation } from "react-i18next";
import { SignupFormData } from "@/constants/signup-types";
import {
	Alert,
	KeyboardAvoidingView,
	Platform,
	ScrollView,
	TouchableOpacity,
	View,
} from "react-native";
import { useSafeAreaInsets } from "react-native-safe-area-context";
import { useAuthContext } from "@/hooks/use-auth-context";
import { AuthBackground } from "@/components/AuthBackground";

// 2. email and password
export default function AccountScreen() {
	const router = useRouter();
	const { t: translate } = useTranslation();
	const insets = useSafeAreaInsets();
	const { signup } = useAuthContext();
	const { handleSubmit, control, getValues } = useFormContext<SignupFormData>();

	const onSubmit = (formData: SignupFormData) => {
		signup(
			formData.name,
			formData.email,
			formData.username,
			formData.password,
			formData.language_preference,
			undefined,
			(msg) => {
				Alert.alert(translate("common.error"), msg);
			},
			() => router.push("/(auth)/signup/photo"),
		);
	};

	const handleCreateAccount = () => {
		const email = getValues("email");
		const password = getValues("password");
		const confirmPassword = getValues("confirm_password");

		if (!email || !password || !confirmPassword) {
			Alert.alert(
				translate("common.error"),
				translate("childProfile.requiredFieldsError"),
			);
			return;
		}
		if (password !== confirmPassword) {
			Alert.alert(
				translate("common.error"),
				translate("onboarding.passwordMismatch"),
			);
			return;
		}
		handleSubmit(onSubmit)();
	};

	return (
		<View className="absolute inset-0">
			<AuthBackground />
			<View className="flex-1" style={{ paddingTop: insets.top }}>
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
								{translate("onboarding.makeAccount")}
							</ThemedText>
						</View>

						{/* Form fields */}
						<View className="px-6 gap-6 pt-20">
							<View className="gap-2">
								<ThemedText className="text-base font-nunito-semibold">
									{translate("onboarding.email")}
								</ThemedText>
								<AuthFormInput
									control={control}
									name="email"
									keyboardType="email-address"
									autoCapitalize="none"
								/>
							</View>

							<View className="gap-2">
								<ThemedText className="text-base font-nunito-semibold">
									{translate("onboarding.password")}
								</ThemedText>
								<AuthFormInput
									control={control}
									name="password"
									secureTextEntry
								/>
							</View>

							<View className="gap-2">
								<ThemedText className="text-base font-nunito-semibold">
									{translate("onboarding.confirmPassword")}
								</ThemedText>
								<AuthFormInput
									control={control}
									name="confirm_password"
									secureTextEntry
								/>
							</View>
						</View>
					</ScrollView>
				</KeyboardAvoidingView>

				{/* Button pinned to bottom */}
				<View
					className="items-center px-6 pt-4"
					style={{ paddingBottom: insets.bottom + 56 }}
				>
					<Button
						label={translate("onboarding.createAccount")}
						onPress={handleCreateAccount}
						disabled={false}
					/>
				</View>
			</View>
		</View>
	);
}
