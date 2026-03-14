import ParallaxScrollView from "@/components/parallax-scroll-view";
import { ThemedText } from "@/components/themed-text";
import { ThemedView } from "@/components/themed-view";
import { IconSymbol } from "@/components/ui/icon-symbol";

export default function TabTwoScreen() {
  return (
    <ParallaxScrollView
      headerBackgroundColor={{ light: "#D0D0D0", dark: "#353636" }}
      headerImage={
        <IconSymbol
          size={310}
          color="#808080"
          name="gearshape"
          style={{
            color: "#808080",
            bottom: -90,
            left: -35,
            position: "absolute",
          }}
        />
      }
    >
      <ThemedView className="flex-col gap-2">
        <ThemedText type="title">Settings</ThemedText>
        <ThemedText>This is where you can change your settings</ThemedText>
      </ThemedView>
    </ParallaxScrollView>
  );
}
