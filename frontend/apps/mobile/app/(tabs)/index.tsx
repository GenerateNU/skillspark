import { Image } from "expo-image";
import {
  ActivityIndicator,
  FlatList,
  View,
  TextInput,
  ScrollView,
  Pressable,
  Text,
} from "react-native";
import { ThemedText } from "@/components/themed-text";
import { ThemedView } from "@/components/themed-view";
import { useGetAllEventOccurrences } from "@skillspark/api-client";
import type { EventOccurrence } from "@skillspark/api-client";
import { useState } from "react";


// ── Helpers ───────────────────────────────────────────────────────────────────

function formatDuration(start: string, end: string) {
  const mins = Math.round(
    (new Date(end).getTime() - new Date(start).getTime()) / 60000
  );
  return mins >= 60 ? `${Math.round(mins / 60)} hr` : `${mins} min`;
}

function getUniqueCategories(events: EventOccurrence[]): string[] {
  const cats = new Set<string>();
  events.forEach((e) => e.event.category?.forEach((c) => cats.add(c)));
  return cats.size > 0 ? Array.from(cats) : ["Sport", "Arts", "Music", "Tech"];
}

// ── Stars ─────────────────────────────────────────────────────────────────────

function StarRating({ rating = 4 }: { rating?: number }) {
  const full = Math.round(rating);
  return (
    <Text style={{ fontSize: 13, letterSpacing: 1, marginTop: 2 }}>
      <Text style={{ color: "#FBBF24" }}>{"★".repeat(full)}</Text>
      <Text style={{ color: "#D1D5DB" }}>{"★".repeat(5 - full)}</Text>
    </Text>
  );
}

// ── Filter Chips ──────────────────────────────────────────────────────────────

function FilterChips({
  filters,
  active,
  onToggle,
}: {
  filters: string[];
  active: string[];
  onToggle: (f: string) => void;
}) {
  return (
    <ScrollView
      horizontal
      showsHorizontalScrollIndicator={false}
      contentContainerStyle={{ paddingHorizontal: 20, gap: 8, paddingVertical: 2 }}
    >
      {filters.map((f) => {
        const isActive = active.includes(f);
        return (
          <Pressable
            key={f}
            onPress={() => onToggle(f)}
            style={{
              flexDirection: "row",
              alignItems: "center",
              paddingHorizontal: 14,
              paddingVertical: 6,
              borderRadius: 999,
              borderWidth: 1.5,
              borderColor: isActive ? "#111" : "#D1D5DB",
              backgroundColor: isActive ? "#111" : "#fff",
            }}
          >
            <Text style={{ fontSize: 13, fontWeight: "500", color: isActive ? "#fff" : "#374151" }}>
              {isActive ? `× ${f}` : f}
            </Text>
          </Pressable>
        );
      })}
      <Pressable
        style={{
          width: 32, height: 32, borderRadius: 999,
          borderWidth: 1.5, borderColor: "#D1D5DB",
          alignItems: "center", justifyContent: "center",
        }}
      >
        <Text style={{ fontSize: 18, color: "#9CA3AF", lineHeight: 22 }}>+</Text>
      </Pressable>
    </ScrollView>
  );
}

// ── Discover Banner ───────────────────────────────────────────────────────────

function DiscoverBanner({ event }: { event: EventOccurrence }) {
  const ageLabel = event.event.age_range_min != null
    ? `${event.event.age_range_min}${event.event.age_range_max != null ? `–${event.event.age_range_max}` : ""}+`
    : null;

  return (
    <View style={{ marginHorizontal: 20, borderRadius: 24, overflow: "hidden", height: 180, backgroundColor: "#111" }}>
      {event.event.presigned_url ? (
        <Image
          source={{ uri: event.event.presigned_url }}
          style={{ position: "absolute", top: 0, left: 0, right: 0, bottom: 0, opacity: 0.5 }}
          contentFit="cover"
        />
      ) : (
        <>
          <View style={{ position: "absolute", width: 140, height: 140, borderRadius: 70, backgroundColor: "#7C3AED", top: -20, left: 20, opacity: 0.95 }} />
          <View style={{ position: "absolute", width: 120, height: 120, borderRadius: 60, backgroundColor: "#2563EB", top: 10, left: 90, opacity: 0.95 }} />
          <View style={{ position: "absolute", width: 100, height: 100, borderRadius: 50, backgroundColor: "#059669", top: -5, left: 170, opacity: 0.95 }} />
          {/* White card */}
          <View style={{
            position: "absolute", width: 88, height: 108, backgroundColor: "#fff",
            borderRadius: 16, top: "50%", left: "50%",
            transform: [{ translateX: -44 }, { translateY: -54 }],
            alignItems: "center", justifyContent: "center", gap: 6, padding: 10,
            shadowColor: "#000", shadowOpacity: 0.25, shadowRadius: 12,
          }}>
            <View style={{ width: 36, height: 36, borderRadius: 18, backgroundColor: "#A7F3D0" }} />
            <View style={{ width: 52, height: 7, borderRadius: 4, backgroundColor: "#E5E7EB" }} />
            <View style={{ width: 38, height: 7, borderRadius: 4, backgroundColor: "#F3F4F6" }} />
          </View>
        </>
      )}
      {/* Price tag */}
      <View style={{
        position: "absolute", top: 18, right: 18,
        backgroundColor: "#fff", borderRadius: 10,
        paddingHorizontal: 11, paddingVertical: 6,
        transform: [{ rotate: "12deg" }],
        shadowColor: "#000", shadowOpacity: 0.15, shadowRadius: 6,
      }}>
        <Text style={{ fontWeight: "700", color: "#111", fontSize: 13 }}>
          {ageLabel ?? event.event.title.slice(0, 6)}
        </Text>
      </View>
    </View>
  );
}

