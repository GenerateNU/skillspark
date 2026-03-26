/**
 * Below are the colors that are used in the app. The colors are defined in the light and dark mode.
 * There are many other ways to style your app. For example, [Nativewind](https://www.nativewind.dev/), [Tamagui](https://tamagui.dev/), [unistyles](https://reactnativeuninstyles.vercel.app), etc.
 */

import { Platform } from "react-native";

const tintColorLight = "#0a7ea4";
const tintColorDark = "#fff";

export const Colors = {
  light: {
    text: "#11181C",
    background: "#fff",
    tint: tintColorLight,
    icon: "#687076",
    tabIconDefault: "#687076",
    tabIconSelected: tintColorLight,
    inputBg: '#F3F4F6',
    dropdownBg: '#FFFFFF',
    borderColor: '#E5E7EB',
  },
  dark: {
    text: "#ECEDEE",
    background: "#151718",
    tint: tintColorDark,
    icon: "#9BA1A6",
    tabIconDefault: "#9BA1A6",
    tabIconSelected: tintColorDark,
    inputBg: '#27272a',
    dropdownBg: '#1c1c1e',
    borderColor: '#3f3f46',
  },
};

/** Static (non-theme-dependent) app-wide color constants */
export const AppColors = {
  primaryText: '#111',
  secondaryText: '#374151',
  mutedText: '#6B7280',
  subtleText: '#9CA3AF',
  borderLight: '#D1D5DB',
  divider: '#E5E7EB',
  primaryBlue: '#2563EB',
  danger: '#EF4444',
  starFilled: '#FBBF24',
  checkboxSelected: '#1F2937',
  placeholderText: '#9CA3AF',
  savedBackground: "#99C0EE4D",
  white: '#fff',
  surfaceGray: '#F3F4F6',
  purple: '#7C3AED',
  violet: '#8B5CF6',
  emerald: '#10B981',
  green: '#059669',
  blue: '#3B82F6',
  amber: '#F59E0B',
  pink: '#EC4899',
  mintGreen: '#A7F3D0',
  badgeGreenBg: '#D1FAE5',
  badgeGreenText: '#065F46',
  categoryFallback: '#6B8888',
  imagePlaceholder: '#D9D9D9',
  cardOverlay: 'rgba(0,0,0,0.42)',
};

/** Tag color palette shared across interest tag components */
export const TAG_COLORS = [
  { bg: '#E6F4EA', border: '#4CAF50', text: '#2E7D32' },
  { bg: '#FFF8E1', border: '#FFC107', text: '#F57F17' },
  { bg: '#FCE4EC', border: '#E91E63', text: '#880E4F' },
  { bg: '#E3F2FD', border: '#2196F3', text: '#0D47A1' },
  { bg: '#F3E5F5', border: '#9C27B0', text: '#4A148C' },
];

export const FontFamilies = {
  regular: 'NunitoSans_400Regular',
  medium: 'NunitoSans_500Medium',
  semiBold: 'NunitoSans_600SemiBold',
  bold: 'NunitoSans_700Bold',
};

export const FontSizes = {
  xs: 11,    // subtle/small labels (location, child name, small badges)
  sm: 12,    // secondary text (date, time, age)
  md: 13,    // avatar initials, age badge label
  base: 14,  // card titles, category labels
  lg: 15,    // section headers
  hero: 30,  // greeting / hero text
};

export const Fonts = Platform.select({
  ios: {
    /** iOS `UIFontDescriptorSystemDesignDefault` */
    sans: "system-ui",
    /** iOS `UIFontDescriptorSystemDesignSerif` */
    serif: "ui-serif",
    /** iOS `UIFontDescriptorSystemDesignRounded` */
    rounded: "ui-rounded",
    /** iOS `UIFontDescriptorSystemDesignMonospaced` */
    mono: "ui-monospace",
  },
  default: {
    sans: "normal",
    serif: "serif",
    rounded: "normal",
    mono: "monospace",
  },
  web: {
    sans: "system-ui, -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif",
    serif: "Georgia, 'Times New Roman', serif",
    rounded:
      "'SF Pro Rounded', 'Hiragino Maru Gothic ProN', Meiryo, 'MS PGothic', sans-serif",
    mono: "SFMono-Regular, Menlo, Monaco, Consolas, 'Liberation Mono', 'Courier New', monospace",
  },
});
