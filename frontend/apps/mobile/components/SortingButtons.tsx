import { AppColors } from "@/constants/theme";
import { useState } from "react";
import { ScrollView, Text, TouchableOpacity } from "react-native";

interface FilterOption<T> {
  label: string;
  value: T;
}

interface FilterTabsProps<T> {
  options: FilterOption<T>[];
  onChange: (value: T) => void;
}

export function FilterTabs<T>({ options, onChange }: FilterTabsProps<T>) {
  const [selected, setSelected] = useState<T>(options[0]?.value);

  function handleSelect(value: T) {
    setSelected(value);
    onChange(value);
  }

  return (
    <ScrollView
      horizontal
      showsHorizontalScrollIndicator={false}
      className="px-5 flex-grow-0 py-3"
      contentContainerStyle={{ gap: 12, paddingHorizontal: 4 }}
    >
      {options.map((opt) => {
        const active = selected === opt.value;
        return (
          <TouchableOpacity
            key={String(opt.value)}
            onPress={() => handleSelect(opt.value)}
            className="px-4 py-2 rounded-full border"
            style={{
              backgroundColor: active ? AppColors.primaryText : "transparent",
              borderColor: active
                ? AppColors.primaryText
                : AppColors.borderLight,
            }}
          >
            <Text
              className="text-sm"
              style={{ color: active ? "#fff" : AppColors.secondaryText }}
            >
              {opt.label}
            </Text>
          </TouchableOpacity>
        );
      })}
    </ScrollView>
  );
}
