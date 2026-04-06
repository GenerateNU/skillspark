import { Image } from "expo-image";
import {
  ActivityIndicator,
  ScrollView,
  Text,
  TouchableOpacity,
  View,
} from "react-native";
import { useSafeAreaInsets } from "react-native-safe-area-context";
import { useLocalSearchParams, useRouter } from "expo-router";
import { useGetEventOccurrencesById } from "@skillspark/api-client";
import type { EventOccurrence } from "@skillspark/api-client";
import MaterialIcons from "@expo/vector-icons/MaterialIcons";
import { AboutPage } from "@/components/AboutPage";
import { formatDuration } from "@/utils/format";

function formatAddress(occurrence: EventOccurrence) {
  const loc = occurrence.location;
  const parts = [loc.address_line1, loc.district].filter(Boolean);
  return parts.join(", ") || "Location";
}

function EventOccurrenceDetail({ occurrence }: { occurrence: EventOccurrence }) {
  const router = useRouter();
  const insets = useSafeAreaInsets();
  const duration = formatDuration(occurrence.start_time, occurrence.end_time);
  const address = formatAddress(occurrence);

  return (
    <View className="flex-1 bg-gray-100">
      <ScrollView
        showsVerticalScrollIndicator={false}
        contentContainerStyle={{ paddingBottom: 120 }}
      >
        {/* Header image */}
        <View className="h-64 bg-gray-900">
          {occurrence.event.presigned_url ? (
            <Image
              source={{ uri: occurrence.event.presigned_url }}
              style={{ width: "100%", height: "100%" }}
              contentFit="cover"
            />
          ) : (
            <View className="flex-1 bg-gray-300" />
          )}
          <TouchableOpacity
            onPress={() => router.back()}
            activeOpacity={0.7}
            style={{ top: insets.top + 8 }}
            className="absolute left-4 z-10 flex-row items-center bg-white rounded-full px-4 py-2.5 shadow-md"
          >
            <MaterialIcons name="chevron-left" size={20} color="#1a1a1a" />
            <Text className="text-base text-gray-900 font-medium">Back</Text>
          </TouchableOpacity>
        </View>

        {/* Content card */}
        <View className="bg-white rounded-t-3xl -mt-7 px-6 pt-6 pb-6">
          {/* Title */}
          <Text className="text-2xl font-bold text-gray-900 tracking-tight mb-1.5">
            {occurrence.event.title}
          </Text>

          {/* Location */}
          <View className="flex-row items-center gap-1.5 mb-5">
            <MaterialIcons name="location-on" size={16} color="#6b7280" />
            <Text className="text-sm text-gray-500">{address}</Text>
          </View>

          {/* Duration + price */}
          <View className="flex-row justify-between items-center mb-5">
            <Text className="text-2xl font-bold text-gray-900">{duration}</Text>
            <View className="items-end">
              <Text className="text-xl font-bold text-gray-900">{occurrence.price} THB</Text>
              <Text className="text-xs text-gray-400">/Session</Text>
            </View>
          </View>

          {/* Divider */}
          <View className="border-b border-dashed border-gray-200 mb-5" />

          {/* About section */}
          <AboutPage
            description={occurrence.event.description}
            links={occurrence.org_links ?? []}
          />
        </View>
      </ScrollView>

      {/* Register button pinned at bottom */}
      <View
        className="absolute bottom-0 left-0 right-0 bg-white px-6 pt-3 border-t border-gray-100"
        style={{ paddingBottom: insets.bottom + 12 }}
      >
        <TouchableOpacity
          onPress={() => {}}
          activeOpacity={0.7}
          className="bg-gray-900 rounded-2xl py-4 items-center"
        >
          <Text className="text-white text-lg font-bold">Register</Text>
        </TouchableOpacity>
      </View>
    </View>
  );
}

export default function EventOccurrenceScreen() {
  const { id } = useLocalSearchParams<{ id: string }>();
  const { data: response, isLoading, error } = useGetEventOccurrencesById(id);

  if (isLoading) {
    return (
      <View className="flex-1 items-center justify-center">
        <ActivityIndicator size="large" />
      </View>
    );
  }

  if (error || !response || response.status !== 200) {
    return (
      <View className="flex-1 items-center justify-center p-6">
        <Text className="text-red-500 font-semibold text-base">
          Event not found
        </Text>
      </View>
    );
  }

  return <EventOccurrenceDetail occurrence={response.data} />;
}