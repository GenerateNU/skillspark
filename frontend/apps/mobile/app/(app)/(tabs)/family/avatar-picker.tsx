import React, { useState } from "react";
import { View, TouchableOpacity, ScrollView } from "react-native";
import { Stack, useRouter, useLocalSearchParams } from "expo-router";
import { useSafeAreaInsets } from "react-native-safe-area-context";
import { ThemedText } from "@/components/themed-text";
import { ThemedView } from "@/components/themed-view";
import { Colors, AppColors } from "@/constants/theme";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { useColorScheme } from "@/hooks/use-color-scheme";
import { AvatarPicker, DEFAULT_AVATAR_COLOR } from "@/components/AvatarPicker";
import { ChildAvatar, getInitials } from "@/components/ChildAvatar";
import { resolvePendingAvatarCallback } from "@/constants/avatarPickerStore";
import { useTranslation } from "react-i18next";

export default function AvatarPickerScreen() {
  const router = useRouter();
  const params = useLocalSearchParams();
  const insets = useSafeAreaInsets();
  const colorScheme = useColorScheme();
  const theme = Colors[colorScheme ?? "light"];
  const { t: translate } = useTranslation();

  const childName = (params.childName as string) || "?";
  const initials = getInitials(childName);

  const [face, setFace] = useState<string | null>(
    params.avatarFace ? (params.avatarFace as string) : null,
  );
  const [background, setBackground] = useState<string>(
    (params.avatarBackground as string) || DEFAULT_AVATAR_COLOR,
  );

  const handleSave = () => {
    resolvePendingAvatarCallback({ face, background });
    router.back();
  };

  const handleCancel = () => {
    router.back();
  };

  return (
    <ThemedView className="flex-1" style={{ paddingTop: insets.top }}>
      <Stack.Screen options={{ headerShown: false }} />
      <ScrollView
        contentContainerStyle={{
          paddingHorizontal: 20,
          paddingBottom: 40,
          paddingTop: 10,
        }}
        showsVerticalScrollIndicator={false}
      >
        {/* Header */}
        <View className="flex-row items-center justify-between mb-6">
          <TouchableOpacity
            onPress={handleCancel}
            className="w-8 h-8 justify-center items-start"
          >
            <IconSymbol name="chevron.left" size={24} color={theme.text} />
          </TouchableOpacity>
          <ThemedText className="text-xl text-center font-nunito-bold">
            {translate("familyInformation.title")}
          </ThemedText>
          <View className="w-8" />
        </View>

        {/* Section title */}
        <ThemedText className="text-[22px] font-nunito-semibold mb-5">
          {translate("childProfile.editTitle")}
        </ThemedText>

        {/* Avatar preview with pencil overlay */}
        <View className="items-center mb-5">
          <View style={{ position: "relative" }}>
            <ChildAvatar
              name={childName}
              avatarFace={face}
              avatarBackground={background}
              size={72}
            />
            <View
              style={{
                position: "absolute",
                bottom: 0,
                right: 0,
                width: 22,
                height: 22,
                borderRadius: 11,
                backgroundColor: theme.text,
                alignItems: "center",
                justifyContent: "center",
              }}
            >
              <IconSymbol name="pencil" size={11} color={theme.background} />
            </View>
          </View>
          <ThemedText
            className="text-sm font-nunito-semibold mt-2"
            style={{ color: AppColors.mutedText }}
          >
            {translate("childProfile.changeProfilePicture")}
          </ThemedText>
        </View>

        {/* Picker */}
        <AvatarPicker
          selectedFace={face}
          selectedBackground={background}
          onFaceChange={setFace}
          onBackgroundChange={setBackground}
          childInitials={initials}
        />

        {/* Cancel / Save buttons */}
        <View className="flex-row gap-3 mt-2">
          <TouchableOpacity
            className="flex-1 py-4 rounded-xl items-center justify-center border"
            style={{ borderColor: theme.borderColor }}
            onPress={handleCancel}
          >
            <ThemedText className="text-base font-nunito-semibold">
              {translate("common.cancel")}
            </ThemedText>
          </TouchableOpacity>
          <TouchableOpacity
            className="flex-1 py-4 rounded-xl items-center justify-center"
            style={{ backgroundColor: theme.text }}
            onPress={handleSave}
          >
            <ThemedText
              className="text-base font-nunito-semibold"
              style={{ color: theme.background }}
            >
              {translate("common.save")}
            </ThemedText>
          </TouchableOpacity>
        </View>
      </ScrollView>
    </ThemedView>
  );
}
