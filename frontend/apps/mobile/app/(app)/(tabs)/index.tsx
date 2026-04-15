import { IconSymbol } from "@/components/ui/icon-symbol";
import {
  ActivityIndicator,
  View,
  TextInput,
  ScrollView,
  Pressable,
  Text,
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
  getTrendingEventOccurrences,
  useGetTrendingEventOccurrences,
} from "@skillspark/api-client";
import { useEffect, useMemo, useState } from "react";
import { AppColors, FontSizes } from "@/constants/theme";
import { useAuthContext } from "@/hooks/use-auth-context";
import { useDebounce } from "use-debounce";
import { isWithinNext7Days } from "@/utils/format";
import { DiscoverBanner } from "@/components/home/DiscoverBanner";
import { UpcomingClassCard } from "@/components/home/UpcomingClassCard";
import { RecommendedCard } from "@/components/home/RecommendedCard";
import { CategoryCard } from "@/components/home/CategoryCard";
import { ThemedText } from "@/components/themed-text";
import { useTranslation } from "react-i18next";
import { TrendingCard } from "@/components/home/TrendingCard";

import * as Location from "expo-location";
import CarouselCard from "@/components/home/CarouselCard";
import { FLOATING_TAB_BAR_SCROLL_PADDING } from "@/components/floating-tab-bar";

export default function HomeScreen() {
  const { t: translate } = useTranslation();
  const { guardianId } = useAuthContext();
  const [searchText, setSearchText] = useState("");
  const [_debouncedSearch] = useDebounce(searchText, 300);
  const { width, height } = useWindowDimensions();

  const [geoLocationLat, setGeoLocationLat] = useState<string | undefined>(
    "13.7563",
  );
  const [geoLocationLong, setGeoLocationLong] = useState<string | undefined>(
    "100.5018",
  );

  useEffect(() => {
    (async () => {
      const { status } = await Location.requestForegroundPermissionsAsync();
      if (status !== "granted") return;
      const loc = await Location.getCurrentPositionAsync({});
      setGeoLocationLat(String(loc.coords.latitude));
      setGeoLocationLong(String(loc.coords.longitude));
    })();
  }, []);

  const { data: localizedOccurrencesResp } = useGetAllEventOccurrences({
    lat: geoLocationLat,
    lng: geoLocationLong,
    radius_km: 50,
    limit: 5,
  });
  const allLocalizedOccurrences: EventOccurrence[] = useMemo(() => {
    const d = localizedOccurrencesResp as unknown as
      | { data: EventOccurrence[] }
      | undefined;
    return Array.isArray(d?.data) ? d.data : [];
  }, [localizedOccurrencesResp]);

  const { data: guardianResp } = useGetGuardianById(guardianId!, {
    query: { enabled: !!guardianId },
  });
  const guardian = (guardianResp as unknown as { data: Guardian } | undefined)
    ?.data;

  const { data: occurrencesResp, isLoading } = useGetAllEventOccurrences();
  const allOccurrences: EventOccurrence[] = useMemo(() => {
    const d = occurrencesResp as unknown as
      | { data: EventOccurrence[] }
      | undefined;
    return Array.isArray(d?.data) ? d.data : [];
  }, [occurrencesResp]);

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
  const children: Child[] = useMemo(() => {
    const d = childrenResp as unknown as { data: Child[] } | undefined;
    return Array.isArray(d?.data) ? d.data : [];
  }, [childrenResp]);

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

  const trendingEvents: EventOccurrence[] = useMemo(() => {
    const d = trendingResp as unknown as
      | { data: EventOccurrence[] }
      | undefined;
    return Array.isArray(d?.data) ? d.data : [];
  }, [trendingResp]);

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
    () => allOccurrences.filter((o) => new Date(o.start_time) > new Date()),
    [allOccurrences],
  );

  const childRecommendations = useMemo(() => {
    const shuffled = [...futureOccurrences].sort(() => Math.random() - 0.5);
    return children
      .map((child, i) => ({
        child,
        occurrence: shuffled[i % shuffled.length],
      }))
      .filter((r) => r.occurrence != null);
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
          <TextInput
            className="flex-1 text-sm font-nunito"
            style={{ color: AppColors.primaryText }}
            placeholder={translate("dashboard.searchPlaceholder")}
            placeholderTextColor={AppColors.placeholderText}
            value={searchText}
            onChangeText={setSearchText}
          />
          <Pressable
            className="w-9 h-9 rounded-full items-center justify-center"
            style={{ backgroundColor: AppColors.primaryText }}
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
            Your Upcoming Classes
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
          <Text
            className="font-nunito-bold px-5 mb-3"
            style={{ fontSize: FontSizes.lg, color: AppColors.primaryText }}
          >
            Trending in Your Area
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
            Recommended for...
          </Text>
          <ScrollView
            horizontal
            showsHorizontalScrollIndicator={false}
            contentContainerStyle={{ paddingHorizontal: 20 }}
          >
            {childRecommendations.map(({ child, occurrence }) => (
              <RecommendedCard
                key={child.id}
                occurrence={occurrence}
                childName={child.name.split(" ")[0]}
              />
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
            Explore by Category
          </Text>
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
