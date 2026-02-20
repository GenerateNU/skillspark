import React from 'react';
import { View, TouchableOpacity, StyleSheet, useColorScheme } from 'react-native';
import { ThemedText } from '@/components/themed-text';
import { Colors } from '@/constants/theme';

type SectionHeaderProps = {
  title: string;
  actionLabel?: string;
  onAction?: () => void;
};

export function SectionHeader({ title, actionLabel, onAction }: SectionHeaderProps) {
  const colorScheme = useColorScheme();
  const theme = Colors[colorScheme ?? 'light'];

  return (
    <View style={styles.sectionHeaderRow}>
      <ThemedText style={styles.sectionTitle}>{title}</ThemedText>
      {actionLabel && (
        <TouchableOpacity onPress={onAction}>
          <ThemedText style={[styles.addLink, { color: theme.tint }]}>
            {actionLabel}
          </ThemedText>
        </TouchableOpacity>
      )}
    </View>
  );
}

const styles = StyleSheet.create({
  sectionHeaderRow: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
    paddingVertical: 14,
  },
  sectionTitle: { fontSize: 17, fontFamily: 'Archivo_700Bold' },
  addLink: { fontSize: 14, fontFamily: 'Archivo_500Medium' },
});