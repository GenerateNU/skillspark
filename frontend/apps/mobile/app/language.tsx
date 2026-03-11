import React, { useState } from 'react';
import {
  StyleSheet,
  View,
  TouchableOpacity,
  useColorScheme,
} from 'react-native';
import { useRouter } from 'expo-router';
import { useSafeAreaInsets } from 'react-native-safe-area-context';
import { ThemedText } from '@/components/themed-text';
import { ThemedView } from '@/components/themed-view';
import { IconSymbol } from '@/components/ui/icon-symbol';
import { Colors } from '@/constants/theme';

const LANGUAGES = [
  { code: 'en', label: 'English', flag: '🇺🇸' },
  { code: 'th', label: 'Thai',    flag: '🇹🇭' },
];

export default function LanguageScreen() {
  const router = useRouter();
  const insets = useSafeAreaInsets();
  const colorScheme = useColorScheme();
  const theme = Colors[colorScheme ?? 'light'];

  const dividerColor = colorScheme === 'dark' ? '#3a3a3c' : '#E5E7EB';

  const [selected, setSelected] = useState('en');

  return (
    <ThemedView style={[styles.container, { paddingTop: insets.top }]}>
      <View style={styles.header}>
        <TouchableOpacity
          onPress={() => router.back()}
          style={styles.backButton}
          hitSlop={{ top: 10, bottom: 10, left: 10, right: 10 }}
        >
          <IconSymbol name="chevron.left" size={24} color={theme.text} />
        </TouchableOpacity>
        <ThemedText style={styles.headerTitle}>Settings</ThemedText>
        <View style={styles.headerRight} />
      </View>
      <ThemedText style={styles.sectionLabel}>Language</ThemedText>
      <View style={styles.list}>
        {LANGUAGES.map((lang, index) => (
          <React.Fragment key={lang.code}>
            <TouchableOpacity
              style={styles.row}
              onPress={() => setSelected(lang.code)}
              activeOpacity={0.6}
            >
              <ThemedText style={styles.flag}>{lang.flag}</ThemedText>
              <ThemedText style={styles.langLabel}>{lang.label}</ThemedText>
              <IconSymbol
                name={selected === lang.code ? 'checkmark.circle.fill' : 'circle'}
                size={26}
                color={selected === lang.code ? theme.text : '#C7C7CC'}
              />
            </TouchableOpacity>
            {index < LANGUAGES.length - 1 && (
              <View style={[styles.divider, { backgroundColor: dividerColor }]} />
            )}
          </React.Fragment>
        ))}
      </View>
    </ThemedView>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
  },
  header: {
    flexDirection: 'row',
    alignItems: 'center',
    justifyContent: 'space-between',
    paddingHorizontal: 20,
    paddingVertical: 14,
  },
  backButton: {
    width: 40,
    justifyContent: 'center',
    alignItems: 'flex-start',
  },
  headerTitle: {
    fontSize: 20,
    fontFamily: 'Archivo_700Bold',
    textAlign: 'center',
  },
  headerRight: {
    width: 40,
  },
  sectionLabel: {
    fontSize: 24,
    fontFamily: 'Archivo_700Bold',
    paddingHorizontal: 20,
    paddingTop: 16,
    paddingBottom: 20,
  },
  list: {
    paddingHorizontal: 20,
  },
  row: {
    flexDirection: 'row',
    alignItems: 'center',
    paddingVertical: 18,
    gap: 14,
  },
  flag: {
    fontSize: 38,
    lineHeight: 46,
  },
  langLabel: {
    flex: 1,
    fontSize: 18,
    fontFamily: 'Archivo_400Regular',
  },
  divider: {
    height: StyleSheet.hairlineWidth,
    marginLeft: 66,
  },
});
