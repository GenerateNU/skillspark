import {
  TouchableOpacity,
  Text,
  DimensionValue,
  ColorValue,
} from "react-native";

interface ButtonProps {
  label: string;
  onPress: () => void;
  disabled: boolean;
  bgColor?: ColorValue;
  width?: DimensionValue;
  textColor?: ColorValue;
}

export const Button = ({
  label,
  onPress,
  disabled,
  bgColor,
  width,
  textColor,
}: ButtonProps) => {
  return (
    <TouchableOpacity
      className={"rounded-3xl py-4 items-center shadow-sm"}
      style={{
        opacity: disabled ? 0.5 : 1,
        backgroundColor: bgColor ?? "#1B1B1B",
        width: width ?? "95%",
      }}
      onPress={onPress}
      activeOpacity={0.7}
      disabled={disabled}
    >
      <Text
        className="text-lg font-nunito-bold"
        style={{
          color: textColor ?? "#FFFFFF",
        }}
      >
        {label}
      </Text>
    </TouchableOpacity>
  );
};
