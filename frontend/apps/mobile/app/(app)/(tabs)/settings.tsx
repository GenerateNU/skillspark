import ParallaxScrollView from "@/components/parallax-scroll-view";
import { ThemedText } from "@/components/themed-text";
import { ThemedView } from "@/components/themed-view";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { useAuthContext } from "@/hooks/use-auth-context";
import { PageRedirectButton } from "@/components/PageRedirectButton";

export default function TabTwoScreen() {
  const { logout } = useAuthContext();

  const handleLogOut = () => {
    logout();
  };

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
        <ThemedText className="flex-col gap-8" type="title">
          Settings
        </ThemedText>
        <ThemedText>This is where you can change your settings</ThemedText>
        <PageRedirectButton label="Log out" onPress={handleLogOut} />
      </ThemedView>
    </ParallaxScrollView>
  );
}
