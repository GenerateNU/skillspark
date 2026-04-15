import React, { useState, useRef } from "react";
import {
  View,
  Text,
  TouchableOpacity,
  TextInput,
  ScrollView,
  Platform,
  KeyboardAvoidingView,
} from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";
import { useRouter, useLocalSearchParams } from "expo-router";
import Slider from "@react-native-community/slider";
import DateTimePicker, {
  type DateTimePickerEvent,
} from "@react-native-community/datetimepicker";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { ThemedText } from "@/components/themed-text";
import { AppColors } from "@/constants/theme";
import { useThemeColor } from "@/hooks/use-theme-color";

const MAX_DISTANCE = 40;
const MAX_AGE = 12;

function minutesToDate(minutes: number): Date {
  const d = new Date();
  d.setHours(Math.floor(minutes / 60), minutes % 60, 0, 0);
  return d;
}

function dateToMinutes(date: Date): number {
  return date.getHours() * 60 + date.getMinutes();
}

function formatTime(minutes: number): string {
  const h = Math.floor(minutes / 60);
  const m = minutes % 60;
  const ampm = h >= 12 ? "PM" : "AM";
  const displayH = h % 12 === 0 ? 12 : h % 12;
  return `${displayH}:${m.toString().padStart(2, "0")} ${ampm}`;
}

