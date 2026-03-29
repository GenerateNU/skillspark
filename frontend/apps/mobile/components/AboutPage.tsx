import { Pressable, Text, View } from "react-native";
import { useCallback, useState } from "react";
import * as WebBrowser from "expo-web-browser";
import type { OrgLink } from "@skillspark/api-client";
import { AppColors } from "@/constants/theme";

function useOrgLinks(links: OrgLink[]) {
  const openLink = useCallback(async (href: string) => {
    await WebBrowser.openBrowserAsync(href);
  }, []);

  return { openLink, hasLinks: links.length > 0 };
}

interface AboutPageProps {
  description: string;
  links: OrgLink[];
}

export function AboutPage({ description, links }: AboutPageProps) {
  const { openLink, hasLinks } = useOrgLinks(links);
  const [expanded, setExpanded] = useState(false);
  const [truncated, setTruncated] = useState(false);

  return (
    <View>
      <Text style={{ fontSize: 18, fontWeight: "700", color: AppColors.primaryText, marginBottom: 10 }}>
        About
      </Text>

      <Text
        numberOfLines={expanded ? undefined : 4}
        onTextLayout={(e) => {
          if (!expanded) setTruncated(e.nativeEvent.lines.length >= 4);
        }}
        style={{ fontSize: 14, color: AppColors.secondaryText, lineHeight: 22, marginBottom: 4 }}
      >
        {description}
      </Text>

      {truncated && (
        <Pressable onPress={() => setExpanded((p) => !p)} style={{ marginBottom: 16 }}>
          <Text style={{ fontSize: 13, color: AppColors.primaryText, fontWeight: "600" }}>
            {expanded ? "See less" : "See more"}
          </Text>
        </Pressable>
      )}

      {!truncated && <View style={{ marginBottom: 16 }} />}

      {hasLinks && (
      <View style={{ flexDirection: "row", flexWrap: "wrap", gap: 10 }}>
        {links.map((link, index) => (
          <Pressable
            key={index}
            onPress={() => openLink(link.href)}
            style={({ pressed }) => ({
              borderRadius: 999,
              paddingHorizontal: 20,
              paddingVertical: 10,
              backgroundColor: pressed ? "#F3F4F6" : "#fff",
              alignItems: "center",
              shadowColor: "#000",
              shadowOpacity: 0.1,
              shadowRadius: 8,
              shadowOffset: { width: 0, height: 2 },
              elevation: 4,
            })}
          >
            <Text style={{ fontSize: 14, color: AppColors.primaryText, fontWeight: "600" }}>
              {link.label}
            </Text>
          </Pressable>
        ))}
      </View>
    )}
    </View>
  );
}