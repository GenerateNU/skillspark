import React from "react";
import { View, TouchableOpacity } from "react-native";
import { ThemedText } from "@/components/themed-text";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { InterestTags } from "./InterestTags";
import { Child } from "@skillspark/api-client";
import { AppColors } from "@/constants/theme";
import { useTranslation } from "react-i18next";
import { ChildAvatar } from "@/components/ChildAvatar";

type ChildListItemProps = {
  child: Child;
  onPress?: () => void;
};

export function ChildListItem({ child, onPress }: ChildListItemProps) {
  const { t: translate } = useTranslation();
  return (
    <TouchableOpacity
      className="flex-row items-start py-3 gap-3"
      onPress={onPress}
      activeOpacity={0.7}
    >
      <ChildAvatar
        name={child.name}
        avatarFace={child.avatar_face}
        avatarBackground={child.avatar_background}
        size={44}
      />
      <View className="flex-1 gap-[2px]">
        <ThemedText className="text-base font-nunito-semibold">
          {child.name}
        </ThemedText>
        <ThemedText
          className="text-[13px] font-nunito"
          style={{ color: AppColors.mutedText }}
        >
          {child.birth_month ? `${child.birth_month}, ` : ""}
          {child.birth_year}
        </ThemedText>
        {child.interests && (
          <ThemedText
            className="text-[13px] font-nunito"
            style={{ color: AppColors.mutedText }}
          >
            {translate("familyInformation.interests")}
          </ThemedText>
        )}
        <InterestTags interests={child.interests} />
      </View>
      <IconSymbol name="chevron.right" size={18} color={AppColors.subtleText} />
    </TouchableOpacity>
  );
}
