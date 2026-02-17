import React from 'react';
import { View, Image, ScrollView, TouchableOpacity, ActivityIndicator } from 'react-native';
import { ThemedText } from '@/components/themed-text';
import { ThemedView } from '@/components/themed-view';
import { IconSymbol } from '@/components/ui/icon-symbol';
import { SafeAreaView } from 'react-native-safe-area-context';
import { useGetUser, useGetChildrenByGuardianId } from '@skillspark/api-client';

const HARDCODED_USER_ID = "123e4567-e89b-12d3-a456-426614174000"; 

export default function ProfileScreen() {
  const { data: userResponse, isLoading: userLoading } = useGetUser(HARDCODED_USER_ID);
  const { data: childrenResponse, isLoading: familyLoading } = useGetChildrenByGuardianId(HARDCODED_USER_ID);

  const user = userResponse?.status === 200 ? userResponse.data : null;
  const children = childrenResponse?.status === 200 ? childrenResponse.data : [];

  if (userLoading || familyLoading) {
    return (
      <ThemedView className="flex-1 justify-center items-center">
        <ActivityIndicator size="large" />
      </ThemedView>
    );
  }

  return (
    <ThemedView className="flex-1">
      <SafeAreaView className="flex-1">
        <ScrollView contentContainerStyle={{ paddingBottom: 40 }}>
          <View className="items-center py-8 bg-neutral-100 dark:bg-neutral-900 mb-6">
            <View className="h-24 w-24 rounded-full bg-gray-300 overflow-hidden mb-4 border-4 border-white dark:border-neutral-800 shadow-sm">
               <Image 
                 source={{ uri: 'https://i.pravatar.cc/300' }} 
                 className="w-full h-full"
               />
            </View>
            <ThemedText type="title" className="text-2xl font-bold">
              {user?.full_name}
            </ThemedText>
            <ThemedText className="text-gray-500 mt-1">
              {user?.email}
            </ThemedText>
            <TouchableOpacity className="mt-4 px-6 py-2 bg-blue-600 rounded-full">
               <ThemedText className="text-white font-semibold">Edit Profile</ThemedText>
            </TouchableOpacity>
          </View>
          <View className="px-4 mb-6">
            <View className="flex-row justify-between items-center mb-4">
              <ThemedText type="subtitle" className="font-semibold">My Family</ThemedText>
              <TouchableOpacity>
                 <ThemedText className="text-blue-600">Manage</ThemedText>
              </TouchableOpacity>
            </View>
            <View className="bg-white dark:bg-neutral-800 rounded-xl overflow-hidden shadow-sm">
              {children.length === 0 ? (
                <View className="p-4 items-center">
                  <ThemedText className="text-gray-400">No family members added yet.</ThemedText>
                </View>
              ) : (
                children.map((child, index) => {
                  const currentYear = new Date().getFullYear();
                  const age = child.birth_year ? currentYear - child.birth_year : '?';

                  return (
                    <View key={child.id} className={`p-4 flex-row items-center justify-between ${index !== 0 ? 'border-t border-gray-100 dark:border-neutral-700' : ''}`}>
                      <View className="flex-row items-center gap-3">
                        <View className="h-10 w-10 rounded-full bg-blue-100 items-center justify-center">
                           <IconSymbol name="person" size={20} color="#2563EB" />
                        </View>
                        <View>
                          <ThemedText className="font-medium">{child.name}</ThemedText>
                          <ThemedText className="text-xs text-gray-500">{age} years old</ThemedText>
                        </View>
                      </View>
                      <IconSymbol name="chevron.right" size={20} color="#9CA3AF" />
                    </View>
                  );
                })
              )}
            </View>
            
            <TouchableOpacity className="mt-4 flex-row items-center justify-center p-3 border border-dashed border-gray-400 rounded-xl">
               <IconSymbol name="plus" size={20} color="#6B7280" />
               <ThemedText className="ml-2 text-gray-600">Add Family Member</ThemedText>
            </TouchableOpacity>
          </View>
          <View className="px-4">
             <ThemedText type="subtitle" className="mb-4 font-semibold">Settings</ThemedText>
             <View className="bg-white dark:bg-neutral-800 rounded-xl overflow-hidden shadow-sm">
                {['Account Security', 'Notifications', 'Help & Support', 'Log Out'].map((item, index) => (
                   <TouchableOpacity key={item} className={`p-4 flex-row justify-between items-center ${index !== 3 ? 'border-b border-gray-100 dark:border-neutral-700' : ''}`}>
                      <ThemedText className={item === 'Log Out' ? 'text-red-500' : ''}>{item}</ThemedText>
                      <IconSymbol name="chevron.right" size={18} color="#9CA3AF" />
                   </TouchableOpacity>
                ))}
             </View>
          </View>
        </ScrollView>
      </SafeAreaView>
    </ThemedView>
  );
}