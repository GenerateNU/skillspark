import { useEffect, useRef } from "react";
import { Animated, PanResponder } from "react-native";

/**
 * Provides a pan responder and animated value for a draggable bottom sheet.
 * Dragging down past 100px calls `onDismiss`; releasing before that snaps back.
 */
export function useDraggableBottomSheet(onDismiss: () => void) {
  const onDismissRef = useRef(onDismiss);
  useEffect(() => {
    onDismissRef.current = onDismiss;
  }, [onDismiss]);

  const translateY = useRef(new Animated.Value(0)).current;

  const panResponder = useRef(
    PanResponder.create({
      onStartShouldSetPanResponder: () => true,
      onMoveShouldSetPanResponder: (_, gs) => gs.dy > 4,
      onPanResponderMove: (_, gs) => {
        if (gs.dy > 0) translateY.setValue(gs.dy);
      },
      onPanResponderRelease: (_, gs) => {
        if (gs.dy > 100) {
          onDismissRef.current();
        } else {
          Animated.spring(translateY, { toValue: 0, useNativeDriver: true }).start();
        }
      },
    })
  ).current;

  const reset = () => translateY.setValue(0);

  return { panResponder, translateY, reset };
}
