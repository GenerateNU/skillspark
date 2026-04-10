import { BottomTabBarProps } from "@react-navigation/bottom-tabs";
import * as Haptics from "expo-haptics";
import { useEffect } from "react";
import Animated, {
  FadeIn,
  FadeOut,
  useAnimatedStyle,
  useSharedValue,
  withSpring,
} from "react-native-reanimated";
import { Pressable, View } from "react-native";
import { useSafeAreaInsets } from "react-native-safe-area-context";

import { IconSymbol } from "@/components/ui/icon-symbol";

const TAB_WIDTH = 68;
const TAB_HEIGHT = 52;
const PILL_PADDING = 8;
const GAP = 4;

const VISIBLE_TABS = ["index", "map", "activity", "profile"];

const TAB_ICONS: Record<string, any> = {
  index: "house.fill",
  map: "map.fill",
  activity: "bolt.fill",
  profile: "person.fill",
};

export function FloatingTabBar({
  state,
  descriptors,
  navigation,
}: BottomTabBarProps) {
  const insets = useSafeAreaInsets();
  const visibleRoutes = state.routes.filter((r) =>
    VISIBLE_TABS.includes(r.name),
  );

  const activeIndex = visibleRoutes.findIndex(
    (r) => r.name === state.routes[state.index]?.name,
  );

  const translateX = useSharedValue(
    PILL_PADDING + Math.max(activeIndex, 0) * (TAB_WIDTH + GAP),
  );

  useEffect(() => {
    if (activeIndex !== -1) {
      translateX.value = withSpring(
        PILL_PADDING + activeIndex * (TAB_WIDTH + GAP),
        { damping: 28, stiffness: 350, mass: 0.8 },
      );
    }
  }, [activeIndex]); // eslint-disable-line react-hooks/exhaustive-deps

  const bubbleStyle = useAnimatedStyle(() => ({
    transform: [{ translateX: translateX.value }],
  }));

  return (
    <View
      className="absolute left-0 right-0 items-center"
      style={{ bottom: Math.max(insets.bottom, 8) + 12 }}
      pointerEvents="box-none"
    >
      <View
        className="flex-row items-center bg-[#1a1a1a] rounded-[50px] p-2 gap-1"
        style={{
          shadowColor: "#000",
          shadowOffset: { width: 0, height: 8 },
          shadowOpacity: 0.22,
          shadowRadius: 18,
          elevation: 12,
        }}
      >
        <Animated.View
          style={[
            {
              position: "absolute",
              top: PILL_PADDING,
              left: 0,
              width: TAB_WIDTH,
              height: TAB_HEIGHT,
              borderRadius: 36,
              backgroundColor: "rgba(235, 237, 255, 0.95)",
            },
            bubbleStyle,
          ]}
          pointerEvents="none"
        />
        {visibleRoutes.map((route) => {
          const isFocused = state.routes[state.index]?.name === route.name;
          const options = descriptors[route.key].options;
          const label =
            typeof options.title === "string" ? options.title : route.name;
          const iconName = TAB_ICONS[route.name] ?? "circle";

          const onPress = () => {
            if (process.env.EXPO_OS === "ios") {
              Haptics.impactAsync(Haptics.ImpactFeedbackStyle.Light);
            }
            const event = navigation.emit({
              type: "tabPress",
              target: route.key,
              canPreventDefault: true,
            });
            if (!isFocused && !event.defaultPrevented) {
              navigation.navigate(route.name);
            }
          };

          const onLongPress = () => {
            navigation.emit({ type: "tabLongPress", target: route.key });
          };

          return (
            <Pressable
              key={route.key}
              onPress={onPress}
              onLongPress={onLongPress}
              accessibilityRole="button"
              accessibilityState={isFocused ? { selected: true } : {}}
              accessibilityLabel={label}
              className="flex-col items-center justify-center gap-[3px]"
              style={{ width: TAB_WIDTH, height: TAB_HEIGHT }}
            >
              <IconSymbol
                name={iconName}
                size={20}
                color={isFocused ? "#1a1a1a" : "rgba(255,255,255,0.6)"}
              />
              {isFocused && (
                <Animated.Text
                  entering={FadeIn.duration(200)}
                  exiting={FadeOut.duration(120)}
                  style={{
                    fontSize: 13,
                    fontWeight: "600",
                    color: "#1a1a1a",
                    letterSpacing: 0.1,
                  }}
                >
                  {label}
                </Animated.Text>
              )}
            </Pressable>
          );
        })}
      </View>
    </View>
  );
}
