import { AppColors, Colors } from "@/constants/theme";
import { useColorScheme } from "@/hooks/use-color-scheme.web";
import { useState } from "react";
import { View, Text, TouchableOpacity, Modal, ViewStyle } from "react-native";

interface DropdownOption {
  label: string;
  value: string;
}

interface DropdownProps {
  value: string;
  onChange: (value: string) => void;
  options: DropdownOption[];
  placeholder?: string;
  style?: ViewStyle;
}

export const Dropdown = ({ value, onChange, options, placeholder, style }: DropdownProps) => {
  const [isOpen, setOpen] = useState(false);
  const selected = options.find(o => o.value === value);
  const colorScheme = useColorScheme();
  const colors = Colors[colorScheme ?? "light"];

  return (
    <>
      <TouchableOpacity
        onPress={() => setOpen(true)}
        className="w-full flex-row justify-between items-center rounded-lg p-[10px]"
        style={[{ borderWidth: 1, borderColor: colors.borderColor }, style]}
      >
        <Text className="text-base" style={{ color: AppColors.placeholderText }}>
          {selected ? selected.label : placeholder}
        </Text>
        <Text style={{ color: "#9ca3af" }}>▼</Text>
      </TouchableOpacity>

      <Modal visible={isOpen} transparent animationType="fade">
        <TouchableOpacity
          className="flex-1 justify-center p-6"
          style={{ backgroundColor: "rgba(0,0,0,0.1)" }}
          onPress={() => setOpen(false)}
        >
          <View 
            className="rounded-lg overflow-hidden"
            style={{ backgroundColor: colors.dropdownBg }}
          >
            {options.map(option => (
              <TouchableOpacity
                key={option.value}
                onPress={() => {
                  onChange(option.value);
                  setOpen(false);
                }}
                className="p-4"
                style={{
                  borderBottomWidth: 1,
                  borderBottomColor: colors.borderColor,
                  backgroundColor: colors.dropdownBg,
                }}
              >
                <Text className="text-base" style={{ color: colors.text }}>
                  {option.label}
                </Text>
              </TouchableOpacity>
            ))}
          </View>
        </TouchableOpacity>
      </Modal>
    </>
  );
};