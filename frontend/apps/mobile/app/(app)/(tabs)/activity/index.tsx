import { ThemedText } from "@/components/themed-text";
import { ThemedView } from "@/components/themed-view";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { AppColors } from "@/constants/theme";
import { useMemo, useRef, useState } from "react";
import { useTranslation } from "react-i18next";
import {
  Animated,
  Modal,
  PanResponder,
  Pressable,
  ScrollView,
  TouchableOpacity,
  View,
} from "react-native";
import { useQueryClient } from "@tanstack/react-query";
import { useSafeAreaInsets } from "react-native-safe-area-context";
import { useRouter } from "expo-router";
import {
  useCancelRegistration,
  getGetRegistrationsByGuardianIdQueryKey,
  type Registration,
} from "@skillspark/api-client";
import { UpcomingRegistrationCard } from "@/components/UpcomingRegistrationCard";
import { PastRegistrationCard } from "@/components/PastRegistrationCard";
import {
  type ChildRegistration,
  type RegistrationCardData,
} from "@/components/RegistrationCard.types";
import { ChildAvatar } from "@/components/ChildAvatar";
import { useActivityData } from "@/hooks/use-activity-data";
import LogoBgWrapper from "@/components/LogoBgWrapper";

type toggleValue = "upcoming" | "past" | undefined;

interface ToggleProps {
  value: toggleValue;
  onChange: (newValue: toggleValue) => void;
}

function Toggle({ value, onChange }: ToggleProps) {
  const { t } = useTranslation();
  return (
    <View className="w-11/12 border bg-white border-gray-200 rounded-md flex flex-row justify-between p-2 mt-2">
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
    </View>
  );
}

