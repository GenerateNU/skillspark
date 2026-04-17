import { View } from "react-native";
import type { StyleProp, ViewStyle } from "react-native";
import { ChildAvatar } from "@/components/ChildAvatar";
import type { Child } from "@skillspark/api-client";

interface ChildAvatarGroupProps {
  children: Child[];
  size?: number;
  style?: StyleProp<ViewStyle>;
}

/** Renders a row of child avatars. Defaults to flex-row with wrap and a 6px gap. */
export function ChildAvatarGroup({
  children,
  size = 32,
  style,
}: ChildAvatarGroupProps) {
  return (
    <View style={[{ flexDirection: "row", flexWrap: "wrap", gap: 6 }, style]}>
      {children.map((child) => (
        <ChildAvatar
          key={child.id}
          name={child.name}
          avatarFace={child.avatar_face}
          avatarBackground={child.avatar_background}
          size={size}
        />
      ))}
    </View>
  );
}
