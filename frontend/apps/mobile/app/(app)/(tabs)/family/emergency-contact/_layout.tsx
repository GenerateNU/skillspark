import { Stack } from "expo-router";
import React from "react";
import { useColorScheme } from "@/hooks/use-color-scheme";
import { Colors } from "@/constants/theme";
import { useTranslation } from "react-i18next";

export default function EmergencyContactLayout() {
  const colorScheme = useColorScheme();
  const theme = Colors[colorScheme ?? 'light'];
  const { t: translate } = useTranslation();

  return (
    <Stack
      screenOptions={{
        headerShown: false,
        headerTintColor: theme.tint,
        headerStyle: { backgroundColor: theme.background },
        headerTitleStyle: { fontFamily: 'NunitoSans_600SemiBold', color: theme.text },
        headerBackTitle: "", 
      }}
    >
      <Stack.Screen 
        name="manage" 
        options={{ 
          title: translate('emergencyContact.manageTitle'),
          headerShown: false
        }} 
      />
    </Stack>
  );
}