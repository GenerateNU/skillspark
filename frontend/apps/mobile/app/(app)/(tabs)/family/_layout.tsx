import { Stack } from "expo-router";
import React from "react";
import { useTranslation } from "react-i18next";
import { Colors, FontFamilies } from "@/constants/theme";

export default function FamilyLayout() {
  const theme = Colors.light;
  const { t: translate } = useTranslation();

  return (
    <Stack
      screenOptions={{
        headerShown: false,
        headerTintColor: theme.tint,
        headerStyle: { backgroundColor: theme.background },
        headerTitleStyle: {
          fontFamily: "NunitoSans_600SemiBold",
          color: theme.text,
        },
        headerBackTitle: "",
      }}
    >
      <Stack.Screen
        name="index"
        options={{
          headerShown: false,
        }}
      />
      <Stack.Screen
        name="manage"
        options={{
          title: translate("nav.manageChild"),
          headerShown: false,
        }}
      />
      <Stack.Screen
        name="avatar-picker"
        options={{
          headerShown: false,
        }}
      />
    </Stack>
  );
}
