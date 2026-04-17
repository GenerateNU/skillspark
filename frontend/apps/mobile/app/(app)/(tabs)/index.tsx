import { FLOATING_TAB_BAR_SCROLL_PADDING } from "@/components/floating-tab-bar";
import CarouselCard from "@/components/home/CarouselCard";
import { CategoryCard } from "@/components/home/CategoryCard";
import { RecommendedCard } from "@/components/home/RecommendedCard";
import { TrendingCard } from "@/components/home/TrendingCard";
import { UpcomingClassCard } from "@/components/home/UpcomingClassCard";
import LogoBgWrapper from "@/components/LogoBgWrapper";
import { ThemedText } from "@/components/themed-text";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { AppColors, FontSizes } from "@/constants/theme";
import { useAuthContext } from "@/hooks/use-auth-context";
import { useFilters } from "@/hooks/use-filters";
import { useGeolocation } from "@/hooks/use-geolocation";
import { extractResponseData, isWithinNext7Days } from "@/utils/format";
import {
  useGetAllEventOccurrences,
  useGetAllEvents,
  useGetChildrenByGuardianId,
  useGetGuardianById,
  useGetRegistrationsByGuardianId,
  useGetTrendingEventOccurrences,
  type Child,
  type Event,
  type EventOccurrence,
  type Guardian,
  type Registration,
} from "@skillspark/api-client";
import { useRouter } from "expo-router";
import { useMemo } from "react";
import { useTranslation } from "react-i18next";
import {
  ActivityIndicator,
  Pressable,
  ScrollView,
  Text,
  useWindowDimensions,
  View,
} from "react-native";

export default function HomeScreen() {
  const { t: translate } = useTranslation();
  const { guardianId } = useAuthContext();
  const { hasActiveFilters } = useFilters();
  const router = useRouter();
  const { width, height } = useWindowDimensions();

  const { lat: geoLocationLat, lng: geoLocationLong } = useGeolocation();

  const { data: localizedOccurrencesResp } = useGetAllEventOccurrences({
    lat: geoLocationLat,
    lng: geoLocationLong,
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

  // Events (templates) — used for carousel and recommendations
  const { data: eventsResp, isLoading } = useGetAllEvents();
  const allEvents: Event[] = useMemo(
    () => extractResponseData<Event>(eventsResp),
    [eventsResp],
  );

  // Occurrences — kept only for matching upcoming registered classes
  const { data: occurrencesResp } = useGetAllEventOccurrences({});
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
      lat: Number(geoLocationLat),
      lng: Number(geoLocationLong),
      radius: 50,
      max_returns: 5,
    },
    {
      query: {
        enabled: !!geoLocationLat && !!geoLocationLong,
      },
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

  const childRecommendations = useMemo(() => {
    const shuffled = [...allEvents].sort(() => Math.random() - 0.5);
    return children
      .map((child, i) => {
        const start = i * 3;
        const slice = shuffled.slice(start, start + 3);
        const events = slice.length > 0 ? slice : shuffled.slice(0, 3);
        return { child, events };
      })
      .filter((r) => r.events.length > 0);
  }, [children, allEvents]);

  const categories = [
    "Sports & Physical Activities",
    "Arts & Creative Expression",
    "Languages",
    "Academics",
    "Personal Development & Life Skills",
    "Music & Performance",
    "Math",
    "Tech & Innovation",
  ];

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
      <LogoBgWrapper verticalOffset={0}>
        {/* Header */}
        <View className="px-5 pt-28 pb-4">
          <Text
            className="font-nunito-bold"
            style={{
              letterSpacing: -0.5,
              fontSize: FontSizes.hero,
              color: AppColors.primaryText,
            }}
          >
            {translate("dashboard.greeting", { name: firstName })}
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
            <Text
              className="font-nunito-bold px-5 mb-3"
              style={{ fontSize: FontSizes.lg, color: AppColors.primaryText }}
            >
              {translate("dashboard.upcomingClasses")}
            </Text>
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
        {allEvents.length > 0 && (
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
                  ? allLocalizedOccurrences.map((o) => o.event)
                  : allEvents.slice(0, 5)
              }
              width={width}
              height={height}
            />
          </View>
        )}

        {/* Trending In Your Area */}
        {trendingEvents && trendingEvents.length > 0 && (
          <View className="mb-6">
            <Text
              className="font-nunito-bold px-5 mb-3"
              style={{ fontSize: FontSizes.lg, color: AppColors.primaryText }}
            >
              {translate("dashboard.trendingInYourArea")}
            </Text>
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
            <Text
              className="font-nunito-bold px-5 mb-3"
              style={{ fontSize: FontSizes.lg, color: AppColors.primaryText }}
            >
              {translate("dashboard.recommendedFor")}
            </Text>
            <ScrollView
              horizontal
              showsHorizontalScrollIndicator={false}
              contentContainerStyle={{ paddingHorizontal: 20 }}
            >
              {childRecommendations.map(({ child, events }) => (
                <RecommendedCard key={child.id} child={child} events={events} />
              ))}
            </ScrollView>
          </View>
        )}

        {/* Explore by Category */}
        {categories.length > 0 && (
          <View className="mb-6">
            <Text
              className="font-nunito-bold px-5 mb-3"
              style={{ fontSize: FontSizes.lg, color: AppColors.primaryText }}
            >
              {translate("dashboard.exploreByCategory")}
            </Text>
            <View className="px-[15px]">
              {categoryPairs.map((pair, idx) => (
                <View key={idx} className="flex-row">
                  {pair.map((cat) => (
                    <CategoryCard key={cat} category={cat} />
                  ))}
                  {pair.length === 1 && <View className="flex-1 m-[5px]" />}
                </View>
              ))}
            </View>
          </View>
        </View>
      )}
      </LogoBgWrapper>
    </ScrollView>
  );
}