export default function MapFilterScreen() {
  const router = useRouter();
  const params = useLocalSearchParams<{
    distanceKm?: string;
    minStartMinutes?: string;
    maxStartMinutes?: string;
    age?: string;
    categories?: string;
  }>();

  const bgColor = useThemeColor({}, "background");
  const borderColor = useThemeColor({}, "borderColor");
  const inputBg = useThemeColor({}, "inputBg");

  // ── Distance ──────────────────────────────────────────────────────────────
  const [distanceKm, setDistanceKm] = useState<number>(
    params.distanceKm ? parseInt(params.distanceKm, 10) : MAX_DISTANCE
  );

  // ── Start-time bounds ─────────────────────────────────────────────────────
  const initMinMins = params.minStartMinutes
    ? parseInt(params.minStartMinutes, 10)
    : null;
  const initMaxMins = params.maxStartMinutes
    ? parseInt(params.maxStartMinutes, 10)
    : null;

  const [minStartMinutes, setMinStartMinutes] = useState<number | null>(
    initMinMins
  );
  const [maxStartMinutes, setMaxStartMinutes] = useState<number | null>(
    initMaxMins
  );
  const [showMinPicker, setShowMinPicker] = useState(false);
  const [showMaxPicker, setShowMaxPicker] = useState(false);

  // ── Age ───────────────────────────────────────────────────────────────────
  const [age, setAge] = useState<number>(
    params.age ? parseInt(params.age, 10) : 0
  );

  // ── Categories ────────────────────────────────────────────────────────────
  const initCategories =
    params.categories && params.categories.length > 0
      ? params.categories.split(",").filter(Boolean)
      : [];
  const [categories, setCategories] = useState<string[]>(initCategories);
  const [categoryInput, setCategoryInput] = useState("");
  const inputRef = useRef<TextInput>(null);

  function addCategory() {
    const trimmed = categoryInput.trim().toLowerCase();
    if (trimmed && !categories.includes(trimmed)) {
      setCategories((prev) => [...prev, trimmed]);
    }
    setCategoryInput("");
  }

  function removeCategory(cat: string) {
    setCategories((prev) => prev.filter((c) => c !== cat));
  }

  // ── Apply ─────────────────────────────────────────────────────────────────
  function handleShowResults() {
    router.navigate({
      pathname: "/(app)/(tabs)/map",
      params: {
        distanceKm: String(distanceKm),
        minStartMinutes:
          minStartMinutes != null ? String(minStartMinutes) : "",
        maxStartMinutes:
          maxStartMinutes != null ? String(maxStartMinutes) : "",
        age: String(age),
        categories: categories.join(","),
      },
    });
  }

  return (
    <SafeAreaView className="flex-1" style={{ backgroundColor: bgColor }}>
      {/* Header */}
      <View
        className="flex-row items-center border-b px-4 py-3"
        style={{ borderBottomColor: borderColor }}
      >
        <TouchableOpacity onPress={() => router.back()} className="p-1 pr-3">
          <IconSymbol name="chevron.left" size={22} color={AppColors.primaryText} />
        </TouchableOpacity>
        <ThemedText className="font-nunito-bold text-xl">Filters</ThemedText>
      </View>

      <KeyboardAvoidingView
        className="flex-1"
        behavior={Platform.OS === "ios" ? "padding" : undefined}
      >
        <ScrollView
          className="flex-1"
          contentContainerStyle={{ paddingHorizontal: 20, paddingTop: 24, paddingBottom: 40 }}
          keyboardShouldPersistTaps="handled"
        >
          {/* ── Distance ── */}
          <Section title="Distance" borderColor={borderColor}>
            <ThemedText className="font-nunito text-sm mb-3" style={{ color: AppColors.mutedText }}>
              {distanceKm >= MAX_DISTANCE
                ? "Any distance"
                : `Within ${distanceKm} km`}
            </ThemedText>
            <Slider
              minimumValue={1}
              maximumValue={MAX_DISTANCE}
              step={1}
              value={distanceKm}
              onValueChange={setDistanceKm}
              minimumTrackTintColor={AppColors.checkboxSelected}
              maximumTrackTintColor={AppColors.borderLight}
              thumbTintColor={AppColors.checkboxSelected}
            />
            <View className="flex-row justify-between mt-1">
              <ThemedText className="font-nunito text-xs" style={{ color: AppColors.mutedText }}>
                1 km
              </ThemedText>
              <ThemedText className="font-nunito text-xs" style={{ color: AppColors.mutedText }}>
                40+ km
              </ThemedText>
            </View>
          </Section>

          {/* ── Start Time ── */}
          <Section title="Start Time" borderColor={borderColor}>
            <ThemedText className="font-nunito text-sm mb-4" style={{ color: AppColors.mutedText }}>
              Show only events that start within this time range.
            </ThemedText>

            {/* Min time */}
            <View className="flex-row items-center justify-between mb-3">
              <ThemedText className="font-nunito-semibold text-sm">Earliest</ThemedText>
              <View className="flex-row items-center gap-2">
                <TouchableOpacity
                  onPress={() => {
                    setShowMinPicker(!showMinPicker);
                    setShowMaxPicker(false);
                  }}
                  className="rounded-lg px-4 py-2"
                  style={{ backgroundColor: inputBg }}
                >
                  <Text className="font-nunito text-sm" style={{ color: AppColors.primaryText }}>
                    {minStartMinutes != null
                      ? formatTime(minStartMinutes)
                      : "Not set"}
                  </Text>
                </TouchableOpacity>
                {minStartMinutes != null && (
                  <TouchableOpacity onPress={() => setMinStartMinutes(null)}>
                    <IconSymbol name="xmark.circle.fill" size={20} color={AppColors.mutedText} />
                  </TouchableOpacity>
                )}
              </View>
            </View>

            {showMinPicker && (
              <DateTimePicker
                value={minutesToDate(minStartMinutes ?? 480)}
                mode="time"
                display={Platform.OS === "ios" ? "spinner" : "default"}
                onChange={(_: DateTimePickerEvent, date?: Date) => {
                  if (Platform.OS === "android") setShowMinPicker(false);
                  if (date) setMinStartMinutes(dateToMinutes(date));
                }}
              />
            )}

            {/* Max time */}
            <View className="flex-row items-center justify-between">
              <ThemedText className="font-nunito-semibold text-sm">Latest</ThemedText>
              <View className="flex-row items-center gap-2">
                <TouchableOpacity
                  onPress={() => {
                    setShowMaxPicker(!showMaxPicker);
                    setShowMinPicker(false);
                  }}
                  className="rounded-lg px-4 py-2"
                  style={{ backgroundColor: inputBg }}
                >
                  <Text className="font-nunito text-sm" style={{ color: AppColors.primaryText }}>
                    {maxStartMinutes != null
                      ? formatTime(maxStartMinutes)
                      : "Not set"}
                  </Text>
                </TouchableOpacity>
                {maxStartMinutes != null && (
                  <TouchableOpacity onPress={() => setMaxStartMinutes(null)}>
                    <IconSymbol name="xmark.circle.fill" size={20} color={AppColors.mutedText} />
                  </TouchableOpacity>
                )}
              </View>
            </View>

            {showMaxPicker && (
              <DateTimePicker
                value={minutesToDate(maxStartMinutes ?? 1200)}
                mode="time"
                display={Platform.OS === "ios" ? "spinner" : "default"}
                onChange={(_: DateTimePickerEvent, date?: Date) => {
                  if (Platform.OS === "android") setShowMaxPicker(false);
                  if (date) setMaxStartMinutes(dateToMinutes(date));
                }}
              />
            )}
          </Section>

          {/* ── Age ── */}
          <Section title="Age" borderColor={borderColor}>
            <ThemedText className="font-nunito text-sm mb-3" style={{ color: AppColors.mutedText }}>
              {age === 0
                ? "Any age"
                : age >= MAX_AGE
                ? "Age 12+"
                : `Age ${age}`}
            </ThemedText>
            <Slider
              minimumValue={0}
              maximumValue={MAX_AGE}
              step={1}
              value={age}
              onValueChange={setAge}
              minimumTrackTintColor={AppColors.checkboxSelected}
              maximumTrackTintColor={AppColors.borderLight}
              thumbTintColor={AppColors.checkboxSelected}
            />
            <View className="flex-row justify-between mt-1">
              <ThemedText className="font-nunito text-xs" style={{ color: AppColors.mutedText }}>
                Any
              </ThemedText>
              <ThemedText className="font-nunito text-xs" style={{ color: AppColors.mutedText }}>
                12+
              </ThemedText>
            </View>
          </Section>

          {/* ── Categories ── */}
          <Section title="Categories" borderColor={borderColor} isLast>
            <ThemedText className="font-nunito text-sm mb-3" style={{ color: AppColors.mutedText }}>
              Show orgs with events matching any of these categories.
            </ThemedText>
            <View className="flex-row items-center gap-2 mb-3">
              <TextInput
                ref={inputRef}
                value={categoryInput}
                onChangeText={setCategoryInput}
                onSubmitEditing={addCategory}
                placeholder="e.g. science, music…"
                placeholderTextColor={AppColors.placeholderText}
                returnKeyType="done"
                className="flex-1 rounded-lg px-3 py-2 font-nunito text-sm"
                style={{
                  backgroundColor: inputBg,
                  color: AppColors.primaryText,
                  borderColor,
                  borderWidth: 1,
                }}
              />
              <TouchableOpacity
                onPress={addCategory}
                className="rounded-lg px-4 py-2"
                style={{ backgroundColor: AppColors.checkboxSelected }}
              >
                <Text className="font-nunito-semibold text-sm text-white">Add</Text>
              </TouchableOpacity>
            </View>

            {categories.length > 0 && (
              <View className="flex-row flex-wrap gap-2">
                {categories.map((cat) => (
                  <View
                    key={cat}
                    className="flex-row items-center gap-1 rounded-full px-3 py-1"
                    style={{ backgroundColor: AppColors.surfaceGray }}
                  >
                    <Text
                      className="font-nunito text-sm"
                      style={{ color: AppColors.primaryText }}
                    >
                      {cat}
                    </Text>
                    <TouchableOpacity onPress={() => removeCategory(cat)}>
                      <IconSymbol
                        name="xmark"
                        size={12}
                        color={AppColors.mutedText}
                      />
                    </TouchableOpacity>
                  </View>
                ))}
              </View>
            )}
          </Section>
        </ScrollView>
      </KeyboardAvoidingView>

      {/* Show Results button */}
      <View
        className="border-t px-5 py-4"
        style={{ borderTopColor: borderColor, backgroundColor: bgColor }}
      >
        <TouchableOpacity
          onPress={handleShowResults}
          className="w-full items-center rounded-full py-4"
          style={{ backgroundColor: AppColors.checkboxSelected }}
        >
          <Text className="font-nunito-bold text-base text-white">
            Show Results
          </Text>
        </TouchableOpacity>
      </View>
    </SafeAreaView>
  );
}

function Section({
  title,
  children,
  borderColor,
  isLast = false,
}: {
  title: string;
  children: React.ReactNode;
  borderColor: string;
  isLast?: boolean;
}) {
  return (
    <View
      className={isLast ? "pb-4" : "pb-6 mb-6 border-b"}
      style={isLast ? undefined : { borderBottomColor: borderColor }}
    >
      <ThemedText className="font-nunito-bold text-base mb-4">{title}</ThemedText>
      {children}
    </View>
  );
}
