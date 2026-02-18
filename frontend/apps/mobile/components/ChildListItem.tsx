import React from 'react';
import { View, TouchableOpacity, StyleSheet } from 'react-native';
import { ThemedText } from '@/components/themed-text';
import { IconSymbol } from '@/components/ui/icon-symbol';
import { InterestTags } from './InterestTags';
import { Child } from '@skillspark/api-client';

type ChildListItemProps = {
  child: Child; 
  onPress?: () => void;
};

export function ChildListItem({ child, onPress }: ChildListItemProps) {
  return (
    <TouchableOpacity style={styles.row} onPress={onPress} activeOpacity={0.7}>
      <View style={styles.initialsCircle}>
        <ThemedText style={styles.initialsText}>
          {child.name?.slice(0, 2).toUpperCase() || '??'}
        </ThemedText>
      </View>
      <View style={styles.rowInfo}>
        <ThemedText style={styles.rowTitle}>{child.name}</ThemedText>
        <ThemedText style={styles.rowSub}>
          {child.birth_month ? `${child.birth_month}, ` : ''}
          {child.birth_year}
        </ThemedText>
        {child.interests && (
          <ThemedText style={styles.rowSub}>Interests</ThemedText>
        )}
        <InterestTags interests={child.interests} />
      </View>
      <IconSymbol name="chevron.right" size={18} color="#9CA3AF" />
    </TouchableOpacity>
  );
}

const styles = StyleSheet.create({
  row: {
    flexDirection: 'row',
    alignItems: 'flex-start',
    paddingVertical: 12,
    gap: 12,
  },
  initialsCircle: {
    width: 44,
    height: 44,
    borderRadius: 22,
    borderWidth: 1.5,
    borderColor: '#9CA3AF',
    justifyContent: 'center',
    alignItems: 'center',
  },
  initialsText: { fontSize: 15, fontFamily: 'Archivo_600SemiBold' },
  rowInfo: { flex: 1, gap: 2 },
  rowTitle: { fontSize: 16, fontFamily: 'Archivo_600SemiBold' },
  rowSub: { fontSize: 13, color: '#6B7280', fontFamily: 'Archivo_400Regular' },
});