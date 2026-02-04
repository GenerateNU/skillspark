import { Image } from "expo-image";
import { StyleSheet } from "react-native";

import ParallaxScrollView from "@/components/parallax-scroll-view";
import { ThemedText } from "@/components/themed-text";
import { ThemedView } from "@/components/themed-view";

export default function HomeScreen() {
  return (
    <ParallaxScrollView
      headerBackgroundColor={{ light: "#BBBBBB", dark: "#333333" }}
      headerImage={
        <Image
          source={require("@/assets/images/generate.png")}
          style={styles.generateLogo}
        />
      }
    >
      <ThemedView style={styles.stepContainer}>
        <ThemedText type="title">Welcome to SkillSpark</ThemedText>
      </ThemedView>
    </ParallaxScrollView>
  );
}

const styles = StyleSheet.create({
  titleContainer: {
    flexDirection: "row",
    alignItems: "center",
    gap: 8,
  },
  stepContainer: {
    gap: 8,
    marginBottom: 8,
  },
  generateLogo: {
    height: 300,
    width: 300,
    position: "absolute",
    bottom: -60,
    left: -35,
  },
});
