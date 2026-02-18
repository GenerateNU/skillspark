import { Stack } from "expo-router";
import React from "react";
import { useColorScheme } from "@/hooks/use-color-scheme";
import { Colors } from "@/constants/theme";

export default function FamilyLayout() {
  const colorScheme = useColorScheme();
  const theme = Colors[colorScheme ?? 'light'];

  return (
    <Stack
      screenOptions={{
        headerShown: true,
        headerTintColor: theme.tint,
        headerStyle: { backgroundColor: theme.background },
        headerTitleStyle: { fontFamily: 'Archivo_600SemiBold', color: theme.text },
        headerBackTitle: "", 
      }}
    >
      <Stack.Screen 
        name="index" 
        options={{ 
          headerShown: false 
        }} 
      />
      <Stack.Screen 
        name="manage" 
        options={{ 
          title: "Manage Child",
          headerShown: false
        }} 
      />
    </Stack>
  );
}