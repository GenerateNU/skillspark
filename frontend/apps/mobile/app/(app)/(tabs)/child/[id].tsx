import {
  ActivityIndicator,
  ScrollView,
  Text,
  TouchableOpacity,
  View,
} from "react-native";

import { SafeAreaView } from "react-native-safe-area-context";
import { useLocalSearchParams, useRouter } from "expo-router";
import { useMemo, useState } from "react";
import { useTranslation } from "react-i18next";
import {
  useGetAllEventOccurrences,
  useGetTrendingEventOccurrences,
  useGetOrganization,
  useGetChildById,
  type EventOccurrence,
  type Organization,
  type Child,
} from "@skillspark/api-client";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { AppColors } from "@/constants/theme";
import {
  formatAgeRange,
  filterFutureOccurrences,
  extractResponseData,
  formatAddress,
} from "@/utils/format";
import { useGeolocation } from "@/hooks/use-geolocation";
import { FLOATING_TAB_BAR_SCROLL_PADDING } from "@/components/floating-tab-bar";
import { TrendingCard } from "@/components/home/TrendingCard";
import { SearchBar } from "@/components/SearchBar";
import { FeaturedOccurrenceCard } from "@/components/FeaturedOccurrenceCard";

export default function ForChildScreen() {
  const { id, name } = useLocalSearchParams<{ id: string; name: string }>();
  const router = useRouter();
  const { t: translate } = useTranslation();
  const [searchText, setSearchText] = useState("");

  const { lat: geoLocationLat, lng: geoLocationLng } = useGeolocation();

  const { data: occurrencesResp, isLoading } = useGetAllEventOccurrences();
  const allOccurrences: EventOccurrence[] = useMemo(
    () => extractResponseData<EventOccurrence>(occurrencesResp),
    [occurrencesResp],
  );

  const futureOccurrences = useMemo(
    () => filterFutureOccurrences(allOccurrences),
    [allOccurrences],
  );

  const filteredOccurrences = useMemo(
    () =>
      searchText.trim()
        ? futureOccurrences.filter((o) =>
            o.event.title.toLowerCase().includes(searchText.toLowerCase()),
          )
        : futureOccurrences,
    [futureOccurrences, searchText],
  );

  const { data: trendingResp } = useGetTrendingEventOccurrences(
    {
      lat: Number(geoLocationLat),
      lng: Number(geoLocationLng),
      radius: 50,
      max_returns: 5,
    },
    { query: { enabled: !!geoLocationLat && !!geoLocationLng } },
  );
  const trendingEvents: EventOccurrence[] = useMemo(
    () => extractResponseData<EventOccurrence>(trendingResp),
    [trendingResp],
  );

  const featuredOccurrence = trendingEvents[0] ?? futureOccurrences[0] ?? null;

  const orgId = featuredOccurrence?.event.organization_id;
  const { data: orgResp } = useGetOrganization(orgId!, {
    query: { enabled: !!orgId },
  });
  const orgName =
    orgResp?.status === 200 ? (orgResp.data as Organization).name : null;

  const categoryMap = useMemo(() => {
    const map: Record<string, EventOccurrence[]> = {};
    filteredOccurrences.forEach((o) => {
      o.event.category?.forEach((c) => {
        if (!map[c]) map[c] = [];
        if (map[c].length < 6) map[c].push(o);
      });
    });
    return map;
  }, [filteredOccurrences]);

  const categories = useMemo(() => Object.keys(categoryMap), [categoryMap]);

  const { data: childResp } = useGetChildById(id!, {
    query: { enabled: !!id },
  });
  const child = (childResp as unknown as { data: Child } | undefined)?.data;
  const childInterests: string[] = Array.isArray(child?.interests)
    ? child.interests
    : [];

  const matchedCategories = useMemo(() => {
    const matched = categories.filter((cat) => childInterests.includes(cat));
    return matched.slice(0, 2);
  }, [categories, childInterests]);

  const matchedEventIds = useMemo(() => {
    const ids = new Set<string>();
    matchedCategories.forEach((cat) => {
      categoryMap[cat]?.forEach((o) => ids.add(o.id));
    });
    return ids;
  }, [matchedCategories, categoryMap]);

  const otherEvents = useMemo(
    () => filteredOccurrences.filter((o) => !matchedEventIds.has(o.id)),
    [filteredOccurrences, matchedEventIds],
  );

  const featuredAddress = featuredOccurrence
    ? formatAddress(featuredOccurrence) || null
    : null;

  const featuredAgeLabel = featuredOccurrence
    ? formatAgeRange(
        featuredOccurrence.event.age_range_min,
        featuredOccurrence.event.age_range_max,
      ) || null
    : null;

  return (
    <SafeAreaView className="flex-1 bg-white" edges={["top"]}>
      <View className="flex-row items-center px-5 pt-3 pb-4">
        <TouchableOpacity onPress={() => router.back()} className="mr-3">
          <IconSymbol
            name="chevron.left"
            size={22}
            color={AppColors.primaryText}
          />
        </TouchableOpacity>
        <Text
          className="font-nunito-bold text-[22px]"
          style={{ color: AppColors.primaryText }}
        >
          {translate("dashboard.forChild", { name })}
        </Text>
      </View>
      <SearchBar
        value={searchText}
        onChangeText={setSearchText}
        placeholder={translate("dashboard.searchPlaceholder")}
        style={{ marginBottom: 16 }}
      />
      {isLoading ? (
        <View className="flex-1 items-center justify-center">
          <ActivityIndicator size="large" />
        </View>
      ) : (
        <ScrollView
          showsVerticalScrollIndicator={false}
          contentContainerStyle={{
            paddingBottom: FLOATING_TAB_BAR_SCROLL_PADDING,
          }}
        >
          {featuredOccurrence && (
            <FeaturedOccurrenceCard
              occurrence={featuredOccurrence}
              orgName={orgName}
              ageLabel={featuredAgeLabel}
              address={featuredAddress}
              childName={name ?? ""}
              child={child}
            />
          )}
          {matchedCategories.map((cat) => (
            <View key={cat} className="mb-5">
              <Text
                className="font-nunito-bold px-5 mb-2 text-xl"
                style={{ color: AppColors.primaryText }}
              >
                {cat.charAt(0).toUpperCase() + cat.slice(1)}
              </Text>
              <ScrollView
                horizontal
                showsHorizontalScrollIndicator={false}
                contentContainerStyle={{ paddingHorizontal: 20 }}
              >
                {categoryMap[cat].map((o) => (
                  <TrendingCard
                    key={o.id}
                    occurrence={o}
                    userLat={
                      geoLocationLat ? Number(geoLocationLat) : undefined
                    }
                    userLng={
                      geoLocationLng ? Number(geoLocationLng) : undefined
                    }
                    style={{ marginRight: 14 }}
                  />
                ))}
              </ScrollView>
            </View>
          ))}
          {otherEvents.length > 0 && (
            <View className="mb-5">
              <Text
                className="font-nunito-bold px-5 mb-2 text-xl"
                style={{ color: AppColors.primaryText }}
              >
                Other Events
              </Text>
              <View className="px-5">
                {otherEvents.map((o) => (
                  <TrendingCard
                    key={o.id}
                    occurrence={o}
                    userLat={
                      geoLocationLat ? Number(geoLocationLat) : undefined
                    }
                    userLng={
                      geoLocationLng ? Number(geoLocationLng) : undefined
                    }
                    width="100%"
                  />
                ))}
              </View>
            </View>
          )}
        </ScrollView>
      )}
    </SafeAreaView>
  );
}
