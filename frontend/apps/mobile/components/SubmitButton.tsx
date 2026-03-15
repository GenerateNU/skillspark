import { TouchableOpacity, Text } from "react-native";

interface SubmitButtonProps {
  label: string;
  onPress: () => void;
}

export const SubmitButton = ({ label, onPress }: SubmitButtonProps) => {
  return (
    <TouchableOpacity
      style={{
        backgroundColor: "#3b82f6",
        borderRadius: 8,
        padding: 10,
        width: "100%",
        alignItems: "center",
      }}
      onPress={onPress}
      activeOpacity={0.5}
    >
      <Text style={{ color: "white", fontSize: 16, fontWeight: "500" }}>
        {label}
      </Text>
    </TouchableOpacity>
  );
};
