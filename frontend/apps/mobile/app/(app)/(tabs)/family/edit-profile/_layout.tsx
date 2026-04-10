import { Stack } from "expo-router";
import React from "react";
import { Colors } from "@/constants/theme";

export default function EditLayout() {
  const theme = Colors.light;

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
          animation: "slide_from_bottom",
        }}
      />
    </Stack>
  );
}
