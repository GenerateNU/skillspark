import React, { useRef, useState } from "react";
import {
  View,
  TextInput,
  TouchableOpacity,
  ScrollView,
  Modal,
} from "react-native";
import { ThemedText } from "@/components/themed-text";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { AppColors, TAG_COLORS, Colors } from "@/constants/theme";
import { SchoolPicker } from "@/components/SchoolPicker";
import { useTranslation } from "react-i18next";
import { DEFAULT_AVATAR_COLOR } from "@/components/AvatarPicker";
import { ChildAvatar } from "@/components/ChildAvatar";

const INTEREST_OPTIONS = [
  "science",
  "math",
  "music",
  "art",
  "sports",
  "technology",
  "language",
  "other",
];

const capitalize = (s: string) => s.charAt(0).toUpperCase() + s.slice(1);

export const MONTHS = [
  "January",
  "February",
  "March",
  "April",
  "May",
  "June",
  "July",
  "August",
  "September",
  "October",
  "November",
  "December",
];

export const YEARS = Array.from({ length: 30 }, (_, i) =>
  String(new Date().getFullYear() - i)
);

type DropdownLayout = { x: number; y: number; width: number };

type DropdownModalProps = {
  visible: boolean;
  onClose: () => void;
  layout: DropdownLayout | null;
  options: string[];
  onSelect: (value: string) => void;
};

function DropdownModal({
  visible,
  onClose,
  layout,
  options,
  onSelect,
}: DropdownModalProps) {
  const theme = Colors.light;
  if (!layout) return null;

  return (
    <Modal
      visible={visible}
      transparent
      animationType="none"
      onRequestClose={onClose}
    >
      <TouchableOpacity
        style={{ flex: 1 }}
        activeOpacity={1}
        onPress={onClose}
      >
        <View
          style={{
            position: "absolute",
            top: layout.y,
            left: layout.x,
            width: layout.width,
            backgroundColor: theme.dropdownBg,
            borderRadius: 10,
            borderWidth: 1,
            borderColor: theme.borderColor,
            shadowColor: "#000",
            shadowOpacity: 0.12,
            shadowRadius: 8,
            shadowOffset: { width: 0, height: 2 },
            elevation: 8,
            maxHeight: 220,
            overflow: "hidden",
          }}
        >
          <ScrollView bounces={false}>
            {options.map((opt) => (
              <TouchableOpacity
                key={opt}
                style={{
                  paddingHorizontal: 16,
                  paddingVertical: 12,
                  borderBottomWidth: 1,
                  borderBottomColor: "#E5E7EB",
                }}
                onPress={() => {
                  onSelect(opt);
                  onClose();
                }}
              >
                <ThemedText>{opt}</ThemedText>
              </TouchableOpacity>
            ))}
          </ScrollView>
        </View>
      </TouchableOpacity>
    </Modal>
  );
}

export type ChildProfileFormProps = {
  firstName: string;
  setFirstName: (v: string) => void;
  lastName: string;
  setLastName: (v: string) => void;
  birthMonth: string;
  setBirthMonth: (v: string) => void;
  birthYear: string;
  setBirthYear: (v: string) => void;
  schoolId: string;
  setSchoolId: (v: string) => void;
  interests: string[];
  setInterests: React.Dispatch<React.SetStateAction<string[]>>;
  searchQuery: string;
  setSearchQuery: (v: string) => void;
  showMonthDrop: boolean;
  setShowMonthDrop: (v: boolean) => void;
  showYearDrop: boolean;
  setShowYearDrop: (v: boolean) => void;
  avatarFace: string | null;
  avatarBackground: string;
  onAvatarPress: () => void;
};

