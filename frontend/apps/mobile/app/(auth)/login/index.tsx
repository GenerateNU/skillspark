import ParallaxScrollView from "@/components/parallax-scroll-view";
import { ThemedText } from "@/components/themed-text";
import { ThemedView } from "@/components/themed-view";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { Button, StyleSheet, TextInput } from "react-native";

export default function LoginScreen() {
  return (
  
        <ThemedView style={styles.titleContainer}>
          <ThemedText type="title">Log In</ThemedText>
          <TextInput
          placeholder="Email"
          />
          <TextInput
          placeholder="Password"
          />
          <Button
            title={"Log In"}
            onPress={() => console.log("smth")}
          />
        </ThemedView>
     );
}

const styles = StyleSheet.create({
    headerImage: {
    color: "#FFFFFF",
    bottom: -90,
    left: -35,
    position: "absolute",
  },
  titleContainer: {
    flexDirection: "row",
    alignItems: "center",
    gap: 8,
  },
  stepContainer: {
    gap: 8,
    marginBottom: 16,
  },
  centerContainer: {
    padding: 16,
    alignItems: "center",
    gap: 8,
  },
  separator: {
    height: 12,
  },
  errorText: {
    color: '#ff4444',
  },
});