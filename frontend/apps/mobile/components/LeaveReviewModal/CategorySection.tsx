import { AppColors } from "@/constants/theme";
import { Text, View } from "react-native";
import { CategoryPill } from "./CategoryPill";

interface Props {
  title: string;
  options: { value: string; labelKey: string }[];
  selected: string[];
  onToggle: (value: string) => void;
  translate: (key: string) => string;
}

export function CategorySection({
  title,
  options,
  selected,
  onToggle,
  translate,
}: Props) {
  return (
    <>
      <Text
        className="text-base font-nunito-bold mb-3"
        style={{ color: AppColors.primaryText }}
      >
        {title}
      </Text>
      <View className="flex-row flex-wrap gap-2 mb-5">
        {options.map(({ value, labelKey }) => (
          <CategoryPill
            key={value}
            label={translate(labelKey)}
            selected={selected.includes(value)}
            onPress={() => onToggle(value)}
          />
        ))}
      </View>
    </>
  );
}
