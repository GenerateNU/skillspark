import React from 'react';
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

export default function SettingsScreen() {
  const router = useRouter();
  const insets = useSafeAreaInsets();
  const colorScheme = useColorScheme();
  const theme = Colors[colorScheme ?? 'light'];

  const cardBg = colorScheme === 'dark' ? '#1c1c1e' : '#EFEFEF';
  const dividerColor = colorScheme === 'dark' ? '#3a3a3c' : '#D1D5DB';

  const handleLogOut = () => {};

  const handleDeleteAccount = () => {};

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
      <View style={styles.content}>
        <View style={[styles.card, { backgroundColor: cardBg }]}>
          <TouchableOpacity style={styles.row} activeOpacity={0.6} onPress={() => router.push('/language')}>
            <ThemedText style={styles.rowText}>Language</ThemedText>
            <IconSymbol name="chevron.right" size={16} color="#9CA3AF" />
          </TouchableOpacity>
          <View style={[styles.divider, { backgroundColor: dividerColor }]} />
          <TouchableOpacity style={styles.row} activeOpacity={0.6}>
            <ThemedText style={styles.rowText}>Terms and Conditions</ThemedText>
            <IconSymbol name="chevron.right" size={16} color="#9CA3AF" />
          </TouchableOpacity>
          <View style={[styles.divider, { backgroundColor: dividerColor }]} />
          <TouchableOpacity style={styles.row} activeOpacity={0.6}>
            <ThemedText style={styles.rowText}>Privacy Policy</ThemedText>
            <IconSymbol name="chevron.right" size={16} color="#9CA3AF" />
          </TouchableOpacity>
          <View style={[styles.divider, { backgroundColor: dividerColor }]} />
          <TouchableOpacity style={styles.row} activeOpacity={0.6} onPress={handleLogOut}>
            <ThemedText style={styles.rowText}>Log Out</ThemedText>
          </TouchableOpacity>
          <View style={[styles.divider, { backgroundColor: dividerColor }]} />
          <TouchableOpacity style={styles.row} activeOpacity={0.6} onPress={handleDeleteAccount}>
            <ThemedText style={styles.rowText}>Delete Account</ThemedText>
          </TouchableOpacity>
        </View>
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
  content: {
    paddingHorizontal: 16,
    paddingTop: 20,
  },
  card: {
    borderRadius: 16,
    overflow: 'hidden',
  },
  row: {
    flexDirection: 'row',
    alignItems: 'center',
    justifyContent: 'space-between',
    paddingHorizontal: 16,
    paddingVertical: 18,
  },
  rowText: {
    fontSize: 17,
    fontFamily: 'Archivo_400Regular',
  },
  divider: {
    height: StyleSheet.hairlineWidth,
  },
});
