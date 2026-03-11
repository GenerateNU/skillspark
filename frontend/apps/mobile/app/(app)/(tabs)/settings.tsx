import { StyleSheet, Text, TouchableOpacity, View } from "react-native";
import ParallaxScrollView from "@/components/parallax-scroll-view";
import { ThemedText } from "@/components/themed-text";
import { ThemedView } from "@/components/themed-view";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { useAuthContext } from "@/hooks/use-auth-context";

export default function TabTwoScreen() {
  const { logout, debug } = useAuthContext();

  const handleLogOut = () => {
    logout();
  };

  const handleDebug = () => {
    debug();
  };

  return (
    <ParallaxScrollView
      headerBackgroundColor={{ light: "#D0D0D0", dark: "#353636" }}
      headerImage={
        <IconSymbol
          size={310}
          color="#808080"
          name="gearshape"
          style={styles.headerImage}
        />
      }
    >
      <ThemedView style={styles.titleContainer}>
        <ThemedText type="title">Settings</ThemedText>
        <ThemedText>This is where you can change your settings</ThemedText>
      </ThemedView>
      <View style={{ width: "100%", paddingHorizontal: 24, gap: 16, alignItems: "center" }}>
        <TouchableOpacity
          style={{
            borderRadius: 8,
            padding: 10,
            width: "100%",
            alignItems: "center",
          }}
          onPress={handleLogOut}
          activeOpacity={0.5}
        >
          <Text style={{ color: "#3b82f6", fontSize: 16, fontWeight: "500" }}>Log out</Text>
        </TouchableOpacity>

        <TouchableOpacity
          style={{
            borderRadius: 8,
            padding: 10,
            width: "100%",
            alignItems: "center",
          }}
          onPress={handleDebug}
          activeOpacity={0.5}
        >
          <Text style={{ color: "#3b82f6", fontSize: 16, fontWeight: "500" }}>DEBUG: Check context values</Text>
        </TouchableOpacity>
      </View>
    </ParallaxScrollView>
  );
}

const styles = StyleSheet.create({
  headerImage: {
    color: "#808080",
    bottom: -90,
    left: -35,
    position: "absolute",
  },
  titleContainer: {
    flexDirection: "column",
    gap: 8,
  },
});
