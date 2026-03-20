import React from 'react';
import { View } from 'react-native';
import { ThemedText } from '@/components/themed-text';
import { IconSymbol } from '@/components/ui/icon-symbol';
import { AppColors, TAG_COLORS } from '@/constants/theme';
import { useTranslation } from 'react-i18next';

const MAX_VISIBLE_TAGS = 3;

export function InterestTags({ interests }: { interests?: string[] | string }) {
  const { t } = useTranslation();
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
              {t(`interests.${tag}`, { defaultValue: tag })}
            </ThemedText>
          </View>
        );
      })}
      {overflow > 0 && (
        <ThemedText className="text-[13px] font-nunito-medium self-center" style={{ color: AppColors.mutedText }}>+{overflow}</ThemedText>
      )}
    </View>
  );
}
