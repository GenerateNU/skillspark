import { TouchableOpacity, Text } from "react-native";

interface PageRedirectButtonProps {
  label: string;
  onPress: () => void;
}

export const PageRedirectButton = ({
  label,
  onPress,
}: PageRedirectButtonProps) => {
  return (
    <TouchableOpacity
      style={{
        borderRadius: 8,
        padding: 10,
        width: "100%",
        alignItems: "center",
      }}
      onPress={onPress}
      activeOpacity={0.5}
    >
      <Text style={{ color: "#3b82f6", fontSize: 16, fontWeight: "500" }}>
        {label}
      </Text>
    </TouchableOpacity>
  );
};
