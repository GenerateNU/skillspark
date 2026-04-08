import { RATING_OPTIONS } from "@/constants/ratings";
import {
  CONTENT_OPTIONS,
  DIFFICULTY_OPTIONS,
  INSTRUCTOR_OPTIONS,
  VALUE_OPTIONS,
} from "@/constants/review-tags";
import { AppColors } from "@/constants/theme";
import { Image } from "expo-image";
import { Text, TouchableOpacity, View } from "react-native";
import { CategorySection } from "./CategorySection";

interface Props {
  rating: number | null;
  setRating: (r: number) => void;
  selectedCategories: string[];
  toggleCategory: (value: string) => void;
  translate: (key: string) => string;
  onNext: () => void;
}

export function ReviewStep1({
  rating,
  setRating,
  selectedCategories,
  toggleCategory,
  translate,
  onNext,
}: Props) {
  return (
    <>
      <Text
        className="text-xl font-nunito-bold mb-5"
        style={{ color: AppColors.primaryText }}
      >
        {translate("review.howWasExperience")}
      </Text>

      <View className="flex-row justify-between mb-7">
        {RATING_OPTIONS.map(({ rating: r, image, labelKey }) => (
          <TouchableOpacity
            key={r}
            onPress={() => setRating(r)}
            className="flex-1 items-center gap-2"
            style={{
              opacity: rating === null || rating === r ? 1 : 0.3,
            }}
          >
            <Image source={image} style={{ width: 55, height: 55 }} />
            <Text
              className="text-sm text-center"
              style={{ color: AppColors.secondaryText }}
            >
              {translate(labelKey)}
            </Text>
          </TouchableOpacity>
        ))}
      </View>

      <CategorySection
        title={translate("review.content")}
        options={CONTENT_OPTIONS}
        selected={selectedCategories}
        onToggle={toggleCategory}
        translate={translate}
      />

      <CategorySection
        title={translate("review.difficulty")}
        options={DIFFICULTY_OPTIONS}
        selected={selectedCategories}
        onToggle={toggleCategory}
        translate={translate}
      />

      <CategorySection
        title={translate("review.instructor")}
        options={INSTRUCTOR_OPTIONS}
        selected={selectedCategories}
        onToggle={toggleCategory}
        translate={translate}
      />

      <CategorySection
        title={translate("review.value")}
        options={VALUE_OPTIONS}
        selected={selectedCategories}
        onToggle={toggleCategory}
        translate={translate}
      />

      <TouchableOpacity
        onPress={onNext}
        disabled={!rating}
        className="py-4 rounded-2xl items-center mt-3"
        style={{
          backgroundColor: rating ? AppColors.primaryText : AppColors.borderLight,
        }}
      >
        <Text
          className="text-base font-nunito-bold"
          style={{ color: rating ? "#fff" : AppColors.subtleText }}
        >
          {translate("review.next")}
        </Text>
      </TouchableOpacity>
    </>
  );
}
