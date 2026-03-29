import { Image } from "expo-image";
import {
  ActivityIndicator,
  Pressable,
  ScrollView,
  Text,
  TouchableOpacity,
  View,
} from "react-native";
import { useState } from "react";
import { useSafeAreaInsets } from "react-native-safe-area-context";
import { useLocalSearchParams, useRouter } from "expo-router";
import { useGetEventOccurrencesById } from "@skillspark/api-client";
import type { EventOccurrence } from "@skillspark/api-client";
import MaterialIcons from "@expo/vector-icons/MaterialIcons";
import { AppColors } from "@/constants/theme";
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
    <View style={{ flex: 1, backgroundColor: "#F4F6F8" }}>
      <ScrollView
        showsVerticalScrollIndicator={false}
        contentContainerStyle={{ paddingBottom: 120 }}
      >
        <View style={{ height: 250, backgroundColor: "#1a1a1a" }}>
          {occurrence.event.presigned_url ? (
            <Image
              source={{ uri: occurrence.event.presigned_url }}
              style={{ width: "100%", height: "100%" }}
              contentFit="cover"
            />
          ) : (
            <View style={{ flex: 1, backgroundColor: "#C5C5C5" }} />
          )}
          <TouchableOpacity
            onPress={() => router.back()}
            activeOpacity={0.7}
            style={{
              position: "absolute",
              top: insets.top + 8,
              left: 16,
              zIndex: 10,
              flexDirection: "row",
              alignItems: "center",
              backgroundColor: "#fff",
              borderRadius: 999,
              paddingHorizontal: 16,
              paddingVertical: 10,
              shadowColor: "#000",
              shadowOpacity: 0.15,
              shadowRadius: 8,
              elevation: 10,
            }}
          >
            <MaterialIcons name="chevron-left" size={20} color={AppColors.primaryText} />
            <Text style={{ fontSize: 15, color: AppColors.primaryText, fontWeight: "500" }}>Back</Text>
          </TouchableOpacity>
        </View>

        <View style={{ backgroundColor: "#fff", borderTopLeftRadius: 28, borderTopRightRadius: 28, marginTop: -28, padding: 22 }}>
          {/* Title */}
          <Text style={{ fontSize: 26, fontWeight: "700", color: AppColors.primaryText, letterSpacing: -0.5, marginBottom: 6 }}>
            {occurrence.event.title}
          </Text>

          <View style={{ flexDirection: "row", alignItems: "center", gap: 6, marginBottom: 20 }}>
            <MaterialIcons name="location-on" size={16} color={AppColors.secondaryText} />
            <Text style={{ fontSize: 13, color: AppColors.secondaryText }}>{address}</Text>
          </View>

          <View style={{ flexDirection: "row", justifyContent: "space-between", alignItems: "center", marginBottom: 20 }}>
            <Text style={{ fontSize: 22, fontWeight: "700", color: AppColors.primaryText }}>{duration}</Text>
            <View style={{ alignItems: "flex-end" }}>
              <Text style={{ fontSize: 20, fontWeight: "700", color: AppColors.primaryText }}>{occurrence.price} THB</Text>
              <Text style={{ fontSize: 12, color: AppColors.subtleText }}>/Session</Text>
            </View>
          </View>

          <View style={{ borderBottomWidth: 1, borderStyle: "dashed", borderColor: AppColors.borderLight, marginBottom: 20 }} />

          <AboutPage
            description={occurrence.event.description}
            links={occurrence.org_links ?? []}
          />
        </View>
      </ScrollView>

      <View style={{
        position: "absolute",
        bottom: 0,
        left: 0,
        right: 0,
        backgroundColor: "#fff",
        paddingHorizontal: 22,
        paddingTop: 12,
        paddingBottom: insets.bottom + 12,
        borderTopWidth: 1,
        borderColor: AppColors.borderLight,
      }}>
        <TouchableOpacity
          onPress={() => {}}
          activeOpacity={0.7}
          style={{
            backgroundColor: AppColors.primaryText,
            borderRadius: 16,
            paddingVertical: 16,
            alignItems: "center",
          }}
        >
          <Text style={{ color: "#fff", fontSize: 17, fontWeight: "700" }}>Register</Text>
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
      <View style={{ flex: 1, alignItems: "center", justifyContent: "center" }}>
        <ActivityIndicator size="large" />
      </View>
    );
  }

  if (error || !response || response.status !== 200) {
    return (
      <View style={{ flex: 1, alignItems: "center", justifyContent: "center", padding: 24 }}>
        <Text style={{ color: AppColors.danger, fontWeight: "600", fontSize: 16 }}>
          Event not found
        </Text>
      </View>
    );
  }

  return <EventOccurrenceDetail occurrence={response.data} />;
}