import React from 'react';
import { StyleSheet, TouchableOpacity, TouchableOpacityProps, useColorScheme } from 'react-native';
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
      style={[
        styles.container, 
        !isLast && { borderBottomWidth: 1, borderBottomColor: borderColor },
        style
      ]} 
      {...rest}
    >
      <ThemedText style={styles.text}>{label}</ThemedText>
      <IconSymbol name="chevron.right" size={14} color="#9CA3AF" />
    </TouchableOpacity>
  );
}

const styles = StyleSheet.create({
  container: {
    flexDirection: 'row',
    alignItems: 'center',
    justifyContent: 'space-between',
    paddingVertical: 12,
    paddingHorizontal: 16,
  },
  text: {
    fontSize: 15,
    fontFamily: 'Archivo_400Regular',
  },
});