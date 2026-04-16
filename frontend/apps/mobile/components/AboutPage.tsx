import { Pressable, Text, View } from "react-native";
import { useState } from "react";
import type { OrgLink } from "@skillspark/api-client";
import { useOrgLinks } from "@/hooks/useOrgLinks";
import { useTranslation } from "react-i18next";

interface AboutPageProps {
  description: string;
  links: OrgLink[];
}

export function AboutPage({ description, links }: AboutPageProps) {
  const { openLink, hasLinks } = useOrgLinks(links);
  const { t } = useTranslation();
  const [expanded, setExpanded] = useState(false);
  const [truncated, setTruncated] = useState(false);

  return (
    <View>
      <Text
        numberOfLines={expanded ? undefined : 4}
        onTextLayout={(e) => {
          if (!expanded) setTruncated(e.nativeEvent.lines.length >= 4);
        }}
        className="text-sm text-gray-500 leading-relaxed mb-1"
      >
        {description}
      </Text>

      {truncated && (
        <Pressable onPress={() => setExpanded((p) => !p)} className="mb-4">
          <Text className="text-sm text-gray-900 font-semibold">
            {expanded ? t("event.seeLess") : t("event.seeMore")}
          </Text>
        </Pressable>
      )}

      {!truncated && <View className="mb-4" />}

      {hasLinks && (
        <View className="flex-row flex-wrap gap-2.5">
          {links.map((link, index) => (
            <Pressable
              key={index}
              onPress={() => openLink(link.href)}
              className="rounded-full px-5 py-2.5 bg-white items-center shadow"
              style={{
                shadowColor: "#000",
                shadowOpacity: 0.1,
                shadowRadius: 8,
                shadowOffset: { width: 0, height: 2 },
                elevation: 4,
              }}
            >
              <Text className="text-sm text-gray-900 font-semibold">
                {link.label}
              </Text>
            </Pressable>
          ))}
        </View>
      )}
    </View>
  );
}
