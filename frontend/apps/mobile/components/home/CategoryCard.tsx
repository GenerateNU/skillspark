import { memo, useEffect, useState } from "react";
import { Image as RNImage, View, Pressable } from "react-native";
import { SvgXml } from "react-native-svg";
import { AppColors } from "@/constants/theme";
import { useRouter } from "expo-router";

const CATEGORY_URIS: Record<string, string> = Object.fromEntries(
  Object.entries({
    art: require("@/assets/images/art.svg"),
    sports: require("@/assets/images/sports.svg"),
    music: require("@/assets/images/music.svg"),
    technology: require("@/assets/images/tech.svg"),
    science: require("@/assets/images/study.svg"),
    math: require("@/assets/images/math.svg"),
    language: require("@/assets/images/talking.svg"),
    other: require("@/assets/images/life_skills.svg"),
  }).map(([cat, moduleId]) => [
    cat,
    RNImage.resolveAssetSource(moduleId as number).uri,
  ]),
);

// Module-level XML content cache — survives remounts, persists for the app's lifetime
const svgXmlCache = new Map<string, string>();

// Eagerly fetch all SVG content when this module is first imported
Object.values(CATEGORY_URIS).forEach((uri) => {
  if (uri && !svgXmlCache.has(uri)) {
    fetch(uri)
      .then((r) => r.text())
      .then((xml) => svgXmlCache.set(uri, xml))
      .catch(() => {});
  }
});

function useCategoryXml(uri: string | undefined): string | null {
  const [xml, setXml] = useState<string | null>(
    uri ? (svgXmlCache.get(uri) ?? null) : null,
  );

  useEffect(() => {
    if (!uri) return;
    const cached = svgXmlCache.get(uri);
    if (cached) {
      setXml(cached);
      return;
    }
    fetch(uri)
      .then((r) => r.text())
      .then((text) => {
        svgXmlCache.set(uri, text);
        setXml(text);
      })
      .catch(() => {});
  }, [uri]);

  return xml;
}

export const CategoryCard = memo(function CategoryCard({
  category,
}: {
  category: string;
}) {
  const router = useRouter();
  const xml = useCategoryXml(CATEGORY_URIS[category]);

  return (
    <Pressable
      onPress={() =>
        router.push({ pathname: "/event-categories", params: { category } })
      }
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
        {xml ? (
          <SvgXml xml={xml} width="100%" height="100%" />
        ) : (
          <View
            className="absolute inset-0"
            style={{ backgroundColor: AppColors.categoryFallback }}
          />
        )}
      </View>
    </Pressable>
  );
});
