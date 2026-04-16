import { Image } from "expo-image";
import {
  ActivityIndicator,
  Pressable,
  ScrollView,
  Text,
  TextInput,
  TouchableOpacity,
  View,
} from "react-native";

import { SafeAreaView } from "react-native-safe-area-context";
import { useLocalSearchParams, useRouter } from "expo-router";
import { useMemo, useState } from "react";
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
import { formatAgeRange } from "@/utils/format";
import { useGeoLocation } from "@/hooks/use-geo-location";
import { FLOATING_TAB_BAR_SCROLL_PADDING } from "@/components/floating-tab-bar";
import { TrendingCard } from "@/components/home/TrendingCard";
import { ChildAvatar } from "@/components/ChildAvatar";

export default function ForChildScreen() {
  const { id, name } = useLocalSearchParams<{ id: string; name: string }>();
  const router = useRouter();
  const [searchText, setSearchText] = useState("");

  const { lat: geoLocationLat, lng: geoLocationLng } = useGeoLocation();

  const { data: occurrencesResp, isLoading } = useGetAllEventOccurrences();
  const allOccurrences: EventOccurrence[] = useMemo(() => {
    const d = occurrencesResp as unknown as
      | { data: EventOccurrence[] }
      | undefined;
    return Array.isArray(d?.data) ? d.data : [];
  }, [occurrencesResp]);

  const futureOccurrences = useMemo(
    () => allOccurrences.filter((o) => new Date(o.start_time) > new Date()),
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
    { lat: geoLocationLat, lng: geoLocationLng, radius: 50, max_returns: 5 },
    { query: { enabled: true } },
  );
  const trendingEvents: EventOccurrence[] = useMemo(() => {
    const d = trendingResp as unknown as
      | { data: EventOccurrence[] }
      | undefined;
    return Array.isArray(d?.data) ? d.data : [];
  }, [trendingResp]);

  const featuredOccurrence =
    trendingEvents[0] ?? futureOccurrences[0] ?? null;

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
  const childInterests: string[] = Array.isArray(child?.interests) ? child.interests : [];

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
    ? [
        featuredOccurrence.location?.address_line1,
        featuredOccurrence.location?.address_line2,
        featuredOccurrence.location?.district,
      ]
        .filter(Boolean)
        .join(", ")
    : null;

  const featuredAgeLabel = featuredOccurrence
    ? formatAgeRange(featuredOccurrence.event.age_range_min, featuredOccurrence.event.age_range_max) || null
    : null;

  return (
    <SafeAreaView className="flex-1 bg-white" edges={["top", "bottom"]}>
      <View className="flex-row items-center px-5 pt-3 pb-4">
        <TouchableOpacity onPress={() => router.back()} className="mr-3">
          <IconSymbol
            name="chevron.left"
            size={22}
            color={AppColors.primaryText}
          />
        </TouchableOpacity>
        <Text className="font-nunito-bold text-[22px]" style={{ color: AppColors.primaryText }}>
          For {name}
        </Text>
      </View>
      <View className="px-5 mb-4">
        <View className="flex-row items-center rounded-full px-4 py-[10px]" style={{ backgroundColor: AppColors.surfaceGray }}>
          <View className="mr-2">
            <IconSymbol name="magnifyingglass" size={18} color={AppColors.subtleText} />
          </View>
          <TextInput
            className="flex-1 text-sm font-nunito"
            style={{ color: AppColors.primaryText }}
            placeholder="Search a class"
            placeholderTextColor={AppColors.placeholderText}
            value={searchText}
            onChangeText={setSearchText}
          />
          <Pressable className="w-9 h-9 rounded-full items-center justify-center" style={{ backgroundColor: AppColors.primaryText }}>
            <IconSymbol
              name="slider.horizontal.3"
              size={16}
              color={AppColors.white}
            />
          </Pressable>
        </View>
      </View>
      {isLoading ? (
        <View className="flex-1 items-center justify-center">
          <ActivityIndicator size="large" />
        </View>
      ) : (
        <ScrollView
          showsVerticalScrollIndicator={false}
          contentContainerStyle={{ paddingBottom: FLOATING_TAB_BAR_SCROLL_PADDING }}
        >
          {featuredOccurrence && (
            <View className="mx-5 mb-5 rounded-3xl p-4" style={{ backgroundColor: AppColors.bluePastel }}>
              <View className="flex-row items-center gap-2 mb-3">
                <Text className="font-nunito-bold text-base text-[#111]">
                  ✦ Trending right now...
                </Text>
              </View>
              <Pressable
                onPress={() =>
                  router.push(`/event/${featuredOccurrence.event.id}`)
                }
                className="bg-white rounded-2xl p-3 flex-row items-center gap-3"
              >
                {featuredOccurrence.event.presigned_url ? (
                  <Image
                    source={{ uri: featuredOccurrence.event.presigned_url }}
                    className="w-[110px] h-[80px] rounded-xl"
                    contentFit="cover"
                  />
                ) : (
                  <View className="w-[110px] h-[80px] rounded-xl bg-[#D9D9D9]" />
                )}
                <View className="flex-1 gap-0.5">
                  <Text className="font-nunito-bold text-base text-[#111]" numberOfLines={1}>
                    {featuredOccurrence.event.title}
                  </Text>
                  {!!orgName && (
                    <Text className="font-nunito text-xs text-gray-500">
                      {orgName}
                    </Text>
                  )}
                  {!!featuredAgeLabel && (
                    <Text className="font-nunito text-xs text-gray-500">
                      {featuredAgeLabel}
                    </Text>
                  )}
                  {!!featuredAddress && (
                    <Text className="font-nunito text-xs text-gray-500" numberOfLines={2}>
                      {featuredAddress}
                    </Text>
                  )}
                </View>
                <View className="items-end justify-between self-stretch gap-2">
                  <View className="flex-row items-center bg-gray-100 rounded-full px-2 py-1 gap-1">
                    <ChildAvatar
                      name={name ?? ""}
                      avatarFace={child?.avatar_face}
                      avatarBackground={child?.avatar_background}
                      size={20}
                    />
                    <Text className="font-nunito-semibold text-xs text-[#111]">
                      {name}
                    </Text>
                  </View>
                  <View className="w-9 h-9 rounded-full items-center justify-center" style={{ backgroundColor: AppColors.slateBlue }}>
                    <IconSymbol name="chevron.right" size={14} color="white" />
                  </View>
                </View>
              </Pressable>
            </View>
          )}
          {matchedCategories.map((cat) => (
            <View key={cat} className="mb-5">
              <Text className="font-nunito-bold px-5 mb-2 text-xl"
                style={{ color: AppColors.primaryText }}>
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
                    userLat={geoLocationLat}
                    userLng={geoLocationLng}
                  />
                ))}
              </ScrollView>
            </View>
          ))}
          {otherEvents.length > 0 && (
            <View className="mb-5">
              <Text className="font-nunito-bold px-5 mb-2 text-xl"
                style={{ color: AppColors.primaryText }}>
                Other Events
              </Text>
              <View className="px-5">
                {otherEvents.map((o) => (
                  <TrendingCard
                    key={o.id}
                    occurrence={o}
                    userLat={geoLocationLat}
                    userLng={geoLocationLng}
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
