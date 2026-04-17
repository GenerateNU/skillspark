import React from "react";
import {
  View,
  TouchableOpacity,
  ScrollView,
  Text,
  ActivityIndicator,
} from "react-native";
import { Image } from "expo-image";
import { Stack, useRouter, useLocalSearchParams } from "expo-router";
import { useSafeAreaInsets } from "react-native-safe-area-context";
import { ThemedView } from "@/components/themed-view";
import { Colors, AppColors, FontSizes, FontFamilies } from "@/constants/theme";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { ErrorScreen } from "@/components/ErrorScreen";
import { useGetAllEventOccurrences } from "@skillspark/api-client";
import { useTranslation } from "react-i18next";
import { useGeolocation } from "@/hooks/use-geolocation";
import { EventCategoriesListItem } from "@/components/EventCategoriesListItem";

const CATEGORY_DB_VALUES: Record<string, string> = {
  "Sports & Physical Activities":
    "sports,soccer,basketball,tennis,swimming,martial arts,yoga,fitness,running,hiking,cycling",
  "Arts & Creative Expression":
    "art,painting,drawing,sculpture,photography,filmmaking,graphic design,fashion,crafts,creative writing",
  Languages: "language,language learning,linguistics",
  Academics: "science,physics,chemistry,biology,astronomy,earth science,history",
  "Personal Development & Life Skills": "other",
  "Music & Performance":
    "music,instrumental music,vocal music,theater,acting,dance,improv",
  Math: "math",
  "Tech & Innovation":
    "technology,robotics,engineering,coding,data science,statistics",
};

export default function EventCategoryScreen() {
  const router = useRouter();
  const params = useLocalSearchParams();
  const insets = useSafeAreaInsets();
  const theme = Colors.light;

  const { t: translate } = useTranslation();
  const category = params.category as string;
  const { lat: geoLocationLat, lng: geoLocationLong } = useGeolocation();

  const dbCategory = CATEGORY_DB_VALUES[category] ?? category;

  const { data: localizedOccurrencesResp, isLoading } =
    useGetAllEventOccurrences({
      category: dbCategory,

      lat: geoLocationLat,
      lng: geoLocationLong,
      radius_km: 50,
      limit: 20,
      soldout: false,
    });

  if (!category) {
    return (
      <ErrorScreen message={translate("eventCategories.noCategorySupplied")} />
    );
  }

  if (!geoLocationLat || !geoLocationLong) {
    return (
      <ErrorScreen message={translate("eventCategories.noLocationFound")} />
    );
  }

  // deduplicates events
  const occurrences =
    localizedOccurrencesResp?.status === 200
      ? localizedOccurrencesResp.data.filter(
          (o, idx, arr) =>
            arr.findIndex((x) => x.event.id === o.event.id) === idx
        )
      : [];
  const displayCategory =
    category.charAt(0).toUpperCase() + category.slice(1);

  return (
    <ThemedView className="flex-1" style={{ paddingTop: insets.top }}>
      <Stack.Screen options={{ headerShown: false }} />

      {/* Header */}
      <View
        style={{
          paddingHorizontal: 20,
          paddingTop: 10,
          paddingBottom: 16,
          overflow: "hidden",
        }}
      >
        <TouchableOpacity
          onPress={() => router.back()}
          style={{
            width: 32,
            height: 32,
            justifyContent: "center",
            alignItems: "flex-start",
          }}
        >
          <IconSymbol name="chevron.left" size={24} color={theme.text} />
        </TouchableOpacity>
        <Text
          style={{
            fontFamily: FontFamilies.bold,
            fontSize: 32,
            color: AppColors.primaryText,
            marginTop: 12,
          }}
        >
          {displayCategory}
        </Text>
      </View>

      {isLoading ? (
        <ActivityIndicator
          size="large"
          style={{ marginTop: 40 }}
          color={AppColors.primaryText}
        />
      ) : occurrences.length === 0 ? (
        <Text
          style={{
            fontFamily: FontFamilies.bold,
            fontSize: FontSizes.lg,
            color: AppColors.primaryText,
            paddingHorizontal: 20,
            marginTop: 20,
          }}
        >
          {translate("eventCategories.noEventsFound")}
        </Text>
      ) : (
        <ScrollView
          contentContainerStyle={{ paddingHorizontal: 20, paddingBottom: 40 }}
          showsVerticalScrollIndicator={false}
        >
          {occurrences.map((o) => (
            <EventCategoriesListItem
              key={o.id}
              eventOccurrence={o}
              onPress={() =>
                router.push({
                  pathname: "/event/[id]",
                  params: {
                    id: o.event.id,
                    from: "event-categories",
                    category,
                  },
                })
              }
            />
          ))}
        </ScrollView>
      )}
    </ThemedView>
  );
}
