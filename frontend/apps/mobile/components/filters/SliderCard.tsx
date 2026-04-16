import { AppColors, FontFamilies, FontSizes } from "@/constants/theme";
import { useRef, useState } from "react";
import { PanResponder, Text, View } from "react-native";

const THUMB_SIZE = 22;
const TRACK_HEIGHT = 3;
const TRACK_TOP = (THUMB_SIZE - TRACK_HEIGHT) / 2;

const CARD_STYLE = {
  backgroundColor: "#F0F0F0",
  borderRadius: 18,
  borderWidth: 1,
  borderColor: "rgba(124, 58, 237, 0.22)",
  paddingHorizontal: 16,
  paddingTop: 14,
  paddingBottom: 12,
} as const;

function snapVal(
  raw: number,
  min: number,
  step: number,
  lo: number,
  hi: number,
) {
  const snapped = Math.round((raw - min) / step) * step + min;
  return Math.min(Math.max(snapped, lo), hi);
}

export type SliderCardProps = {
  label: string;
  valueLabel: string;
  value: number | [number, number];
  onValueChange: (val: number[]) => void;
  min: number;
  max: number;
  step?: number;
  minLabel: string;
  maxLabel: string;
};

export function SliderCard({
  label,
  valueLabel,
  value,
  onValueChange,
  min,
  max,
  step = 1,
  minLabel,
  maxLabel,
}: SliderCardProps) {
  const [trackWidth, setTrackWidth] = useState(0);
  const trackWidthRef = useRef(0);

  const isRange = Array.isArray(value);
  const v0 = isRange ? (value as [number, number])[0] : (value as number);
  const v1 = isRange ? (value as [number, number])[1] : max;

  // Refs so PanResponder callbacks always see the latest values
  const v0Ref = useRef(v0);
  const v1Ref = useRef(v1);
  const startPosRef = useRef(0);
  const startPos1Ref = useRef(0);
  const onChangeRef = useRef(onValueChange);

  v0Ref.current = v0;
  v1Ref.current = v1;
  onChangeRef.current = onValueChange;

  const pan0 = useRef(
    PanResponder.create({
      onStartShouldSetPanResponder: () => true,
      onStartShouldSetPanResponderCapture: () => true,
      onPanResponderGrant: () => {
        const tw = trackWidthRef.current;
        startPosRef.current = ((v0Ref.current - min) / (max - min)) * tw;
      },
      onPanResponderMove: (_, g) => {
        const tw = trackWidthRef.current;
        if (!tw) return;
        const newPos = Math.min(Math.max(startPosRef.current + g.dx, 0), tw);
        const raw = (newPos / tw) * (max - min) + min;
        const hi = isRange ? v1Ref.current : max;
        const newVal = snapVal(raw, min, step, min, hi);
        v0Ref.current = newVal;
        onChangeRef.current(isRange ? [newVal, v1Ref.current] : [newVal]);
      },
    }),
  ).current;

  const pan1 = useRef(
    PanResponder.create({
      onStartShouldSetPanResponder: () => true,
      onStartShouldSetPanResponderCapture: () => true,
      onPanResponderGrant: () => {
        const tw = trackWidthRef.current;
        startPos1Ref.current = ((v1Ref.current - min) / (max - min)) * tw;
      },
      onPanResponderMove: (_, g) => {
        const tw = trackWidthRef.current;
        if (!tw) return;
        const newPos = Math.min(Math.max(startPos1Ref.current + g.dx, 0), tw);
        const raw = (newPos / tw) * (max - min) + min;
        const newVal = snapVal(raw, min, step, v0Ref.current, max);
        v1Ref.current = newVal;
        onChangeRef.current([v0Ref.current, newVal]);
      },
    }),
  ).current;

  // Pixel offset of each thumb's left edge within the container.
  // The track occupies [THUMB_SIZE/2 … THUMB_SIZE/2 + trackWidth].
  // A thumb at position `pos` has its center at THUMB_SIZE/2 + pos, which
  // lands exactly on the corresponding point of the track.
  const pos0 = trackWidth > 0 ? ((v0 - min) / (max - min)) * trackWidth : 0;
  const pos1 =
    trackWidth > 0 ? ((v1 - min) / (max - min)) * trackWidth : trackWidth;

  const thumbBase = {
    position: "absolute" as const,
    top: 0,
    width: THUMB_SIZE,
    height: THUMB_SIZE,
    borderRadius: THUMB_SIZE / 2,
    backgroundColor: AppColors.primaryText,
    elevation: 3,
    shadowColor: "#000",
    shadowOpacity: 0.2,
    shadowRadius: 3,
    shadowOffset: { width: 0, height: 1 },
  };

  return (
    <View style={CARD_STYLE}>
      {/* Label row */}
      <View
        style={{
          flexDirection: "row",
          justifyContent: "space-between",
          marginBottom: 12,
        }}
      >
        <Text
          style={{
            fontFamily: FontFamilies.bold,
            fontSize: FontSizes.lg,
            color: AppColors.primaryText,
          }}
        >
          {label}
        </Text>
        <Text
          style={{
            fontFamily: FontFamilies.regular,
            fontSize: FontSizes.sm,
            color: AppColors.mutedText,
          }}
        >
          {valueLabel}
        </Text>
      </View>

      {/* Track + thumbs */}
      <View style={{ height: THUMB_SIZE, marginBottom: 4 }}>
        {/* Gray background track — onLayout gives us the usable track width */}
        <View
          onLayout={(e) => {
            const w = e.nativeEvent.layout.width;
            trackWidthRef.current = w;
            setTrackWidth(w);
          }}
          style={{
            position: "absolute",
            top: TRACK_TOP,
            left: THUMB_SIZE / 2,
            right: THUMB_SIZE / 2,
            height: TRACK_HEIGHT,
            backgroundColor: AppColors.borderLight,
            borderRadius: TRACK_HEIGHT / 2,
          }}
        />

        {trackWidth > 0 && (
          <>
            {/* Filled segment */}
            <View
              pointerEvents="none"
              style={{
                position: "absolute",
                top: TRACK_TOP,
                left: THUMB_SIZE / 2 + (isRange ? pos0 : 0),
                width: isRange ? pos1 - pos0 : pos0,
                height: TRACK_HEIGHT,
                backgroundColor: AppColors.primaryText,
                borderRadius: TRACK_HEIGHT / 2,
              }}
            />

            {/* Thumb 0 */}
            <View {...pan0.panHandlers} style={{ ...thumbBase, left: pos0 }} />

            {/* Thumb 1 — range only */}
            {isRange && (
              <View
                {...pan1.panHandlers}
                style={{ ...thumbBase, left: pos1 }}
              />
            )}
          </>
        )}
      </View>

      {/* Min / max labels */}
      <View style={{ flexDirection: "row", justifyContent: "space-between" }}>
        <Text
          style={{
            fontFamily: FontFamilies.regular,
            fontSize: FontSizes.xs,
            color: AppColors.subtleText,
          }}
        >
          {minLabel}
        </Text>
        <Text
          style={{
            fontFamily: FontFamilies.regular,
            fontSize: FontSizes.xs,
            color: AppColors.subtleText,
          }}
        >
          {maxLabel}
        </Text>
      </View>
    </View>
  );
}
