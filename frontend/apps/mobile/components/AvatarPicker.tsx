import React, { useState } from "react";
import { View, TouchableOpacity } from "react-native";
import { SvgXml } from "react-native-svg";
import { ThemedText } from "@/components/themed-text";
import { Colors } from "@/constants/theme";
import { useColorScheme } from "@/hooks/use-color-scheme";
import {
  UNIQUE_AVATAR_FACES,
  getSvgWithColor,
  getAvatarSvg,
} from "@/constants/avatarFaces";

export { UNIQUE_AVATAR_FACES, getSvgWithColor, getAvatarSvg };

export const AVATAR_COLORS = [
  "#9CAF7D",
  "#F5D878",
  "#F5A85A",
  "#B5A8E0",
  "#F0A0BC",
  "#E07070",
  "#6B9FD4",
  "#A8D4E8",
  "#B0B0B0",
];

export const DEFAULT_AVATAR_COLOR = AVATAR_COLORS[3];

const CIRCLE_SIZE = 62;
const RING_WIDTH = 2.5;
const RING_GAP = 2;
const RING_RADIUS = CIRCLE_SIZE / 2 + RING_GAP + RING_WIDTH;

type AvatarPickerProps = {
  selectedFace: string | null;
  selectedBackground: string;
  onFaceChange: (face: string | null) => void;
  onBackgroundChange: (color: string) => void;
  childInitials?: string;
};

type Tab = "Colors" | "Avatar";

function RingWrapper({
  selected,
  onPress,
  children,
}: {
  selected: boolean;
  onPress: () => void;
  children: React.ReactNode;
}) {
  return (
    <TouchableOpacity
      onPress={onPress}
      activeOpacity={0.75}
      style={{
        padding: RING_GAP,
        borderRadius: RING_RADIUS,
        borderWidth: RING_WIDTH,
        borderColor: selected ? "#6B7FC8" : "transparent",
      }}
    >
      {children}
    </TouchableOpacity>
  );
}

function toRows<T>(items: T[], perRow = 3): T[][] {
  const rows: T[][] = [];
  for (let i = 0; i < items.length; i += perRow) {
    rows.push(items.slice(i, i + perRow));
  }
  return rows;
}

export function AvatarPicker({
  selectedFace,
  selectedBackground,
  onFaceChange,
  onBackgroundChange,
  childInitials = "?",
}: AvatarPickerProps) {
  const [activeTab, setActiveTab] = useState<Tab>("Colors");
  const colorScheme = useColorScheme();
  const theme = Colors[colorScheme ?? "light"];
  const isDark = colorScheme === "dark";

  const selectedBg = selectedBackground || DEFAULT_AVATAR_COLOR;

  const colorRows = toRows(AVATAR_COLORS);
  const avatarItems: (string | null)[] = [
    null,
    ...UNIQUE_AVATAR_FACES.map((f) => f.key),
  ];
  const avatarRows = toRows(avatarItems);

  return (
    <View
      className="rounded-2xl border mb-5 overflow-hidden"
      style={{ borderColor: theme.borderColor }}
    >
      {/* Tab switcher */}
      <View
        className="flex-row p-1.5 gap-1.5"
        style={{ backgroundColor: isDark ? "#1c1c1e" : "#F3F4F6" }}
      >
        {(["Colors", "Avatar"] as Tab[]).map((tab) => (
          <TouchableOpacity
            key={tab}
            onPress={() => setActiveTab(tab)}
            className="flex-1 py-2 rounded-[10px] items-center"
            style={{
              backgroundColor:
                activeTab === tab ? "#8AA0CC" : "transparent",
            }}
          >
            <ThemedText
              className="text-sm font-nunito-semibold"
              style={{ color: activeTab === tab ? "#fff" : theme.text }}
            >
              {tab}
            </ThemedText>
          </TouchableOpacity>
        ))}
      </View>

      {/* Grid */}
      <View
        className="px-4 pt-3 pb-4"
        style={{ backgroundColor: isDark ? "#1c1c1e" : "#fff" }}
      >
        {activeTab === "Colors"
          ? colorRows.map((row, ri) => (
              <View
                key={ri}
                className="flex-row justify-between"
                style={{ marginBottom: ri < colorRows.length - 1 ? 14 : 0 }}
              >
                {row.map((color) => (
                  <RingWrapper
                    key={color}
                    selected={selectedBg === color}
                    onPress={() => onBackgroundChange(color)}
                  >
                    <View
                      className="rounded-full"
                      style={{
                        width: CIRCLE_SIZE,
                        height: CIRCLE_SIZE,
                        backgroundColor: color,
                      }}
                    />
                  </RingWrapper>
                ))}
              </View>
            ))
          : avatarRows.map((row, ri) => (
              <View
                key={ri}
                className="flex-row justify-between"
                style={{ marginBottom: ri < avatarRows.length - 1 ? 14 : 0 }}
              >
                {row.map((faceKey) => {
                  const isSelected = selectedFace === faceKey;

                  if (faceKey === null) {
                    return (
                      <RingWrapper
                        key="initials"
                        selected={isSelected}
                        onPress={() => onFaceChange(null)}
                      >
                        <View
                          className="rounded-full items-center justify-center"
                          style={{
                            width: CIRCLE_SIZE,
                            height: CIRCLE_SIZE,
                            backgroundColor: selectedBg,
                          }}
                        >
                          <ThemedText
                            className="text-lg font-nunito-bold"
                            style={{ color: "#fff" }}
                          >
                            {childInitials}
                          </ThemedText>
                        </View>
                      </RingWrapper>
                    );
                  }

                  const face = UNIQUE_AVATAR_FACES.find(
                    (f) => f.key === faceKey,
                  )!;
                  return (
                    <RingWrapper
                      key={faceKey}
                      selected={isSelected}
                      onPress={() => onFaceChange(faceKey)}
                    >
                      {/* overflow:hidden on this inner View only — keeps the
                          SVG clipped to a circle without affecting the ring */}
                      <View
                        className="rounded-full overflow-hidden"
                        style={{ width: CIRCLE_SIZE, height: CIRCLE_SIZE }}
                      >
                        <SvgXml
                          xml={getSvgWithColor(face.svg, selectedBg)}
                          width={CIRCLE_SIZE}
                          height={CIRCLE_SIZE}
                        />
                      </View>
                    </RingWrapper>
                  );
                })}
              </View>
            ))}
      </View>
    </View>
  );
}