// ── Event Card ────────────────────────────────────────────────────────────────

function EventCard({ item }: { item: EventOccurrence }) {
  const duration = formatDuration(item.start_time, item.end_time);
  const ageLabel = item.event.age_range_min != null
    ? `${item.event.age_range_min}${item.event.age_range_max != null ? `–${item.event.age_range_max}` : ""}+`
    : null

  return (
    <View style={{ marginHorizontal: 20, marginBottom: 20 }}>
      {/* Photo */}
      <View style={{ borderRadius: 20, overflow: "hidden", height: 185 }}>
        {item.event.presigned_url ? (
          <Image
            source={{ uri: item.event.presigned_url }}
            style={{ width: "100%", height: "100%" }}
            contentFit="cover"
          />
        ) : (
          <View style={{ width: "100%", height: "100%", backgroundColor: "#E5E7EB" }} />
        )}

        {/* Pill overlay */}
        <View style={{
          position: "absolute", bottom: 12, left: 12,
          flexDirection: "row", alignItems: "center",
          backgroundColor: "rgba(255,255,255,0.92)",
          borderRadius: 999, paddingHorizontal: 12, paddingVertical: 5,
        }}>
          {ageLabel && (
            <>
              <Text style={{ fontSize: 12, color: "#374151", fontWeight: "500" }}>🧑 {ageLabel}</Text>
              <View style={{ width: 1, height: 14, backgroundColor: "#D1D5DB", marginHorizontal: 10 }} />
            </>
          )}
          <Text style={{ fontSize: 12, color: "#374151", fontWeight: "500" }}>🚌 {duration}</Text>
        </View>
      </View>

      {/* Below image row */}
      <View style={{ flexDirection: "row", alignItems: "center", justifyContent: "space-between", marginTop: 10, paddingHorizontal: 4 }}>
        <View style={{ flex: 1, marginRight: 16 }}>
          <Text style={{ fontSize: 16, fontWeight: "600", color: "#111" }} numberOfLines={1}>
            {item.event.title}
          </Text>
          <StarRating />
        </View>
        <View style={{ backgroundColor: "#111", borderRadius: 999, paddingHorizontal: 20, paddingVertical: 10 }}>
          <Text style={{ color: "#fff", fontWeight: "700", fontSize: 14 }}>
            {item.curr_enrolled}/{item.max_attendees}
          </Text>
        </View>
      </View>
    </View>
  );
}

// ── Main Feed ─────────────────────────────────────────────────────────────────