export default function ActivityScreen() {
  const insets = useSafeAreaInsets();
  const router = useRouter();
  const { t } = useTranslation();
  const [selection, setSelection] = useState<"upcoming" | "past">("upcoming");

  const {
    guardianId,
    registrations,
    eventOccurrencesMap,
    children,
    childMap,
    reviewedEventIds,
  } = useActivityData();

  const queryClient = useQueryClient();
  const { mutate: cancelRegistration } = useCancelRegistration();

  // ── Cancel sheet ─────────────────────────────────────────────────────────────

  const [cancelTarget, setCancelTarget] = useState<ChildRegistration[] | null>(
    null,
  );
  const [cancelSelections, setCancelSelections] = useState<Set<string>>(
    new Set(),
  );
  const cancelSheetTranslateY = useRef(new Animated.Value(0)).current;

  const cancelPanResponder = useRef(
    PanResponder.create({
      onStartShouldSetPanResponder: () => true,
      onMoveShouldSetPanResponder: (_, gs) => gs.dy > 4,
      onPanResponderMove: (_, gs) => {
        if (gs.dy > 0) cancelSheetTranslateY.setValue(gs.dy);
      },
      onPanResponderRelease: (_, gs) => {
        if (gs.dy > 100) {
          setCancelTarget(null);
        } else {
          Animated.spring(cancelSheetTranslateY, {
            toValue: 0,
            useNativeDriver: true,
          }).start();
        }
      },
    }),
  ).current;

  const resetCancelSheet = () => cancelSheetTranslateY.setValue(0);

  const getOnRemove = (childRegistrations: ChildRegistration[]) => () => {
    setCancelSelections(new Set(childRegistrations.map((cr) => cr.child.id)));
    resetCancelSheet();
    setCancelTarget(childRegistrations);
  };

  const toggleCancelSelection = (childId: string) => {
    setCancelSelections((prev) => {
      const next = new Set(prev);
      next.has(childId) ? next.delete(childId) : next.add(childId);
      return next;
    });
  };

  const confirmCancellation = () => {
    if (!cancelTarget) return;
    const toCancel = cancelTarget.filter((cr) =>
      cancelSelections.has(cr.child.id),
    );
    if (toCancel.length === 0) {
      setCancelTarget(null);
      return;
    }
    const registrationIds = toCancel.map((cr) => cr.registrationId);
    setCancelTarget(null);
    const queryKey = getGetRegistrationsByGuardianIdQueryKey(guardianId!);
    queryClient.setQueryData(queryKey, (old: unknown) => {
      const prev = old as
        | { data: { registrations: Registration[] } }
        | undefined;
      if (!prev?.data?.registrations) return old;
      const idSet = new Set(registrationIds);
      return {
        ...prev,
        data: {
          ...prev.data,
          registrations: prev.data.registrations.map((r) =>
            idSet.has(r.id) ? { ...r, status: "cancelled" } : r,
          ),
        },
      };
    });
    registrationIds.forEach((id) =>
      cancelRegistration(
        { id },
        {
          onError: () => {
            queryClient.invalidateQueries({ queryKey });
          },
        },
      ),
    );
  };

  // ── Registration card data ───────────────────────────────────────────────────

  const { upcomingRegistrations, pastRegistrations } = useMemo(() => {
    const now = new Date();
    const grouped: Record<string, RegistrationCardData> = {};

    registrations
      .filter((r) => r.status === "registered")
      .forEach((r) => {
        const occurrence = eventOccurrencesMap[r.event_occurrence_id];
        if (!occurrence) return;
        const child = childMap[r.child_id];
        if (grouped[r.event_occurrence_id]) {
          if (child)
            grouped[r.event_occurrence_id].childRegistrations.push({
              child,
              registrationId: r.id,
            });
        } else {
          grouped[r.event_occurrence_id] = {
            event_id: occurrence.event.id,
            event_occurrence_id: r.event_occurrence_id,
            image_uri: occurrence.event.presigned_url,
            start_time: new Date(occurrence.start_time),
            end_time: new Date(occurrence.end_time),
            title: occurrence.event.title,
            childRegistrations: child ? [{ child, registrationId: r.id }] : [],
            location: occurrence.location.address_line1,
            price: occurrence.price / 100,
            hasReviewed: false,
            onClickRemove: () => {},
            onClickReview: () => {},
          };
        }
      });

    Object.values(grouped).forEach((g) => {
      g.hasReviewed = reviewedEventIds.has(g.event_id);
      g.onClickRemove = getOnRemove(g.childRegistrations);
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
        });
      };
    });

    const all = Object.values(grouped);
    return {
      upcomingRegistrations: all
        .filter((r) => r.start_time >= now)
        .sort((a, b) => a.start_time.getTime() - b.start_time.getTime()),
      pastRegistrations: all
        .filter((r) => r.start_time < now)
        .sort((a, b) => b.start_time.getTime() - a.start_time.getTime()),
    };
  }, [registrations, eventOccurrencesMap, childMap, reviewedEventIds]);

  // ── Filter sheet ─────────────────────────────────────────────────────────────

  const [filterOpen, setFilterOpen] = useState(false);
  const [activeFilter, setActiveFilter] = useState<Set<string>>(new Set());
  const [pendingFilter, setPendingFilter] = useState<Set<string>>(new Set());
  const sheetTranslateY = useRef(new Animated.Value(0)).current;

  const panResponder = useRef(
    PanResponder.create({
      onStartShouldSetPanResponder: () => true,
      onMoveShouldSetPanResponder: (_, gs) => gs.dy > 4,
      onPanResponderMove: (_, gs) => {
        if (gs.dy > 0) sheetTranslateY.setValue(gs.dy);
      },
      onPanResponderRelease: (_, gs) => {
        if (gs.dy > 100) {
          setFilterOpen(false);
        } else {
          Animated.spring(sheetTranslateY, {
            toValue: 0,
            useNativeDriver: true,
          }).start();
        }
      },
    }),
  ).current;

  const resetFilterSheet = () => sheetTranslateY.setValue(0);

  const openFilter = () => {
    const allIds = new Set(children.map((c) => c.id));
    setPendingFilter(activeFilter.size === 0 ? allIds : new Set(activeFilter));
    resetFilterSheet();
    setFilterOpen(true);
  };

  const togglePending = (id: string) => {
    setPendingFilter((prev) => {
      const next = new Set(prev);
      next.has(id) ? next.delete(id) : next.add(id);
      return next;
    });
  };

  const applyFilter = () => {
    const allSelected = children.every((c) => pendingFilter.has(c.id));
    setActiveFilter(allSelected ? new Set() : new Set(pendingFilter));
    setFilterOpen(false);
  };

  const resetFilter = () => {
    setPendingFilter(new Set(children.map((c) => c.id)));
  };

  const filterActive = activeFilter.size > 0;

  // ── Displayed list ───────────────────────────────────────────────────────────

  const baseDisplayed =
    selection === "upcoming" ? upcomingRegistrations : pastRegistrations;
  const displayed =
    activeFilter.size === 0
      ? baseDisplayed
      : baseDisplayed.filter((reg) =>
          reg.childRegistrations.some((cr) => activeFilter.has(cr.child.id)),
        );

  const toggleSelection = (newValue: toggleValue) => {
    setSelection(newValue!);
  };

  return (
    <ThemedView className="w-full flex-1" style={{ paddingTop: insets.top }}>
      <LogoBgWrapper className="flex-1">
        <View className="w-full flex items-center">
          <ThemedText className="py-3">{t("nav.activity")}</ThemedText>
        </View>

        <View className="w-full flex items-center">
          <Toggle value={selection} onChange={toggleSelection} />
        </View>

        {children.length > 0 && (
          <View className="w-11/12 self-center flex flex-row items-center justify-between py-3">
            <View className="flex flex-row flex-wrap gap-1.5">
              {children.map((child) => (
                <ChildAvatar
                  key={child.id}
                  name={child.name}
                  avatarFace={child.avatar_face}
                  avatarBackground={child.avatar_background}
                  size={32}
                />
              ))}
            </View>
            <TouchableOpacity onPress={openFilter} activeOpacity={0.7}>
              <IconSymbol
                name="line.3.horizontal.decrease"
                size={22}
                color={
                  filterActive ? (AppColors.primaryText ?? "#7C3AED") : "black"
                }
              />
            </TouchableOpacity>
          </View>
        )}

        <ScrollView
          contentContainerStyle={{
            alignItems: "center",
            paddingTop: 16,
            paddingBottom: 32,
          }}
          showsVerticalScrollIndicator={false}
        >
          {displayed.length === 0 ? (
            <ThemedText className="mt-8" style={{ color: AppColors.mutedText }}>
              {selection === "upcoming"
                ? t("activity.noUpcomingRegistrations")
                : t("activity.noPastRegistrations")}
            </ThemedText>
          ) : (
            displayed.map((reg) =>
              selection === "upcoming" ? (
                <UpcomingRegistrationCard
                  key={reg.event_occurrence_id}
                  data={reg}
                />
              ) : (
                <PastRegistrationCard key={reg.event_occurrence_id} data={reg} />
              ),
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
              borderTopLeftRadius: 20,
              borderTopRightRadius: 20,
              paddingBottom: insets.bottom + 16,
              paddingHorizontal: 24,
              paddingTop: 12,
            }}
          >
            <View
              {...cancelPanResponder.panHandlers}
              className="items-center pb-3"
            >
              <View
                style={{
                  width: 36,
                  height: 4,
                  borderRadius: 2,
                  backgroundColor: AppColors.borderLight,
                }}
              />
            </View>

            <ThemedText type="defaultSemiBold" className="text-lg mb-1">
              {t("activity.cancelTitle")}
            </ThemedText>
            <ThemedText
              className="text-sm mb-5"
              style={{ color: AppColors.mutedText }}
            >
              {t("activity.cancelSubtitle")}
            </ThemedText>

            {(cancelTarget ?? []).map((cr) => {
              const checked = cancelSelections.has(cr.child.id);
              return (
                <Pressable
                  key={cr.child.id}
                  onPress={() => toggleCancelSelection(cr.child.id)}
                  className="flex flex-row items-center justify-between py-3"
                  style={{
                    borderBottomWidth: 1,
                    borderBottomColor: AppColors.borderLight,
                  }}
                >
                  <View className="flex flex-row items-center gap-3">
                    <ChildAvatar
                      name={cr.child.name}
                      avatarFace={cr.child.avatar_face}
                      avatarBackground={cr.child.avatar_background}
                      size={36}
                    />
                    <ThemedText className="text-base">{cr.child.name}</ThemedText>
                  </View>
                  <View
                    className="w-5 h-5 rounded border justify-center items-center"
                    style={{
                      borderColor: checked ? "#000" : AppColors.borderLight,
                      backgroundColor: checked ? "#000" : "transparent",
                    }}
                  >
                    {checked && (
                      <ThemedText
                        lightColor="white"
                        className="text-xs leading-none"
                      >
                        ✓
                      </ThemedText>
                    )}
                  </View>
                </Pressable>
              );
            })}

            <View className="flex flex-row gap-3 mt-5">
              <TouchableOpacity
                onPress={() => setCancelTarget(null)}
                activeOpacity={0.7}
                className="flex-1 py-3 rounded-xl items-center"
                style={{ borderWidth: 1, borderColor: AppColors.borderLight }}
              >
                <ThemedText className="font-medium">
                  {t("activity.keep")}
                </ThemedText>
              </TouchableOpacity>
              <TouchableOpacity
                onPress={confirmCancellation}
                activeOpacity={0.7}
                className="flex-1 py-3 rounded-xl items-center"
                style={{
                  backgroundColor:
                    cancelSelections.size === 0
                      ? AppColors.borderLight
                      : "#EF4444",
                }}
                disabled={cancelSelections.size === 0}
              >
                <ThemedText lightColor="white" className="font-medium">
                  {t("activity.remove")}
                  {cancelSelections.size > 0 ? ` (${cancelSelections.size})` : ""}
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
              <IconSymbol
                name="line.3.horizontal.decrease"
                size={20}
                color="black"
              />
              <ThemedText type="defaultSemiBold" className="text-lg">
                {t("activity.filterByChild")}
              </ThemedText>
            </View>

            {children.map((child) => {
              const checked = pendingFilter.has(child.id);
              return (
                <Pressable
                  key={child.id}
                  onPress={() => togglePending(child.id)}
                  className="flex flex-row items-center justify-between py-3"
                  style={{
                    borderBottomWidth: 1,
                    borderBottomColor: AppColors.borderLight,
                  }}
                >
                  <View className="flex flex-row items-center gap-3">
                    <ChildAvatar
                      name={child.name}
                      avatarFace={child.avatar_face}
                      avatarBackground={child.avatar_background}
                      size={36}
                    />
                    <ThemedText className="text-base">{child.name}</ThemedText>
                  </View>
                  <View
                    className="w-5 h-5 rounded border justify-center items-center"
                    style={{
                      borderColor: checked ? "#000" : AppColors.borderLight,
                      backgroundColor: checked ? "#000" : "transparent",
                    }}
                  >
                    {checked && (
                      <ThemedText
                        lightColor="white"
                        className="text-xs leading-none"
                      >
                        ✓
                      </ThemedText>
                    )}
                  </View>
                </Pressable>
              );
            })}

            <View className="flex flex-row gap-3 mt-5">
              <TouchableOpacity
                onPress={resetFilter}
                activeOpacity={0.7}
                className="flex-1 py-3 rounded-xl items-center"
                style={{ borderWidth: 1, borderColor: AppColors.borderLight }}
              >
                <ThemedText className="font-medium">
                  {t("activity.reset")}
                </ThemedText>
              </TouchableOpacity>
              <TouchableOpacity
                onPress={applyFilter}
                activeOpacity={0.7}
                className="flex-1 py-3 rounded-xl items-center bg-black"
              >
                <ThemedText lightColor="white" className="font-medium">
                  {t("activity.apply")}
                </ThemedText>
              </TouchableOpacity>
            </View>
          </Animated.View>
        </Modal>
      </LogoBgWrapper>
    </ThemedView>
  );
}
