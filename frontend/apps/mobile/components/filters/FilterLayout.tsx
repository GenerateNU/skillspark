import { AppColors, FontFamilies, FontSizes } from "@/constants/theme";
import { Text, View } from "react-native";

export function SectionLabel({ label }: { label: string }) {
  return (
    <Text
      className="mb-3"
      style={{
        fontFamily: FontFamilies.bold,
        fontSize: FontSizes.lg,
        color: AppColors.primaryText,
      }}
    >
      {label}
    </Text>
  );
}

export function Divider() {
  return (
    <View
      className="my-5 h-px"
      style={{ backgroundColor: AppColors.divider }}
    />
  );
}
