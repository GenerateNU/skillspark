import React from "react";
import { View } from "react-native";
import { SvgXml } from "react-native-svg";
import { ThemedText } from "@/components/themed-text";
import { DEFAULT_AVATAR_COLOR, DARK_AVATAR_COLORS } from "@/constants/avatarColors";
import { getAvatarSvg, getSvgWithColor } from "@/constants/avatarFaces";

type ChildAvatarProps = {
  name: string;
  avatarFace?: string | null;
  avatarBackground?: string | null;
  size?: number;
};

/** First letter of first word + first letter of last word (if present). */
export function getInitials(name: string): string {
  const parts = name.trim().split(/\s+/).filter(Boolean);
  if (parts.length === 0) return "?";
  if (parts.length === 1) return parts[0][0].toUpperCase();
  return (parts[0][0] + parts[parts.length - 1][0]).toUpperCase();
}

export function ChildAvatar({
  name,
  avatarFace,
  avatarBackground,
  size = 44,
}: ChildAvatarProps) {
  const bgColor = avatarBackground || DEFAULT_AVATAR_COLOR;
  const svgTemplate = avatarFace ? getAvatarSvg(avatarFace) : null;
  const borderRadius = size / 2;

  if (svgTemplate) {
    return (
      <View
        style={{
          width: size,
          height: size,
          borderRadius,
          overflow: "hidden",
        }}
      >
        <SvgXml
          xml={getSvgWithColor(svgTemplate, bgColor)}
          width={size}
          height={size}
        />
      </View>
    );
  }

  return (
    <View
      style={{
        width: size,
        height: size,
        borderRadius,
        backgroundColor: bgColor,
        alignItems: "center",
        justifyContent: "center",
      }}
    >
      <ThemedText
        style={{
          fontSize: size * 0.32,
          fontFamily: "Nunito-SemiBold",
          color: DARK_AVATAR_COLORS[bgColor] ?? "#5A5A5A",
        }}
      >
        {getInitials(name)}
      </ThemedText>
    </View>
  );
}
