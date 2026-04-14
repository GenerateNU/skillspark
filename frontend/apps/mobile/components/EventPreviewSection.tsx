import { AppColors, FontFamilies } from "@/constants/theme";
import { formatModalTime } from "@/utils/format";
import { EventOccurrence } from "@skillspark/api-client";
import { Image } from "expo-image";
import { useTranslation } from "react-i18next";
import { Text, View } from "react-native";

export function EventPreviewSection({
  occurrence,
  titleOverride,
}: {
  occurrence: EventOccurrence;
  titleOverride?: string;
}) {
  const { t: translate } = useTranslation();
  const timeLabel = translate("occurrence.classTime", {
    time: formatModalTime(occurrence.start_time),
  });
  return (
    <View className="items-center mb-6">
      <View
        className="w-[120px] h-[120px] rounded-2xl overflow-hidden mb-4"
        style={{ backgroundColor: AppColors.imagePlaceholder }}
      >
        {occurrence.event.presigned_url ? (
          <Image
            source={{ uri: occurrence.event.presigned_url }}
            style={{ width: "100%", height: "100%" }}
            contentFit="cover"
          />
        ) : null}
      </View>
      <Text
        className="text-xl text-center mb-1"
        style={{ fontFamily: FontFamilies.bold, color: AppColors.primaryText }}
      >
        {titleOverride ?? occurrence.event.title}
      </Text>
      <Text
        className="text-sm text-center mb-3"
        style={{ fontFamily: FontFamilies.regular, color: AppColors.mutedText }}
      >
        {timeLabel}
      </Text>
      {!!occurrence.event.description && (
        <Text
          className="text-sm text-center leading-5"
          style={{
            fontFamily: FontFamilies.regular,
            color: AppColors.secondaryText,
          }}
        >
          {occurrence.event.description}
        </Text>
      )}
    </View>
  );
}
