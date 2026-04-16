import { useState } from "react";
import { Pressable, Text } from "react-native";
import { useTranslation } from "react-i18next";
import { AppColors } from "@/constants/theme";

interface ExpandableTextProps {
  text: string;
  maxLines?: number;
}

/** Text block that truncates at `maxLines` with a toggleable See More / See Less button. */
export function ExpandableText({ text, maxLines = 4 }: ExpandableTextProps) {
  const { t } = useTranslation();
  const [expanded, setExpanded] = useState(false);
  const [truncated, setTruncated] = useState(false);

  return (
    <>
      <Text
        numberOfLines={expanded ? undefined : maxLines}
        onTextLayout={(e) => {
          if (!expanded) setTruncated(e.nativeEvent.lines.length >= maxLines);
        }}
        className={`text-sm leading-[22px] font-nunito ${truncated ? "mb-1" : "mb-4"}`}
        style={{ color: AppColors.secondaryText }}
      >
        {text}
      </Text>
      {truncated && (
        <Pressable onPress={() => setExpanded((prev) => !prev)} className="mb-4">
          <Text className="text-[13px] font-semibold" style={{ color: AppColors.primaryText }}>
            {expanded ? t("event.seeLess") : t("event.seeMore")}
          </Text>
        </Pressable>
      )}
    </>
  );
}
