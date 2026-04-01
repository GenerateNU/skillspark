import { useAuthContext } from "@/hooks/use-auth-context";
import { useRouter } from "expo-router";
import {
	ScrollView,
	TouchableOpacity,
	useColorScheme,
	View,
} from "react-native";
import { useSafeAreaInsets } from "react-native-safe-area-context";
import { AppColors, Colors } from "@/constants/theme";
import { ThemedView } from "@/components/themed-view";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { ThemedText } from "@/components/themed-text";
import { useState } from "react";
import { useForm } from "react-hook-form";
import { ErrorScreen } from "@/components/ErrorScreen";
import { AuthFormInput } from "@/components/AuthFormInput";
import { Button } from "@/components/Button";
import { ErrorMessage } from "@/components/ErrorMessage";
import { ImageSelector } from "@/components/ImageSelector";
import { useGuardian } from "@/hooks/use-guardian";

type EditFormData = {
	name: string;
	username: string;
};

export default function EditProfileScreen() {
	const router = useRouter();
	const insets = useSafeAreaInsets();
	const colorScheme = useColorScheme();
	const theme = Colors[colorScheme ?? "light"];

	const [errorText, setErrorText] = useState("");
	const { update, guardianId, langPref } = useAuthContext();
	const { guardian } = useGuardian(guardianId);
	const [image, setImage] = useState<string | undefined>(undefined);
	const [pfp, setPfp] = useState<string | undefined>(
		guardian?.profile_picture_s3_key,
	);

	const { control, handleSubmit } = useForm<EditFormData>({
		defaultValues: {
			name: "",
			username: "",
		},
	});

	if (!guardianId) {
		return <ErrorScreen message="Illegal state: no guardian ID found" />;
	}

	const onSubmit = (formData: EditFormData) => {
		if (!guardian) {
			setErrorText("No guardian ID found");
			return;
		}

		const id = guardian.id;
		const email = guardian.email;
		const language_preference = langPref!;
		const name = formData.name === "" ? guardian.name : formData.name;
		const username =
			formData.username === "" ? guardian.username : formData.username;
		update(
			() => router.back(),
			setErrorText,
			id,
			email,
			language_preference,
			name,
			username,
			image ?? pfp,
		);
	};

	const clearImage = () => {
		setImage(undefined);
		setPfp(undefined);
	};

	return (
		<ThemedView className="flex-1" style={{ paddingTop: insets.top }}>
			<View className="flex-row items-center justify-between px-5 py-3">
				<TouchableOpacity
					onPress={() => router.back()}
					className="w-10 justify-center items-start"
					hitSlop={{ top: 10, bottom: 10, left: 10, right: 10 }}
				>
					<IconSymbol name="chevron.left" size={24} color={theme.text} />
				</TouchableOpacity>
				<ThemedText className="text-xl text-center font-nunito-bold">
					Family Information
				</ThemedText>
				<View className="w-10" />
			</View>

			<ScrollView
				className="flex-1 px-5"
				contentContainerStyle={{ paddingBottom: insets.bottom + 24 }}
				keyboardShouldPersistTaps="handled"
			>
				<ThemedText className="text-lg font-nunito-bold mt-4 mb-6">
					Edit Profile
				</ThemedText>

				<View className="items-center mb-8">
					<View className="relative w-40">
						<ImageSelector
							setImage={setImage}
							image={image ?? pfp}
							width={160}
							height={160}
							className="items-center gap-2"
						/>
						{(image || pfp) && (
							<TouchableOpacity
								onPress={clearImage}
								className="absolute top-1 right-1 w-7 h-7 rounded-full items-center justify-center"
								style={{ backgroundColor: AppColors.danger }}
							>
								<IconSymbol name="xmark" size={14} color="white" />
							</TouchableOpacity>
						)}
					</View>
				</View>

				<View className="gap-5">
					<View className="gap-1">
						<ThemedText className="text-sm font-nunito-bold">Name</ThemedText>
						<AuthFormInput
							control={control}
							name="name"
							placeholder="Name"
							autoCapitalize="none"
						/>
					</View>

					<View className="gap-1">
						<ThemedText className="text-sm font-nunito-bold">
							Username
						</ThemedText>
						<AuthFormInput
							control={control}
							name="username"
							placeholder="Username"
							autoCapitalize="none"
						/>
					</View>
				</View>

				<ErrorMessage message={errorText} />

				<View className="mt-8">
					<Button label="Save Changes" onPress={handleSubmit(onSubmit)} />
				</View>
			</ScrollView>
		</ThemedView>
	);
}
