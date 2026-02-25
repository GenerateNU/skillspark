import React, { useState } from 'react';
import { StyleSheet, View, TextInput, TouchableOpacity, Alert, useColorScheme, ScrollView } from 'react-native';
import { Stack, useRouter, useLocalSearchParams } from 'expo-router';
import { ThemedText } from '@/components/themed-text';
import { ThemedView } from '@/components/themed-view';
import { Colors } from '@/constants/theme';
import { useCreateChild, useUpdateChild, useDeleteChild } from '@skillspark/api-client';

const GUARDIAN_ID = '88888888-8888-8888-8888-888888888888';

export default function ManageChildScreen() {
  const router = useRouter();
  const params = useLocalSearchParams();
  const colorScheme = useColorScheme();
  const theme = Colors[colorScheme ?? 'light'];
  
  const isEditing = !!params.id;
  
  const [name, setName] = useState(params.name as string || '');
  const [birthYear, setBirthYear] = useState(params.birth_year as string || '');
  const [birthMonth, setBirthMonth] = useState(params.birth_month as string || '');
  const [schoolId, setSchoolId] = useState(params.school_id as string || '');
  
  const [interests, setInterests] = useState(
    Array.isArray(params.interests) 
      ? params.interests.join(', ') 
      : (params.interests as string || '')
  );

  const [isSubmitting, setIsSubmitting] = useState(false);

  const createChildMutation = useCreateChild();
  const updateChildMutation = useUpdateChild();
  const deleteChildMutation = useDeleteChild();

  const handleSave = async () => {
    if (!name || !birthYear || !birthMonth || !schoolId) {
      Alert.alert('Error', 'Please fill in all required fields (Name, DOB, School ID)');
      return;
    }

    if (isNaN(Number(birthYear)) || birthYear.length !== 4) {
      Alert.alert('Error', 'Please enter a valid 4-digit year');
      return;
    }

    if (isNaN(Number(birthMonth)) || Number(birthMonth) < 1 || Number(birthMonth) > 12) {
      Alert.alert('Error', 'Please enter a valid month (1-12)');
      return;
    }

    const interestsArray = interests.split(',').map(s => s.trim()).filter(s => s.length > 0);

    setIsSubmitting(true);
    try {
      const childData = {
        name,
        birth_year: parseInt(birthYear, 10),
        birth_month: parseInt(birthMonth, 10),
        guardian_id: GUARDIAN_ID,
        school_id: schoolId,
        interests: interestsArray
      };

      if (isEditing) {
        await updateChildMutation.mutateAsync({
          id: params.id as string,
          data: childData
        });
      } else {
        await createChildMutation.mutateAsync({
          data: childData
        });
      }
      router.back();
    } catch (error) {
      console.error(error);
      Alert.alert('Error', 'Failed to save information. Please try again.');
    } finally {
      setIsSubmitting(false);
    }
  };

  const handleDelete = async () => {
    Alert.alert(
      'Delete Child',
      'Are you sure you want to delete this family member? This action cannot be undone.',
      [
        { text: 'Cancel', style: 'cancel' },
        { 
          text: 'Delete', 
          style: 'destructive',
          onPress: async () => {
            setIsSubmitting(true);
            try {
              await deleteChildMutation.mutateAsync({ id: params.id as string });
              router.back();
            } catch (error) {
              console.error(error);
              Alert.alert('Error', 'Failed to delete. Please try again.');
              setIsSubmitting(false);
            }
          }
        }
      ]
    );
  };

  const inputStyle = [
    styles.input, 
    { 
      color: theme.text, 
      backgroundColor: colorScheme === 'dark' ? '#27272a' : '#F3F4F6',
      borderColor: colorScheme === 'dark' ? '#3f3f46' : '#E5E7EB'
    }
  ];

  return (
    <ThemedView style={styles.container}>
      <Stack.Screen 
        options={{ 
          title: isEditing ? 'Edit Child' : 'Add Child',
          headerRight: () => isEditing && (
            <TouchableOpacity onPress={handleDelete}>
              <ThemedText style={{ color: '#EF4444', fontFamily: 'Archivo_600SemiBold' }}>Delete</ThemedText>
            </TouchableOpacity>
          )
        }} 
      />
      
      <ScrollView contentContainerStyle={styles.scrollContent}>
        <View style={styles.form}>
          <View style={styles.inputGroup}>
            <ThemedText style={styles.label}>Full Name</ThemedText>
            <TextInput
              style={inputStyle}
              value={name}
              onChangeText={setName}
              placeholder="e.g. Sarah Smith"
              placeholderTextColor="#9CA3AF"
            />
          </View>

          <View style={styles.row}>
            <View style={[styles.inputGroup, { flex: 1 }]}>
              <ThemedText style={styles.label}>Birth Month</ThemedText>
              <TextInput
                style={inputStyle}
                value={birthMonth}
                onChangeText={setBirthMonth}
                placeholder="1-12"
                placeholderTextColor="#9CA3AF"
                keyboardType="number-pad"
                maxLength={2}
              />
            </View>
            <View style={{ width: 10 }} />
            <View style={[styles.inputGroup, { flex: 1 }]}>
              <ThemedText style={styles.label}>Birth Year</ThemedText>
              <TextInput
                style={inputStyle}
                value={birthYear}
                onChangeText={setBirthYear}
                placeholder="e.g. 2015"
                placeholderTextColor="#9CA3AF"
                keyboardType="number-pad"
                maxLength={4}
              />
            </View>
          </View>

          <View style={styles.inputGroup}>
            <ThemedText style={styles.label}>School ID</ThemedText>
            <TextInput
              style={inputStyle}
              value={schoolId}
              onChangeText={setSchoolId}
              placeholder="Enter School UUID"
              placeholderTextColor="#9CA3AF"
              autoCapitalize="none"
              autoCorrect={false}
            />
          </View>

          <View style={styles.inputGroup}>
            <ThemedText style={styles.label}>Interests</ThemedText>
            <TextInput
              style={inputStyle}
              value={interests}
              onChangeText={setInterests}
              placeholder="soccer, music, coding"
              placeholderTextColor="#9CA3AF"
            />
            <ThemedText style={styles.helperText}>Comma separated</ThemedText>
          </View>

          <TouchableOpacity 
            style={[styles.saveButton, { backgroundColor: theme.tint, opacity: isSubmitting ? 0.7 : 1 }]}
            onPress={handleSave}
            disabled={isSubmitting}
          >
            <ThemedText style={styles.saveButtonText}>
              {isSubmitting ? 'Saving...' : 'Save Information'}
            </ThemedText>
          </TouchableOpacity>
        </View>
      </ScrollView>
    </ThemedView>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
  },
  scrollContent: {
    paddingBottom: 40,
  },
  form: {
    padding: 20,
    gap: 20,
  },
  inputGroup: {
    gap: 8,
  },
  row: {
    flexDirection: 'row',
  },
  label: {
    fontSize: 16,
    fontFamily: 'Archivo_600SemiBold',
  },
  input: {
    borderWidth: 1,
    borderRadius: 12,
    padding: 16,
    fontSize: 16,
    fontFamily: 'Archivo_400Regular',
  },
  helperText: {
    fontSize: 12,
    color: '#6B7280',
    fontFamily: 'Archivo_400Regular',
  },
  saveButton: {
    marginTop: 20,
    padding: 16,
    borderRadius: 12,
    alignItems: 'center',
  },
  saveButtonText: {
    color: '#FFF',
    fontSize: 16,
    fontFamily: 'Archivo_600SemiBold',
  },
});