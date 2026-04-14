import { AppColors } from "@/constants/theme";
import React from "react";
import { ScrollView, Text, TouchableOpacity } from "react-native";

type FilterTabOption<T> = {
  label: string;
  value: T;
};

type FilterTabsProps<T> = {
  options: FilterTabOption<T>[];
  value: T;
  onChange: (value: T) => void;
};

export function FilterTabs<T>({ options, value, onChange }: FilterTabsProps<T>) {
  return (
    <ScrollView
      horizontal
      showsHorizontalScrollIndicator={false}
      className="px-5 flex-grow-0 py-3"
      contentContainerStyle={{ gap: 12, paddingHorizontal: 4 }}
    >
      {options.map((opt) => {
        const active = value === opt.value;

        return (
          <TouchableOpacity
            key={String(opt.value)}
            onPress={() => onChange(opt.value)}
            className="px-4 py-2 rounded-full border"
            style={{
              backgroundColor: active
                ? AppColors.primaryText
                : "transparent",
              borderColor: active
                ? AppColors.primaryText
                : AppColors.borderLight,
            }}
          >
            <Text
              className="text-sm"
              style={{
                color: active ? "#fff" : AppColors.secondaryText,
              }}
            >
              {opt.label}
            </Text>
          </TouchableOpacity>
        );
      })}
    </ScrollView>
  );
}