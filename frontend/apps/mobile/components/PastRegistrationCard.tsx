import { ThemedText } from "@/components/themed-text";
import { AppColors } from "@/constants/theme";
import { useTranslation } from "react-i18next";
import { Image, TouchableOpacity, View } from "react-native";
import {
  type RegistrationCardProps,
  formatTime,
  formatDate,
} from "@/components/RegistrationCard.types";
import { ChildAvatar } from "@/components/ChildAvatar";

export function PastRegistrationCard({ data }: RegistrationCardProps) {
  const { t } = useTranslation();
  const children = data.childRegistrations.map((cr) => cr.child);
  const priceDisplay = `฿${(data.price / 100).toLocaleString()}`;

  return (
    <View
      className="w-11/12 rounded-xl overflow-hidden mb-4 flex-row"
      style={{
        height: 150,
        borderWidth: 1,
        borderColor: AppColors.borderLight,
        backgroundColor: AppColors.white,
        shadowColor: "#000",
        shadowOffset: { width: 0, height: 1 },
        shadowOpacity: 0.06,
        shadowRadius: 4,
        elevation: 2,
      }}
    >
      <View className="py-3 pl-3">
        <Image
          source={{ uri: data.image_uri }}
          style={{ width: 100, flex: 1, borderRadius: 8 }}
          resizeMode="cover"
        />
      </View>

      <View className="flex-1 px-3 py-3 justify-between">
        <ThemedText type="subtitle" numberOfLines={2}>
          {data.title}
        </ThemedText>
        <View>
          <ThemedText
            className="text-sm"
            style={{ color: AppColors.mutedText }}
          >
            {formatDate(data.start_time)}
          </ThemedText>
          <ThemedText
            className="text-sm"
            style={{ color: AppColors.mutedText }}
          >
            {formatTime(data.start_time)} – {formatTime(data.end_time)}
          </ThemedText>
        </View>
        <ThemedText className="text-sm" style={{ color: AppColors.mutedText }}>
          {t("activity.payment", { price: priceDisplay })}
        </ThemedText>
      </View>

      <View className="py-3 pr-3 items-end justify-between">
        <View
          className="flex-row flex-wrap gap-1 justify-end"
          style={{ maxWidth: 80 }}
        >
          {children.map((child) => (
            <ChildAvatar
              key={child.id}
              name={child.name}
              avatarFace={child.avatar_face}
              avatarBackground={child.avatar_background}
              size={28}
            />
          ))}
        </View>
        {data.hasReviewed ? (
          <View
            className="px-4 py-2 rounded-full"
            style={{ backgroundColor: AppColors.borderLight }}
          >
            <ThemedText
              className="text-sm font-medium"
              style={{ color: AppColors.mutedText }}
            >
              {t("activity.reviewed")}
            </ThemedText>
          </View>
        ) : (
          <TouchableOpacity
            onPress={data.onClickReview}
            className="px-6 py-2 rounded-full bg-black"
            activeOpacity={0.7}
          >
            <ThemedText lightColor="white" className="text-sm font-medium">
              {t("activity.review")}
            </ThemedText>
          </TouchableOpacity>
        )}
      </View>
    </View>
  );
}
