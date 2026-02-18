import React from 'react';
import { StyleSheet, View, useColorScheme } from 'react-native';
import { ThemedText } from '@/components/themed-text';

type FamilyCardProps = {
  initials: string;
  name: string;
  date: string;
};

export function FamilyCard({ initials, name, date }: FamilyCardProps) {
  const colorScheme = useColorScheme();
  
  const backgroundColor = colorScheme === 'dark' ? '#27272a' : '#F3F4F6';
  const borderColor = colorScheme === 'dark' ? '#3f3f46' : '#E5E7EB';

  return (
    <View style={[styles.container, { backgroundColor }]}>
      <View style={[styles.initialsCircle, { borderColor }]}>
        <ThemedText style={styles.initialsText}>{initials}</ThemedText>
      </View>
      <View>
        <ThemedText style={styles.name}>{name}</ThemedText>
        <ThemedText style={styles.date}>{date}</ThemedText>
      </View>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    width: '48%', 
    borderRadius: 12,
    padding: 10,
    flexDirection: 'row',
    alignItems: 'center',
  },
  initialsCircle: {
    width: 32,
    height: 32,
    borderRadius: 16,
    borderWidth: 1,
    justifyContent: 'center',
    alignItems: 'center',
    marginRight: 8,
  },
  initialsText: {
    fontSize: 12,
    fontFamily: 'Archivo_600SemiBold',
  },
  name: {
    fontSize: 14,
    fontFamily: 'Archivo_500Medium',
  },
  date: {
    fontSize: 10,
    fontFamily: 'Archivo_400Regular',
    color: '#6B7280',
  },
});