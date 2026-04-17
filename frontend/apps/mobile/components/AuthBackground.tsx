import { Dimensions, StyleSheet } from "react-native";
import { SvgXml } from "react-native-svg";
import { BACKGROUND } from "@/constants/background";

const { width, height } = Dimensions.get("screen");

export function AuthBackground() {
  return (
    <SvgXml
      xml={BACKGROUND}
      width={width}
      height={height}
      preserveAspectRatio="xMidYMid slice"
      style={StyleSheet.absoluteFill}
    />
  );
}
