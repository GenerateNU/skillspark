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

export default function PaymentScreen() {
  const router = useRouter();
  const insets = useSafeAreaInsets();
  const colorScheme = useColorScheme();
  const theme = Colors[colorScheme ?? 'light'];

  const handleUpdateBilling = () => {};
  const handleDelete = () => {};

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
        <ThemedText style={styles.headerTitle}>Payment</ThemedText>
        <View style={styles.headerRight} />
      </View>
      <View style={styles.content}>
        <ThemedText style={styles.sectionTitle}>Manage Billing</ThemedText>
        <ThemedText style={styles.infoText}>Credit Card</ThemedText>
        <ThemedText style={styles.infoText}>Name</ThemedText>
        <ThemedText style={styles.cardNumber}>**** **** **** XXXX</ThemedText>
        <View style={styles.buttonRow}>
          <TouchableOpacity
            style={styles.updateBtn}
            onPress={handleUpdateBilling}
            activeOpacity={0.8}
          >
            <ThemedText style={styles.updateBtnText}>Update Billing</ThemedText>
          </TouchableOpacity>

          <TouchableOpacity
            style={[styles.deleteBtn, { borderColor: theme.text }]}
            onPress={handleDelete}
            activeOpacity={0.8}
          >
            <ThemedText style={[styles.deleteBtnText, { color: theme.text }]}>Delete</ThemedText>
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
    paddingHorizontal: 20,
    paddingTop: 20,
  },
  sectionTitle: {
    fontSize: 22,
    fontFamily: 'Archivo_700Bold',
    marginBottom: 20,
  },
  infoText: {
    fontSize: 16,
    fontFamily: 'Archivo_400Regular',
    marginBottom: 6,
  },
  cardNumber: {
    fontSize: 16,
    fontFamily: 'Archivo_400Regular',
    marginBottom: 32,
    letterSpacing: 1,
  },
  buttonRow: {
    flexDirection: 'row',
    gap: 16,
  },
  updateBtn: {
    flex: 1,
    backgroundColor: '#2563EB',
    paddingVertical: 14,
    borderRadius: 8,
    alignItems: 'center',
    justifyContent: 'center',
  },
  updateBtnText: {
    color: '#FFFFFF',
    fontSize: 15,
    fontFamily: 'Archivo_600SemiBold',
  },
  deleteBtn: {
    flex: 1,
    backgroundColor: 'transparent',
    paddingVertical: 14,
    borderRadius: 8,
    borderWidth: 1.5,
    alignItems: 'center',
    justifyContent: 'center',
  },
  deleteBtnText: {
    fontSize: 15,
    fontFamily: 'Archivo_400Regular',
  },
});
