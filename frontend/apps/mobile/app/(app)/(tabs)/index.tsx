import {
  ActivityIndicator,
  View,
  ScrollView,
  Text,
  Pressable,
  useWindowDimensions,
} from "react-native";
import {
  useGetAllEventOccurrences,
  useGetGuardianById,
  useGetRegistrationsByGuardianId,
  useGetChildrenByGuardianId,
  type EventOccurrence,
  type Guardian,
  type Registration,
  type Child,
  useGetTrendingEventOccurrences,
} from "@skillspark/api-client";
import { useMemo, useState } from "react";
import { AppColors, FontSizes } from "@/constants/theme";
import { useAuthContext } from "@/hooks/use-auth-context";
import { useFilters } from "@/hooks/use-filters";
import { useRouter } from "expo-router";
import { isWithinNext7Days, filterFutureOccurrences, extractResponseData } from "@/utils/format";
import { HomeSectionHeader } from "@/components/SectionHeader";
import { DiscoverBanner } from "@/components/home/DiscoverBanner";
import { UpcomingClassCard } from "@/components/home/UpcomingClassCard";
import { RecommendedCard } from "@/components/home/RecommendedCard";
import { CategoryCard } from "@/components/home/CategoryCard";
import { ThemedText } from "@/components/themed-text";
import { useTranslation } from "react-i18next";
import { TrendingCard } from "@/components/home/TrendingCard";
import { useGeoLocation } from "@/hooks/use-geo-location";
import CarouselCard from "@/components/home/CarouselCard";
import { FLOATING_TAB_BAR_SCROLL_PADDING } from "@/components/floating-tab-bar";
import { SearchBar } from "@/components/SearchBar";
import { IconSymbol } from "@/components/ui/icon-symbol";