export function ChildProfileForm({
  firstName,
  setFirstName,
  lastName,
  setLastName,
  birthMonth,
  setBirthMonth,
  birthYear,
  setBirthYear,
  schoolId,
  setSchoolId,
  interests,
  setInterests,
  searchQuery,
  setSearchQuery,
  showMonthDrop,
  setShowMonthDrop,
  showYearDrop,
  setShowYearDrop,
  avatarFace,
  avatarBackground,
  onAvatarPress,
}: ChildProfileFormProps) {
  const theme = Colors.light;
  const { t: translate } = useTranslation();

  const monthTriggerRef = useRef<View>(null);
  const yearTriggerRef = useRef<View>(null);
  const [monthDropLayout, setMonthDropLayout] = useState<DropdownLayout | null>(null);
  const [yearDropLayout, setYearDropLayout] = useState<DropdownLayout | null>(null);

  const openMonthDrop = () => {
    monthTriggerRef.current?.measure((_fx: number, _fy: number, width: number, height: number, px: number, py: number) => {
      setMonthDropLayout({ x: px, y: py + height + 4, width });
      setShowMonthDrop(true);
      setShowYearDrop(false);
    });
  };

  const openYearDrop = () => {
    yearTriggerRef.current?.measure((_fx: number, _fy: number, width: number, height: number, px: number, py: number) => {
      setYearDropLayout({ x: px, y: py + height + 4, width });
      setShowYearDrop(true);
      setShowMonthDrop(false);
    });
  };

  const removeInterest = (tag: string) =>
    setInterests((prev) => prev.filter((i) => i !== tag));

  const toggleInterest = (item: string) => {
    setInterests((prev) =>
      prev.includes(item) ? prev.filter((i) => i !== item) : [...prev, item]
    );
  };

  const filteredOptions = INTEREST_OPTIONS.filter((o) =>
    o.toLowerCase().includes(searchQuery.toLowerCase())
  );

  return (
    <>
      {/* Avatar selection */}
      <View className="items-center mb-5">
        <TouchableOpacity
          onPress={onAvatarPress}
          className="items-center gap-2"
          activeOpacity={0.7}
        >
          <View style={{ position: "relative" }}>
            <ChildAvatar
              name={[firstName, lastName].filter(Boolean).join(" ") || "?"}
              avatarFace={avatarFace}
              avatarBackground={avatarBackground || DEFAULT_AVATAR_COLOR}
              size={72}
            />
            <View
              style={{
                position: "absolute",
                bottom: 0,
                right: 0,
                width: 22,
                height: 22,
                borderRadius: 11,
                backgroundColor: theme.text,
                alignItems: "center",
                justifyContent: "center",
              }}
            >
              <IconSymbol name="pencil" size={11} color={theme.background} />
            </View>
          </View>
          <ThemedText
            className="text-sm font-nunito-semibold"
            style={{ color: AppColors.mutedText }}
          >
            {translate("childProfile.changeProfilePicture", {
              defaultValue: "Change Profile Picture",
            })}
          </ThemedText>
        </TouchableOpacity>
      </View>
      <TextInput
        className="rounded-[10px] px-4 py-[14px] text-base font-nunito mb-3 bg-[#F3F4F6] text-[#11181C]"
        value={firstName}
        onChangeText={setFirstName}
        placeholder={translate("childProfile.firstName")}
        placeholderTextColor={AppColors.placeholderText}
      />
      <TextInput
        className="rounded-[10px] px-4 py-[14px] text-base font-nunito mb-3 bg-[#F3F4F6] text-[#11181C]"
        value={lastName}
        onChangeText={setLastName}
        placeholder={translate("childProfile.lastName")}
        placeholderTextColor={AppColors.placeholderText}
      />

      {/* Month / Year row */}
      <View className="flex-row gap-3 mb-6">
        {/* Month trigger */}
        <View className="flex-1">
          <View ref={monthTriggerRef}>
            <TouchableOpacity
              className="rounded-[10px] px-4 py-[14px] flex-row items-center justify-between bg-[#F3F4F6]"
              onPress={openMonthDrop}
            >
              <ThemedText className={birthMonth ? "" : "font-nunito text-[#9CA3AF]"}>
                {birthMonth || translate("childProfile.month")}
              </ThemedText>
              <IconSymbol name="chevron.down" size={16} color={AppColors.mutedText} />
            </TouchableOpacity>
          </View>
        </View>

        {/* Year trigger */}
        <View className="flex-1">
          <View ref={yearTriggerRef}>
            <TouchableOpacity
              className="rounded-[10px] px-4 py-[14px] flex-row items-center justify-between bg-[#F3F4F6]"
              onPress={openYearDrop}
            >
              <ThemedText className={birthYear ? "" : "font-nunito text-[#9CA3AF]"}>
                {birthYear || translate("childProfile.year")}
              </ThemedText>
              <IconSymbol name="chevron.down" size={16} color={AppColors.mutedText} />
            </TouchableOpacity>
          </View>
        </View>
      </View>

      {/* Month dropdown modal */}
      <DropdownModal
        visible={showMonthDrop}
        onClose={() => setShowMonthDrop(false)}
        layout={monthDropLayout}
        options={MONTHS}
        onSelect={setBirthMonth}
      />

      {/* Year dropdown modal */}
      <DropdownModal
        visible={showYearDrop}
        onClose={() => setShowYearDrop(false)}
        layout={yearDropLayout}
        options={YEARS}
        onSelect={setBirthYear}
      />

      <SchoolPicker value={schoolId} onChange={setSchoolId} />

      <ThemedText className="text-base font-nunito-semibold mb-3">
        {translate("familyInformation.interests")}
      </ThemedText>

      {interests.length > 0 && (
        <ScrollView
          horizontal
          showsHorizontalScrollIndicator={false}
          className="mb-3"
          contentContainerStyle={{ gap: 8, paddingRight: 4 }}
        >
          {interests.map((tag, idx) => {
            const color = TAG_COLORS[idx % TAG_COLORS.length];
            return (
              <TouchableOpacity
                key={tag}
                className="flex-row items-center px-2 py-1 rounded-full border gap-1"
                style={{ backgroundColor: color.bg, borderColor: color.border }}
                onPress={() => removeInterest(tag)}
              >
                <IconSymbol name="camera.filters" size={13} color={color.border} />
                <ThemedText
                  className="text-xs font-nunito-medium"
                  style={{ color: color.text }}
                >
                  {translate(`interests.${tag}`, { defaultValue: capitalize(tag) })}
                </ThemedText>
              </TouchableOpacity>
            );
          })}
        </ScrollView>
      )}

      <View className="border rounded-[10px] overflow-hidden mb-6 border-[#E5E7EB]">
        <View className="flex-row items-center px-4 py-3 gap-2">
          <TextInput
            className="flex-1 text-base font-nunito text-[#11181C]"
            value={searchQuery}
            onChangeText={setSearchQuery}
            placeholder={translate("childProfile.searchInterests")}
            placeholderTextColor={AppColors.placeholderText}
          />
          <IconSymbol name="magnifyingglass" size={20} color={AppColors.mutedText} />
        </View>
        <View className="h-px bg-[#E5E7EB]" />
        <View
          onStartShouldSetResponder={() => true}
          onMoveShouldSetResponder={() => true}
        >
          <ScrollView
            nestedScrollEnabled
            showsVerticalScrollIndicator
            className="max-h-[150px]"
          >
            {filteredOptions.map((item) => (
              <TouchableOpacity
                key={item}
                className="flex-row items-center justify-between px-4 py-4 border-b border-b-[#F3F4F6]"
                onPress={() => toggleInterest(item)}
              >
                <ThemedText className="text-base font-nunito">
                  {translate(`interests.${item}`, { defaultValue: capitalize(item) })}
                </ThemedText>
                <View
                  className="w-[22px] h-[22px] rounded-[4px] border-[1.5px] items-center justify-center"
                  style={{
                    borderColor: interests.includes(item)
                      ? AppColors.checkboxSelected
                      : AppColors.subtleText,
                  }}
                >
                  {interests.includes(item) && (
                    <IconSymbol
                      name="checkmark"
                      size={12}
                      color={AppColors.checkboxSelected}
                    />
                  )}
                </View>
              </TouchableOpacity>
            ))}
          </ScrollView>
        </View>
      </View>
    </>
  );
}