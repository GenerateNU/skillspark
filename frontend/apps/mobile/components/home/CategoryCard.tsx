import { memo, useCallback, useMemo } from "react";
import { View, Pressable, Text } from "react-native";
import { AppColors } from "@/constants/theme";
import { useRouter } from "expo-router";
import { Image } from "expo-image";
import { useTranslation } from "react-i18next";

const sportsImage = require("@/assets/images/categories/sports-and-physical-activities.png");
const artsAndCreativeImage = require("@/assets/images/categories/arts-and-creative-expression.png");
const languagesImage = require("@/assets/images/categories/languages.png");
const academicsImage = require("@/assets/images/categories/academics.png");
const lifeSkillsImage = require("@/assets/images/categories/life-skills.png");
const musicImage = require("@/assets/images/categories/music-and-performance.png");
const mathImage = require("@/assets/images/categories/math.png");
const techImage = require("@/assets/images/categories/tech-and-innovation.png");

const CATEGORY_SOURCES: Record<
  string,
  { uri: number; translationKey: string }
> = {
  "Sports & Physical Activities": {
    uri: sportsImage,
    translationKey: "dashboard.categories.sports",
  },
  "Arts & Creative Expression": {
    uri: artsAndCreativeImage,
    translationKey: "dashboard.categories.art",
  },
  Languages: {
    uri: languagesImage,
    translationKey: "dashboard.categories.languages",
  },
  Academics: {
    uri: academicsImage,
    translationKey: "dashboard.categories.academics",
  },
  "Personal Development & Life Skills": {
    uri: lifeSkillsImage,
    translationKey: "dashboard.categories.lifeSkills",
  },
  "Music & Performance": {
    uri: musicImage,
    translationKey: "dashboard.categories.music",
  },
  Math: { uri: mathImage, translationKey: "dashboard.categories.math" },
  "Tech & Innovation": {
    uri: techImage,
    translationKey: "dashboard.categories.technology",
  },
};

export const CategoryCard = memo(function CategoryCard({
  category,
}: {
  category: string;
}) {
  const router = useRouter();
  const { t: translate } = useTranslation();
  const source = useMemo(() => CATEGORY_SOURCES[category], [category]);
  const handlePress = useCallback(
    () => router.push({ pathname: "/event-categories", params: { category } }),
    [category, router]
  );

  return (
    <Pressable
      onPress={handlePress}
      className="flex-1 m-[5px]"
      style={{
        shadowColor: "#000",
        shadowOpacity: 0.25,
        shadowRadius: 4,
        shadowOffset: { width: 0, height: 4 },
        elevation: 3,
      }}
    >
      <View className="h-[80px] rounded-[15px] overflow-hidden">
        {source != null ? (
          <Image
            source={source.uri}
            style={{ width: "100%", height: "100%" }}
          />
        ) : (
          <View
            className="absolute inset-0"
            style={{ backgroundColor: AppColors.categoryFallback }}
          />
        )}
        <View className="absolute inset-0 justify-end p-3">
          <Text className="text-white text-xs font-bold">
            {source != null ? translate(source.translationKey) : category}
          </Text>
        </View>
      </View>
    </Pressable>
  );
});
