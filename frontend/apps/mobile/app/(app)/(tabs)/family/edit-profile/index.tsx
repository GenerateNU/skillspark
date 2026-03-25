import { useAuthContext } from "@/hooks/use-auth-context";
import { useRouter } from "expo-router";
import { TouchableOpacity, useColorScheme, View } from "react-native";
import { useSafeAreaInsets } from "react-native-safe-area-context";
import { Colors, AppColors } from '@/constants/theme';
import { ThemedView } from "@/components/themed-view";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { ThemedText } from "@/components/themed-text";
import { useState } from "react";
import { useForm } from "react-hook-form";
import { ErrorScreen } from "@/components/ErrorScreen";
import { Guardian, useGetGuardianById } from "@skillspark/api-client";
import { AuthFormInput } from "@/components/AuthFormInput";
import { Button } from "@/components/Button";
import { ErrorMessage } from "@/components/ErrorMessage";

type EditFormData = {
  name: string;
  username: string;
};

export default function EditProfileScreen() {
    const router = useRouter();
    const insets = useSafeAreaInsets();
    const colorScheme = useColorScheme();
    const theme = Colors[colorScheme ?? 'light'];
    
    const [errorText, setErrorText] = useState("");
    const { update, guardianId, langPref } = useAuthContext();

    const guardianData = useGetGuardianById(guardianId!, {
        query: {
            enabled: !!guardianId,
        }
    })
    
    const { control, handleSubmit } = useForm<EditFormData>({
    defaultValues: {
        name: "",
        username: "",
    },
    });

    if (!guardianId) {
        return <ErrorScreen message="Illegal state: no guardian ID found"></ErrorScreen>;
    }
    
    const onSubmit = (formData: EditFormData) => {
        const guardian = (guardianData as unknown as { data: Guardian })?.data;
        const id = guardianId!;
        const email = guardian.email;
        const language_preference = langPref!;
        const name = formData.name === "" ? guardian.name : formData.name;
        const username = formData.username === "" ? guardian.username : formData.username;
        update(setErrorText, id, email, language_preference, name, username);
    };

    return (
        <ThemedView className="flex-1" style={{ paddingTop: insets.top }}>
            <View className="flex-row items-center justify-between px-5 py-3">
                <TouchableOpacity
                  onPress={() => router.navigate('/profile')}
                  className="w-10 justify-center items-start"
                  hitSlop={{ top: 10, bottom: 10, left: 10, right: 10 }}
                >
                  <IconSymbol name="chevron.left" size={24} color={theme.text} />
                </TouchableOpacity>
                <ThemedText className="text-xl text-center font-nunito-bold">Family Information</ThemedText>
                <View className="w-10" />
                <AuthFormInput
                      control={control}
                      name="name"
                      placeholder="Name"
                      autoCapitalize="none"
                    />
                    <AuthFormInput
                      control={control}
                      name="username"
                      placeholder="Username"
                    />
                    <Button label="Log In" onPress={handleSubmit(onSubmit)} />
                    <ErrorMessage message={errorText} />
            </View>
        </ThemedView>
    )
}