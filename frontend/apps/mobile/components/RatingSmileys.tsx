import { ThemedText } from "@/components/themed-text";
import { ThemedView } from "@/components/themed-view";
import { RATING_OPTIONS } from "@/constants/ratings";
import { useState } from "react";
import { useTranslation } from "react-i18next";
import { Image, TouchableOpacity } from "react-native";

export function RatingSmileys({
  onSelect,
}: {
  onSelect: (rating: number) => void;
}) {
  const { t: translate } = useTranslation();
  const [selected, setSelected] = useState<number | null>(null);

  function handleSelect(rating: number) {
    setSelected(rating);
    onSelect(rating);
  }

  return (
    <ThemedView className="flex-row justify-between">
      {RATING_OPTIONS.map(({ rating, image, labelKey }) => (
        <TouchableOpacity
          key={rating}
          onPress={() => handleSelect(rating)}
          className="flex-1 items-center gap-2"
          style={{
            opacity: selected === null || selected === rating ? 1 : 0.3,
          }}
        >
          <Image source={image} style={{ width: 55, height: 55 }} />
          <ThemedText className="text-sm text-center">
            {translate(labelKey)}
          </ThemedText>
        </TouchableOpacity>
      ))}
    </ThemedView>
  );
}
