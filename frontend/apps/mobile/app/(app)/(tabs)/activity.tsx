import { ThemedText } from "@/components/themed-text";
import { ThemedView } from "@/components/themed-view";
import { useSafeAreaInsets } from "react-native-safe-area-context";

export default function ActivityScreen() {
  const insets = useSafeAreaInsets();
  return (
    <ThemedView style={{ paddingTop: insets.top }}>
      <ThemedText>hellooooooo</ThemedText>
    </ThemedView>
  );
}
