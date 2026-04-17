import { Pressable, Text, View } from "react-native";
import { AppColors, FontFamilies } from "@/constants/theme";

export function CircleIcon({
  bg,
  children,
  label,
  onPress,
}: {
  bg: string;
  children: React.ReactNode;
  label: string;
  onPress: () => void;
}) {
  return (
    <Pressable
      onPress={onPress}
      className="items-center gap-1.5"
      style={{ width: 64 }}
    >
      <View
        className="w-[54px] h-[54px] rounded-full items-center justify-center"
        style={{ backgroundColor: bg }}
      >
        {children}
      </View>
      <Text
        numberOfLines={1}
        style={{
          fontFamily: FontFamilies.regular,
          fontSize: 11,
          color: AppColors.secondaryText,
        }}
      >
        {label}
      </Text>
    </Pressable>
  );
}
