import { ThemedText } from "@/components/themed-text";
import { ThemedView } from "@/components/themed-view";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { AppColors } from "@/constants/theme";
import { useMemo, useRef, useState } from "react";
import { useTranslation } from "react-i18next";
import { Animated, Image, Modal, PanResponder, Pressable, ScrollView, TouchableOpacity, View } from "react-native";
import { useQueryClient } from "@tanstack/react-query";
import { useSafeAreaInsets } from "react-native-safe-area-context";
import { useRouter } from "expo-router";
import {
  useGetAllEventOccurrences,
  useGetRegistrationsByGuardianId,
  useGetChildrenByGuardianId,
  useGetReviewByGuardianId,
  useCancelRegistration,
  getGetRegistrationsByGuardianIdQueryKey,
  type EventOccurrence,
  type Registration,
  type Child,
  type Review,
} from "@skillspark/api-client";

import { useAuthContext } from "@/hooks/use-auth-context";

type toggleValue = "upcoming" | "past" | undefined

interface ChildRegistration {
  child: Child
  registrationId: string
}

interface ToggleProps {
  value: toggleValue
  onChange: (newValue: toggleValue) => void
}

function Toggle({ value, onChange }: ToggleProps) {
  const { t } = useTranslation();
  return (
    <ThemedView className="w-11/12 border border-gray-200 rounded-md flex flex-row justify-between p-2 mt-2">
      <Pressable
        className={`${value === "upcoming" ? "bg-gray-200" : ""} w-6/12 flex items-center rounded-md py-1.5`}
        onPress={() => onChange("upcoming")}
      >
        <ThemedText>{t("activity.upcoming")}</ThemedText>
      </Pressable>

      <Pressable
        className={`${value === "past" ? "bg-gray-200" : ""} w-6/12 flex items-center rounded-md py-1.5`}
        onPress={() => onChange("past")}
      >
        <ThemedText>{t("activity.past")}</ThemedText>
      </Pressable>
    </ThemedView>
  );
}

interface RegistrationCardData {
  event_id: string
  event_occurrence_id: string
  image_uri: string
  start_time: Date
  end_time: Date
  title: string
  childRegistrations: ChildRegistration[]
  childColorMap: Record<string, string>
  location: string
  price: number
  hasReviewed: boolean
  onClickRemove: () => void
  onClickReview: () => void
}

interface RegistrationCardProps {
  data: RegistrationCardData
}

const formatTime = (d: Date) =>
  d.toLocaleTimeString([], { hour: "2-digit", minute: "2-digit" })
const formatDate = (d: Date) =>
  d.toLocaleDateString([], { weekday: "short", month: "short", day: "numeric" })

const CHILD_BUBBLE_COLORS = [
  "#E53E3E", "#DD6B20", "#D69E2E", "#276749", "#2B6CB0",
  "#553C9A", "#B83280", "#00695C", "#C53030", "#2C5282",
]

function getInitials(name: string) {
  return name.split(" ").map((n) => n[0]).join("").toUpperCase().slice(0, 2)
}

function UpcomingRegistrationCard({ data }: RegistrationCardProps) {
  const { t } = useTranslation();
  const children = data.childRegistrations.map((cr) => cr.child)

  return (
    <View
      className="w-11/12 rounded-xl overflow-hidden mb-4"
      style={{
        borderWidth: 1,
        borderColor: AppColors.borderLight,
        backgroundColor: AppColors.white,
        shadowColor: "#000",
        shadowOffset: { width: 0, height: 1 },
        shadowOpacity: 0.06,
        shadowRadius: 4,
        elevation: 2,
      }}
    >
      <Image
        source={{ uri: data.image_uri }}
        className="w-full h-44 px-4 pt-4 rounded-md"
      />
      <View className="flex flex-row justify-between items-center p-1">
        <View className="px-4 pb-4 gap-1 flex flex-col justify-between">
          <ThemedText type="subtitle">
            {data.title}
          </ThemedText>
          <View className="flex flex-row gap-2 items-center ">
            <IconSymbol name="clock" color="black" size={18}/>
            <ThemedText className="text-sm" >
              {formatTime(data.start_time)} – {formatTime(data.end_time)}
            </ThemedText>
          </View>
          <View className="flex flex-row gap-2 items-center ">
            <IconSymbol name="location" color="black" size={18}/>
            <ThemedText className="text-sm" >
              {data.location}
            </ThemedText>
          </View>
        </View>
        <View className="flex flex-col justify-center items-center bg-[#E69BF040] w-20 h-20 mr-2 rounded-md">
          <ThemedText type="subtitle" className="font-bold leading-none mr-1"> {data.start_time.getDate() < 10 ? "0" + data.start_time.getDate().toString() : data.start_time.getDate().toString()}</ThemedText>
          <ThemedText type="subtitle" className=" font-semibold tracking-widest "> {data.start_time.toLocaleString('default', { month: 'short' }) }</ThemedText>
        </View>
      </View>
      <View
        className="flex flex-row justify-between items-center px-4 py-3"
      >
        <View className="flex flex-row gap-2">
          {children.map((child) => {
            const color = data.childColorMap[child.id] ?? CHILD_BUBBLE_COLORS[0]
            return (
              <View
                key={child.id}
                className="w-8 h-8 rounded-full justify-center items-center"
                style={{ backgroundColor: `${color}33`, borderWidth: 1, borderColor: color }}
              >
                <ThemedText className="text-xs font-semibold" style={{ color }}>
                  {getInitials(child.name)}
                </ThemedText>
              </View>
            )
          })}
        </View>
        <TouchableOpacity
          onPress={data.onClickRemove}
          className="px-6 py-2 rounded-full bg-black"
          activeOpacity={0.7}
        >
          <ThemedText lightColor="white" className="text-sm font-medium">{t("activity.remove")}</ThemedText>
        </TouchableOpacity>
      </View>
    </View>
  );
}

