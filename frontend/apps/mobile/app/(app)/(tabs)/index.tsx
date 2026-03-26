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
import { useGetAllEventOccurrences } from "@skillspark/api-client";
import type { EventOccurrence } from "@skillspark/api-client";
import { useState } from "react";
import { useRouter } from "expo-router";
import { AppColors } from "@/constants/theme";
import { StarRating } from "@/components/StarRating";
import { formatDuration } from "@/utils/format";
import { useTranslation } from "react-i18next";


// ── Helpers ───────────────────────────────────────────────────────────────────

function getUniqueCategories(events: EventOccurrence[]): string[] {
  const cats = new Set<string>();
  events.forEach((e) => e.event.category?.forEach((c) => cats.add(c)));
  return cats.size > 0 ? Array.from(cats) : ["Sport", "Arts", "Music", "Tech"];
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
  const { t: translate } = useTranslation();
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
            className="flex-row items-center px-3.5 py-1.5 rounded-full border-[1.5px]"
            style={{
              borderColor: isActive ? AppColors.primaryText : AppColors.borderLight,
              backgroundColor: isActive ? AppColors.primaryText : "#fff",
            }}
          >
            <Text
              className="text-[13px] font-medium"
              style={{ color: isActive ? "#fff" : AppColors.secondaryText }}
            >
              {isActive ? `× ${translate(`interests.${f}`, { defaultValue: f })}` : translate(`interests.${f}`, { defaultValue: f })}
            </Text>
          </Pressable>
        );
      })}
      <Pressable
        className="w-8 h-8 rounded-full border-[1.5px] items-center justify-center"
        style={{ borderColor: AppColors.borderLight }}
      >
        <Text className="text-lg leading-[22px]" style={{ color: AppColors.subtleText }}>+</Text>
      </Pressable>
    </ScrollView>
  );
}

// ── Discover Banner ───────────────────────────────────────────────────────────

function DiscoverBanner({ event }: { event: EventOccurrence }) {
  const router = useRouter();
  const ageLabel = event.event.age_range_min != null
    ? `${event.event.age_range_min}${event.event.age_range_max != null ? `–${event.event.age_range_max}` : ""}+`
    : null;

  return (
    <Pressable
      onPress={() => router.push(`/event/${event.id}`)}
      className="mx-5 rounded-3xl overflow-hidden h-[180px]"
      style={{ backgroundColor: AppColors.primaryText }}
    >
      {event.event.presigned_url ? (
        <Image
          source={{ uri: event.event.presigned_url }}
          className="absolute inset-0 opacity-50"
          contentFit="cover"
        />
      ) : (
        <>
          <View className="absolute w-[140px] h-[140px] rounded-full opacity-95" style={{ backgroundColor: "#7C3AED", top: -20, left: 20 }} />
          <View className="absolute w-[120px] h-[120px] rounded-full opacity-95" style={{ backgroundColor: AppColors.primaryBlue, top: 10, left: 90 }} />
          <View className="absolute w-[100px] h-[100px] rounded-full opacity-95" style={{ backgroundColor: "#059669", top: -5, left: 170 }} />
          {/* White card */}
          <View className="absolute w-[88px] h-[108px] bg-white rounded-2xl items-center justify-center gap-1.5 p-2.5"
            style={{
              top: "50%", left: "50%",
              transform: [{ translateX: -44 }, { translateY: -54 }],
              shadowColor: "#000", shadowOpacity: 0.25, shadowRadius: 12,
            }}
          >
            <View className="w-9 h-9 rounded-full bg-[#A7F3D0]" />
            <View className="w-[52px] h-[7px] rounded-sm" style={{ backgroundColor: AppColors.divider }} />
            <View className="w-[38px] h-[7px] rounded-sm bg-[#F3F4F6]" />
          </View>
        </>
      )}
      {/* Price tag */}
      <View
        className="absolute top-[18px] right-[18px] bg-white rounded-[10px] px-[11px] py-1.5"
        style={{ transform: [{ rotate: "12deg" }], shadowColor: "#000", shadowOpacity: 0.15, shadowRadius: 6 }}
      >
        <Text className="font-bold text-[13px]" style={{ color: AppColors.primaryText }}>
          {ageLabel ?? event.event.title.slice(0, 6)}
        </Text>
      </View>
    </Pressable>
  );
}

// ── Event Card ────────────────────────────────────────────────────────────────

function EventCard({ item }: { item: EventOccurrence }) {
  const router = useRouter();
  const { t: translate } = useTranslation();
  const duration = formatDuration(item.start_time, item.end_time, {
    hr: translate('event.hr'),
    min: translate('event.min'),
  });
  const ageLabel = item.event.age_range_min != null
    ? `${item.event.age_range_min}${item.event.age_range_max != null ? `–${item.event.age_range_max}` : ""}+`
    : null;

  return (
    <Pressable
      onPress={() => router.push(`/event/${item.id}`)}
      className="mx-5 mb-5"
    >
      {/* Photo */}
      <View className="rounded-[20px] overflow-hidden h-[185px]">
        {item.event.presigned_url ? (
          <Image
            source={{ uri: item.event.presigned_url }}
            className="w-full h-full"
            contentFit="cover"
          />
        ) : (
          <View className="w-full h-full" style={{ backgroundColor: AppColors.divider }} />
        )}

        {/* Pill overlay */}
        <View className="absolute bottom-3 left-3 flex-row items-center bg-white/90 rounded-full px-3 py-[5px]">
          {ageLabel && (
            <>
              <Text className="text-xs font-medium" style={{ color: AppColors.secondaryText }}>🧑 {ageLabel}</Text>
              <View className="w-px h-3.5 mx-2.5" style={{ backgroundColor: AppColors.borderLight }} />
            </>
          )}
          <Text className="text-xs font-medium" style={{ color: AppColors.secondaryText }}>{duration}</Text>
        </View>
      </View>

      {/* Below image row */}
      <View className="flex-row items-center justify-between mt-2.5 px-1">
        <View className="flex-1 mr-4">
          <Text className="text-base font-semibold" style={{ color: AppColors.primaryText }} numberOfLines={1}>
            {item.event.title}
          </Text>
          <StarRating />
        </View>
        <View className="rounded-full px-5 py-2.5" style={{ backgroundColor: AppColors.primaryText }}>
          <Text className="text-white font-bold text-sm">
            {item.curr_enrolled}/{item.max_attendees}
          </Text>
        </View>
      </View>
    </Pressable>
  );
}

