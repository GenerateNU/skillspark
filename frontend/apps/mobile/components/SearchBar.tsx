import { Pressable, TextInput, View, type ViewStyle } from "react-native";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { AppColors } from "@/constants/theme";

export function SearchBar({
  value,
  onChangeText,
  placeholder,
  style,
}: {
  value: string;
  onChangeText: (text: string) => void;
  placeholder?: string;
  style?: ViewStyle;
}) {
  return (
    <View className="px-5" style={style}>
      <View
        className="flex-row items-center rounded-full px-4 py-[10px]"
        style={{ backgroundColor: AppColors.surfaceGray }}
      >
        <IconSymbol
          name="magnifyingglass"
          size={18}
          color={AppColors.subtleText}
          style={{ marginRight: 8 }}
        />
        <TextInput
          className="flex-1 text-sm font-nunito"
          style={{ color: AppColors.primaryText }}
          placeholder={placeholder}
          placeholderTextColor={AppColors.placeholderText}
          value={value}
          onChangeText={onChangeText}
        />
        <Pressable
          className="w-9 h-9 rounded-full items-center justify-center"
          style={{ backgroundColor: AppColors.primaryText }}
        >
          <IconSymbol name="slider.horizontal.3" size={16} color={AppColors.white} />
        </Pressable>
      </View>
    </View>
  );
}
