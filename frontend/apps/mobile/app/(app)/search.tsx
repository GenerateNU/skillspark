import { useSearchEvents, type Event } from "@skillspark/api-client";
import { useLocalSearchParams, useRouter } from "expo-router";
import { useEffect, useMemo, useState } from "react";
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
import { AppColors } from "@/constants/theme";
import { FLOATING_TAB_BAR_SCROLL_PADDING } from "@/components/floating-tab-bar";

export default function SearchScreen() {
  const router = useRouter();
  const insets = useSafeAreaInsets();
  const { q } = useLocalSearchParams<{ q?: string }>();
  const { filters } = useFilters();
  const { t: translate } = useTranslation();

  const [searchText, setSearchText] = useState(q ?? "");
  const [debouncedSearch] = useDebounce(searchText.toLowerCase(), 300);

  const {
    data: resp,
    isLoading,
    error,
  } = useSearchEvents(
    { q: debouncedSearch, limit: 5 },
    { query: { enabled: !!debouncedSearch } },
  );

  const results: Event[] = useMemo(() => {
    const d = resp as unknown as { data: Event[] } | undefined;
    return Array.isArray(d?.data) ? d.data : [];
  }, [resp]);

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
          // placeholder - we will change the SearchCard to accomodate event data instead
          renderItem={({ item }) => (
            <Text className="font-nunito text-sm text-[#111] p-3">
              {item.title}
            </Text>
          )}
          keyboardShouldPersistTaps="handled"
          keyboardDismissMode="on-drag"
        />
      )}
    </View>
  );
}
