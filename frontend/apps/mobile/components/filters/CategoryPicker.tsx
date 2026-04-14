import { AppColors, FontFamilies, FontSizes } from "@/constants/theme";
import { useState } from "react";
import { Text, TouchableOpacity, View } from "react-native";
import { useTranslation } from "react-i18next";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { PillRow } from "./PillRow";
import { SectionLabel } from "./FilterLayout";

const COLLAPSED_COUNT = 8;

type Props = {
  label: string;
  categoryLabels: string[];
  activeIndex: number;
  onSelect: (idx: number) => void;
};

export function CategoryPicker({ label, categoryLabels, activeIndex, onSelect }: Props) {
  const { t } = useTranslation();
  const [expanded, setExpanded] = useState(false);

  const visibleLabels = expanded ? categoryLabels : categoryLabels.slice(0, COLLAPSED_COUNT);
  const visibleActiveIndex = expanded ? activeIndex : activeIndex < COLLAPSED_COUNT ? activeIndex : -1;

  return (
    <View>
      <SectionLabel label={label} />
      <PillRow
        options={visibleLabels}
        activeIndex={visibleActiveIndex}
        onSelect={onSelect}
      />
      <TouchableOpacity
        onPress={() => setExpanded((v) => !v)}
        activeOpacity={0.7}
        className="flex-row items-center gap-1 mt-2 self-start"
      >
        <Text
          style={{
            fontFamily: FontFamilies.semiBold,
            fontSize: FontSizes.sm,
            color: AppColors.primaryBlue,
          }}
        >
          {expanded ? t("filters.showLess") : t("filters.showMore")}
        </Text>
        <IconSymbol
          name={expanded ? "chevron.up" : "chevron.down"}
          size={14}
          color={AppColors.primaryBlue}
        />
      </TouchableOpacity>
    </View>
  );
}