// ── Main Feed ─────────────────────────────────────────────────────────────────

function EventOccurrencesList() {
  const { data: response, isLoading, error } = useGetAllEventOccurrences();
  const [activeFilters, setActiveFilters] = useState<string[]>([]);
  const [search, setSearch] = useState("");
  const { t: translate } = useTranslation();

  const toggleFilter = (f: string) =>
    setActiveFilters((prev) =>
      prev.includes(f) ? prev.filter((x) => x !== f) : [...prev, f]
    );

  if (isLoading) {
    return (
      <View className="flex-1 items-center justify-center gap-2">
        <ActivityIndicator size="large" />
        <ThemedText>{translate('common.loadingEvents')}</ThemedText>
      </View>
    );
  }

  if (error) {
    return (
      <View className="flex-1 items-center justify-center p-4">
        <ThemedText className="font-semibold" style={{ color: AppColors.danger }}>{translate('common.errorLoadingEvents')}</ThemedText>
        <ThemedText>{error.detail || translate('common.errorOccurred')}</ThemedText>
      </View>
    );
  }

  if (!response || !Array.isArray(response.data)) {
    return (
      <View className="flex-1 items-center justify-center p-4">
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
        <View className="bg-white">
          {/* Nav bar */}
          <View className="flex-row items-center justify-between px-5 pt-14 pb-5">
            <View className="w-9 h-9 items-center justify-center">
              <Text className="text-[22px]">⚡</Text>
            </View>
            <View className="gap-[5px] p-1">
              <View className="w-[22px] h-0.5 rounded-sm" style={{ backgroundColor: AppColors.primaryText }} />
              <View className="w-[22px] h-0.5 rounded-sm" style={{ backgroundColor: AppColors.primaryText }} />
            </View>
          </View>

          {/* Title */}
          <View className="px-5 mb-4">
            <Text className="text-[28px] font-bold tracking-tight" style={{ color: AppColors.primaryText }}>
              {translate('dashboard.title')}
            </Text>
          </View>

          {/* Filter chips */}
          <FilterChips filters={allCategories} active={activeFilters} onToggle={toggleFilter} />

          {/* Search */}
          <View
            className="mx-5 mt-3.5 mb-5 flex-row items-center rounded-full px-[18px] py-[11px]"
            style={{ backgroundColor: "#F3F4F6" }}
          >
            <Text className="text-sm mr-2.5" style={{ color: AppColors.subtleText }}>🔍</Text>
            <TextInput
              className="flex-1 text-sm"
              style={{ color: AppColors.primaryText }}
              placeholder={translate('dashboard.searchPlaceholder')}
              placeholderTextColor={AppColors.placeholderText}
              value={search}
              onChangeText={setSearch}
            />
          </View>

          {/* Discover Weekly */}
          <View className="flex-row items-center px-5 mb-3">
            <Text className="text-[#7C3AED] text-[13px] mr-1.5">✦</Text>
            <Text className="text-[15px] font-semibold" style={{ color: AppColors.primaryText }}>{translate('dashboard.discoverWeekly')}</Text>
          </View>
          {featuredEvent && <DiscoverBanner event={featuredEvent} />}

          {/* For You */}
          {listEvents.length > 0 && (
            <View className="px-5 mt-6 mb-1">
              <View className="flex-row items-center gap-1.5 mb-1">
                <View className="w-[26px] h-[26px] rounded-full bg-[#3B82F6] items-center justify-center">
                  <Text className="text-white text-[11px] font-bold">A</Text>
                </View>
                <Text className="text-sm font-semibold" style={{ color: AppColors.primaryText }}>{translate('dashboard.forYou')}</Text>
                <View className="flex-row ml-1">
                  {["#10B981", "#6366F1"].map((c, i) => (
                    <View key={i} className="w-[22px] h-[22px] rounded-full border-2 border-white" style={{ backgroundColor: c, marginLeft: i > 0 ? -8 : 0 }} />
                  ))}
                </View>
              </View>
              <View className="flex-row items-center gap-1">
                <Text className="text-[#7C3AED] text-xs">✦</Text>
                <Text className="text-[13px]" style={{ color: AppColors.mutedText }}>{translate('dashboard.basedOn')} </Text>
                <Text className="text-[13px] text-[#3B82F6]">{translate('dashboard.upcomingEvents')}</Text>
              </View>
            </View>
          )}

          <View className="h-4" />
        </View>
      }
      renderItem={({ item }) => <EventCard item={item} />}
      ListEmptyComponent={
        <View className="items-center p-8">
          <ThemedText style={{ color: AppColors.subtleText }}>{translate('common.noUpcomingEvents')}</ThemedText>
        </View>
      }
    />
  );
}

// ── Screen ────────────────────────────────────────────────────────────────────

export default function HomeScreen() {
  return (
    <View className="flex-1 bg-white">
      <EventOccurrencesList />
    </View>
  );
}