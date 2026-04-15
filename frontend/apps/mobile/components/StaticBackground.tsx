import { BottomSheetBackdropProps } from "@gorhom/bottom-sheet";
import { Pressable, StyleSheet } from "react-native";
import Animated, {
  SharedValue,
  useAnimatedReaction,
  useAnimatedStyle,
  useSharedValue,
} from "react-native-reanimated";

// Backdrop that appears instantly when opening and disappears instantly when
// closing starts — no fade animation in either direction.
export function StaticBackdrop({
  animatedIndex,
  style,
  onPress,
}: {
  animatedIndex: SharedValue<number>;
  style: BottomSheetBackdropProps["style"];
  onPress: () => void;
}) {
  const opacity = useSharedValue(animatedIndex.value > -1 ? 1 : 0);

  useAnimatedReaction(
    () => animatedIndex.value,
    (current, previous) => {
      if (previous === null) return;
      if (current > -1 && previous <= -1) {
        // Sheet just started opening — appear immediately.
        opacity.value = 1;
      } else if (current < previous && current < 0) {
        // Sheet just started closing — disappear immediately.
        opacity.value = 0;
      }
    }
  );

  const animatedStyle = useAnimatedStyle(() => ({ opacity: opacity.value }));

  return (
    <Animated.View
      style={[style, animatedStyle, { backgroundColor: "rgba(0,0,0,0.4)" }]}
    >
      <Pressable style={StyleSheet.absoluteFill} onPress={onPress} />
    </Animated.View>
  );
}