function EventOccurrencesList() {
  const { data: response, isLoading, error } = useGetAllEventOccurrences();
  const [activeFilters, setActiveFilters] = useState<string[]>([]);
  const [search, setSearch] = useState("");

  console.log("response:", JSON.stringify(response));

  const toggleFilter = (f: string) =>
    setActiveFilters((prev) =>
      prev.includes(f) ? prev.filter((x) => x !== f) : [...prev, f]
    );

  if (isLoading) {
    return (
      <View style={{ flex: 1, alignItems: "center", justifyContent: "center", gap: 8 }}>
        <ActivityIndicator size="large" />
        <ThemedText>Loading events...</ThemedText>
      </View>
    );
  }

  if (error) {
    return (
      <View style={{ flex: 1, alignItems: "center", justifyContent: "center", padding: 16 }}>
        <ThemedText style={{ color: "#EF4444", fontWeight: "600" }}>Error loading events</ThemedText>
        <ThemedText>{error.message || "An error occurred"}</ThemedText>
      </View>
    );
  }

  if (!response || !Array.isArray(response.data)) {
    return (
      <View style={{ flex: 1, alignItems: "center", justifyContent: "center", padding: 16 }}>
        <ThemedText>No events available</ThemedText>
      </View>
    );
  }

  const upcomingEvents = response.data
    .filter((o) => new Date(o.start_time) >= new Date())
    .sort((a, b) => new Date(a.start_time).getTime() - new Date(b.start_time).getTime());

  const allCategories = getUniqueCategories(response.data);
  const featuredEvent = upcomingEvents[0];
  const listEvents = upcomingEvents.slice(1);

  return (
    <FlatList
      data={listEvents}
      keyExtractor={(item) => item.id}
      showsVerticalScrollIndicator={false}
      contentContainerStyle={{ paddingBottom: 40 }}
      ListHeaderComponent={
        <View style={{ backgroundColor: "#fff" }}>
          {/* Nav bar */}
          <View style={{ flexDirection: "row", alignItems: "center", justifyContent: "space-between", paddingHorizontal: 20, paddingTop: 56, paddingBottom: 20 }}>
            <View style={{ width: 36, height: 36, alignItems: "center", justifyContent: "center" }}>
              <Text style={{ fontSize: 22 }}>⚡</Text>
            </View>
            <View style={{ gap: 5, padding: 4 }}>
              <View style={{ width: 22, height: 2, backgroundColor: "#111", borderRadius: 2 }} />
              <View style={{ width: 22, height: 2, backgroundColor: "#111", borderRadius: 2 }} />
            </View>
          </View>

          {/* Title */}
          <View style={{ paddingHorizontal: 20, marginBottom: 16 }}>
            <Text style={{ fontSize: 28, fontWeight: "700", color: "#111", letterSpacing: -0.5 }}>
              My Dashboard
            </Text>
          </View>

          {/* Filter chips */}
          <FilterChips filters={allCategories} active={activeFilters} onToggle={toggleFilter} />

          {/* Search */}
          <View style={{
            marginHorizontal: 20, marginTop: 14, marginBottom: 20,
            flexDirection: "row", alignItems: "center",
            backgroundColor: "#F3F4F6", borderRadius: 999,
            paddingHorizontal: 18, paddingVertical: 11,
          }}>
            <Text style={{ fontSize: 14, marginRight: 10, color: "#9CA3AF" }}>🔍</Text>
            <TextInput
              style={{ flex: 1, fontSize: 14, color: "#111" }}
              placeholder="Search for a class"
              placeholderTextColor="#9CA3AF"
              value={search}
              onChangeText={setSearch}
            />
          </View>

          {/* Discover Weekly */}
          <View style={{ flexDirection: "row", alignItems: "center", paddingHorizontal: 20, marginBottom: 12 }}>
            <Text style={{ color: "#7C3AED", fontSize: 13, marginRight: 6 }}>✦</Text>
            <Text style={{ fontSize: 15, fontWeight: "600", color: "#111" }}>Discover Weekly</Text>
          </View>
          {featuredEvent && <DiscoverBanner event={featuredEvent} />}

          {/* For You */}
          {listEvents.length > 0 && (
            <View style={{ paddingHorizontal: 20, marginTop: 24, marginBottom: 4 }}>
              <View style={{ flexDirection: "row", alignItems: "center", gap: 6, marginBottom: 4 }}>
                <View style={{ width: 26, height: 26, borderRadius: 13, backgroundColor: "#3B82F6", alignItems: "center", justifyContent: "center" }}>
                  <Text style={{ color: "#fff", fontSize: 11, fontWeight: "700" }}>A</Text>
                </View>
                <Text style={{ fontSize: 14, fontWeight: "600", color: "#111" }}>For You</Text>
                <View style={{ flexDirection: "row", marginLeft: 4 }}>
                  {["#10B981", "#6366F1"].map((c, i) => (
                    <View key={i} style={{
                      width: 22, height: 22, borderRadius: 11,
                      backgroundColor: c,
                      borderWidth: 2, borderColor: "#fff",
                      marginLeft: i > 0 ? -8 : 0,
                    }} />
                  ))}
                </View>
              </View>
              <View style={{ flexDirection: "row", alignItems: "center", gap: 4 }}>
                <Text style={{ color: "#7C3AED", fontSize: 12 }}>✦</Text>
                <Text style={{ fontSize: 13, color: "#6B7280" }}>Based on </Text>
                <Text style={{ fontSize: 13, color: "#3B82F6" }}>upcoming events</Text>
              </View>
            </View>
          )}

          <View style={{ height: 16 }} />
        </View>
      }
      renderItem={({ item }) => <EventCard item={item} />}
      ListEmptyComponent={
        <View style={{ alignItems: "center", padding: 32 }}>
          <ThemedText style={{ color: "#9CA3AF" }}>No upcoming events</ThemedText>
        </View>
      }
    />
  );
}

// ── Screen ────────────────────────────────────────────────────────────────────

export default function HomeScreen() {
  return (
    <View style={{ flex: 1, backgroundColor: "#fff" }}>
      <EventOccurrencesList />
    </View>
  );
}