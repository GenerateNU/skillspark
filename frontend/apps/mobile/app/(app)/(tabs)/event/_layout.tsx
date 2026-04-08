import { Stack } from "expo-router";

export default function EventLayout() {
  return (
    <Stack screenOptions={{ headerShown: false }}>
      <Stack.Screen name="[id]" />
      <Stack.Screen name="review" />
    </Stack>
  );
}
