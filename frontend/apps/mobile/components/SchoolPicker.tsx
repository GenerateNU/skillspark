import React, { useState } from "react";
import { View, TouchableOpacity, ScrollView, Modal } from "react-native";
import { ThemedText } from "@/components/themed-text";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { AppColors, Colors } from "@/constants/theme";
import { useGetAllSchools, School } from "@skillspark/api-client";
import { useTranslation } from "react-i18next";

type SchoolPickerProps = {
  value: string;
  onChange: (schoolId: string) => void;
};

export function SchoolPicker({ value, onChange }: SchoolPickerProps) {
  const [showDrop, setShowDrop] = useState(false);
  const theme = Colors.light;
  const { t: translate } = useTranslation();

  const { data, isLoading, isError } = useGetAllSchools();
  const schools = Array.isArray(data?.data) ? data.data : [];
  const selectedSchool = schools.find((s: School) => s.id === value);

  const placeholderLabel = isLoading
    ? translate("childProfile.loadingSchools")
    : isError
      ? translate("childProfile.failedToLoadSchools")
      : translate("childProfile.selectSchool");

  return (
    <View className="mb-6">
      <TouchableOpacity
        className="rounded-[10px] px-4 py-[14px] flex-row items-center justify-between bg-[#F3F4F6]"
        onPress={() => setShowDrop(true)}
        disabled={isLoading || isError}
      >
        <ThemedText
          className={`font-nunito ${selectedSchool ? "" : "text-[#9CA3AF]"}`}
        >
          {selectedSchool ? selectedSchool.name : placeholderLabel}
        </ThemedText>
        <IconSymbol
          name={showDrop ? "chevron.up" : "chevron.down"}
          size={16}
          color={AppColors.mutedText}
        />
      </TouchableOpacity>

      <Modal
        visible={showDrop}
        transparent
        animationType="fade"
        onRequestClose={() => setShowDrop(false)}
      >
        <TouchableOpacity
          className="flex-1 justify-center p-6 bg-black/10"
          onPress={() => setShowDrop(false)}
          activeOpacity={1}
        >
          <View
            className="rounded-[10px] overflow-hidden border"
            style={{
              backgroundColor: theme.dropdownBg,
              borderColor: theme.borderColor,
              shadowColor: "#000",
              shadowOpacity: 0.1,
              shadowRadius: 8,
              shadowOffset: { width: 0, height: 2 },
            }}
          >
            <ScrollView nestedScrollEnabled className="max-h-[200px]">
              {schools.map((school) => (
                <TouchableOpacity
                  key={school.id}
                  className="px-4 py-3 border-b border-b-[#E5E7EB]"
                  onPress={() => {
                    onChange(school.id);
                    setShowDrop(false);
                  }}
                >
                  <ThemedText>{school.name}</ThemedText>
                </TouchableOpacity>
              ))}
              {schools.length === 0 && !isLoading && (
                <View className="px-4 py-3">
                  <ThemedText className="text-[#6B7280]">
                    {translate("childProfile.noSchoolsFound")}
                  </ThemedText>
                </View>
              )}
            </ScrollView>
          </View>
        </TouchableOpacity>
      </Modal>
    </View>
  );
}
