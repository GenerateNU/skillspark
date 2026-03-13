import React from 'react';
import { View } from 'react-native';
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
    <View className="flex-row flex-wrap gap-[6px] mt-1">
      {visible.map((tag, i) => {
        const c = TAG_COLORS[i % TAG_COLORS.length];
        return (
          <View
            key={tag}
            className="flex-row items-center gap-1 px-2 py-1 rounded-full border"
            style={{ backgroundColor: c.bg, borderColor: c.border }}
          >
            <IconSymbol name="camera.filters" size={13} color={c.border} />
            <ThemedText className="text-xs font-nunito-medium" style={{ color: c.text }}>
              {tag}
            </ThemedText>
          </View>
        );
      })}
      {overflow > 0 && (
        <ThemedText className="text-[13px] text-[#6B7280] font-nunito-medium self-center">+{overflow}</ThemedText>
      )}
    </View>
  );
}
