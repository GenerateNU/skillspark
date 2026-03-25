import { Ionicons } from "@expo/vector-icons";
import {
  ActivityIndicator,
  View,
  TextInput,
  ScrollView,
  Pressable,
  Text,
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
} from "@skillspark/api-client";
import { useMemo, useState } from "react";
import { AppColors, FontSizes } from "@/constants/theme";
import { useAuthContext } from "@/hooks/use-auth-context";
import { useDebounce } from "use-debounce";
import { isWithinNext7Days } from "@/utils/format";
import { DiscoverBanner } from "@/components/home/DiscoverBanner";
import { UpcomingClassCard } from "@/components/home/UpcomingClassCard";
import { TrendingCard } from "@/components/home/TrendingCard";
import { RecommendedCard } from "@/components/home/RecommendedCard";
import { CategoryCard } from "@/components/home/CategoryCard";

export default function HomeScreen() {
  const { guardianId } = useAuthContext();
  const [searchText, setSearchText] = useState("");
  const [_debouncedSearch] = useDebounce(searchText, 300);

  const { data: guardianResp } = useGetGuardianById(guardianId!, {
    query: { enabled: !!guardianId },
  });
  const guardian = (guardianResp as unknown as { data: Guardian } | undefined)?.data;

  const { data: occurrencesResp, isLoading } = useGetAllEventOccurrences();
  const allOccurrences: EventOccurrence[] = useMemo(() => {
    const d = occurrencesResp as unknown as { data: EventOccurrence[] } | undefined;
    return Array.isArray(d?.data) ? d.data : [];
  }, [occurrencesResp]);

  const { data: registrationsResp } = useGetRegistrationsByGuardianId(guardianId!, {
    query: { enabled: !!guardianId },
  });
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

  const upcomingClasses = useMemo(() => {
    const upcomingIds = new Set(
      registrations
        .filter((r) => r.status === "registered" && isWithinNext7Days(r.occurrence_start_time))
        .map((r) => r.event_occurrence_id)
    );
    return allOccurrences.filter((o) => upcomingIds.has(o.id));
  }, [registrations, allOccurrences]);

  const futureOccurrences = useMemo(
    () => allOccurrences.filter((o) => new Date(o.start_time) > new Date()),
    [allOccurrences]
  );

  const trendingEvents = useMemo(
    () => [...futureOccurrences].sort(() => Math.random() - 0.5).slice(0, 5),
    [futureOccurrences]
  );

  const childRecommendations = useMemo(() => {
    const shuffled = [...futureOccurrences].sort(() => Math.random() - 0.5);
    return children.map((child, i) => ({
      child,
      occurrence: shuffled[i % shuffled.length],
    })).filter((r) => r.occurrence != null);
  }, [children, futureOccurrences]);

  const categories = useMemo(() => {
    const cats = new Set<string>();
    allOccurrences.forEach((o) => o.event.category?.forEach((c) => cats.add(c)));
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
      <View style={{ flex: 1, alignItems: "center", justifyContent: "center", backgroundColor: AppColors.white }}>
        <ActivityIndicator size="large" />
      </View>
    );
  }

  const categoryPairs: string[][] = [];
  for (let i = 0; i < categories.length; i += 2) {
    categoryPairs.push(categories.slice(i, i + 2));
  }

  return (
    <ScrollView
      style={{ flex: 1, backgroundColor: AppColors.white }}
      showsVerticalScrollIndicator={false}
      contentContainerStyle={{ paddingBottom: 40 }}
    >
      {/* Header */}
      <View style={{ paddingHorizontal: 20, paddingTop: 56, paddingBottom: 16 }}>
        <Text style={{ fontSize: FontSizes.hero, fontFamily: "NunitoSans_700Bold", color: AppColors.primaryText, letterSpacing: -0.5 }}>
          Hello, {firstName}
        </Text>
      </View>

      {/* Search row */}
      <View style={{ paddingHorizontal: 20, marginBottom: 22 }}>
        <View
          style={{
            flexDirection: "row",
            alignItems: "center",
            backgroundColor: AppColors.surfaceGray,
            borderRadius: 50,
            paddingHorizontal: 16,
            paddingVertical: 10,
          }}
        >
          <Ionicons name="search" size={18} color={AppColors.subtleText} style={{ marginRight: 8 }} />
          <TextInput
            style={{ flex: 1, fontSize: FontSizes.base, color: AppColors.primaryText, fontFamily: "NunitoSans_400Regular" }}
            placeholder="Search for a class"
            placeholderTextColor={AppColors.placeholderText}
            value={searchText}
            onChangeText={setSearchText}
          />
          <Pressable
            style={{
              width: 36,
              height: 36,
              borderRadius: 18,
              backgroundColor: AppColors.primaryText,
              alignItems: "center",
              justifyContent: "center",
            }}
          >
            <Ionicons name="options" size={16} color={AppColors.white} />
          </Pressable>
        </View>
      </View>

      {/* Your Upcoming Classes — conditional */}
      {upcomingClasses.length > 0 && (
        <View style={{ marginBottom: 24 }}>
          <Text style={{ fontSize: FontSizes.lg, fontFamily: "NunitoSans_700Bold", color: AppColors.primaryText, paddingHorizontal: 20, marginBottom: 12 }}>
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
        <View style={{ marginBottom: 24 }}>
          <View style={{ flexDirection: "row", alignItems: "center", paddingHorizontal: 20, marginBottom: 12 }}>
            <Text style={{ color: AppColors.purple, fontSize: FontSizes.md, marginRight: 6, fontFamily: "NunitoSans_400Regular" }}>✦</Text>
            <Text style={{ fontSize: FontSizes.lg, fontFamily: "NunitoSans_700Bold", color: AppColors.primaryText }}>Discover Weekly</Text>
          </View>
          <DiscoverBanner event={futureOccurrences[0]} />
        </View>
      )}

      {/* Trending In Your Area */}
      {trendingEvents.length > 0 && (
        <View style={{ marginBottom: 24 }}>
          <Text style={{ fontSize: FontSizes.lg, fontFamily: "NunitoSans_700Bold", color: AppColors.primaryText, paddingHorizontal: 20, marginBottom: 12 }}>
            Trending in Your Area
          </Text>
          <ScrollView
            horizontal
            showsHorizontalScrollIndicator={false}
            contentContainerStyle={{ paddingHorizontal: 20, paddingTop: 0 }}
          >
            {trendingEvents.map((o, i) => (
              <TrendingCard key={o.id} occurrence={o} index={i} />
            ))}
          </ScrollView>
        </View>
      )}

      {/* Recommended For... */}
      {childRecommendations.length > 0 && (
        <View style={{ marginBottom: 24 }}>
          <Text style={{ fontSize: FontSizes.lg, fontFamily: "NunitoSans_700Bold", color: AppColors.primaryText, paddingHorizontal: 20, marginBottom: 12 }}>
            Recommended for...
          </Text>
          <ScrollView
            horizontal
            showsHorizontalScrollIndicator={false}
            contentContainerStyle={{ paddingHorizontal: 20 }}
          >
            {childRecommendations.map(({ child, occurrence }) => (
              <RecommendedCard key={child.id} occurrence={occurrence} childName={child.name.split(" ")[0]} />
            ))}
          </ScrollView>
        </View>
      )}

      {/* Explore by Category */}
      {categories.length > 0 && (
        <View style={{ marginBottom: 24 }}>
          <Text style={{ fontSize: FontSizes.lg, fontFamily: "NunitoSans_700Bold", color: AppColors.primaryText, paddingHorizontal: 20, marginBottom: 12 }}>
            Explore by Category
          </Text>
          <View style={{ paddingHorizontal: 15 }}>
            {categoryPairs.map((pair, idx) => (
              <View key={idx} style={{ flexDirection: "row" }}>
                {pair.map((cat) => (
                  <CategoryCard key={cat} category={cat} occurrence={categoryEventMap[cat]} />
                ))}
                {pair.length === 1 && <View style={{ flex: 1, margin: 5 }} />}
              </View>
            ))}
          </View>
        </View>
      )}
    </ScrollView>
  );
}