function PastRegistrationCard({ data }: RegistrationCardProps) {
  const { t } = useTranslation();
  const children = data.childRegistrations.map((cr) => cr.child)
  const priceDisplay = `฿${(data.price / 100).toLocaleString()}`

  return (
    <View
      className="w-11/12 rounded-xl overflow-hidden mb-4 flex-row"
      style={{
        height: 150,
        borderWidth: 1,
        borderColor: AppColors.borderLight,
        backgroundColor: AppColors.white,
        shadowColor: "#000",
        shadowOffset: { width: 0, height: 1 },
        shadowOpacity: 0.06,
        shadowRadius: 4,
        elevation: 2,
      }}
    >
      <View className="py-3 pl-3">
        <Image
          source={{ uri: data.image_uri }}
          style={{ width: 100, flex: 1, borderRadius: 8 }}
          resizeMode="cover"
        />
      </View>

      <View className="flex-1 px-3 py-3 justify-between">
        <ThemedText type="subtitle" numberOfLines={2}>
          {data.title}
        </ThemedText>
        <View>
          <ThemedText className="text-sm" style={{ color: AppColors.mutedText }}>
            {formatDate(data.start_time)}
          </ThemedText>
          <ThemedText className="text-sm" style={{ color: AppColors.mutedText }}>
            {formatTime(data.start_time)} – {formatTime(data.end_time)}
          </ThemedText>
        </View>
        <ThemedText className="text-sm" style={{ color: AppColors.mutedText }}>
          {t("activity.payment", { price: priceDisplay })}
        </ThemedText>
      </View>

      <View className="py-3 pr-3 items-end justify-between">
        <View className="flex-row flex-wrap gap-1 justify-end" style={{ maxWidth: 80 }}>
          {children.map((child) => {
            const color = data.childColorMap[child.id] ?? CHILD_BUBBLE_COLORS[0]
            return (
              <View
                key={child.id}
                className="w-7 h-7 rounded-full justify-center items-center"
                style={{ backgroundColor: `${color}33`, borderWidth: 1, borderColor: color }}
              >
                <ThemedText className="text-xs font-semibold" style={{ color }}>
                  {getInitials(child.name)}
                </ThemedText>
              </View>
            )
          })}
        </View>
        {data.hasReviewed ? (
          <View className="px-4 py-2 rounded-full" style={{ backgroundColor: AppColors.borderLight }}>
            <ThemedText className="text-sm font-medium" style={{ color: AppColors.mutedText }}>{t("activity.reviewed")}</ThemedText>
          </View>
        ) : (
          <TouchableOpacity
            onPress={data.onClickReview}
            className="px-6 py-2 rounded-full bg-black"
            activeOpacity={0.7}
          >
            <ThemedText lightColor="white" className="text-sm font-medium">{t("activity.review")}</ThemedText>
          </TouchableOpacity>
        )}
      </View>
    </View>
  );
}

