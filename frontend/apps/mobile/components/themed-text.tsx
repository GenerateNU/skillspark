import { Text, type TextProps } from 'react-native';

import { useThemeColor } from '@/hooks/use-theme-color';

export type ThemedTextProps = TextProps & {
  lightColor?: string;
  darkColor?: string;
  className?: string;
  type?: 'default' | 'title' | 'defaultSemiBold' | 'subtitle' | 'link';
};

export function ThemedText({
  style,
  className,
  lightColor,
  darkColor,
  type = 'default',
  ...rest
}: ThemedTextProps) {
  const color = useThemeColor({ light: lightColor, dark: darkColor }, 'text');

  const typeClassName =
    type === 'default'
      ? 'text-base leading-6'
      : type === 'title'
        ? 'text-[32px] leading-[32px] font-bold'
        : type === 'defaultSemiBold'
          ? 'text-base leading-6 font-semibold'
          : type === 'subtitle'
            ? 'text-xl font-bold'
            : type === 'link'
              ? 'leading-[30px] text-base text-[#0a7ea4]'
              : '';

  return (
    <Text
      style={[
        { color },
        style,
      ]}
      className={[typeClassName, className].filter(Boolean).join(' ')}
      {...rest}
    />
  );
}
