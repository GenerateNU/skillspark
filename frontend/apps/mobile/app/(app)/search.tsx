import { getAllEvents, searchEvents, type Event } from "@skillspark/api-client";
import { useLocalSearchParams, useRouter } from "expo-router";
import { useMemo, useState } from "react";
import {
  ActivityIndicator,
  FlatList,
  Pressable,
  Text,
  TextInput,
  View,
} from "react-native";
import { useSafeAreaInsets } from "react-native-safe-area-context";
import { useDebounce } from "use-debounce";
import { useTranslation } from "react-i18next";
import { SearchResultCard } from "@/components/home/SearchResultCard";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { useFilters } from "@/hooks/use-filters";
import { AppColors } from "@/constants/theme";
import { FLOATING_TAB_BAR_SCROLL_PADDING } from "@/components/floating-tab-bar";
import { useInfiniteQuery } from "@tanstack/react-query";

const PAGE_SIZE = 20;

export default function SearchScreen() {
  const router = useRouter();
  const insets = useSafeAreaInsets();
  const { q } = useLocalSearchParams<{ q?: string }>();
  const { filters } = useFilters();
  const { t: translate } = useTranslation();

  const [searchText, setSearchText] = useState(q ?? "");
  const [debouncedSearch] = useDebounce(searchText, 300);

  const hasQuery = debouncedSearch.trim().length > 0;

  // Fuzzy OpenSearch — only when there is a search query
  const {
    data: searchData,
    isLoading: searchLoading,
    fetchNextPage: fetchNextSearchPage,
    hasNextPage: hasNextSearch,
    isFetchingNextPage: isFetchingNextSearch,
  } = useInfiniteQuery({
    queryKey: ["search", "events", debouncedSearch],
    queryFn: ({ pageParam }) =>
      searchEvents({ q: debouncedSearch, page: pageParam, limit: PAGE_SIZE }),
    initialPageParam: 1,
    getNextPageParam: (lastPage, allPages) => {
      const items = lastPage.data;
      if (Array.isArray(items) && items.length === PAGE_SIZE) {
        return allPages.length + 1;
      }
      return undefined;
    },
    enabled: hasQuery,
  });

  // Postgres filtered fetch — only when there is no search query
  const {
    data: allEventsData,
    isLoading: allEventsLoading,
    fetchNextPage: fetchNextAllPage,
    hasNextPage: hasNextAll,
    isFetchingNextPage: isFetchingNextAll,
  } = useInfiniteQuery({
    queryKey: ["events", "all", filters.category, filters.min_age, filters.max_age],
    queryFn: ({ pageParam }) =>
      getAllEvents({
        category: filters.category,
        min_age: filters.min_age,
        max_age: filters.max_age,
        page: pageParam,
        limit: PAGE_SIZE,
      }),
    initialPageParam: 1,
    getNextPageParam: (lastPage, allPages) => {
      const items = lastPage.data;
      if (Array.isArray(items) && items.length === PAGE_SIZE) {
        return allPages.length + 1;
      }
      return undefined;
    },
    enabled: !hasQuery,
  });

  const isLoading = hasQuery ? searchLoading : allEventsLoading;
  const fetchNextPage = hasQuery ? fetchNextSearchPage : fetchNextAllPage;
  const hasNextPage = hasQuery ? hasNextSearch : hasNextAll;
  const isFetchingNextPage = hasQuery ? isFetchingNextSearch : isFetchingNextAll;

  const results: Event[] = useMemo(() => {
    if (hasQuery) {
      const raw =
        searchData?.pages.flatMap((page) =>
          Array.isArray(page.data) ? (page.data as Event[]) : [],
        ) ?? [];

      // Client-side apply active filters on top of fuzzy results
      return raw.filter((event) => {
        if (filters.category) {
          const dbCats = filters.category.split(",");
          if (!event.category?.some((c) => dbCats.includes(c))) return false;
        }
        if (
          filters.min_age != null &&
          event.age_range_max != null &&
          event.age_range_max < filters.min_age
        )
          return false;
        if (
          filters.max_age != null &&
          event.age_range_min != null &&
          event.age_range_min > filters.max_age
        )
          return false;
        return true;
      });
    }

    return (
      allEventsData?.pages.flatMap((page) =>
        Array.isArray(page.data) ? (page.data as Event[]) : [],
      ) ?? []
    );
  }, [hasQuery, searchData, allEventsData, filters]);

  return (
    <View className="flex-1 bg-white" style={{ paddingTop: insets.top }}>
      {/* Search bar */}
      <View className="px-5 pb-3 flex-row items-center gap-2">
        <Pressable onPress={() => router.back()} hitSlop={12}>
          <IconSymbol
            name="chevron.left"
            size={22}
            color={AppColors.primaryText}
          />
        </Pressable>
        <View className="flex-1 flex-row items-center rounded-full px-4 py-[10px] bg-[#F3F4F6]">
          <IconSymbol
            name="magnifyingglass"
            size={18}
            color={AppColors.subtleText}
            className="mr-2"
          />
          <TextInput
            className="flex-1 font-nunito text-sm text-[#111]"
            placeholder={translate("search.placeholder")}
            placeholderTextColor={AppColors.placeholderText}
            value={searchText}
            onChangeText={setSearchText}
            autoFocus
            returnKeyType="search"
            clearButtonMode="while-editing"
          />
        </View>
      </View>

      {/* Results */}
      {isLoading ? (
        <View className="flex-1 items-center justify-center">
          <ActivityIndicator size="large" />
        </View>
      ) : results.length === 0 ? (
        <View className="flex-1 items-center justify-center gap-2">
          <Text className="font-nunito-bold text-[15px] text-[#111]">
            {translate("search.noResults")}
          </Text>
          <Text className="font-nunito text-sm text-[#6B7280]">
            {translate("search.tryDifferent")}
          </Text>
        </View>
      ) : (
        <FlatList
          data={results}
          keyExtractor={(item) => item.id}
          contentContainerStyle={{
            paddingHorizontal: 15,
            paddingTop: 8,
            paddingBottom: FLOATING_TAB_BAR_SCROLL_PADDING,
            gap: 12,
          }}
          renderItem={({ item }) => <SearchResultCard event={item} />}
          keyboardShouldPersistTaps="handled"
          keyboardDismissMode="on-drag"
          onEndReached={() => {
            if (hasNextPage && !isFetchingNextPage) {
              fetchNextPage();
            }
          }}
          onEndReachedThreshold={0.3}
          ListFooterComponent={
            isFetchingNextPage ? (
              <View className="py-4 items-center">
                <ActivityIndicator size="small" />
              </View>
            ) : null
          }
        />
      )}
    </View>
  );
}
