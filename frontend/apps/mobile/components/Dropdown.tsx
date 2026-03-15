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

  return (
    <>
      <TouchableOpacity
        onPress={() => setOpen(true)}
        style={[
          {
            width: "100%",
            borderWidth: 1,
            borderColor: "#d1d5db",
            borderRadius: 8,
            padding: 10,
            flexDirection: "row",
            justifyContent: "space-between",
            alignItems: "center",
          },
          style,
        ]}
      >
        <Text style={{ fontSize: 16, color: "#9ca3af" }}>
          {selected ? selected.label : placeholder}
        </Text>
        <Text style={{ color: "#9ca3af" }}>▼</Text>
      </TouchableOpacity>

      <Modal visible={isOpen} transparent animationType="fade">
        <TouchableOpacity
          style={{ flex: 1, backgroundColor: "rgba(0,0,0,0.1)", justifyContent: "center", padding: 24 }}
          onPress={() => setOpen(false)}
        >
          <View style={{ backgroundColor: "white", borderRadius: 8, overflow: "hidden" }}>
            {options.map(option => (
              <TouchableOpacity
                key={option.value}
                onPress={() => {
                  onChange(option.value);
                  setOpen(false);
                }}
                style={{
                  padding: 16,
                  borderBottomWidth: 1,
                  borderBottomColor: "#f3f4f6",
                  backgroundColor: "white",
                }}
              >
                <Text style={{ fontSize: 16, color: "#000" }}>
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