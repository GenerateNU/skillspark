import React from 'react';
import { StyleSheet, View } from 'react-native';
import { ThemedText } from '@/components/themed-text';
import { IconSymbol } from '@/components/ui/icon-symbol';

const TAG_COLORS = [
  { bg: '#E6F4EA', border: '#4CAF50', text: '#2E7D32' },
  { bg: '#FFF8E1', border: '#FFC107', text: '#F57F17' },
  { bg: '#FCE4EC', border: '#E91E63', text: '#880E4F' },
  { bg: '#E3F2FD', border: '#2196F3', text: '#0D47A1' },
  { bg: '#F3E5F5', border: '#9C27B0', text: '#4A148C' },
];

const MAX_VISIBLE_TAGS = 3;

export function InterestTags({ interests }: { interests?: string[] | string }) {
  const tags: string[] = Array.isArray(interests)
    ? interests
    : typeof interests === 'string' && interests
    ? interests.split(',').map((s) => s.trim()).filter(Boolean)
    : [];

  if (!tags.length) return null;

  const visible = tags.slice(0, MAX_VISIBLE_TAGS);
  const overflow = tags.length - MAX_VISIBLE_TAGS;

  return (
    <View style={styles.tagsRow}>
      {visible.map((tag, i) => {
        const c = TAG_COLORS[i % TAG_COLORS.length];
        return (
          <View
            key={tag}
            style={[
              styles.tag,
              { backgroundColor: c.bg, borderColor: c.border },
            ]}
          >
            <IconSymbol name="camera.filters" size={13} color={c.border} />
            <ThemedText style={[styles.tagText, { color: c.text }]}>
              {tag}
            </ThemedText>
          </View>
        );
      })}
      {overflow > 0 && (
        <ThemedText style={styles.overflowText}>+{overflow}</ThemedText>
      )}
    </View>
  );
}

const styles = StyleSheet.create({
  tagsRow: { flexDirection: 'row', flexWrap: 'wrap', gap: 6, marginTop: 4 },
  tag: {
    flexDirection: 'row',
    alignItems: 'center',
    gap: 4,
    paddingHorizontal: 8,
    paddingVertical: 4,
    borderRadius: 20,
    borderWidth: 1,
  },
  tagText: { fontSize: 12, fontFamily: 'Archivo_500Medium' },
  overflowText: {
    fontSize: 13,
    color: '#6B7280',
    fontFamily: 'Archivo_500Medium',
    alignSelf: 'center',
  },
});