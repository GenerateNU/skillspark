import React from "react";
import { View, TextInput, TouchableOpacity, ScrollView } from "react-native";
import { ThemedText } from "@/components/themed-text";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { AppColors, TAG_COLORS, Colors } from "@/constants/theme";
import { useColorScheme } from "@/hooks/use-color-scheme";
import { SchoolPicker } from "@/components/SchoolPicker";
import { useTranslation } from "react-i18next";
import { AvatarPicker, DEFAULT_AVATAR_COLOR } from "@/components/AvatarPicker";
import { ChildAvatar, getInitials } from "@/components/ChildAvatar";

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

export const YEARS = Array.from({ length: 20 }, (_, i) =>
  String(new Date().getFullYear() - i),
);

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
  setAvatarFace: (v: string | null) => void;
  avatarBackground: string;
  setAvatarBackground: (v: string) => void;
  showAvatarPicker: boolean;
  setShowAvatarPicker: (v: boolean) => void;
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
  setAvatarFace,
  avatarBackground,
  setAvatarBackground,
  showAvatarPicker,
  setShowAvatarPicker,
}: ChildProfileFormProps) {
  const colorScheme = useColorScheme();
  const theme = Colors[colorScheme ?? "light"];
  const { t: translate } = useTranslation();

  const removeInterest = (tag: string) =>
    setInterests((prev) => prev.filter((i) => i !== tag));

  const toggleInterest = (item: string) => {
    setInterests((prev) =>
      prev.includes(item) ? prev.filter((i) => i !== item) : [...prev, item],
    );
  };

  const filteredOptions = INTEREST_OPTIONS.filter((o) =>
    o.toLowerCase().includes(searchQuery.toLowerCase()),
  );

  return (
    <>
      {/* Avatar selection */}
      {(() => {
        const fullName = [firstName, lastName].filter(Boolean).join(" ");
        const initials = getInitials(fullName || "?");
        return (
          <>
            <View className="items-center mb-5">
              <TouchableOpacity
                onPress={() => setShowAvatarPicker(!showAvatarPicker)}
                className="items-center gap-2"
                activeOpacity={0.7}
              >
                <ChildAvatar
                  name={fullName || "?"}
                  avatarFace={avatarFace}
                  avatarBackground={avatarBackground || DEFAULT_AVATAR_COLOR}
                  size={72}
                />
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
            {showAvatarPicker && (
              <AvatarPicker
                selectedFace={avatarFace}
                selectedBackground={avatarBackground || DEFAULT_AVATAR_COLOR}
                onFaceChange={setAvatarFace}
                onBackgroundChange={setAvatarBackground}
                childInitials={initials}
              />
            )}
          </>
        );
      })()}
      <TextInput
        className="rounded-[10px] px-4 py-[14px] text-base font-nunito mb-3 bg-[#F3F4F6] dark:bg-[#27272a] text-[#11181C] dark:text-[#ECEDEE]"
        value={firstName}
        onChangeText={setFirstName}
        placeholder={translate("childProfile.firstName")}
        placeholderTextColor={AppColors.placeholderText}
      />
      <TextInput
        className="rounded-[10px] px-4 py-[14px] text-base font-nunito mb-3 bg-[#F3F4F6] dark:bg-[#27272a] text-[#11181C] dark:text-[#ECEDEE]"
        value={lastName}
        onChangeText={setLastName}
        placeholder={translate("childProfile.lastName")}
        placeholderTextColor={AppColors.placeholderText}
      />
      <View className="flex-row gap-3 mb-6 z-10">
        {/* Month picker */}
        <View className="flex-1 z-10">
          <TouchableOpacity
            className="rounded-[10px] px-4 py-[14px] flex-row items-center justify-between bg-[#F3F4F6] dark:bg-[#27272a]"
            onPress={() => {
              setShowMonthDrop(!showMonthDrop);
              setShowYearDrop(false);
            }}
          >
            <ThemedText
              className={birthMonth ? "" : "font-nunito text-[#9CA3AF]"}
            >
              {birthMonth || translate("childProfile.month")}
            </ThemedText>
            <IconSymbol
              name="chevron.down"
              size={16}
              color={AppColors.mutedText}
            />
          </TouchableOpacity>
          {showMonthDrop && (
            <View
              className="absolute left-0 right-0 top-[52px] rounded-[10px] border z-[100] elevation-5"
              style={{
                backgroundColor: theme.dropdownBg,
                borderColor: theme.borderColor,
                shadowColor: "#000",
                shadowOpacity: 0.1,
                shadowRadius: 8,
                shadowOffset: { width: 0, height: 2 },
              }}
            >
              <ScrollView nestedScrollEnabled className="max-h-[180px]">
                {MONTHS.map((m) => (
                  <TouchableOpacity
                    key={m}
                    className="px-4 py-3 border-b border-b-[#E5E7EB] dark:border-b-[#3f3f46]"
                    onPress={() => {
                      setBirthMonth(m);
                      setShowMonthDrop(false);
                    }}
                  >
                    <ThemedText>{m}</ThemedText>
                  </TouchableOpacity>
                ))}
              </ScrollView>
            </View>
          )}
        </View>
        {/* Year picker */}
        <View className="flex-1 z-10">
          <TouchableOpacity
            className="rounded-[10px] px-4 py-[14px] flex-row items-center justify-between bg-[#F3F4F6] dark:bg-[#27272a]"
            onPress={() => {
              setShowYearDrop(!showYearDrop);
              setShowMonthDrop(false);
            }}
          >
            <ThemedText
              className={birthYear ? "" : "font-nunito text-[#9CA3AF]"}
            >
              {birthYear || translate("childProfile.year")}
            </ThemedText>
            <IconSymbol
              name="chevron.down"
              size={16}
              color={AppColors.mutedText}
            />
          </TouchableOpacity>
          {showYearDrop && (
            <View
              className="absolute left-0 right-0 top-[52px] rounded-[10px] border z-[100] elevation-5"
              style={{
                backgroundColor: theme.dropdownBg,
                borderColor: theme.borderColor,
                shadowColor: "#000",
                shadowOpacity: 0.1,
                shadowRadius: 8,
                shadowOffset: { width: 0, height: 2 },
              }}
            >
              <ScrollView nestedScrollEnabled className="max-h-[180px]">
                {YEARS.map((y) => (
                  <TouchableOpacity
                    key={y}
                    className="px-4 py-3 border-b border-b-[#E5E7EB] dark:border-b-[#3f3f46]"
                    onPress={() => {
                      setBirthYear(y);
                      setShowYearDrop(false);
                    }}
                  >
                    <ThemedText>{y}</ThemedText>
                  </TouchableOpacity>
                ))}
              </ScrollView>
            </View>
          )}
        </View>
      </View>
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
                <IconSymbol
                  name="camera.filters"
                  size={13}
                  color={color.border}
                />
                <ThemedText
                  className="text-xs font-nunito-medium"
                  style={{ color: color.text }}
                >
                  {translate(`interests.${tag}`, {
                    defaultValue: capitalize(tag),
                  })}
                </ThemedText>
              </TouchableOpacity>
            );
          })}
        </ScrollView>
      )}
      <View className="border rounded-[10px] overflow-hidden mb-6 border-[#E5E7EB] dark:border-[#3f3f46]">
        <View className="flex-row items-center px-4 py-3 gap-2">
          <TextInput
            className="flex-1 text-base font-nunito text-[#11181C] dark:text-[#ECEDEE]"
            value={searchQuery}
            onChangeText={setSearchQuery}
            placeholder={translate("childProfile.searchInterests")}
            placeholderTextColor={AppColors.placeholderText}
          />
          <IconSymbol
            name="magnifyingglass"
            size={20}
            color={AppColors.mutedText}
          />
        </View>
        <View className="h-px bg-[#E5E7EB] dark:bg-[#3f3f46]" />
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
                className="flex-row items-center justify-between px-4 py-4 border-b border-b-[#F3F4F6] dark:border-b-[#27272a]"
                onPress={() => toggleInterest(item)}
              >
                <ThemedText className="text-base font-nunito">
                  {translate(`interests.${item}`, {
                    defaultValue: capitalize(item),
                  })}
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
