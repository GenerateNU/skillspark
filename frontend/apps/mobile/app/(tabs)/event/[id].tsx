import { Image } from "expo-image";
import {
  ActivityIndicator,
  Pressable,
  ScrollView,
  Text,
  View,
} from "react-native";
import { useSafeAreaInsets } from "react-native-safe-area-context";
import { useLocalSearchParams, useRouter } from "expo-router";
import { useGetEventOccurrencesById } from "@skillspark/api-client";
import type { EventOccurrence } from "@skillspark/api-client";
import MaterialIcons from "@expo/vector-icons/MaterialIcons";
import { AppColors } from "@/constants/theme";
import { StarRating } from "@/components/StarRating";
import { formatDuration } from "@/utils/format";

function formatAddress(occurrence: EventOccurrence) {
  const loc = occurrence.location;
  const parts = [loc.address_line1, loc.district].filter(Boolean);
  return parts.join(", ") || "Location";
}

function BookmarkIcon() {
  return (
    <View style={{ alignItems: "center" }}>
      <View style={{ width: 22, height: 26, backgroundColor: AppColors.primaryText }} />
      <View
        style={{
          width: 0,
          height: 0,
          borderLeftWidth: 11,
          borderRightWidth: 11,
          borderTopWidth: 9,
          borderLeftColor: "transparent",
          borderRightColor: "transparent",
          borderTopColor: AppColors.primaryText,
        }}
      />
    </View>
  );
}