export default function HomeScreen() {
  const { t: translate } = useTranslation();
  const { guardianId } = useAuthContext();
  const { filters, hasActiveFilters } = useFilters();
  const router = useRouter();
  const { width, height } = useWindowDimensions();

  const { lat: geoLocationLat, lng: geoLocationLong } = useGeoLocation();

  const { data: localizedOccurrencesResp } = useGetAllEventOccurrences({
    lat: String(geoLocationLat),
    lng: String(geoLocationLong),
    radius_km: 50,
    limit: 5,
  });
  const allLocalizedOccurrences: EventOccurrence[] = useMemo(
    () => extractResponseData<EventOccurrence>(localizedOccurrencesResp),
    [localizedOccurrencesResp],
  );

  const { data: guardianResp } = useGetGuardianById(guardianId!, {
    query: { enabled: !!guardianId },
  });
  const guardian = (guardianResp as unknown as { data: Guardian } | undefined)
    ?.data;

  const { data: occurrencesResp, isLoading } = useGetAllEventOccurrences({});
  const allOccurrences: EventOccurrence[] = useMemo(
    () => extractResponseData<EventOccurrence>(occurrencesResp),
    [occurrencesResp],
  );

  const { data: registrationsResp } = useGetRegistrationsByGuardianId(
    guardianId!,
    {
      query: { enabled: !!guardianId },
    },
  );
  const registrations: Registration[] = useMemo(() => {
    const d = registrationsResp as unknown as
      | { data: { registrations: Registration[] } }
      | undefined;
    return d?.data?.registrations ?? [];
  }, [registrationsResp]);

  const { data: childrenResp } = useGetChildrenByGuardianId(guardianId!, {
    query: { enabled: !!guardianId },
  });
  const children: Child[] = useMemo(
    () => extractResponseData<Child>(childrenResp),
    [childrenResp],
  );

  const { data: trendingResp } = useGetTrendingEventOccurrences(
    {
      lat: geoLocationLat,
      lng: geoLocationLong,
      radius: 50,
      max_returns: 5,
    },
    {
      query: { enabled: true },
    },
  );

  const trendingEvents: EventOccurrence[] = useMemo(
    () => extractResponseData<EventOccurrence>(trendingResp),
    [trendingResp],
  );

  const upcomingClasses = useMemo(() => {
    const upcomingIds = new Set(
      registrations
        .filter(
          (r) =>
            r.status === "registered" &&
            isWithinNext7Days(r.occurrence_start_time),
        )
        .map((r) => r.event_occurrence_id),
    );
    return allOccurrences.filter((o) => upcomingIds.has(o.id));
  }, [registrations, allOccurrences]);

  const futureOccurrences = useMemo(
    () => filterFutureOccurrences(allOccurrences),
    [allOccurrences],
  );

  const childRecommendations = useMemo(() => {
    const shuffled = [...futureOccurrences].sort(() => Math.random() - 0.5);
    return children
      .map((child, i) => {
        const start = i * 3;
        const slice = shuffled.slice(start, start + 3);
        const occurrences = slice.length > 0 ? slice : shuffled.slice(0, 3);
        return { child, occurrences };
      })
      .filter((r) => r.occurrences.length > 0);
  }, [children, futureOccurrences]);

  const categories = useMemo(() => {
    const cats = new Set<string>();
    allOccurrences.forEach((o) =>
      o.event.category?.forEach((c) => cats.add(c)),
    );
    return cats.size > 0
      ? Array.from(cats)
      : ["Sport", "Arts", "Music", "Tech", "Activity", "Tutoring"];
  }, [allOccurrences]);

  const categoryEventMap = useMemo(() => {
    const map: Record<string, EventOccurrence> = {};
    allOccurrences.forEach((o) => {
      o.event.category?.forEach((c) => {
        if (!map[c] && o.event.presigned_url) map[c] = o;
      });
    });
    return map;
  }, [allOccurrences]);

  const firstName = guardian?.name?.split(" ")[0] ?? "there";

  if (isLoading) {
    return (
      <View className="flex-1 items-center justify-center bg-white">
        <ActivityIndicator size="large" />
        <ThemedText>{translate("common.loadingEvents")}</ThemedText>
      </View>
    );
  }

  const categoryPairs: string[][] = [];
  for (let i = 0; i < categories.length; i += 2) {
    categoryPairs.push(categories.slice(i, i + 2));
  }

  return (
    <ScrollView
      className="flex-1 bg-white"
      showsVerticalScrollIndicator={false}
      contentContainerStyle={{ paddingBottom: FLOATING_TAB_BAR_SCROLL_PADDING }}
    >
      {/* Header */}
      <View className="px-5 pt-14 pb-4">
        <Text
          className="font-nunito-bold"
          style={{
            letterSpacing: -0.5,
            fontSize: FontSizes.hero,
            color: AppColors.primaryText,
          }}
        >
          Hello, {firstName}
        </Text>
      </View>

      {/* Search row */}
      <View className="px-5 mb-[22px]">
        <View
          className="flex-row items-center rounded-full px-4 py-[10px]"
          style={{ backgroundColor: AppColors.surfaceGray }}
        >
          <IconSymbol
            name="magnifyingglass"
            size={18}
            color={AppColors.subtleText}
            style={{ marginRight: 8 }}
          />
          <Pressable
            className="flex-1"
            onPress={() => router.push("/(app)/search")}
          >
            <Text
              style={{
                fontFamily: "NunitoSans_400Regular",
                fontSize: 14,
                color: AppColors.placeholderText,
              }}
            >
              {translate("dashboard.searchPlaceholder")}
            </Text>
          </Pressable>
          <Pressable
            onPress={() => router.push("/(app)/(tabs)/filters")}
            className="w-9 h-9 rounded-full items-center justify-center"
            style={{
              backgroundColor: hasActiveFilters
                ? AppColors.primaryBlue
                : AppColors.primaryText,
            }}
          >
            <IconSymbol
              name="slider.horizontal.3"
              size={16}
              color={AppColors.white}
            />
          </Pressable>
        </View>
      </View>

      {/* Your Upcoming Classes — conditional */}
      {upcomingClasses.length > 0 && (
        <View className="mb-6">
          <HomeSectionHeader title="Your Upcoming Classes" />
          <ScrollView
            horizontal
            showsHorizontalScrollIndicator={false}
            contentContainerStyle={{ paddingHorizontal: 20 }}
          >
            {upcomingClasses.map((o) => (
              <UpcomingClassCard key={o.id} occurrence={o} />
            ))}
          </ScrollView>
        </View>
      )}

      {/* Discover Weekly */}
      {futureOccurrences.length > 0 && (
        <View className="mb-6">
          <View className="flex-row items-center px-5 mb-3">
            <Text
              className="mr-1.5 font-nunito"
              style={{ color: AppColors.purple, fontSize: FontSizes.md }}
            >
              ✦
            </Text>
            <Text
              className="font-nunito-bold"
              style={{ fontSize: FontSizes.lg, color: AppColors.primaryText }}
            >
              {translate("dashboard.discoverWeekly")}
            </Text>
          </View>
          <CarouselCard
            events={
              allLocalizedOccurrences.length > 0
                ? allLocalizedOccurrences
                : futureOccurrences.slice(0, 5)
            }
            width={width}
            height={height}
          />
        </View>
      )}

      {/* Trending In Your Area */}
      {trendingEvents && trendingEvents.length > 0 && (
        <View className="mb-6">
          <HomeSectionHeader title="Trending in Your Area" />
          <ScrollView
            horizontal
            showsHorizontalScrollIndicator={false}
            contentContainerStyle={{ paddingHorizontal: 20 }}
          >
            {trendingEvents.map((o) => (
              <TrendingCard key={o.id} occurrence={o} />
            ))}
          </ScrollView>
        </View>
      )}

      {/* Recommended For... */}
      {childRecommendations.length > 0 && (
        <View className="mb-6">
          <HomeSectionHeader title="Recommended for..." />
          <View>
            {childRecommendations.map(({ child, occurrences }) => (
              <RecommendedCard
                key={child.id}
                child={child}
                occurrences={occurrences}
              />
            ))}
          </View>
        </View>
      )}

      {/* Explore by Category */}
      {categories.length > 0 && (
        <View className="mb-6">
          <HomeSectionHeader title="Explore by Category" />
          <View className="px-[15px]">
            {categoryPairs.map((pair, idx) => (
              <View key={idx} className="flex-row">
                {pair.map((cat) => (
                  <CategoryCard
                    key={cat}
                    category={cat}
                    occurrence={categoryEventMap[cat]}
                  />
                ))}
                {pair.length === 1 && <View className="flex-1 m-[5px]" />}
              </View>
            ))}
          </View>
        </View>
      )}
    </ScrollView>
  );
}
