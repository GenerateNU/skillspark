import React from 'react';
import { TouchableOpacity, TouchableOpacityProps, useColorScheme } from 'react-native';
import { ThemedText } from '@/components/themed-text';
import { IconSymbol } from '@/components/ui/icon-symbol';

type ListItemProps = TouchableOpacityProps & {
  label: string;
  isLast?: boolean;
};

export function ListItem({ label, isLast, style, ...rest }: ListItemProps) {
  const colorScheme = useColorScheme();
  const borderColor = colorScheme === 'dark' ? '#3f3f46' : '#E5E7EB';

  return (
    <TouchableOpacity
      className="flex-row items-center justify-between py-3 px-4"
      style={[!isLast && { borderBottomWidth: 1, borderBottomColor: borderColor }, style]}
      {...rest}
    >
      <ThemedText className="text-[15px] font-nunito">{label}</ThemedText>
      <IconSymbol name="chevron.right" size={14} color="#9CA3AF" />
    </TouchableOpacity>
  );
}
