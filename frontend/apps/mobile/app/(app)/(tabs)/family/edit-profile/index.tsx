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
import { useEffect, useState } from "react";
import { useForm } from "react-hook-form";
import { ErrorScreen } from "@/components/ErrorScreen";
import { AuthFormInput } from "@/components/AuthFormInput";
import { Button } from "@/components/Button";
import { ErrorMessage } from "@/components/ErrorMessage";
import { ImageSelector } from "@/components/ImageSelector";
import { useGuardian } from "@/hooks/use-guardian";
import { useTranslation } from "react-i18next";

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
	const ogPfp = guardian?.profile_picture_s3_key ?? undefined;
	const [imgCleared, setImgCleared] = useState(false);
	const { t: translate } = useTranslation();

	const { control, handleSubmit, watch } = useForm<EditFormData>({
		defaultValues: {
			name: "",
			username: "",
		},
	});

	const [canUpdate, setCanUpdate] = useState(false);
	const formValues = watch();

	useEffect(() => {
		const hasTextChanges = formValues.name !== "" || formValues.username !== "";
		const currentImage = imgCleared ? undefined : image || ogPfp;
		const hasImgChanges = currentImage !== ogPfp;
		setCanUpdate(hasTextChanges || hasImgChanges);
		console.log({
			image,
			imgCleared,
			ogPfp,
			currentImage,
			hasImgChanges,
			hasTextChanges,
		});
	}, [formValues.name, formValues.username, image, ogPfp, imgCleared]);

	if (!guardianId || !guardian) {
		// change to reroute to login
		return <ErrorScreen message="Illegal state: no guardian ID found" />;
	}

	const onSubmit = (formData: EditFormData) => {
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
			imgCleared ? undefined : (image ?? ogPfp),
		);
	};

	const clearImage = () => {
		setImage(undefined);
		setImgCleared(true);
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
					{translate("editProfile.title")}
				</ThemedText>
				<View className="w-10" />
			</View>

			<ScrollView
				className="flex-1 px-5"
				contentContainerStyle={{ paddingBottom: insets.bottom + 24 }}
				keyboardShouldPersistTaps="handled"
			>
				<ThemedText className="text-lg font-nunito-bold mt-4 mb-6">
					{translate("editProfile.subtitle")}
				</ThemedText>

				<View className="items-center mb-8">
					<View className="relative w-40">
						<ImageSelector
							setImage={(e) => {
								setImage(e);
								setImgCleared(false);
							}}
							image={imgCleared ? undefined : (image ?? ogPfp)}
							width={160}
							height={160}
							className="items-center gap-2"
						/>
						{(imgCleared ? undefined : (image ?? ogPfp)) && (
							<TouchableOpacity
								onPress={clearImage}
								className="absolute top-0 right-0 w-10 h-10 rounded-full items-center justify-center"
								style={{ backgroundColor: AppColors.danger }}
							>
								<IconSymbol name="xmark" size={18} color="white" />
							</TouchableOpacity>
						)}
					</View>
				</View>

				<View className="gap-5">
					<View className="gap-1">
						<ThemedText className="text-sm font-nunito-bold">
							{translate("editProfile.name")}
						</ThemedText>
						<AuthFormInput
							control={control}
							name="name"
							placeholder={guardian.name || translate("editProfile.name")}
							autoCapitalize="none"
						/>
					</View>

					<View className="gap-1">
						<ThemedText className="text-sm font-nunito-bold">
							{translate("editProfile.username")}
						</ThemedText>
						<AuthFormInput
							control={control}
							name="username"
							placeholder={
								guardian.username || translate("editProfile.username")
							}
							autoCapitalize="none"
						/>
					</View>
				</View>

				<ErrorMessage message={errorText} />

				<View className="mt-8">
					<Button
						label={translate("editProfile.saveChanges")}
						onPress={handleSubmit(onSubmit)}
						disabled={!canUpdate}
					/>
				</View>
			</ScrollView>
		</ThemedView>
	);
}