export default function ActivityScreen() {
  const insets = useSafeAreaInsets();
  const router = useRouter();
  const { t } = useTranslation();
  const [selection, setSelection] = useState<"upcoming" | "past">("upcoming");

  const { guardianId } = useAuthContext();
  // const guardianId = "11111111-1111-1111-1111-111111111111";

  const queryClient = useQueryClient()
  const { mutate: cancelRegistration } = useCancelRegistration()

  const [cancelTarget, setCancelTarget] = useState<ChildRegistration[] | null>(null)
  const [cancelSelections, setCancelSelections] = useState<Set<string>>(new Set())
  const cancelSheetTranslateY = useRef(new Animated.Value(0)).current

  const cancelPanResponder = useRef(
    PanResponder.create({
      onStartShouldSetPanResponder: () => true,
      onMoveShouldSetPanResponder: (_, gs) => gs.dy > 4,
      onPanResponderMove: (_, gs) => {
        if (gs.dy > 0) cancelSheetTranslateY.setValue(gs.dy)
      },
      onPanResponderRelease: (_, gs) => {
        if (gs.dy > 100) {
          setCancelTarget(null)
        } else {
          Animated.spring(cancelSheetTranslateY, { toValue: 0, useNativeDriver: true }).start()
        }
      },
    })
  ).current

  const getOnRemove = (childRegistrations: ChildRegistration[]) => {
    return function () {
      setCancelSelections(new Set(childRegistrations.map((cr) => cr.child.id)))
      cancelSheetTranslateY.setValue(0)
      setCancelTarget(childRegistrations)
    }
  }

  const toggleCancelSelection = (childId: string) => {
    setCancelSelections((prev) => {
      const next = new Set(prev)
      next.has(childId) ? next.delete(childId) : next.add(childId)
      return next
    })
  }

  const confirmCancellation = () => {
    if (!cancelTarget) return
    const toCancel = cancelTarget.filter((cr) => cancelSelections.has(cr.child.id))
    if (toCancel.length === 0) {
      setCancelTarget(null)
      return
    }
    const registrationIds = toCancel.map((cr) => cr.registrationId)
    setCancelTarget(null)
    const queryKey = getGetRegistrationsByGuardianIdQueryKey(guardianId)
    queryClient.setQueryData(queryKey, (old: unknown) => {
      const prev = old as { data: { registrations: Registration[] } } | undefined
      if (!prev?.data?.registrations) return old
      const idSet = new Set(registrationIds)
      return {
        ...prev,
        data: {
          ...prev.data,
          registrations: prev.data.registrations.map((r) =>
            idSet.has(r.id) ? { ...r, status: "cancelled" } : r
          ),
        },
      }
    })
    registrationIds.forEach((id) =>
      cancelRegistration(
        { id },
        {
          onError: () => {
            queryClient.invalidateQueries({ queryKey })
          },
        },
      )
    )
  }

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
    // console.log(d?.data?.registrations)
    return d?.data?.registrations ?? [];
  }, [registrationsResp]);

  const { data: occurrencesResp } = useGetAllEventOccurrences({ limit: 100 });
  const allOccurrences: EventOccurrence[] = useMemo(() => {
    const d = occurrencesResp as unknown as
      | { data: EventOccurrence[] }
      | undefined;
    return Array.isArray(d?.data) ? d.data : [];
  }, [occurrencesResp]);

  const eventOccurrencesMap: Record<string, EventOccurrence> = useMemo(() => {
    const map: Record<string, EventOccurrence> = {}
    allOccurrences.forEach((o) => {
      map[o.id] = o
    })
    return map
  }, [allOccurrences])

  const { data: childrenResp } = useGetChildrenByGuardianId(
    guardianId!,
    {
      query: { enabled: !!guardianId },
    },
  )

  const children: Child[] = useMemo(() => {
    const d = childrenResp as unknown as
      | { data: Child[] }
      | undefined;
    return Array.isArray(d?.data) ? d.data : [];
  }, [childrenResp]);

  const childMap = useMemo(() => {
    const map: Record<string, Child> = {}
    children.forEach((c) => {
      map[c.id] = c
    })
    return map
  }, [children])

  const childColorMap = useMemo(() => {
    const map: Record<string, string> = {}
    children.forEach((c, i) => {
      map[c.id] = CHILD_BUBBLE_COLORS[i % CHILD_BUBBLE_COLORS.length]
    })
    return map
  }, [children])

  const { data: guardianReviewsResp } = useGetReviewByGuardianId(
    guardianId!,
    undefined,
    { query: { enabled: !!guardianId } },
  )
  const reviewedEventIds = useMemo(() => {
    const list =
      guardianReviewsResp?.status === 200
        ? (guardianReviewsResp.data as Review[])
        : []
    return new Set(list.map((r) => r.event_id))
  }, [guardianReviewsResp])

  // Get upcoming and past registrations
  const { upcomingRegistrations, pastRegistrations } = useMemo(() => {
    const now = new Date()
    const grouped: Record<string, RegistrationCardData> = {}
    registrations.filter((r) => r.status === "registered").forEach((r) => {
      const occurrence = eventOccurrencesMap[r.event_occurrence_id]
      if (!occurrence) return
      const child = childMap[r.child_id]
      if (grouped[r.event_occurrence_id]) {
        if (child) grouped[r.event_occurrence_id].childRegistrations.push({ child, registrationId: r.id })
      } else {
        const startDate = new Date(occurrence.start_time)
        const endDate = new Date(occurrence.end_time)
        grouped[r.event_occurrence_id] = {
          event_id: occurrence.event.id,
          event_occurrence_id: r.event_occurrence_id,
          image_uri: "https://picsum.photos/800/300",
          start_time: startDate,
          end_time: endDate,
          title: occurrence.event.title,
          childRegistrations: child ? [{ child, registrationId: r.id }] : [],
          childColorMap,
          location: occurrence.location.address_line1,
          price: occurrence.price,
          hasReviewed: false,
          onClickRemove: () => {},
          onClickReview: () => {},
        }
      }
    })
    // Wire up callbacks and reviewed state after all child registrations are collected
    Object.values(grouped).forEach((g) => {
      g.hasReviewed = reviewedEventIds.has(g.event_id)
      g.onClickRemove = getOnRemove(g.childRegistrations)
      g.onClickReview = () => {
        router.push({
          pathname: "/(app)/(tabs)/activity/[id]/reviews",
          params: {
            id: g.event_id,
            canReview: "true",
            occurrenceId: g.event_occurrence_id,
            eventName: g.title,
            eventLocation: g.location,
            eventImageUrl: g.image_uri,
          },
        })
      }
    })
    const all = Object.values(grouped)
    return {
      upcomingRegistrations: all.filter((r) => r.start_time >= now).sort((a, b) => a.start_time.getTime() - b.start_time.getTime()),
      pastRegistrations: all.filter((r) => r.start_time < now).sort((a, b) => b.start_time.getTime() - a.start_time.getTime()),
    }
  }, [registrations, eventOccurrencesMap, childMap, reviewedEventIds])

  const toggleSelection = (newValue: toggleValue) => {
    setSelection(newValue!);
  };

  const [filterOpen, setFilterOpen] = useState(false)
  const [activeFilter, setActiveFilter] = useState<Set<string>>(new Set())
  const [pendingFilter, setPendingFilter] = useState<Set<string>>(new Set())

  const sheetTranslateY = useRef(new Animated.Value(0)).current

  const panResponder = useRef(
    PanResponder.create({
      onStartShouldSetPanResponder: () => true,
      onMoveShouldSetPanResponder: (_, gs) => gs.dy > 4,
      onPanResponderMove: (_, gs) => {
        if (gs.dy > 0) sheetTranslateY.setValue(gs.dy)
      },
      onPanResponderRelease: (_, gs) => {
        if (gs.dy > 100) {
          setFilterOpen(false)
        } else {
          Animated.spring(sheetTranslateY, { toValue: 0, useNativeDriver: true }).start()
        }
      },
    })
  ).current

  const openFilter = () => {
    const allIds = new Set(children.map((c) => c.id))
    setPendingFilter(activeFilter.size === 0 ? allIds : new Set(activeFilter))
    sheetTranslateY.setValue(0)
    setFilterOpen(true)
  }

  const togglePending = (id: string) => {
    setPendingFilter((prev) => {
      const next = new Set(prev)
      next.has(id) ? next.delete(id) : next.add(id)
      return next
    })
  }

  const applyFilter = () => {
    const allSelected = children.every((c) => pendingFilter.has(c.id))
    setActiveFilter(allSelected ? new Set() : new Set(pendingFilter))
    setFilterOpen(false)
  }

  const resetFilter = () => {
    setPendingFilter(new Set(children.map((c) => c.id)))
  }

  const baseDisplayed = selection === "upcoming" ? upcomingRegistrations : pastRegistrations;
  const displayed = activeFilter.size === 0
    ? baseDisplayed
    : baseDisplayed.filter((reg) =>
        reg.childRegistrations.some((cr) => activeFilter.has(cr.child.id))
      )

  const filterActive = activeFilter.size > 0

  return (
    <ThemedView className="w-full flex-1" style={{ paddingTop: insets.top }}>
      <ThemedView className="w-full flex items-center border-b border-black/[0.5]">
        <ThemedText className="py-3">{t("nav.activity")}</ThemedText>
      </ThemedView>

      <ThemedView className="w-full flex items-center">
        <Toggle value={selection} onChange={toggleSelection} />
      </ThemedView>

      {children.length > 0 && (
        <ThemedView className="w-11/12 self-center flex flex-row items-center justify-between py-3">
          <View className="flex flex-row flex-wrap gap-1.5">
            {children.map((child) => {
              const color = childColorMap[child.id] ?? CHILD_BUBBLE_COLORS[0]
              return (
                <View
                  key={child.id}
                  className="w-8 h-8 rounded-full justify-center items-center"
                  style={{ backgroundColor: `${color}33`, borderWidth: 1, borderColor: color }}
                >
                  <ThemedText className="text-xs font-semibold" style={{ color }}>
                    {getInitials(child.name)}
                  </ThemedText>
                </View>
              )
            })}
          </View>
          <TouchableOpacity onPress={openFilter} activeOpacity={0.7}>
            <IconSymbol
              name="line.3.horizontal.decrease"
              size={22}
              color={filterActive ? AppColors.primaryText ?? "#7C3AED" : "black"}
            />
          </TouchableOpacity>
        </ThemedView>
      )}

      <ScrollView
        contentContainerStyle={{ alignItems: "center", paddingTop: 16, paddingBottom: 32 }}
        showsVerticalScrollIndicator={false}
      >
        {displayed.length === 0 ? (
          <ThemedText className="mt-8" style={{ color: AppColors.mutedText }}>
            {selection === "upcoming" ? t("activity.noUpcomingRegistrations") : t("activity.noPastRegistrations")}
          </ThemedText>
        ) : (
          displayed.map((reg) =>
            selection === "upcoming"
              ? <UpcomingRegistrationCard key={reg.event_occurrence_id} data={reg} />
              : <PastRegistrationCard key={reg.event_occurrence_id} data={reg} />
          )
        )}
      </ScrollView>

      <Modal
        visible={cancelTarget !== null}
        transparent
        animationType="slide"
        onRequestClose={() => setCancelTarget(null)}
      >
        <Pressable
          className="flex-1"
          style={{ backgroundColor: "rgba(0,0,0,0.4)" }}
          onPress={() => setCancelTarget(null)}
        />
        <Animated.View
          style={{
            transform: [{ translateY: cancelSheetTranslateY }],
            backgroundColor: AppColors.white,
            borderTopLeftRadius: 20,
            borderTopRightRadius: 20,
            paddingBottom: insets.bottom + 16,
            paddingHorizontal: 24,
            paddingTop: 12,
          }}
        >
          <View {...cancelPanResponder.panHandlers} className="items-center pb-3">
            <View style={{ width: 36, height: 4, borderRadius: 2, backgroundColor: AppColors.borderLight }} />
          </View>

          <ThemedText type="defaultSemiBold" className="text-lg mb-1">{t("activity.cancelTitle")}</ThemedText>
          <ThemedText className="text-sm mb-5" style={{ color: AppColors.mutedText }}>
            {t("activity.cancelSubtitle")}
          </ThemedText>

          {(cancelTarget ?? []).map((cr) => {
            const color = childColorMap[cr.child.id] ?? CHILD_BUBBLE_COLORS[0]
            const checked = cancelSelections.has(cr.child.id)
            return (
              <Pressable
                key={cr.child.id}
                onPress={() => toggleCancelSelection(cr.child.id)}
                className="flex flex-row items-center justify-between py-3"
                style={{ borderBottomWidth: 1, borderBottomColor: AppColors.borderLight }}
              >
                <View className="flex flex-row items-center gap-3">
                  <View
                    className="w-9 h-9 rounded-full justify-center items-center"
                    style={{ backgroundColor: `${color}33`, borderWidth: 1, borderColor: color }}
                  >
                    <ThemedText className="text-xs font-semibold" style={{ color }}>
                      {getInitials(cr.child.name)}
                    </ThemedText>
                  </View>
                  <ThemedText className="text-base">{cr.child.name}</ThemedText>
                </View>
                <View
                  className="w-5 h-5 rounded border justify-center items-center"
                  style={{
                    borderColor: checked ? "#000" : AppColors.borderLight,
                    backgroundColor: checked ? "#000" : "transparent",
                  }}
                >
                  {checked && <ThemedText lightColor="white" className="text-xs leading-none">✓</ThemedText>}
                </View>
              </Pressable>
            )
          })}

          <View className="flex flex-row gap-3 mt-5">
            <TouchableOpacity
              onPress={() => setCancelTarget(null)}
              activeOpacity={0.7}
              className="flex-1 py-3 rounded-xl items-center"
              style={{ borderWidth: 1, borderColor: AppColors.borderLight }}
            >
              <ThemedText className="font-medium">{t("activity.keep")}</ThemedText>
            </TouchableOpacity>
            <TouchableOpacity
              onPress={confirmCancellation}
              activeOpacity={0.7}
              className="flex-1 py-3 rounded-xl items-center"
              style={{ backgroundColor: cancelSelections.size === 0 ? AppColors.borderLight : "#EF4444" }}
              disabled={cancelSelections.size === 0}
            >
              <ThemedText lightColor="white" className="font-medium">
                {t("activity.remove")}{cancelSelections.size > 0 ? ` (${cancelSelections.size})` : ""}
              </ThemedText>
            </TouchableOpacity>
          </View>
        </Animated.View>
      </Modal>

      <Modal
        visible={filterOpen}
        transparent
        animationType="slide"
        onRequestClose={() => setFilterOpen(false)}
      >
        <Pressable
          className="flex-1"
          style={{ backgroundColor: "rgba(0,0,0,0.4)" }}
          onPress={() => setFilterOpen(false)}
        />
        <Animated.View
          style={{
            transform: [{ translateY: sheetTranslateY }],
            backgroundColor: AppColors.white,
            borderTopLeftRadius: 20,
            borderTopRightRadius: 20,
            paddingBottom: insets.bottom + 16,
            paddingHorizontal: 24,
            paddingTop: 12,
          }}
        >
          <View {...panResponder.panHandlers} className="items-center pb-3">
            <View
              style={{
                width: 36,
                height: 4,
                borderRadius: 2,
                backgroundColor: AppColors.borderLight,
              }}
            />
          </View>

          <View className="flex flex-row items-center gap-2 mb-5">
            <IconSymbol name="line.3.horizontal.decrease" size={20} color="black" />
            <ThemedText type="defaultSemiBold" className="text-lg">{t("activity.filterByChild")}</ThemedText>
          </View>

          {children.map((child) => {
            const color = childColorMap[child.id] ?? CHILD_BUBBLE_COLORS[0]
            const checked = pendingFilter.has(child.id)
            return (
              <Pressable
                key={child.id}
                onPress={() => togglePending(child.id)}
                className="flex flex-row items-center justify-between py-3"
                style={{ borderBottomWidth: 1, borderBottomColor: AppColors.borderLight }}
              >
                <View className="flex flex-row items-center gap-3">
                  <View
                    className="w-9 h-9 rounded-full justify-center items-center"
                    style={{ backgroundColor: `${color}33`, borderWidth: 1, borderColor: color }}
                  >
                    <ThemedText className="text-xs font-semibold" style={{ color }}>
                      {getInitials(child.name)}
                    </ThemedText>
                  </View>
                  <ThemedText className="text-base">{child.name}</ThemedText>
                </View>
                <View
                  className="w-5 h-5 rounded border justify-center items-center"
                  style={{
                    borderColor: checked ? "#000" : AppColors.borderLight,
                    backgroundColor: checked ? "#000" : "transparent",
                  }}
                >
                  {checked && <ThemedText lightColor="white" className="text-xs leading-none">✓</ThemedText>}
                </View>
              </Pressable>
            )
          })}

          <View className="flex flex-row gap-3 mt-5">
            <TouchableOpacity
              onPress={resetFilter}
              activeOpacity={0.7}
              className="flex-1 py-3 rounded-xl items-center"
              style={{ borderWidth: 1, borderColor: AppColors.borderLight }}
            >
              <ThemedText className="font-medium">{t("activity.reset")}</ThemedText>
            </TouchableOpacity>
            <TouchableOpacity
              onPress={applyFilter}
              activeOpacity={0.7}
              className="flex-1 py-3 rounded-xl items-center bg-black"
            >
              <ThemedText lightColor="white" className="font-medium">{t("activity.apply")}</ThemedText>
            </TouchableOpacity>
          </View>
        </Animated.View>
      </Modal>
    </ThemedView>
  );
}