function EventOccurrenceDetail({ occurrence }: { occurrence: EventOccurrence }) {
  const router = useRouter();
  const insets = useSafeAreaInsets();
  const duration = formatDuration(occurrence.start_time, occurrence.end_time);
  const address = formatAddress(occurrence);

  return (
    <View style={{ flex: 1, backgroundColor: "#F4F6F8" }}>
      <View
        style={{
          flex: 1,
          marginTop: insets.top + 8,
          marginHorizontal: 14,
          marginBottom: insets.bottom - 20,
          borderRadius: 32,
          overflow: "hidden",
          backgroundColor: "#fff",
          shadowColor: "#000",
          shadowOpacity: 0.12,
          shadowRadius: 20,
          elevation: 8,
        }}
      >
      <ScrollView
        showsVerticalScrollIndicator={false}
        bounces={false}
        contentContainerStyle={{ flexGrow: 1 }}
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
          <Pressable
            onPress={() => router.navigate('/')}
            style={{
              position: "absolute",
              top: 16,
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
          </Pressable>
        </View>
        <View
          style={{
            backgroundColor: "#fff",
            borderTopLeftRadius: 28,
            borderTopRightRadius: 28,
            marginTop: -28,
            paddingHorizontal: 22,
            paddingBottom: 24,
            shadowColor: "#000",
            shadowOpacity: 0.06,
            shadowRadius: 12,
            elevation: 2,
          }}
        >
          <View
            style={{
              width: 38,
              height: 4,
              backgroundColor: AppColors.borderLight,
              borderRadius: 2,
              alignSelf: "center",
              marginTop: 12,
              marginBottom: 14,
            }}
          />
          <View style={{ position: "absolute", left: 20 }}>
            <BookmarkIcon />
          </View>
          <Text
            style={{
              fontSize: 28,
              fontWeight: "700",
              color: AppColors.primaryText,
              letterSpacing: -0.5,
              marginBottom: 8,
            }}
          >
            {occurrence.event.title}
          </Text>
          <View
            style={{
              flexDirection: "row",
              alignItems: "center",
              gap: 6,
              marginBottom: 14,
            }}
          >
            <MaterialIcons name="location-on" size={16} color={AppColors.primaryText} />
            <Text
              style={{ fontSize: 13, color: AppColors.secondaryText, flex: 1 }}
              numberOfLines={1}
            >
              {address}
            </Text>
            <StarRating size={17} />
          </View>
          <Text
            style={{
              fontSize: 14,
              color: AppColors.secondaryText,
              lineHeight: 22,
              marginBottom: 18,
            }}
          >
            {occurrence.event.description}
          </Text>
          <View
            style={{
              flexDirection: "row",
              alignItems: "center",
              justifyContent: "space-between",
            }}
          >
            <View style={{ flexDirection: "row", gap: 8, flex: 1, flexWrap: "wrap" }}>
              {occurrence.event.category?.map((cat) => (
                <View
                  key={cat}
                  style={{
                    borderWidth: 1.5,
                    borderColor: AppColors.borderLight,
                    borderRadius: 999,
                    paddingHorizontal: 16,
                    paddingVertical: 7,
                  }}
                >
                  <Text style={{ fontSize: 13, color: AppColors.secondaryText }}>{cat}</Text>
                </View>
              ))}
            </View>
            <View style={{ alignItems: "flex-end", marginLeft: 14 }}>
              <Text style={{ fontSize: 20, fontWeight: "700", color: AppColors.primaryText }}>$40.00</Text>
              <Text style={{ fontSize: 12, color: AppColors.subtleText }}>/Session</Text>
            </View>
          </View>
        </View>
        <View
          style={{
            marginHorizontal: 0,
            borderBottomWidth: 1,
            borderStyle: "dashed",
            borderColor: AppColors.borderLight,
          }}
        />
        <View
          style={{
            flex: 1,
            backgroundColor: "#fff",
            paddingHorizontal: 22,
            paddingTop: 22,
            paddingBottom: 28,
            shadowColor: "#000",
            shadowOpacity: 0.06,
            shadowRadius: 12,
            elevation: 2,
          }}
        >
          <View
            style={{
              flexDirection: "row",
              alignItems: "center",
              justifyContent: "space-between",
              marginBottom: 22,
            }}
          >
            <Text
              style={{
                fontSize: 30,
                fontWeight: "700",
                color: AppColors.primaryText,
                letterSpacing: -0.5,
              }}
            >
              {duration}
            </Text>
            <View
              style={{
                flexDirection: "row",
                alignItems: "center",
                gap: 5,
                backgroundColor: "#F3F4F6",
                borderRadius: 999,
                paddingHorizontal: 12,
                paddingVertical: 7,
              }}
            >
              <MaterialIcons name="directions-walk" size={14} color={AppColors.secondaryText} />
              <Text style={{ fontSize: 13, color: AppColors.secondaryText }}>8 min</Text>
              <MaterialIcons name="arrow-forward" size={12} color={AppColors.subtleText} />
              <MaterialIcons name="directions-bus" size={16} color={AppColors.secondaryText} />
            </View>
          </View>
          <View style={{ flexDirection: "row", alignItems: "center", justifyContent: "space-between" }}>
            <View style={{ gap: 0 }}>
              <View style={{ flexDirection: "row", alignItems: "center", gap: 12 }}>
                <View
                  style={{
                    width: 16,
                    height: 16,
                    borderRadius: 8,
                    backgroundColor: AppColors.primaryText,
                    alignItems: "center",
                    justifyContent: "center",
                  }}
                >
                  <View style={{ width: 6, height: 6, borderRadius: 3, backgroundColor: "#fff" }} />
                </View>
                <Text style={{ fontSize: 14, color: AppColors.secondaryText, fontWeight: "500" }}>Home</Text>
              </View>
              <View style={{ paddingLeft: 6, paddingVertical: 2 }}>
                <Text style={{ fontSize: 14, color: AppColors.subtleText, lineHeight: 10 }}>•</Text>
                <Text style={{ fontSize: 14, color: AppColors.subtleText, lineHeight: 10 }}>•</Text>
                <Text style={{ fontSize: 14, color: AppColors.subtleText, lineHeight: 10 }}>•</Text>
              </View>
              <View style={{ flexDirection: "row", alignItems: "center", gap: 10 }}>
                <MaterialIcons name="location-on" size={16} color={AppColors.secondaryText} />
                <Text style={{ fontSize: 14, color: AppColors.secondaryText, fontWeight: "500" }}>Location</Text>
              </View>
            </View>
            <Pressable
              onPress={() => {}}
              style={{
                backgroundColor: AppColors.primaryText,
                borderRadius: 16,
                paddingHorizontal: 26,
                paddingVertical: 14,
              }}
            >
              <Text style={{ color: "#fff", fontSize: 17, fontWeight: "700" }}>Register</Text>
            </Pressable>
          </View>
        </View>
      </ScrollView>
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
